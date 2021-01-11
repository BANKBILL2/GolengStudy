package main

import (
	"fmt"
)

func main() {
	catchMe()
}

func catchMe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	s := []int{}

	panic("capture")
	fmt.Println(s[1])
}

func doSomething(n int) {
	defer fmt.Println(n)
	defer func() { fmt.Println(n) }()
	n *= n
	fmt.Println(n)
}
