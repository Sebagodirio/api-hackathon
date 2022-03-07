package cron

import (
	"api-hackathon/services"
	"os"

	"github.com/robfig/cron"
)

func init() {
	c := cron.New()
	_ = c.AddFunc(os.Getenv("ActualizationPeriod"), func() {
		go services.CreateHackathon()
	})
	c.Start()
}
