package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.ReadFile("keyboard/input.txt")
	if err != nil {
		fmt.Printf("Unable open file: %v", err)
		os.Exit(1)
	}
	data := string(f[:])
	// fmt.Printf("file: %v\n", data)

	splitted := strings.Split(data, "\n")

	for i, x := range splitted {
		splitted[i] = strings.TrimSpace(x)
	}

	// n, err := strconv.Atoi(splitted[0])
	// n, err := sliceAtoi(strings.Split(splitted[0], ""))
	c, err := sliceAtoi(strings.Split(splitted[1], " "))
	r, err := sliceAtoi(strings.Split(splitted[2], " "))
	// k, err := sliceAtoi(strings.Split(splitted[3], ""))
	s, err := sliceAtoi(strings.Split(splitted[4], " "))

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// fmt.Println(n, c, r, k, s)

	keyboard := map[int]int{}

	for i, val := range c {
		keyboard[val] = r[i]
	}

	// fmt.Println(keyboard)
	ans := 0
	for i, val := range s {
		if i == 0 {
			continue
		}
		if keyboard[val] != keyboard[s[i-1]] {
			ans++
		}
	}

	fmt.Println(ans)
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
