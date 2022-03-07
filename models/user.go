package models

import (
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) CreateToken() (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = u.ID
	atClaims["email"] = u.Email
	atClaims["password"] = u.Password
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractToken(ctx *gin.Context) string {
	bearToken := ctx.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractClaims(tokenStr string) User {
	var user User
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})

	if err != nil {
		return user
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user.Email = claims["email"].(string)
		user.Password = claims["password"].(string)
		return user
	}
	return user
}
