package main

import (
	"fmt"
	"math"
)

func main() {
	var a int64 = 600851475143
	for b := squareRoot(a); b >= 1; b-- {
		if a%b == 0 && isPrime(b) {
			fmt.Println(b)
			break
		}
	}
}
func squareRoot(a int64) int64 {
	return int64(math.Sqrt(float64(a)))
}
func isPrime(a int64) bool {
	if a <= 1 {
		return false
	}
	for b := squareRoot(a); b >= 1; b-- {
		if b == 1 {
			return true
		}
		if a%b == 0 {
			return false
		}
	}
	return true
}
