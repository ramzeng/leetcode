package main

import (
	"fmt"
	"time"
)

var c = make(chan struct{}, 10)

func main() {
	for i := 0; i < 100; i++ {
		go invoke()
	}

	time.Sleep(time.Minute)
}

func invoke() {
	c <- struct{}{}

	fmt.Print("invoking function...\n")
	time.Sleep(time.Second * 2)

	<-c
}
