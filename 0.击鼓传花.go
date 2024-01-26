package main

import (
	"fmt"
	"time"
)

var workersCount = 3

func main() {
	queues := make([]chan struct{}, workersCount)

	for i := 0; i < len(queues); i++ {
		queues[i] = make(chan struct{})

		go func(i int) {
			for {
				<-queues[i]

				fmt.Printf("this is worker %d \n", i+1)
				time.Sleep(time.Second)

				queues[(i+1)%workersCount] <- struct{}{}
			}
		}(i)
	}

	queues[0] <- struct{}{}

	select {}
}
