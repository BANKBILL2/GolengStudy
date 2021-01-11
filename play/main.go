package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5}

	slice := array[:]
	slice2 := append(slice[1:3], 7, 8)

	fmt.Println(slice2)
	fmt.Println(slice)
	fmt.Println(array)

}
