package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("isPalindrome: %v\n", isPalindrome)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
