package main

import (
	"fmt"
	"math"
)

// 15 -> 7, 8 -> 3, 4, 4, 4 -> 2, 1, 2, 2, 2, 2, 2 -> 1, ,1, 1,, 1,1, ,1 ,1, 1,, 1,1 ,

func main() {
	var a float64
	fmt.Scan(&a)
	ans := math.Log2(a)
	if a == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(math.Ceil(ans))
	}
}
