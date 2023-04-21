package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Unable open file: %v", err)
		os.Exit(1)
	}

	data := string(f[:])

	data = strings.TrimSpace(data)

	strs := strings.Split(data, "\n")
	for i := range strs {
		strs[i] = strings.TrimSpace(strs[i])
	}

	tmp := strings.Split(strs[0], " ")

	n, _ := strconv.ParseInt(tmp[0], 10, 32)
	m, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))

	a := strings.Split(strs[1], " ")
	towers := make([]int, n)
	for i, v := range a {
		towers[i], err = strconv.Atoi(v)
	}

	c := strings.Split(strs[2], " ")
	mans := make([]int, m)
	for i, v := range c {
		mans[i], err = strconv.Atoi(v)
	}

	platforms := []int{0}
	for i := len(towers) - 1; i >= 0; i-- {
		if towers[i] > platforms[len(platforms)-1] {
			platforms = append(platforms, towers[i])
		}
	}

	sun := []int{}
	for i, v := range platforms {
		if i == 0 {
			continue
		}
		sun = append(sun, v-platforms[i-1])
	}

	sort.Slice(mans, func(i, j int) bool {
		return mans[j] > mans[i]
	})
	sort.Slice(sun, func(i, j int) bool {
		return sun[j] > sun[i]
	})
	answer := 0
	cur := 0
	for _, v := range mans {
		flag := false
		if cur > len(sun)-1 {
			break
		}
		for v > sun[cur] {
			cur++
			if cur > len(sun)-1 {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		answer++
		cur++
	}
	fmt.Println(answer)
}
