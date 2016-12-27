package main

import (
	"flag"
	"github.com/kataras/iris"
	"github.com/reflect/reflect-go"
	"os"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/basicauth"
	"net/http"
)


// A very simple user object
type User struct {
	Username        string `json:"username"`
	TopicOfInterest string `json:"topic"`
	ReflectApiToken string `json:"apiToken"`
}

var (
	//thisUser User

	reflectSecretKey = flag.String("reflect-secret-key", os.Getenv("REFLECT_SECRET_KEY"), "Secret")

	// A list of username/password combos that will satisfy basic auth
	authorizedUsers = map[string]string{
		"user": "pass",
	}
)

// A handler for the /user endpoint
func UserHandler(ctx *iris.Context) {
	username, password := ctx.GetString("user"), ctx.GetString("password")

	if username == "user" && password == "pass" {

		user := User{
			Username: "noah",
			TopicOfInterest: "linux",
		}

		usernameParam := reflect.Parameter{
			Field: "username",
			Op:    reflect.EqualsOperation,
			Value: user.Username,
		}

		topicParam := reflect.Parameter{
			Field: "topic",
			Op:    reflect.EqualsOperation,
			Value: user.TopicOfInterest,
		}

		tokenGenParams := []reflect.Parameter{usernameParam, topicParam}
		generatedToken := reflect.GenerateToken(*reflectSecretKey, tokenGenParams)
		user.ReflectApiToken = generatedToken

		ctx.JSON(200, user)
	}

	ctx.Error("", http.StatusUnauthorized)
}

func init() {
	flag.Parse()

	if *reflectSecretKey == "" {
		panic("You must supply a secret key!")
	}
}

func main() {
	iris.Logger.Printf("Using secret key: %s", *reflectSecretKey)

	authMiddleware := basicauth.Default(authorizedUsers)
	iris.Use(authMiddleware)

	iris.Use(cors.Default())
	iris.Static("/static", "./static", 1)
	iris.Get("/users", UserHandler)

	iris.Listen(":8080")
}
