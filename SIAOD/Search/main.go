package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int{1, 2, 3, 5, 7, 19, 23, 30, 42}

	tree := new(TreeNode)
	for _, v := range arr {
		tree.Insert(v)
	}

	target, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("target: %v\n", target)
	fmt.Printf("%v ", arr)

	fmt.Printf("BinSearch(): %v\n", BinSearch(arr, target))
	fmt.Printf("tree.Search(): %v\n", tree.Search(target))
	fmt.Printf("FibSearch(): %v\n", FibSearch(arr, target))
	fmt.Printf("Interpol(): %v\n", Interpol(arr, target))

}
