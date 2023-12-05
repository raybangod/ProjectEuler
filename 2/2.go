package main

import "fmt"

func main() {

	limit := 4000000

	a, b := 1, 2
	sum := 0

	for b <= limit {
		if b%2 == 0 {
			sum += b
		}

		a, b = b, a+b
		fmt.Println(a, b)
	}
	fmt.Println(sum)
}
