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

## Generating Authentication Tokens

At the moment, this library is used for generating authentication tokens for use in Reflect views. To generate new tokens:

```go
accessKey := "d232c1e5-6083-4aa7-9042-0547052cc5dd"
secretKey := "74678a9b-685c-4c14-ac45-7312fe29de06"

token, err := reflect.NewProjectTokenBuilder(accessKey).
    WithAttribute("user-id", 1234).
    WithAttribute("user-name", "Billy Bob").
    WithParameter(reflect.Parameter{
        Field: "My Field",
        Op: reflect.EqualsOperation,
        Value: "My Value",
    }).
    Build(secretKey)
```
