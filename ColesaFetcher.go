package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ColesaRentInfo struct {
	Car Car `json:"car"`
}

type ColesaRentResponse struct {
	Data []ColesaRentInfo `json:"data"`
}

type ColesaResponse struct {
	Data []Car `json:"data"`
}

type ColesaFetcher struct {
	url string
}

func (f *ColesaFetcher) fetch() ([]Car, error) {
	resp, err := http.Get(f.url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := ColesaResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func InitColesaFetcher(url string) *ColesaFetcher {
	return &ColesaFetcher{url}
}
