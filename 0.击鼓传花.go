package main

import (
	"fmt"
	"time"
)

var workersCount = 3

func main() {
	queue := make([]chan struct{}, workersCount)

	for i := 0; i < workersCount; i++ {
		queue[i] = make(chan struct{})

		go func(i int) {
			for {
				<-queue[i]
				fmt.Println(i + 1)
				time.Sleep(time.Second)
				queue[(i+1)%workersCount] <- struct{}{}
			}
		}(i)
	}

	queue[0] <- struct{}{}
	select {}
}
