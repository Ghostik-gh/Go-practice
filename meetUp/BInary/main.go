package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Unable open file: %v", err)
		os.Exit(1)
	}
	data := string(f[:])

	splitted := strings.Split(data, "\n")
	// n := 4
	s := splitted[1]

	sum := [][]int{}
	sum = append(sum, []int{0, 0})
	zeros, ones := 0, 0

	dp := make([]int, len(s))

	for i, v := range s {
		if v == '0' {
			zeros++
			if zeros > ones {
				dp[i] = 1
			}
		} else {
			ones++
			if ones > zeros {
				dp[i] = 1
			}

		}
		sum = append(sum, []int{zeros, ones})
	}

	// fmt.Printf("dp: %v\n", dp)

	ans := make([]int, len(s))
	ans = dp
	ans[0] = -1

	for i := range sum {
		if i == 0 || i == len(sum)-1 {
			continue
		}
		if ans[i] == 1 {
			continue
		}
		if s[i-1] == s[i] {
			ans[i] = i - 1
			continue
		}

		if ans[i] <= 0 {

			flag := false
			for j := i - 2; j >= 0; j-- {

				if s[i] == '0' && s[j] == '0' && ans[j] == -1 {
					if sum[i+1][0]-sum[j][0] > sum[i+1][1]-sum[j][1] {
						ans[i] = j + 1
						flag = true
					}
					break
				}
				if s[i] == '0' && s[j] == '0' && ans[j] != -1 {
					pointer := ans[j]
					for cur := 0; pointer != -1 && cur < 10; cur++ {

						if ans[pointer] == -1 {
							break
						}
						pointer = ans[pointer]
					}
					if sum[i+1][0]-sum[pointer][0] > sum[i+1][1]-sum[pointer][1] {
						ans[i] = pointer + 1
						flag = true
					}
					break
				}

				if s[i] == '1' && s[j] == '1' && ans[j] == -1 {
					if sum[i+1][0]-sum[j][0] < sum[i+1][1]-sum[j][1] {
						ans[i] = j + 1
						flag = true
					}
					break
				}
				if s[i] == '1' && s[j] == '1' && ans[j] != -1 {
					pointer := ans[j]
					for cur := 0; pointer != -1 && cur < 10; cur++ {
						if ans[pointer] == -1 {
							break
						}
						pointer = ans[pointer]
					}
					if sum[i+1][0]-sum[pointer][0] < sum[i+1][1]-sum[pointer][1] {
						ans[i] = pointer + 1
						flag = true
					}
					break
				}
				// if s[i] == '0' && s[j] == '0' {
				// 	if sum[i+1][0]-sum[j][0] > sum[i+1][1]-sum[j][1] {
				// 		ans[i] = j + 1
				// 		// ans = append(ans, j+1)
				// 		flag = true
				// 		break
				// 	}
				// }
				// if s[i] == '1' && s[j] == '1' {
				// 	if sum[i+1][0]-sum[j][0] < sum[i+1][1]-sum[j][1] {
				// 		ans[i] = j + 1
				// 		// ans = append(ans, j+1)
				// 		flag = true
				// 		break
				// 	}
				// }
			}
			if flag == false {
				ans[i] = -1
				// ans = append(ans, -1)
			}
		}

	}
	for _, v := range ans {
		fmt.Print(v, " ")
	}
}
