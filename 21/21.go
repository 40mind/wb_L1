package main

import "fmt"

type car interface {
	FillPetrolTank()
}

type petrolCar struct {
}

func (c *petrolCar) FillPetrolTank() {
	fmt.Println("Tank filled with petrol")
}

type dieselCar struct {
}

func (c *dieselCar) FillDieselTank() {
	fmt.Println("Tank filled with diesel")
}

type dieselAdapter struct {
	DieselCar *dieselCar
}

func (c *dieselAdapter) FillPetrolTank() {
	c.DieselCar.FillDieselTank()
}

func main() {
	car1 := &petrolCar{}
	car2 := &dieselCar{}
	car2Adapter := &dieselAdapter{DieselCar: car2}
	car1.FillPetrolTank()
	car2Adapter.FillPetrolTank()
}
