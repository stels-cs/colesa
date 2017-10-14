package main

type Car struct {
	Number string `json:"number"`
	Name   string `json:"name"`
	Model struct {
		Name string `json:"name"`
	} `json:"car_model"`
}


func (car Car) toString() string {
	return car.Name
}
