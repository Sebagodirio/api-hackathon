package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Developer contains the struct of the user in the DB
type Developer struct {
	ID            uint
	FirstName     string `json:"firstName" gorm:"column:first_name"`
	LastName      string `json:"lastName" gorm:"column:last_name"`
	Email         string `json:"email"`
	CreatedAt     string `json:"-" gorm:"column:created_at"`
	DevelopmentID uint   `gorm:"column:development_id"`
}

//RandomUser contains the struct for allocating de responde of the API that generates random users
type RandomUser struct {
	Results []Result `json:"results"`
}

//Result contains the struct for allocating de responde of the API that generates random users
type Result struct {
	Name struct {
		First string
		Last  string
	}
	Email string
}

//DeveloperModelToResponse converts the struct Developer to DevelopersResponse
func DeveloperModelToResponse(developers []Developer) []DevelopersResponse {
	var response []DevelopersResponse
	for _, developer := range developers {
		response = append(response, DevelopersResponse{
			FirstName: developer.FirstName,
			LastName:  developer.LastName,
			Email:     developer.Email,
		})
	}
	return response
}

//GenerateRandomDevelopers calls an API for creating random users
func GenerateRandomDevelopers() (RandomUser, error) {
	var randomUsers RandomUser
	response, err := http.Get("https://randomuser.me/api/?results=10")

	if err != nil {
		log.Fatal(err)
		return randomUsers, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return randomUsers, err
	}
	err = json.Unmarshal(responseData, &randomUsers)
	if err != nil {
		log.Fatal(err)
		return randomUsers, err
	}
	return randomUsers, nil
}
