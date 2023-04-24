package main

import (
	"fmt"
	"math"
)

/*
Разделяет массив на две подмассива f[n-2] f[n-1]
где f - это последовательность числе Фибоначчи

O(log2n)

В теории может быть быстрее потому что нет операций деления
используется только + и -

на магнитной ленте работает быстрее чем бинарный :)
*/
func FibSearchMore(arr []int, key int) bool {
	fibArr := FibArrayCreate(&arr)
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("fibArr: %v\n", fibArr)
	k := len(fibArr) - 2
	fmt.Printf("k: %v\n", k)
	i, p, q := fibArr[k], fibArr[k-1], fibArr[k-2]

	for len(arr) <= fibArr[len(fibArr)-1] {
		arr = append(arr, math.MaxInt)
	}

	for {
		fmt.Printf("arr[%v]: %v\n", i, arr[i])

		fmt.Printf("i: %v, p: %v, q: %v\n", i, p, q)

		if key > arr[i] {
			if p == 1 {
				return false
			}

			i, p, q = i+q, p-q, 2*q-p
			continue
		}
		if key < arr[i] {
			if q == 0 {
				return false
			}
			i -= q
			p, q = q, p-q
			continue
		}

		if key == arr[i] {
			return true
		}
	}
}

/*
Создает последовательность Фибоначчи
*/
func FibArrayCreate(arr *[]int) []int {
	lenght := len(*arr) + 1
	fibArr := []int{0, 1}
	for fibArr[len(fibArr)-1] <= lenght {
		fibArr = append(fibArr, fibArr[len(fibArr)-1]+fibArr[len(fibArr)-2])
	}

	return fibArr
}
