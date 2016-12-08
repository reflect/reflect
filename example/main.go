package main

import (
	"flag"
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
	"github.com/reflect/reflect-go"
	"os"
)

var (
	// Pull an access key and secret key for Reflect a into the application either through a CLI flag or an env var
	// reflectAccessKey = flag.String("reflect-access-key", os.Getenv("REFLECT_ACCESS_KEY"), "An API token for your Reflect account")
	reflectSecretKey = flag.String("reflect-secret-key", os.Getenv("REFLECT_SECRET_KEY"), "Secret")

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
		Op:    reflect.EqualsOperation,
		Value: username,
	}

	tokenParams := []reflect.Parameter{usernameParam}

	// Now we generate a user-specific token using the global API token and a parameter
	generatedToken := reflect.GenerateToken(*reflectSecretKey, tokenParams)

	// Now we return a JSON object to the client with information about the user, including the
	// user-specific token
	user := User{
		Username:        username,
		ReflectApiToken: generatedToken,
	}

	ctx.JSON(200, user)
}

func init() {
	flag.Parse()

	if *reflectSecretKey == "" {
		panic("You must supply a secret key!")
	}
}

func main() {
	router := iris.New()

	authMiddleware := basicauth.Default(users)

	router.Get("/user", authMiddleware, UserHandler)

	router.Listen(":8080")
}
