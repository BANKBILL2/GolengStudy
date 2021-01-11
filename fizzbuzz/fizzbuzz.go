package fizzbuzz

import (
	"fmt"
	"strconv"
)

func Say(n int) (s string) {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	}
	return strconv.Itoa(n)
}

type RandFunc func(int) int

func (fn RandFunc) Intn(n int) int {
	return fn(n)
}

func Format(randFunction RandFunc) string {
	// src := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(src)
	return fmt.Sprintf("%s-%s-%s-%s", Say(randFunction(100)), Say(randFunction(100)), Say(randFunction(100)), Say(randFunction(100)))
}

func FormatGetInterface(random interface{ Intn(n int) int }) string {
	return fmt.Sprintf("%s-%s-%s-%s", Say(random.Intn(100)+1), Say(random.Intn(100)+1), Say(random.Intn(100)+1), Say(random.Intn(100)+1))
}
