# The Reflect Go Client

The Go library for [Reflect](https://reflect.io)

**Note**: At the moment, the Reflect Go client is used *only* for generating authentication tokens to pass to

## Installation

```bash
$ go get github.com/reflect/reflect-go
```

## Usage

```go
import (
        "github.com/reflect/reflect-go"
)
```

```go
reflectApiToken := "<Your API token>"
generatedToken := reflect.GenerateToken(reflectApiToken, )
```
