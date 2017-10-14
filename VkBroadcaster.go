package main

import (
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"strings"
)

type VkBroadcaster struct {
	url string
}

type VkBroadcasterError struct {
	message string
}

func (err *VkBroadcasterError) Error() string {
	return err.message
}

func (vk *VkBroadcaster) broadcast(event BroadcastEvent) error {
	data, err := http.PostForm(vk.url,
		url.Values{
			"message[message]" : {event.message},
			"list_ids[]" : {strconv.Itoa(event.listId)},
			"run_now": {"1"},
		},
	)
	if err != nil {
		return err
	}
	defer data.Body.Close()

	json, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return err
	}

	if strings.Index( string(json), "response" ) != -1 {
		return nil
	} else {
		return &VkBroadcasterError{"Bad response: " + string(json) }
	}

}

func InitVkBroadcaster(url string) *VkBroadcaster {
	return &VkBroadcaster{url}
}
