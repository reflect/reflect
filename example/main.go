package main

import (
	"flag"
	"github.com/iris-contrib/middleware/basicauth"
	"github.com/kataras/iris"
	"github.com/reflect/reflect-go"
	"os"
)


// A very simple user object
type User struct {
	Username        string `json:"username"`
	CompanyId       string `json:"companyId"`
	ReflectApiToken string `json:"apiToken"`
}

var (
	thisUser User

	reflectSecretKey = flag.String("reflect-secret-key", os.Getenv("REFLECT_SECRET_KEY"), "Secret")

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
)

// A handler for the /user endpoint
func UserHandler(ctx *iris.Context) {
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

	usernameParam := reflect.Parameter{
		Field: "username",
		Op:    reflect.EqualsOperation,
		Value: thisUser.Username,
	}

	companyParam := reflect.Parameter{
		Field: "companyId",
		Op:    reflect.EqualsOperation,
		Value: thisUser.CompanyId,
	}

	tokenGenParams := []reflect.Parameter{usernameParam, companyParam}
	generatedToken := reflect.GenerateToken(*reflectSecretKey, tokenGenParams)
	thisUser.ReflectApiToken = generatedToken

	ctx.JSON(200, thisUser)
}

func init() {
	flag.Parse()

	if *reflectSecretKey == "" {
		panic("You must supply a secret key!")
	}
}

func main() {
	router := iris.New()

	authMiddleware := basicauth.Default(usersForAuth)

	router.StaticServe("./webapp", "/")
	router.Get("/user", authMiddleware, UserHandler)

	router.Listen(":8080")
}
