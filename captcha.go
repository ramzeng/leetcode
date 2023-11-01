package main

import (
	"fmt"
	"leetcode/captcha"
	"time"
)

func main() {
	captcha.Initialize(captcha.NewMemoryStore())

	key := "captcha"

	if code, err := captcha.Create(key, 10, time.Second*10); err != nil {
		fmt.Println("create captcha code error", err)
		return
	} else {
		fmt.Printf("your captcha code is %s\n", code)

		for {
			if result, err := captcha.Check(key, code); err != nil {
				fmt.Println("check captcha code error", err)
			} else {
				fmt.Printf("your captcha code is %t\n", result)
			}

			time.Sleep(time.Second * 3)
		}
	}
}
