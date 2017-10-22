package main

type Car struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	Model  struct {
		Name string `json:"name"`
	} `json:"car_model"`
	Color struct {
		Name string `json:"name"`
	} `json:"color"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (car Car) toString() string {
	return car.Color.Name + " " + car.Model.Name
}
