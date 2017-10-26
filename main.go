package main

import (
	"log"
	"os"
	"time"
)

var defaultChannelId int = 194659

var modelToListId map[string]int = map[string]int{
	"Hyundai Solaris": 194650,
	"Skoda Rapid":     194656,
	"Smart Fortwo":    194648,
	"Kia Rio":         194651,
	"Fiat 500":        194647,
	"Ford Fiesta":     194652,
	"Skoda Octavia":   194654,
	"Mini Cooper":     194646,
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
