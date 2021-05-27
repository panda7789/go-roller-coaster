package main

func main() {
	var cars [1]*Car = [1]*Car{newCar()}
	for _, car := range cars {
		car.unload()
	}
}
