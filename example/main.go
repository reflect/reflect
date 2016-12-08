package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/reflect/reflect-go"
	"flag"
	"os"
)

var (
	// Pull a Reflect API token into the application either through a CLI flag or an env var
	reflectApiToken = flag.String("reflect-api-token", os.Getenv("REFLECT_API_TOKEN"), "An API token for your Reflect account")

	// A list of username/password combos that will satisfy basic auth
	users = map[string]string{
		"geoff": "beefsticks",
	}
)

// A very simple user object
type User struct {
	Username        string `json:"username"`
	ReflectApiToken string `json:"reflectApiToken"`
}

// A handler for the /user endpoint
func UserHandler(ctx *iris.Context) {
	username := ctx.GetString("user")

	// We'll build a token-generating parameter out of the user's username
	usernameParam := reflect.Parameter{
		Field: "Username",
		Op: reflect.EqualsOperation,
		Value: username,
	}

	tokenParams := []reflect.Parameter{usernameParam}

	// Now we generate a user-specific token using the global API token and a parameter
	generatedToken := reflect.GenerateToken(*reflectApiToken, []reflect.Parameter{tokenParams})

	// Now we return a JSON object to the client with information about the user, including the
	// user-specific token
	user := User{
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

	router.Get("/user", authMiddleware, UserHandler)

	router.Listen(":8080")
}
