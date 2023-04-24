package main

import (
	"fmt"
	"math/rand"
)

type Coords struct {
	x, y int
}

func main() {

	count := 0
	coords := []Coords{}
	board := [][]int{}
	for count != 8 {
		ans := Coords{}
		coords = []Coords{}
		count = 0
		board = [][]int{}
		for i := 0; i < 8; i++ {
			board = append(board, []int{0, 0, 0, 0, 0, 0, 0, 0})
		}
		for i := 0; i < 100; i++ {
			x := rand.Intn(8)
			y := rand.Intn(8)
			if board[x][y] == 0 {
				for cordX := 0; cordX < 8; cordX++ {
					board[cordX][y] = 1
				}
				for cordY := 0; cordY < 8; cordY++ {
					board[x][cordY] = 1
				}
				for cordX, cordY := x, y; cordX < 8 && cordY < 8; cordX, cordY = cordX+1, cordY+1 {
					board[cordX][cordY] = 1
				}
				for cordX, cordY := x, y; cordX >= 0 && cordY >= 0; cordX, cordY = cordX-1, cordY-1 {
					board[cordX][cordY] = 1
				}
				for cordX, cordY := x, y; cordX < 8 && cordY >= 0; cordX, cordY = cordX+1, cordY-1 {
					board[cordX][cordY] = 1
				}
				for cordX, cordY := x, y; cordX >= 0 && cordY < 8; cordX, cordY = cordX-1, cordY+1 {
					board[cordX][cordY] = 1
				}
				board[x][y] = 1
				ans.x = x + 1
				ans.y = y + 1
				coords = append(coords, ans)
				count++
			}
		}
	}
	fmt.Printf("count: %v\n", count)
	fmt.Printf("ans: %v\n", coords)

	board = [][]int{}
	for i := 0; i < 8; i++ {
		board = append(board, []int{0, 0, 0, 0, 0, 0, 0, 0})
	}

	for _, cor := range coords {
		x := cor.x - 1
		y := cor.y - 1
		board[x][y] = 1
	}
	for _, v := range board {
		fmt.Printf("%v\n", v)
	}
}
