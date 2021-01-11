package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	total := 10
	now := time.Now()

	wg.Add(total)
	for i := 0; i < total; i++ {
		go printout(i)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
	// time.Sleep(time.Millisecond * 100)
	fmt.Println(i)
	wg.Done()
}
