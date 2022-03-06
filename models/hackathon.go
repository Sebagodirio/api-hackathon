package models

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
