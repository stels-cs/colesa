package main

import (
	"log"
)

type State struct {
	isInit bool
	available []Car
	logger *log.Logger
	fetcher Fetcher
	broadcaster Broadcaster
	header string
	footer string
	listMap map[string]int
	defaultListId int
	heatbeat int
}


func (state *State) Tick() {

	updates, err := state.fetcher.fetch()

	if err != nil {
		state.logger.Println("Fetch fail:", err.Error())
		return
	}

	if state.isInit {
		events := state.GetEventsFromUpdate(updates)
		for _, event := range events {
			err := state.broadcaster.broadcast(event)
			if err != nil {
				state.logger.Fatalln("Fail broadcast:", err.Error())
			} else {
				state.logger.Println("Broadcast: ", event.message, "to", event.listId)
			}
		}
		state.available = updates

		if len(events) > 0 {
			state.heatbeat = 1
		} else {
			state.heatbeat++
		}

		if state.heatbeat % 30 == 0 {
			state.logger.Println("Updated")
		}

	} else {
		state.InitWithUpdates(updates)
	}
}

func (state *State) GetEventsFromUpdate(cars []Car) []BroadcastEvent {
		events := make([]BroadcastEvent, 0, 5)
		for _, car := range cars {
			if !state.IsLiveCar(car) {
				events = append(events, state.GetEventForCar(car))
			}
		}
		return events
}

func (state *State) GetEventForCar(car Car) BroadcastEvent {
	listId := state.getListIdByModelName(car.Model.Name)
	message := state.getMessageForCar(car)

	return BroadcastEvent{listId, message}
}

func (state *State) getListIdByModelName(model string) int {
	id, has := state.listMap[model]
	if has {
		return id
	} else {
		return state.defaultListId
	}
}

func (state *State) getMessageForCar(car Car) string {
		return state.header  + car.toString()  + state.footer
}

func (state *State) InitWithUpdates(cars []Car) {
	state.isInit = true
	state.available = cars
	state.logger.Println(len(cars), "cars online")
}

func (state *State) IsLiveCar(car Car) bool {
	if state.available == nil {
		return false
	}
	for _, _car := range state.available {
		if _car.Number == car.Number {
			return true
		}
	}
	return false
}
