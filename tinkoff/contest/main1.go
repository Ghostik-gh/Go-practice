package main

import (
	"fmt"
	"math"
)

func Second() {
	var a float64
	fmt.Scan(&a)

	ans := math.Log2(a)
	fmt.Printf("ans: %v\n", ans)
	fmt.Printf("ans: %v\n", int(ans))
}
