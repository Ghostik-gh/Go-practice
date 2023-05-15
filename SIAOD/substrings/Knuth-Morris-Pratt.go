package main

import (
	"strings"
)

// алгоритм Кнута-Морриса-Пратта
func KnuthMorrisPratt(fs FindSubstr) int {

	var (
		k = 0
		l = 0
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

	pi := prefixFunc(s)
	for k < len(t) {
		if l == len(s) {
			return k - l
		}
		if t[k] == s[l] {
			k++
			l++
			continue
		}

		if t[k] != s[l] && l != 0 {
			l = pi[l-1]
		} else {
			k++
		}
	}
	if l == len(s) {
		return k - l
	}
	return -1
}

func prefixFunc(s string) []int {
	pi := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

// vector<int> prefix_function (string s) {
// 	int n = (int) s.length();
// 	vector<int> pi (n);
// 	for (int i=1; i<n; ++i) {
// 		int j = pi[i-1];
// 		while (j > 0 && s[i] != s[j])
// 			j = pi[j-1];
// 		if (s[i] == s[j])  ++j;
// 		pi[i] = j;
// 	}
// 	return pi;
// }
