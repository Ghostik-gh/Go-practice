package main

import "testing"

func FuzzFoo(f *testing.F) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 23, 25, 26, 30, 39, 49, 1123, 124124}

	f.Fuzz(func(t *testing.T, key int) {
		flagBin := BinSearch(arr, key)
		flagFib := IsFibSearch(arr, key)

		if flagFib != flagBin {
			t.Errorf("Bin: %v, Fib: %v", flagBin, flagFib)
		}
	})
}
