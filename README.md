jwt-maker
=========

A little commandline utility wrapper around [jwt-go][] to generate tokens.

[jwt-go]: https://github.com/dgrijalva/jwt-go

Usage
-----

```
Usage: ./jwt-maker <key> <json encoded data>
  -expires-in int
    The number of minutes for which the token should be good (default 1440)
  -signing-method string
    The signing method to use (default "HS256")
```

