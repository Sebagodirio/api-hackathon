package models

import (
	"api-hackathon/helpers"
	"strconv"
)

//Development contains the struct of the developments in the database
type Development struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Score       float64
	CreatedAt   string      `json:"-" gorm:"column:created_at"`
	HackathonID uint        `gorm:"column:hackathon_id"`
	Developers  []Developer `json:"developers" gorm:"foreignKey:DevelopmentID"`
}

//GenerateRandomDevelopment generates a random development by using the helpers for creating random fields
func GenerateRandomDevelopment(developers []Developer) Development {
	return Development{
		Name:        helpers.GetRandomDevelopmentsName(),
		Description: helpers.GetRandomDevelopmentsName(),
		Score:       helpers.GenerateScore(500),
		Developers:  developers,
	}
}

//DevelopmentModelToResponse converts the struct Development to DevelopmentResponse
func DevelopmentModelToResponse(developments []Development) []DevelopmentResponse {
	var response []DevelopmentResponse
	for i, development := range developments {
		response = append(response, DevelopmentResponse{
			Position:    strconv.Itoa(i+1) + "°",
			Name:        development.Name,
			Description: development.Description,
			Score:       development.Score,
			Developers:  DeveloperModelToResponse(development.Developers),
		})
	}
	return response
}
