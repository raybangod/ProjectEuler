package main

import "fmt"

func main() {
	palindrome := 0
	var i2, j2 int
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalindromic(product) && product > palindrome {
				palindrome = product
				i2, j2 = i, j
				break
			}

		}
	}
	fmt.Println(i2, j2, palindrome)

}
func isPalindromic(a int) bool {
	original := a
	reversed := 0

	for a > 0 {
		remainder := a % 10
		reversed = (reversed * 10) + remainder
		a /= 10
	}

	return original == reversed
}
