package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Passager struct {
	index int
}

func (passager Passager) board(queue chan Passager, loadSignal *sync.Cond, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	loadSignal.L.Lock()
	loadSignal.Wait()
	fmt.Printf("Passager %d boarded\n", passager.index)
	queue <- passager
	loadSignal.L.Unlock()
}

func (passager Passager) unboard(unloadSignal *sync.Cond, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	unloadSignal.L.Lock()
	unloadSignal.Wait()
	unloadSignal.L.Unlock()
	fmt.Printf("Passager %d UNboarded\n", passager.index)
}

func (passager Passager) enjoyRide(exitChannel chan bool, unloadSignal *sync.Cond, unloadWaitGroup *sync.WaitGroup) {
	for {
		quoteNum := rand.Intn(3)
		var quote string
		switch quoteNum {
		case 0:
			quote = "👍Wuuuu!👍"
		case 1:
			quote = "🤣🤣AHHHHHHHH🤣🤣"
		case 2:
			quote = "🤪WOOOOW🤪"
		}
		select {
		case <-exitChannel:
			go passager.unboard(unloadSignal, unloadWaitGroup)
			return
		case <-time.After((time.Duration(rand.Intn(5)) * time.Second)):
			fmt.Printf("%d: %s\n", passager.index, quote)
		}
	}
}
