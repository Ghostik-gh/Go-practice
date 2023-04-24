package main

func BinSearch(arr []int, key int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2

		if arr[middle] < key {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if left == len(arr) || arr[left] != key {
		return false
	}
	return true
}
