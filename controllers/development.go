package controllers

import (
	"api-hackathon/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDevelopments GET godoc
// @Summary It returns the top 10 developments of all the hackathons
// @Accept json
// @Produce json
// @Param BearerAuthorization header string true "Token for entering to the endpoint"
// @Success 200 {object} []models.DevelopmentResponse
// @Failure 503 {object} string
// @Router /hackathons/top-ten/ [get]
func GetDevelopments(ctx *gin.Context) {
	developments, err := models.GetTopTenDevelopments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "An internal server error have ocurred"})
		return
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
