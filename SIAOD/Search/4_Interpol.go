package main

func Interpol(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right && target >= arr[left] && target <= arr[right] {
		pos := left + ((target - arr[left]) * (right - left) / (arr[right] - arr[left]))
		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	return -1
}
