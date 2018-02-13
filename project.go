package reflect

import (
	"time"

	"github.com/pborman/uuid"
	jose "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type ProjectClaims struct {
	ViewIdentifiers []string               `json:"http://reflect.io/s/v3/vid,omitempty"`
	Parameters      []Parameter            `json:"http://reflect.io/s/v3/p,omitempty"`
	Attributes      map[string]interface{} `json:"http://reflect.io/s/v3/a,omitempty"`
}

type ProjectTokenBuilder struct {
	accessKey string

	expiration time.Time
	claims     ProjectClaims
}

// WithExpiration sets the time at which this token will no longer be valid.
// All requests made using an expired token will fail.
func (ptb *ProjectTokenBuilder) WithExpiration(when time.Time) *ProjectTokenBuilder {
	ptb.expiration = when
	return ptb
}

// WithViewIdentifier restricts the views that can be loaded with this token.
// This method can be called multiple times, and will permit only the specified
// views to be loaded.
func (ptb *ProjectTokenBuilder) WithViewIdentifier(id string) *ProjectTokenBuilder {
	ptb.claims.ViewIdentifiers = append(ptb.claims.ViewIdentifiers, id)
	return ptb
}

// WithParameter adds a data-filtering parameter to the token. This method can
// be called multiple times to add many parameters.
func (ptb *ProjectTokenBuilder) WithParameter(ps Parameter) *ProjectTokenBuilder {
	ptb.claims.Parameters = append(ptb.claims.Parameters, ps)
	return ptb
}

// WithAttribute adds the given attribute to the token.
func (ptb *ProjectTokenBuilder) WithAttribute(name string, value interface{}) *ProjectTokenBuilder {
	if ptb.claims.Attributes == nil {
		ptb.claims.Attributes = make(map[string]interface{})
	}

	ptb.claims.Attributes[name] = value
	return ptb
}

// Build constructs a new encrypted token that encapsulates the secret
// information provided to this builder.
func (ptb *ProjectTokenBuilder) Build(secretKey string) (string, error) {
	secret := uuid.Parse(secretKey)
	if secret == nil {
		return "", ErrInvalidSecretKey
	}

	jwk := jose.JSONWebKey{
		KeyID: ptb.accessKey,
		Key:   []byte(secret),
	}

	encrypter, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{Algorithm: jose.DIRECT, Key: jwk},
		(&jose.EncrypterOptions{Compression: jose.DEFLATE}).WithType("JWT"),
	)
	if err != nil {
		return "", err
	}

	t := jwt.NewNumericDate(time.Now())

	builder := jwt.Encrypted(encrypter).
		Claims(jwt.Claims{
			IssuedAt:  t,
			NotBefore: t,
		}).
		Claims(ptb.claims)

	if !ptb.expiration.IsZero() {
		builder = builder.Claims(jwt.Claims{
			Expiry: jwt.NewNumericDate(ptb.expiration),
		})
	}

	return builder.CompactSerialize()
}

// NewProjectTokenBuilder creates a new builder that can be configured with
// secret data to be passed to the Reflect API.
func NewProjectTokenBuilder(accessKey string) *ProjectTokenBuilder {
	return &ProjectTokenBuilder{
		accessKey: accessKey,
	}
}
