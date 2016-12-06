package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	key := "def456"

	params := []Parameter{
		{Field: "Field1", Op: EqualsOperation, Value: "abc123"},
		{Field: "Field2", Op: EqualsOperation, AnyValues: []string{"ghi789", "jkl123"}},
	}

	tok := GenerateToken(key, params)
	assert.Equal(t, "=2=xpXaD9BvlZ+5BIB7kx+J10QBbRVqhqEiGdmbGd/lkKE=", tok)
}
