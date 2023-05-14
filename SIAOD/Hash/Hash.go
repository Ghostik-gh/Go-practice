package main

import (
	"fmt"
)

type elem struct {
	key   string
	value any
	next  *elem
}

type LinkedList struct {
	head *elem
}

type myMap struct {
	size int
	list []LinkedList
}

func NewList(size int) myMap {
	list := make([]LinkedList, size)
	return myMap{size: size, list: list}
}

func HashFunc(key string, size int) int {
	id := 0
	for i, v := range key {
		id += int(v) * (i + 0) * 13
		id %= size
	}
	return id
}

func (m *myMap) Insert(key string, value any) {
	index := HashFunc(key, m.size)
	temp1 := &elem{key, value, nil}

	if m.list[index].head == nil {
		m.list[index].head = temp1
	} else {
		temp2 := m.list[index].head
		for temp2.next != nil {
			temp2 = temp2.next
		}
		temp2.next = temp1
	}
}

func (m *myMap) Search(key string) any {
	index := HashFunc(key, m.size)
	if m.list[index].head == nil {
		return "not found"
	}
	temp1 := m.list[index].head
	for temp1.next != nil {
		if temp1.key == key {
			return temp1.value
		}
		temp1 = temp1.next
	}
	if temp1.key == key {
		return temp1.value
	}
	return "not found"
}

func (m *myMap) Delete(key string) {
	index := HashFunc(key, m.size)
	// Delete Head
	if m.list[index].head.key == key {
		m.list[index].head = m.list[index].head.next
	}
	// Delete at middle and tail
	temp1 := m.list[index].head
	temp2 := m.list[index].head
	for temp1.key != key && temp1.next != nil {
		temp2 = temp1
		temp1 = temp1.next
	}
	if temp1.key == key {
		temp2.next = temp1.next
	}
}

func main() {
	size := 10
	myMap := NewList(size)
	myMap.Insert("a", "hello")
	myMap.Insert("j", []int{1, 3})
	myMap.Insert("x", rune(322))
	myMap.Insert("o", 4)
	myMap.Insert("e", 5.012309)

	myMap.Delete("e")

	fmt.Printf("myMap.list[0].head: %v\n", myMap.list[0].head)
	fmt.Printf("myMap.list[0].head: %v\n", myMap.list[0].head.next)
	fmt.Printf("myMap.list[0].head: %v\n", myMap.list[0].head.next.next)
	fmt.Printf("myMap.list[0].head: %v\n", myMap.list[0].head.next.next.next)
	fmt.Printf("myMap.list[0].head: %v\n", myMap.list[0].head.next.next.next.next)

	value := myMap.Search("a")
	fmt.Println(value)
}
