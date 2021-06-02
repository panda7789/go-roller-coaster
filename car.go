package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	passagersCount    int
	bufferForRide     chan Passager
	loadSignal        *sync.Cond
	unloadSignal      *sync.Cond
	bufferForRideWG   *sync.WaitGroup
	bufferForUnloadWG *sync.WaitGroup
}

func newCar(passagersCount int) Car {
	loadLock := sync.Mutex{}
	unloadLock := sync.Mutex{}
	var car = Car{
		passagersCount,
		make(chan Passager, passagersCount),
		sync.NewCond(&loadLock),
		sync.NewCond(&unloadLock),
		&sync.WaitGroup{},
		&sync.WaitGroup{}}
	return car
}

func (car Car) load() {
	car.bufferForRideWG.Add(car.passagersCount)
	car.bufferForUnloadWG.Add(car.passagersCount)
	fmt.Println("#Boarding started")
	time.Sleep(2 * time.Second)
	car.loadSignal.L.Lock()
	for i := 0; i < car.passagersCount; i++ {
		car.loadSignal.Signal()
	}
	car.loadSignal.L.Unlock()
	car.bufferForRideWG.Wait()
	fmt.Println("#Boarding completed")
}

func (car Car) run() {
	exitChannel := make(chan bool)
	fmt.Println("#Ride started")
	for i := 0; i < cap(car.bufferForRide); i++ {
		var passager = <-car.bufferForRide
		go passager.enjoyRide(exitChannel, car.unloadSignal, car.bufferForUnloadWG)
	}
	time.Sleep(10 * time.Second)
	close(exitChannel)
	fmt.Println("#Ride completed")
}

func (car Car) unload() {
	fmt.Println("#Unboarding started")
	car.unloadSignal.L.Lock()
	car.unloadSignal.Broadcast()
	car.unloadSignal.L.Unlock()
	car.bufferForUnloadWG.Wait()
	fmt.Println("#Unboarding completed")
}
