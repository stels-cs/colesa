package main

import (
	"bytes"
	"log"
	"testing"
)

func getCar(number string, name string, model string) Car {
	return Car{
		Number: number,
		Name:   name,
		Model: struct {
			Name string `json:"name"`
		}{Name: model},
	}
}

func getEvent(id int, message string) BroadcastEvent {
	return BroadcastEvent{id, message}
}

func isEqualEvent(ev1 BroadcastEvent, ev2 BroadcastEvent) bool {
	return (ev1.listId == ev2.listId) && (ev1.message == ev2.message)
}

func getLogger() *log.Logger {
	var buf bytes.Buffer
	return log.New(&buf, "colesa: ", log.Lshortfile)
}

func TestMainCase(t *testing.T) {

	unpossbaleListId := 999
	unpossableMessage := "UNPOSSABLE MESSAGE"
	shkodaRadpidListId := 555
	defaultChannelId := 19

	car := getCar("A943TE178", "Hyndai Solaris 21", "Hyndai Solaris")
	car2 := getCar("A555TE178", "Shkoda Rapid 19", "Shkoda Rapid")

	fetcher := FakeFetcher{
		nextResponse:      []Car{car},
		nextResponseError: nil,
	}

	broadcaster := FakeBroadcaster{
		lastEvent: getEvent(unpossbaleListId, unpossableMessage),
	}

	state := State{
		logger:        getLogger(),
		fetcher:       &fetcher,
		broadcaster:   &broadcaster,
		defaultListId: defaultChannelId,
		listMap: map[string]int{
			"Shkoda Rapid": shkodaRadpidListId,
		},
	}

	state.Tick()

	if state.isInit != true {
		t.Error("State must be already init, but is not")
	}

	if len(state.available) != 1 {
		t.Error("State must receive one car, but recive", len(state.available))
	}

	if broadcaster.lastEvent.listId != unpossbaleListId {
		t.Error("State send broadast on init")
	}

	state.Tick()

	if broadcaster.lastEvent.listId != unpossbaleListId {
		t.Error("State send broadast:\n", broadcaster.lastEvent.message, "\nbut not new cars was spawn")
	}

	fetcher.nextResponse = []Car{car2, car}

	state.Tick()

	if broadcaster.lastEvent.listId == unpossbaleListId {
		t.Error("State NOT send event then new cars was spawn")
	}

	if broadcaster.lastEvent.listId != shkodaRadpidListId {
		t.Error("State send event to wrong list must be", shkodaRadpidListId, " but ", broadcaster.lastEvent.listId, "expoected")
	}

	if broadcaster.lastEvent.message != car2.toString() {
		t.Error("State send wrong message event, must be", car2.toString(), "but receve", broadcaster.lastEvent.message)
	}

	broadcaster.lastEvent = getEvent(unpossbaleListId, unpossableMessage)

	state.Tick()

	if broadcaster.lastEvent.listId != unpossbaleListId {
		t.Error("State send broadast:\n", broadcaster.lastEvent.message, "\nbut not new cars was spawn")
	}

	fetcher.nextResponse = []Car{car2}

	state.Tick()

	if broadcaster.lastEvent.listId != unpossbaleListId {
		t.Error("State send broadast:\n", broadcaster.lastEvent.message, "\nbut not new cars was spawn")
	}

	fetcher.nextResponse = []Car{car2, car}

	state.Tick()

	if broadcaster.lastEvent.listId == unpossbaleListId {
		t.Error("State NOT send event then new cars was spawn")
	}

	if broadcaster.lastEvent.listId != defaultChannelId {
		t.Error("State send event to wrong list, must be ", defaultChannelId, " but ", broadcaster.lastEvent.listId)
	}

	if broadcaster.lastEvent.message != car.toString() {
		t.Error("State send wrong message event, must be", car.toString(), "but receve", broadcaster.lastEvent.message)
	}
}
