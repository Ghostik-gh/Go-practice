package main

import (
	"strings"
)

// Алгоритм Бойера-Мура
func BoyerMoore(fs FindSubstr) int {
	var (
		t string
		s string
	)
	if !fs.regSens {
		t = strings.ToLower(fs.text)
		s = strings.ToLower(fs.substr)
	} else {
		t = fs.text
		s = fs.substr
	}
	off := CreateOffset(s)
	length := len(s)
	k := length
	l := length
	// fmt.Printf("off: %v\n", off)
	// fmt.Printf("t: %v\n", t)
	// fmt.Printf("s: %v\n", s)
	for k <= len(t) {
		// fmt.Printf("k: %v, l: %v\n", k, l)
		if t[k-1] == s[l-1] {
			l--
			k--
		} else {
			// fmt.Printf("off[s[l-1]-1]: %v\n", off[t[k-1]])
			k += off[t[k-1]]
			l = off[t[k-1]]
		}
		if l == 0 {
			return k
		}
	}
	return -1
}

func CreateOffset(s string) []int {
	d := make([]int, 256)

	cur := 1
	for i := len(s) - 2; i >= 0; i-- {
		if d[s[i]] == 0 {
			d[s[i]] = cur
		}
		cur++
	}
	if d[s[len(s)-1]] == 0 {
		d[s[len(s)-1]] = cur
	}
	for i, v := range d {
		if v == 0 {
			d[i] = len(s)
		}
	}
	return d
}
