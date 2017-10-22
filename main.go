package main

import (
	"log"
	"os"
	"time"
)

var defaultChannelId int = 183998

var modelToListId map[string]int = map[string]int{
	"Hyundai Solaris": 183992,
	"Skoda Rapid":     183993,
	"Smart Fortwo":    183994,
	"Kia Rio":         183995,
	"Fiat 500":        183996,
	"Ford Fiesta":     183997,
	"Skoda Octavia":   184320,
	"Mini Cooper":     189442,
}

func main() {

	var logger = log.New(os.Stdout, "colesa: ", log.Ldate|log.Ltime)

	colesaUrl := os.Getenv("COLESA_URL")
	broadcastUrl := os.Getenv("BROADCAST_URL")

	logger.Println("COLESA:", colesaUrl)
	logger.Println("BROADCAST:", broadcastUrl)

	state := State{
		available:     make([]Car, 0),
		logger:        logger,
		fetcher:       InitColesaFetcher(colesaUrl),
		broadcaster:   InitVkBroadcaster(broadcastUrl),
		header:        "",
		footer:        "",
		listMap:       modelToListId,
		defaultListId: defaultChannelId,
	}

	ticker := time.NewTicker(60 * time.Second)

	state.Tick()

	for {
		<-ticker.C
		state.Tick()
	}

}
