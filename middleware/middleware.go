package middleware

import (
	"api-hackathon/models"
	"api-hackathon/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthRequired will get the token the user sends and it will check if the credentials are correct
func AuthRequired(ctx *gin.Context) {
	token := models.ExtractToken(ctx)
	user := models.ExtractClaims(token)
	err := services.CheckUserInDB(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "You are not allowed to use this endpoint")
		ctx.Abort()
	}
	ctx.Next()
}
