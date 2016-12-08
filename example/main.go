package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/reflect/reflect-go"
	"flag"
	"os"
)
var (
	reflectApiToken = flag.String("reflect-api-token", os.Getenv("REFLECT_API_TOKEN"), "An API token for your Reflect account")

	users = map[string]string{
		"geoff": "beefsticks",
	}
)

type user struct {
	Username        string `json:"username"`
	ReflectApiToken string `json:"reflectApiToken"`
}

func userHandler(ctx *iris.Context) {
	username := ctx.GetString("user")

	usernameParam := reflect.Parameter{
		Field: "Username",
		Op: reflect.EqualsOperation,
		Value: username,
	}

	generatedToken := reflect.GenerateToken(*reflectApiToken, []reflect.Parameter{usernameParam})

	user := user{
		Username: username,
		ReflectApiToken: generatedToken,
	}

	ctx.JSON(200, user)
}

func init() {
	flag.Parse()

	if *reflectApiToken == "" {
		panic("You must supply an API token!")
	}
}

func main() {
	router := iris.New()

	authMiddleware := basicauth.Default(users)

	router.Get("/user", authMiddleware, userHandler)

	router.Listen(":8080")
}
