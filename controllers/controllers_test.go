package controllers

import (
	"api-hackathon/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTToken(t *testing.T) {
	user := models.User{
		ID:       1,
		Email:    "someemail@gmail.com",
		Password: "123456",
	}

	token, err := user.CreateToken()
	assert.NotEmpty(t, token)
	assert.Equal(t, err, nil)

	user = models.ExtractClaims(token)
	assert.Equal(t, user.Email, "someemail@gmail.com")
	assert.Equal(t, user.Password, "123456")
}
