# The Reflect Go Library

**GoDoc**: https://godoc.org/github.com/reflect/reflect-go

**Note**: At the moment, the Reflect Go client is used *only* for generating authentication tokens to pass to client-side Reflect views. In the future we will release a more fully-featured Go client for the [Reflect REST API](https://reflect.io/docs/reference/rest-api).

## Installation

You can install the `reflect-go` library using `go get`:

```bash
$ go get github.com/reflect/reflect-go
```

## Importing

```go
import (
        "github.com/reflect/reflect-go"
)
```

It's important to note that there is also a core Go library named [`reflect`](https://golang.org/pkg/reflect/). To avoid conflicts with this library, you can rename your `reflect-go` import:

```go
import (
        rf "github.com/reflect/reflect-go"
)
```

## Generating User Tokens

At the moment, this library is used for generating auth tokens for use in Reflect views. To generate new tokens, pass in a [token](https://app.reflect.io/tokens) to the [`GenerateToken`](https://godoc.org/github.com/reflect/reflect-go#GenerateToken) function along with any number of [parameters](#Parameters):

```go
reflectApiToken := "<Your API token>"
generatedToken := reflect.GenerateToken(reflectApiToken, params)
```
