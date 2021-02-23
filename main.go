package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"standup_timer/config"
	"time"
)

var cfg config.Config

var isStandUpTimer bool

var sitTimeDuration time.Duration
var standTimeDuration time.Duration

func main() {
	cfg = config.Read("config.yml")
	sitTimeDuration = time.Duration(cfg.Timer.SitTimeInMinutes) * time.Minute
	standTimeDuration = time.Duration(cfg.Timer.StandTimeInMinutes) * time.Minute

	isStandUpTimer = true

	startTicker()
}

func startTicker(){
	ticker := time.NewTicker(sitTimeDuration)
	for range ticker.C {
		fmt.Println("New tick!")

		if isStandUpTimer {
			sitDownNotification()

			isStandUpTimer = false
			ticker.Reset(sitTimeDuration)
		}else {
			standUpNotification()

			isStandUpTimer = true
			ticker.Reset(standTimeDuration)
		}
	}
}

func sitDownNotification()  {
	err := beeep.Alert("Desk timer", "Time to sit down now!", "assets/warning.png")
	if err != nil {
		panic(err)
	}
}

func standUpNotification(){
	err := beeep.Alert("Desk timer", "Time to stand up now! Lift up your desk!", "assets/warning.png")
	if err != nil {
		panic(err)
	}
}