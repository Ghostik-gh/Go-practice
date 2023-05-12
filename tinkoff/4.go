package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		arr = append(arr, tmp)
	}
	var freq map[int]int
	boring := false
	l := 2
	for !boring && l <= n {
		for k := range freq {
			delete(freq, k)
		}
		for i := 0; i < l; i++ {
			freq[arr[l-i]] = freq[arr[i]] + 1
		}
		for i := l; i < n; i++ {
			freq[arr[i-1]] -= 1
			if freq[arr[i-1]] == 0 {
				delete(freq, arr[i-1])
			}
			freq[arr[i]] = freq[arr[i]] + 1
		}
		max := -1
		min := 1000000000
		for _, element := range arr {
			if element > max {
				max = element
			}
			if element < min {
				min = element
			}
		}
		if max-min <= 1 {
			boring = true
			break
		}
		l += 1
	}
	fmt.Println(l - 1)
}

/*
del freq[a[i-l]]
freq[a[i]] = freq.get(a[i], 0) + 1

counts = list(freq.values())
if max(counts) - min(counts) <= 1:
boring = True
break

l += 1

print(l-1)
*/
