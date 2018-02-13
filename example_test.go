package reflect_test

import (
	"fmt"
	"time"

	"github.com/reflect/reflect-go"
)

func ExampleProjectTokenBuilder() {
	accessKey := "d232c1e5-6083-4aa7-9042-0547052cc5dd"
	secretKey := "74678a9b-685c-4c14-ac45-7312fe29de06"

	token, err := reflect.NewProjectTokenBuilder(accessKey).
		WithExpiration(time.Now().Add(1*time.Hour)).
		WithAttribute("user-id", 1234).
		WithAttribute("user-name", "Billy Bob").
		Build(secretKey)

	fmt.Println(token, err)
}
