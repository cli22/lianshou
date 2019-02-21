package main

import (
	"log"
	"math/rand"
	"sync"
)

const (
	MAXSENDERS   = 1
	MAXRANDOMNUM = 10000000
)

func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < MAXSENDERS; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					close(dataCh)
					return
				case dataCh <- rand.Intn(MAXRANDOMNUM):
				}

			}
		}()
	}

	go func() {
		defer wgReceivers.Done()

		for value := range dataCh {
			if value == 1 {
				close(stopCh)
				return
			}
			log.Println(value)
		}
	}()

	wgReceivers.Wait()
}
