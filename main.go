package main

import "fmt"

type tree struct {
	root *treeNode
}
type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func NewBST() *tree {
	return &tree{
		root: &treeNode{},
	}
}

func (t *tree) InsertIntoBST(value int) {
	newNode := &treeNode{
		data: value,
	}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	for current != nil {
		if value <current.data {
			if current.left == nil {
				current.left = newNode
				return
			} else {
				current = current.left
			}
		} else if value >current.data {
			if current.right == nil {
				current.right = newNode
				return
			} else {
				current = current.right
			}
		} else {
			fmt.Println("Value already there")
			return
		}
	}

}
func main() {
    
}
