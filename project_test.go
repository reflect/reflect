package reflect_test

import (
	"testing"
	"time"

	"github.com/pborman/uuid"
	reflect "github.com/reflect/reflect-go"
	"github.com/stretchr/testify/require"
	"gopkg.in/square/go-jose.v2/jwt"
)

type projectKey struct {
	accessKey string
	secretKey string
}

func newProjectKey(t *testing.T) projectKey {
	return projectKey{
		accessKey: uuid.New(),
		secretKey: uuid.New(),
	}
}

func testProjectTokenBuilder(t *testing.T, fn func(b *reflect.ProjectTokenBuilder)) (jwt.Claims, reflect.ProjectClaims) {
	key := newProjectKey(t)

	b := reflect.NewProjectTokenBuilder(key.accessKey)
	fn(b)

	et, err := b.Build(key.secretKey)
	require.NoError(t, err)

	token, err := jwt.ParseEncrypted(et)
	require.NoError(t, err)

	require.Len(t, token.Headers, 1)
	require.Equal(t, key.accessKey, token.Headers[0].KeyID)

	var ic jwt.Claims
	var pc reflect.ProjectClaims
	require.NoError(t, token.Claims([]byte(uuid.Parse(key.secretKey)), &ic, &pc))

	return ic, pc
}

func TestProjectTokenBuilderSimple(t *testing.T) {
	testProjectTokenBuilder(t, func(b *reflect.ProjectTokenBuilder) {})
}

func TestProjectTokenBuilderExpiration(t *testing.T) {
	ic, _ := testProjectTokenBuilder(t, func(b *reflect.ProjectTokenBuilder) {
		b.WithExpiration(time.Now().Add(15 * time.Minute))
	})

	require.NoError(t, ic.Validate(jwt.Expected{
		Time: time.Now(),
	}))
	require.Equal(t, jwt.ErrExpired, ic.Validate(jwt.Expected{
		Time: time.Now().Add(30 * time.Minute),
	}))
	require.Equal(t, jwt.ErrNotValidYet, ic.Validate(jwt.Expected{
		Time: time.Now().Add(-5 * time.Minute),
	}))
}

func TestProjectTokenBuilderClaims(t *testing.T) {
	p := reflect.Parameter{
		Field: "user-id",
		Op:    reflect.EqualsOperation,
		Value: "1234",
	}

	_, pc := testProjectTokenBuilder(t, func(b *reflect.ProjectTokenBuilder) {
		b.WithViewIdentifier("SecUr3View1D")
		b.WithAttribute("user-id", 1234)
		b.WithAttribute("user-name", "Billy Bob")
		b.WithParameter(p)
	})

	require.Equal(t, []string{"SecUr3View1D"}, pc.ViewIdentifiers)
	require.Equal(t, []reflect.Parameter{p}, pc.Parameters)
	require.Len(t, pc.Attributes, 2)
	require.Equal(t, float64(1234), pc.Attributes["user-id"])
	require.Equal(t, "Billy Bob", pc.Attributes["user-name"])
}
