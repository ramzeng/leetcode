package main

import (
	"fmt"
	"time"
)

var throttle chan struct{}

func main() {
	throttle = make(chan struct{}, 10)

	for i := 0; i < 10; i++ {
		throttle <- struct{}{}
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			select {
			case <-time.After(time.Second * 5):
				fmt.Printf("goroutine %d abort\n", i)
			case <-throttle:
				call(i)
			}
		}(i)
	}

	time.Sleep(time.Second * 100)
}

func call(i int) {
	fmt.Printf("goroutine %d call function\n", i)
	time.Sleep(time.Second * 5)
	throttle <- struct{}{}
}
