package main

import "sync"

func main() {
	var cars [10]*Car
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		cars[i] = newCar(&wg)
	}
	for _, car := range cars {
		go car.unload()
	}
	wg.Wait()
}
