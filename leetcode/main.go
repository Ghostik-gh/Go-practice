package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("longestPalindromeSubseq(): %v\n", longestPalindromeSubseq("asss"))

}

func longestPalindromeSubseq(s string) int {
	dp := [][]float64{}
	for i := 0; i < len(s); i++ {
		tmp := make([]float64, len(s))
		dp = append(dp, tmp)
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}
	for i := 0; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				dp[i][j] = dp[i-1][j+1] + 2
			} else {
				dp[i][j] = math.Max(dp[i-1][j], dp[i][j+1])
			}
		}
	}

	return int(dp[len(s)-1][0])
}

// func isPalindrome(s string) bool {

// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 	}
// 	return true
// }

func PrDp(dp [][]float64) {
	for i := 0; i < len(dp); i++ {
		fmt.Printf("dp: %v\n", dp[i])
	}
}
