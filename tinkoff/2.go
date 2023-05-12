package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if b > a {
		b = a
	}

	fmt.Println(math.Ceil((c * a) / b))

}
