package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/reflect/reflect-go"
)

const (
	reflectApiToken = "644adcaa-6da2-4570-976a-5c2b9fa63be8"
)

var (
	users = map[string]string{
		"colby": "pass",
	}
)

type user struct {
	Username        string `json:"username"`
	ReflectApiToken string `json:"reflectApiToken"`
}

func userHandler(ctx *iris.Context) {
	username := ctx.Param("username")

	usernameParam := reflect.Parameter{
		Field: "Username",
		Op: reflect.EqualsOperation,
		Value: username,
	}

	generatedToken := reflect.GenerateToken(reflectApiToken, []reflect.Parameter{usernameParam})

	user := user{
		Username: username,
		ReflectApiToken: generatedToken,
	}

	ctx.JSON(200, user)
}

func main() {
	router := iris.New()

	authMiddleware := basicauth.Default(users)

	router.Get("/user/:username", authMiddleware, userHandler)

	router.Listen(":8080")
}
