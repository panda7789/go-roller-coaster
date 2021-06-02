package main

import "fmt"

const passagersInCar = 5

var waitingPlace = make(chan Passager, 20)
var alreadyRidedPassagers = 0

func main() {
	var car Car = newCar(passagersInCar)
	fillQueue(&car)
	for alreadyRidedPassagers+car.passagersCount <= cap(waitingPlace) {
		car.load()
		car.run()
		car.unload()
		fmt.Println("Ride sucesfully completed 😊")
		alreadyRidedPassagers += car.passagersCount
	}
}

func fillQueue(car *Car) {
	for i := 0; i < cap(waitingPlace); i++ {
		var pass = Passager{i}
		go pass.board(car.bufferForRide, car.loadSignal, car.bufferForRideWG)
		waitingPlace <- pass
	}
}
