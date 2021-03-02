package main

import (
	"fmt"
	"time"
)

func main() {
	timer2 := time.NewTimer(time.Second)
	go func() {
		select {
		case <-timer2.C:
			fmt.Println("Timer 2 expired")
		}
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	fmt.Println("vim-go")
	timer2.Reset(time.Second)
	time.Sleep(3 * time.Second)
	stop := timer2.Stop()
	if !stop {
		fmt.Println("Timer 2 stopped err")
	}
	timer2.Reset(time.Second)
	stop3 := timer2.Stop()
	if stop3 {
		fmt.Println("Timer 2 stopped")
	}
}
