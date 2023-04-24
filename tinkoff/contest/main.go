package main

import (
	"fmt"
	"math"
)

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
