package main

import (
	"fmt"
	"math"
	"os"
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

	splitted := strings.Split(data, "\n")

	for i := range splitted {
		splitted[i] = strings.TrimSpace(splitted[i])
	}

	tmp := strings.Split(splitted[0], " ")
	n, _ := strconv.Atoi(tmp[0])
	m, _ := strconv.Atoi(tmp[1])
	q, _ := strconv.Atoi(tmp[2])

	arr := [][]int{}

	for i := 0; i < n; i++ {
		arr = append(arr, []int{})
		for j := 0; j < m; j++ {
			arr[i] = append(arr[i], 1)
		}
	}

	reloads := make([]int, n)

	req := [][]string{}

	for i := 0; i < q; i++ {
		splitted[i+1] = strings.TrimSpace(splitted[i+1])
		req = append(req, strings.Split(splitted[i+1], " "))
	}
	for i := 0; i < q; i++ {
		if req[i][0] == "DISABLE" {
			center, _ := strconv.Atoi(req[i][1])
			serve, _ := strconv.Atoi(req[i][2])
			arr[center-1][serve-1] = 0
			continue
		}
		if req[i][0] == "RESET" {
			center, _ := strconv.Atoi(req[i][1])
			reloads[center-1]++
			for j := 0; j < m; j++ {
				arr[center-1][j] = 1
			}
			continue
		}
		if req[i][0] == "GETMAX" {
			maxim := 0
			ans := 1
			for i, v := range reloads {
				cur := 0
				if v != 0 {
					for j := 0; j < m; j++ {
						cur += arr[i][j]
					}
					cur *= v
					if cur > maxim {
						maxim = cur
						ans = i + 1
					}
				}
			}
			fmt.Printf("%v\n", ans)
			continue
		}
		if req[i][0] == "GETMIN" {
			mins := math.MaxInt
			ans := 1
			for i, v := range reloads {
				cur := 0
				if v != 0 {
					for j := 0; j < m; j++ {
						cur += arr[i][j]
					}
					cur *= v
					if cur < mins {
						mins = cur
						ans = i + 1
					}
				}
			}
			fmt.Printf("%v\n", ans)
			continue
		}
	}
}
