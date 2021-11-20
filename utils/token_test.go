package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	t.Run("test token", func(t *testing.T) {
		token := GenerateToken("admin@gmail.com")
		assert.NotEqual(t, token, "","token success generated")
	})
}
func TestValidateToken(t *testing.T) {
	t.Run("test token", func(t *testing.T) {
		token := GenerateToken("admin@gmail.com")
		isValidated, email := ValidateToken(token)
		assert.Equal(t, isValidated,true,"token is validated")
		assert.Equal(t, email,"admin@gmail.com","generated token email is same with value")
	})
}
