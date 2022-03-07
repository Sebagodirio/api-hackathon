package controllers

import (
	"api-hackathon/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := models.CheckUserInDB(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := user.CreateToken()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
