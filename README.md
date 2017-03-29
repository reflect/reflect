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

At the moment, this library is used for generating auth tokens for use in Reflect views. To generate new tokens, pass in a [token](https://app.reflect.io/tokens) to the [`GenerateToken`](https://godoc.org/github.com/reflect/reflect-go#GenerateToken) function along with any number of [parameters](#Token-arameters):

```go
reflectSecretKey := "<Your project's secret key>"
generatedToken := reflect.GenerateToken(reflectApiToken, params)
```

## Token parameters

In addition to an API token, the [`GenerateToken`](https://godoc.org/github.com/reflect/reflect-go#GenerateToken) function also takes an array of [`Parameter`](https://godoc.org/github.com/reflect/reflect-go#Parameter)s that you can use to generate signed auth tokens. Here's an example:

```go
username, hobbies := "jane", []string{"fishing", "painting"}

params := []reflect.Param{
        {Field: "Username", Op: reflect.EqualsOperation, Value: username},
        {Field: "Hobbies", Op: reflect.EqualsOperation, AnyValue: hobbies},
}
generatedToken := reflect.GenerateToken(reflectApiToken, params)
```

There are currently [five parameter-building operations available](https://godoc.org/github.com/reflect/reflect-go#pkg-constants).

## Example

For a basic example of the Go Reflect library in action, see [`example/main.go`](/example/main.go).
