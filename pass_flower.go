package main

import (
	"fmt"
	"time"
)

// N 个 Goroutine 顺序打印 1 到 N
func main() {
	workersCount := 4

	workerChannels := make([]chan struct{}, workersCount)

	for i := 0; i < workersCount; i++ {
		workerChannels[i] = make(chan struct{})
	}

	for i := 0; i < workersCount; i++ {
		go func(i int) {
			for {
				<-workerChannels[i]

				fmt.Println(i + 1)

				time.Sleep(time.Second)

				workerChannels[(i+1)%workersCount] <- struct{}{}
			}
		}(i)
	}

	workerChannels[0] <- struct{}{}

	select {}
}
