package main

import (
	"net/http"
	"net/url"
	"strconv"
)

type VkBroadcaster struct {
	url string
}

func (vk *VkBroadcaster) broadcast(event BroadcastEvent) error {
	_, err := http.PostForm(vk.url,
		url.Values{
			"message.message" : {event.message},
			"list_ids[]" : {strconv.Itoa(event.listId)},
		},
	)
	return err
}

func InitVkBroadcaster(url string) *VkBroadcaster {
	return &VkBroadcaster{url}
}
