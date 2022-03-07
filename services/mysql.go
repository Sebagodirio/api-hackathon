package services

import (
	"api-hackathon/models"
	"fmt"
	"log"
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

//GetTopTenDevelopments gets the top ten developments from all the hackathons, it comes with its developers
func GetTopTenDevelopments() ([]models.Development, error) {
	var developments []models.Development
	err := db.Debug().Preload("Developers").Order("score DESC").Limit(10).Find(&developments).Error
	if err != nil {
		fmt.Println("ERROR", err)
		return developments, err
	}
	return developments, nil
}

//GetAllHackathons gets all the hackathons, with its developments, which comes with its developers (hasMany -> hasMany ->)
func GetAllHackathons() ([]models.Hackathon, error) {
	var hackathons []models.Hackathon
	err := db.Debug().Preload("Developments.Developers").Preload(clause.Associations).Find(&hackathons).Error
	if err != nil {
		log.Fatal("ERROR", err)
		return hackathons, err
	}

	for _, hackathon := range hackathons {
		models.OrderDevelopmentByScore(hackathon.Developments)
	}

	return hackathons, nil
}

//CreateHackathon creates a hackathon and saves it in the database
func CreateHackathon() {
	var developers []models.Developer
	results, _ := models.GenerateRandomDevelopers()

	for _, result := range results.Results {
		developers = append(developers, models.Developer{
			FirstName: result.Name.First,
			LastName:  result.Name.Last,
			Email:     result.Email,
		})
	}

	hackathon := models.GenerateRandomHackathon(developers)

	err := db.Create(&hackathon).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------- Hackathon Successfully created ----------------------")
}

//CheckUserInDB checks if the user's credential are ok
func CheckUserInDB(user models.User) error {
	return db.Where("email = ? and password = ?", user.Email, user.Password).First(&user).Error
}
