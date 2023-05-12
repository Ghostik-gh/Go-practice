package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	var a int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&s)

	a_pos := -1
	b_pos := -1
	c_pos := -1
	d_pos := -1

	dist := 10000000
	if strings.Contains(s, "a") && strings.Contains(s, "b") && strings.Contains(s, "c") && strings.Contains(s, "d") {
		for i, v := range s {
			if v == 'a' {
				a_pos = i
			}
			if v == 'b' {
				b_pos = i
			}
			if v == 'c' {
				c_pos = i
			}
			if v == 'd' {
				d_pos = i
			}
			if a_pos != -1 && b_pos != -1 && c_pos != -1 && d_pos != -1 {
				tmp := Maxs(a_pos, b_pos, c_pos, d_pos) - Mins(a_pos, b_pos, c_pos, d_pos)
				if tmp <= dist {
					dist = tmp
				}
			}
		}
		fmt.Println(dist + 1)
	} else {
		fmt.Println(-1)
	}

}

func Maxs(a, b, c, d int) int {
	q := float64(a)
	w := float64(b)
	e := float64(c)
	r := float64(d)
	return int(math.Max(math.Max(q, w), math.Max(e, r)))
}

func Mins(a, b, c, d int) int {
	q := float64(a)
	w := float64(b)
	e := float64(c)
	r := float64(d)
	return int(math.Min(math.Min(q, w), math.Min(e, r)))
}
