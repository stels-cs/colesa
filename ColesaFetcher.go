package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

func (f *ColesaFetcher) GetMiniCooperInfo() ([]Car, error) {
	_url := "https://colesa.com/api/v2/rents?client_id=cs-web&client_version=0.0.1&access_token=skazGuNeq1a0n_QW_qRndr3oMgXMubX-uV-RuMZTcCA&limit=100&offset=0&page=1"
	resp, err := http.Get(_url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := ColesaRentResponse{}
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

	lockMap := make(map[string]int)
	cars := make([]Car, 0, 5)

	for _, rent := range response.Data {
		if rent.Car.Model.Name == "Mini Cooper" {
			if _, empty := lockMap[rent.Car.Number]; empty == false {
				cars = append(cars, rent.Car)
				lockMap[rent.Car.Number] = 1
			}
		}
	}

	return cars, nil
}
