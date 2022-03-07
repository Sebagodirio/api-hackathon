package controllers

import (
	"api-hackathon/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHackathons GET godoc
// @Summary It returns all the hackathons with its delopments, which contains its developers
// @Accept json
// @Produce json
// @Param BearerAuthorization header string true "Token for entering to the endpoint"
// @Success 200 {object} []models.HackathonResponse
// @Failure 503 {object} string
// @Router /hackathons/ [get]
func GetHackathons(ctx *gin.Context) {
	hackathons, err := models.GetAllHackathons()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "An internal server error have ocurred")
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
