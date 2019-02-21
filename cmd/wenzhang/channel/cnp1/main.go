package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const (
	MAXRECEIVERS = 10
	MAXRANDOMNUM = 10000000
)

func main() {
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(MAXRECEIVERS)

	dataCh := make(chan int, 100)

	go func() {
		for {
			if value := rand.Intn(MAXRANDOMNUM); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	for i := 0; i < MAXRECEIVERS; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				fmt.Println(value)
			}
		}()
	}

	wgReceivers.Wait()

}
