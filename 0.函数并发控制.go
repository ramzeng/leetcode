package main

import (
	"fmt"
	"time"
)

var c = make(chan struct{}, 10)

func main() {
	for {
		for i := 0; i < 100; i++ {
			go func(i int) {
				invoke(i)
			}(i)
		}

		time.Sleep(time.Second * 30)
	}
}

func invoke(i int) {
	c <- struct{}{}

	fmt.Printf("caller %d invoke function...\n", i)
	time.Sleep(time.Second)

	<-c
}
