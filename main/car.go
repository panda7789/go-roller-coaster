package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	semaphoreForRide chan int
	waitGroup        *sync.WaitGroup
}

func newCar(wg *sync.WaitGroup) *Car {
	return &Car{
		make(chan int, 10), wg}
}

func (Car) load(car *Car) {

}

func (Car) run() {

}

func (car Car) unload() {
	fmt.Println("Ahoj")
	time.Sleep(10 * time.Second)
	fmt.Println("Tak zase ahoj")
	car.waitGroup.Done()
}
