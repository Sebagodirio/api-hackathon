package controllers

import (
	"api-hackathon/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDevelopments(ctx *gin.Context) {
	developments, err := models.GetTopTenDevelopments()
	if err != nil {
		panic("Database error")
	}
	var response []models.DevelopmentResponse
	for i, development := range developments {
		response = append(response, models.DevelopmentResponse{
			Position:    strconv.Itoa(i+1) + "°",
			Name:        development.Name,
			Description: development.Description,
			Score:       development.Score,
			Developers:  models.DeveloperModelToResponse(development.Developers),
		})
	}
	ctx.JSON(http.StatusOK, response)
	return
}
