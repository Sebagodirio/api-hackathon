package cron

import (
	"api-hackathon/models"
	"os"

	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	_ = c.AddFunc(os.Getenv("ActualizationPeriod"), func() {
		go models.CreateHackathon()
	})
	c.Start()
}
