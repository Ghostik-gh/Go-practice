package main

import (
	"encoding/json"
	"fmt"
)

type IdTree struct {
	Id       uint32
	Children []IdTree
}

type Comment struct {
	Id    uint32
	Title string
}

type CommentTree struct {
	Comment  Comment
	Children []CommentTree
}

func main() {
	tree := IdTree{
		Id: 1,
		Children: []IdTree{
			{
				Id: 2,
			},
			{
				Id: 3,
				Children: []IdTree{
					{
						Id: 4,
					},
					{
						Id: 5,
					},
				},
			},
			{
				Id: 13,
				Children: []IdTree{
					{
						Id: 14,
					},
					{
						Id: 15,
					},
				},
			},
		},
	}
	// fmt.Printf("tree: %s\n", tree)

	withSpaces, _ := json.MarshalIndent(tree, "", "  ")

	bytes, _ := json.Marshal(tree)

	fmt.Printf("b: %s\n", withSpaces)

	fmt.Printf("b: %v\n", bytes)
}
