package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Developer struct {
	//gorm.Model
	ID        uint
	FirstName string `json:"firstName" gorm:"column:first_name"`
	LastName  string `json:"lastName" gorm:"column:last_name"`
	Email     string `json:"email"`
	//CreatedAt     mysql.NullTime `json:"-" gorm:"column:created_at"`
	CreatedAt     string `json:"-" gorm:"column:created_at"`
	DevelopmentID uint   `gorm:"column:development_id"`
}

type RandomUser struct {
	Results []Result `json:"results"`
}

type Result struct {
	Name struct {
		First string
		Last  string
	}
	Email string
}

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

func GenerateDevelopers() (RandomUser, error) {
	var randomUsers RandomUser
	response, err := http.Get("https://randomuser.me/api/?results=10")

	if err != nil {
		fmt.Print(err.Error())
		return randomUsers, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(responseData, &randomUsers)
	if err != nil {
		log.Println("Error", err)
		return randomUsers, err
	}
	return randomUsers, nil
}
