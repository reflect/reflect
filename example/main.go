package main

import (
	"flag"
	"github.com/kataras/iris"
	"github.com/reflect/reflect-go"
	"os"
	"github.com/iris-contrib/middleware/cors"
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

	/*
	// A list of username/password combos that will satisfy basic auth
	usersForAuth = map[string]string{
		"geoff": "beefsticks",
		"brad":  "beefsticks",
		"alex":  "beefsticks",
	}

	users = []User{
		{ Username: "geoff", CompanyId: "acme" },
		{ Username: "brad",  CompanyId: "microsoft" },
		{ Username: "alex",  CompanyId: "apple" },
	}
	*/
)

// A handler for the /user endpoint
func UserHandler(ctx *iris.Context) {
	/*
	username := ctx.GetString("user")

	userExists := false

	for k, v := range users {
		if v.Username == username {
			userExists = true
			thisUser = users[k]
		}
	}

	if !userExists {
		ctx.Text(404, "")
	}
	*/

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

func init() {
	flag.Parse()

	if *reflectSecretKey == "" {
		panic("You must supply a secret key!")
	}
}

func main() {
	//authMiddleware := basicauth.Default(usersForAuth)

	iris.Use(cors.Default())
	iris.Static("/static", "./static", 1)
	//iris.Get("/users", authMiddleware, UserHandler)
	iris.Get("/users", UserHandler)

	iris.Listen(":8080")
}
