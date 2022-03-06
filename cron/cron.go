package cron

import (
	"api-hackathon/models"
	"fmt"
	"os"

	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	fmt.Println("ENTRO AL CRON")
	go models.CreateHackathon()
	_ = c.AddFunc(os.Getenv("ActualizationPeriod"), func() {
		go models.CreateHackathon()
	})
	c.Start()
}
