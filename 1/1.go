package main

import "fmt"

func main() {
	var belowNumber int
	fmt.Scanln(&belowNumber)
	sumResult := 0
	for i := 3; i < belowNumber; i++ {
		if i%3 == 0 || i%5 == 0 {
			sumResult += i
		}
	}

	fmt.Println(sumResult, 1, 2, 3)
}
