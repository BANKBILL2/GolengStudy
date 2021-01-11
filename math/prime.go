package math

import "fmt"

func prime(n int) {
	for i := 1; i <= n; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 {
			fmt.Print(i, ",")
		}
	}
	fmt.Println()
}
