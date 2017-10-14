package main

type BroadcastEvent struct {
	listId int
	message string
}

type Broadcaster interface {
	broadcast(event BroadcastEvent) error
}


type FakeBroadcaster struct {
	lastEvent BroadcastEvent
	responseError error
}

func (fb *FakeBroadcaster) broadcast(event BroadcastEvent) error {
	fb.lastEvent = event
	if fb.responseError != nil {
		return fb.responseError
	} else {
		return nil
	}
}