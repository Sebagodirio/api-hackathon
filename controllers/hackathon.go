package controllers

import (
	"api-hackathon/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHackathones(ctx *gin.Context) {
	if ctx.Query("create") != "" {
		fmt.Println("entered create")
		models.CreateHackathon()
	}
	fmt.Println("error ocurre aca")
	hackathons, err := models.GetAll()
	if err != nil {
		fmt.Println("error ocurre aca X2")
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
