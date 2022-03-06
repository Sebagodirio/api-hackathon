package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(os.Getenv("SQL_CONNECTION")), &gorm.Config{})
	db.Name()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se conectó con éxito a Mysql")
	}
}

func GetHackathons() []Hackathon {
	var hackathons []Hackathon
	err := db.Debug().Preload("Developments").Find(&hackathons).Error
	if err != nil {
		fmt.Println("ERROR", err)
	}
	return hackathons
}

func GetTopTenDevelopments() ([]Development, error) {
	var developments []Development
	err := db.Debug().Preload("Developers").Order("score DESC").Limit(10).Find(&developments).Error
	if err != nil {
		fmt.Println("ERROR", err)
		return developments, err
	}
	return developments, nil
}

func GetAll() ([]Hackathon, error) {
	var hackathons []Hackathon
	err := db.Debug().Preload("Developments.Developers").Preload(clause.Associations).Find(&hackathons).Error
	if err != nil {
		fmt.Println("ERROR", err)
		return hackathons, err
	}
	return hackathons, nil
}

func CreateHackathon() {
	var developers []Developer
	results, _ := GenerateRandomDevelopers()

	for _, result := range results.Results {
		developers = append(developers, Developer{
			FirstName: result.Name.First,
			LastName:  result.Name.Last,
			Email:     result.Email,
		})
	}

	hackathon := GenerateRandomHackathon(developers)

	err := db.Create(&hackathon).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully created")
}
