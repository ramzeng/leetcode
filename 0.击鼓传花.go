package main

import (
	"fmt"
	"time"
)

func main() {
	workersCount := 3

	queues := make([]chan struct{}, workersCount)

	for i := 0; i < workersCount; i++ {
		queues[i] = make(chan struct{})

		go func(i int) {
			for {
				queue := queues[i]

				<-queue

				fmt.Println(i + 1)

				time.Sleep(time.Second)

				queues[(i+1)%workersCount] <- struct{}{}
			}
		}(i)
	}

	queues[0] <- struct{}{}

	select {}
}
