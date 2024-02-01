package main

import (
	"fmt"
	"time"
)

var c = make(chan struct{}, 10)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				invoke(i)
			}
		}(i)
	}

	select {}
}

func invoke(i int) {
	c <- struct{}{}
	fmt.Printf("this is caller %d invoke\n", i)
	time.Sleep(time.Second)
	<-c
}
