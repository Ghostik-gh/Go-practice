package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {

	if t.val == 0 && t.left == nil && t.right == nil {
		t.val = value
	}

	if t == nil {
		return errors.New("this tree is nil")
	}

	if t.val == value {
		return errors.New("this node value already exists")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}
	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)

	}

	return nil
}

func (t *TreeNode) PrintInOrder() {
	if t == nil {
		return
	}
	t.left.PrintInOrder()
	fmt.Printf("%v ", t.val)
	t.right.PrintInOrder()
}

func (t *TreeNode) Search(key int) bool {

	if t == nil {
		return false

	}

	if t.val == key {
		return true
	}
	if t.left == nil && t.right == nil {
		return false
	}

	if key < t.val {
		return t.left.Search(key)
	} else {
		return t.right.Search(key)
	}

}
