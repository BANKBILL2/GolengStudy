package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	go fibonacci(ch, sig)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	sig <- struct{}{}
	<-sig
	fmt.Println("graceful")
}

func fibonacci(ch chan int, sig chan struct{}) {
	f1, f2 := 0, 1

	for {
		select {
		case ch <- f1:
			f1, f2 = f2, f1+f2
		case <-sig:
			sig <- struct{}{}
			return
		}
	}
}
