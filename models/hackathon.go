package models

import "api-hackathon/helpers"

//Hackathon contains the struct of a hackathon in the database
type Hackathon struct {
	ID            uint          `gorm:"primary_key"`
	Country       string        `json:"name"`
	State         string        `json:"state"`
	HackathonYear string        `json:"hackathonYear" gorm:"column:hackathon_year"`
	CreatedAt     string        `json:"-" gorm:"column:created_at"`
	Developments  []Development `json:"developments" gorm:"foreignKey:HackathonID"`
}

//GenerateRandomHackathon generates a random hackathon
func GenerateRandomHackathon() Hackathon {
	return Hackathon{
		Country:       "United States",
		State:         helpers.GetRandomState(),
		HackathonYear: "2022",
		Developments:  GenerateRandomDevelopments(),
	}
}
