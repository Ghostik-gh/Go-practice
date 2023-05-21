package main

import (
	"fmt"
)

func main() {
	sl2 := make([]int, 1)
	EditSlice(&sl2)
	fmt.Printf("cap(sl2): %v\n", cap(sl2))
	fmt.Println(sl2)

}

func EditSlice(slice *[]int) {
	*slice = append(*slice, 1)
	fmt.Printf("cap(slice): %v\n", cap(*slice))
	fmt.Printf("len(slice): %v\n", len(*slice))
}
