package controllers

import (
	"api-hackathon/models"
	"api-hackathon/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHackathons GET godoc
// @Summary It returns all the hackathons with its delopments, which contains its developers
// @Accept json
// @Produce json
// @Param Request body models.User true "Email and password of the user"
// @Success 200 {object} string
// @Success 401 {object} string
// @Failure 503 {object} string
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := services.CheckUserInDB(user)
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
