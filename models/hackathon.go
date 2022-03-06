package models

import "api-hackathon/helpers"

type Hackathon struct {
	//gorm.Model
	ID            uint   `gorm:"primary_key"`
	Country       string `json:"name"`
	State         string `json:"state"`
	HackathonYear string `json:"hackathonYear" gorm:"column:hackathon_year"`
	//CreatedAt     mysql.NullTime `json:"-" gorm:"column:created_at"`
	CreatedAt    string        `json:"-" gorm:"column:created_at"`
	Developments []Development `json:"developments" gorm:"foreignKey:HackathonID"`
}

func GenerateRandomHackathon(developers []Developer) Hackathon {
	return Hackathon{
		Country:       "United States",
		State:         helpers.GetRandomState(),
		HackathonYear: "2022",
		Developments:  []Development{GenerateRandomDevelopment(developers)},
	}
}

func GenerateRandomDevelopment(developers []Developer) Development {
	return Development{
		Name:        helpers.GetRandomDevelopmentsName(),
		Description: helpers.GetRandomDevelopmentsName(),
		Score:       helpers.GenerateScore(500),
		Developers:  developers,
	}
}
