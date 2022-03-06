package controllers

import (
	"api-hackathon/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHackathones(ctx *gin.Context) {
	hackathons, err := models.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "An error ocurred")
		return
	}
	var response []models.HackathonResponse
	for _, hackathon := range hackathons {
		response = append(response, models.HackathonResponse{
			Country:       hackathon.Country,
			State:         hackathon.State,
			HackathonYear: hackathon.HackathonYear,
			Developments:  models.DevelopmentModelToResponse(hackathon.Developments),
		})
	}
	ctx.JSON(http.StatusOK, response)
	return
}
