package main

import (
	"fmt"
	"time"
)

type Car struct {
	semaphoreForRide chan int
}

func newCar() *Car {
	return &Car{make(chan int, 10)}
}

func (Car) load(car *Car) {

}

func (Car) run() {

}

func (Car) unload() {
	fmt.Println("Ahoj")
	time.Sleep(10 * time.Second)
	fmt.Println("Tak zase ahoj")
}
