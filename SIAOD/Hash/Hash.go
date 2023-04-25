package main

import (
	"fmt"
	"math/rand"
)

var size = 5
var hashTable = make([]int, size)

func main() {
	insert(9, 1)
	insert(21, 3)
	insert(23, 7)
	insert(4, 11)
	insert(99, 12)
	hashFunctionRand()
	fmt.Printf("hashTable: %v\n", hashTable)
	// fmt.Printf("search(): %v\n", search(11))

	// Коллизия
	// insert(12, 923, &hashTable)
}

// Простая функция кэширования
func hashFunction(x int) int {
	return x % (size)
}

func hashFunctionRand() {
	var random = rand.Intn(99) + 1
	for i, v := range hashTable {
		hashTable[i] = (v * random) % (100)
	}
}

func insert(key int, value int) {
	index := hashFunction(key)
	for hashTable[index] != 0 {
		index = (index + 1) % size
	}
	hashTable[index] = value
	fmt.Printf("hashTable: %v\n", hashTable)
}

func search(key int) int {
	index := hashFunction(key)
	startIndex := index
	for hashTable[index] != 0 {
		if hashTable[index] != 0 {
			return hashTable[index]
		}
		index = (index + 1) % size
		if index == startIndex {
			return -1
		}
	}
	return -1
}
