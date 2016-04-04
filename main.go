package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s <key> <json encoded data>\n", os.Args[0])
	flag.PrintDefaults()
}

var expiry_minutes int
var signing_method string

func init() {
	flag.IntVar(&expiry_minutes, "expires-in", 24*60, "The number of minutes for which the token should be good")
	flag.StringVar(&signing_method, "signing-method", "HS256", "The signing method to use")
}

func main() {
	if len(os.Args) != 3 {
		Usage()
		os.Exit(-1)
	}

	flag.Parse()

	key := os.Args[1]
	dataString := os.Args[2]

	var data map[string]interface{}

	if err := json.Unmarshal([]byte(dataString), &data); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to parse json;", err)
		os.Exit(-1)
	}

	token := jwt.New(jwt.GetSigningMethod(signing_method))

	token.Claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiry_minutes)).Unix()

	for k, v := range data {
		token.Claims[k] = v
	}

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to generate string;", err)
		os.Exit(-1)
	}

	fmt.Println(tokenString)
}
