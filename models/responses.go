package models

type HackathonResponse struct {
	Country       string `json:"country"`
	State         string `json:"state"`
	HackathonYear string `json:"hackathon_year"`
	Developments  []DevelopmentResponse
}

type DevelopmentResponse struct {
	Position    string  `json:"position"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Score       float64 `json:"score"`
	Developers  []DevelopersResponse
}

type DevelopersResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
