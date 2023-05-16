package main

import "fmt"

func main() {
	pos := make([]int, 16)
	for i := 0; i < 16; i++ {
		var t int
		fmt.Scan(&t)
		pos[i] = t
	}
	fmt.Printf("pos: %v\n", pos)
	if IsSolving(pos) {
		fmt.Println("Решается")
	} else {
		fmt.Println("ne Решается")
	}
}

func IsSolving(pos []int) bool {
	ans := 0
	for i := 0; i < 15; i++ {
		for j := i + 1; j < 16; j++ {
			if pos[j] != 0 && pos[i] > pos[j] {
				ans++
			}
		}
	}
	return ans%2 == 0
}
