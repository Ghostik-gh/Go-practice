package main

import "fmt"

func main() {
	insert(123, 324)
}

// Простая функция кэширования
func hashFunction(x int) int {
	return x%3 + 1
}

func insert(key, value int) {
	size := 5
	hashTable := make([]int, size)
	index := hashFunction(key)
	for hashTable[index] != 0 {
		index = (index + 1) % size
	}
	hashTable[index] = value
	fmt.Printf("hashTable: %v\n", hashTable)
}

func Hashing() {

}
