package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("findDiff(): %v\n", findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))

}
func findDifference(nums1 []int, nums2 []int) [][]int {
	ans := [][]int{[]int{}, []int{}}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] <= nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] <= nums2[j]
	})
	prev := -100000
	for _, v := range nums1 {
		if v == prev {
			continue
		}
		flag := true
		for _, v2 := range nums2 {
			if v == v2 {
				flag = false
			}
		}
		if flag {
			ans[0] = append(ans[0], v)
			prev = v
		}
	}
	for _, v := range nums2 {
		if v == prev {
			continue
		}
		flag := true
		for _, v2 := range nums1 {
			if v == v2 {
				flag = false
			}
		}
		if flag {
			ans[1] = append(ans[1], v)
			prev = v
		}
	}

	return ans
}
