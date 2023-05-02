package main

import "fmt"

func main() {
	fmt.Printf("average(): %v\n", average([]int{4000, 3000, 1000, 2000}))

}
func average(salary []int) float64 {
	min := salary[0]
	max := salary[0]
	sum := 0.
	for _, v := range salary {
		min = Min(min, v)
		max = Max(max, v)
		sum += float64(v)
	}
	fmt.Printf("sum: %v\n", sum)

	return (sum - float64(min) - float64(max)) / float64(len(salary)-2)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
