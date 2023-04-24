package main

/*
Разделяет массив на две подмассива f[n-2] f[n-1]
где f - это последовательность числе Фибоначчи

O(log2n)

В теории может быть быстрее потому что нет операций деления
используется только + и -

на магнитной ленте работает быстрее чем бинарный :)
*/
func FibSearch(arr []int, key int) int {

	p, q := 1, 1
	i := q + p

	for i < len(arr) {
		p = q
		q = i
		i = p + q
	}

	offset := 0

	for i > 1 {
		i = Min(offset+p, len(arr)-1)

		if key > arr[i] {
			i = q
			q = p
			p = i - q
			offset = i
		} else if key < arr[i] {
			i = p
			q = q - p
			p = i - q
		} else {
			return i
		}
	}
	if q >= 0 && offset < (len(arr)-1) && arr[offset+1] == key {
		return offset + 1
	}
	// if q == 1 && arr[offset+1] == key {
	// 	return offset + 1
	// }
	return -1
}

func IsFibSearch(arr []int, key int) bool {

	p, q := 0, 1
	i := q + p

	for i < len(arr) {
		p = q
		q = i
		i = p + q
	}

	offset := 0

	for i > 1 {

		i = Min(offset+p, len(arr)-1)

		if key > arr[i] {
			i = q
			q = p
			p = i - q
			offset = i
		} else if key < arr[i] {
			i = p
			q = q - p
			p = i - q
		} else {
			return true
		}
	}
	if q == 1 && arr[offset+1] == key {
		return true
	}
	return false
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
