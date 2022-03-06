package models

import (
	"strconv"
)

type Development struct {
	//gorm.Model
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
	Score       float64
	//CreatedAt   mysql.NullTime `json:"-" gorm:"column:created_at"`
	CreatedAt   string      `json:"-" gorm:"column:created_at"`
	HackathonID uint        `gorm:"column:hackathon_id"`
	Developers  []Developer `json:"developers" gorm:"foreignKey:DevelopmentID"`
}

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
