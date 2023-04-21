package main

import (
	"encoding/json"
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

	fmt.Printf("tmp: %v\n", tmp)

	n, _ := strconv.Atoi(tmp[0])
	m, _ := strconv.Atoi(tmp[1])

	fmt.Printf("n: %v, %v\n", n, m)

	req := [][]string{}

	for i := 0; i < n; i++ {
		tmp := strings.Split(splitted[i+1], " ")
		req = append(req, tmp)
	}

	fmt.Printf("splitted[n]: %v\n", splitted[n])

	// jsonic := splitted[n]

	var msg Message

	eeeee := json.Unmarshal([]byte(splitted[n+1]), &msg)

	fmt.Printf("eeeee: %v\n", eeeee)

	fmt.Printf("msg: %v\n", msg)
}

type Message struct {
	id string `json:"title"`
	// "type": "object",
	// "properties": {
	//   "trace_id": {
	// 	"type": "string"
	//   },
	//   "offer": {
	// 	"$ref": "offer.schema.json"
	//   }
	// },
	// "required": [
	//   "trace_id",
	//   "offer"
	// ]
}
