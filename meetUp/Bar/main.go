package main

import (
	"fmt"
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

	n, _ := strconv.ParseInt(tmp[0], 10, 32)
	// m, _ := strconv.ParseInt(tmp[1], 10, 32)

	cap := splitted[1 : n+1]

	request := splitted[n+2:]

	curLayer := 1
	for _, v := range request {
		layer := strings.Split(v, " ")

		count, _ := strconv.ParseInt(layer[1], 10, 32)

		tmp := []rune(layer[2])
		simb := tmp[0]

		for count > 0 {
			str := []rune{}
			for _, ch := range cap[len(cap)-1-curLayer] {
				if ch == ' ' {
					str = append(str, simb)
				} else {
					str = append(str, ch)
				}
			}
			cap[len(cap)-1-curLayer] = string(str)
			count--
			curLayer++
		}

	}

	for _, v := range cap {
		fmt.Printf("%v\n", v)
	}

}
