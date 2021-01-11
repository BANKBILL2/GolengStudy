package main

import (
	"fmt"
	"strconv"
)

// Int represents integer with extension
type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

func (i *Int) Set(n int) {
	*i = Int(n)
}

func (i Int) Int() int {
	return int(i)
}

func main() {
	var i Int = 99
	fmt.Printf("%q\n", i.String())

	i.Set(100)
	fmt.Println(i.Int())
}
