package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	if (a <= b && b <= c && c <= d) || a >= b && b >= c && c >= d {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")
	}
}
