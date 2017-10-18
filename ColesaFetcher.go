package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
)

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

	if strings.Index(string(data), "Mini") != -1 {
		println("Mini at response")
	}

	if strings.Index(string(data), "mini") != -1 {
		println("Mini at response")
	}

	return response.Data, nil
}

func InitColesaFetcher(url string) *ColesaFetcher {
	return &ColesaFetcher{url}
}