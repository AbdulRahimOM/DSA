package tree

import (
	"fmt"
)

func (bst *tree) InsertIntoBST(value int) {
	newNode := &treeNode{
		data: value,
	}
	if bst.root == nil {
		bst.root = newNode
		return
	}

	current := bst.root
	for current != nil {
		if value < current.data {
			if current.left == nil {
				current.left = newNode
				return
			} else {
				current = current.left
			}
		} else if value > current.data {
			if current.right == nil {
				current.right = newNode
				return
			} else {
				current = current.right
			}
		} else { //case of 'equal value'
			return
		}
	}
}

func (root *treeNode) deleteNode(key int) *treeNode {
	if root == nil {
		return nil
	}
	if key < root.data {
		root.left = root.left.deleteNode(key)
	} else if key > root.data {
		root.right = root.right.deleteNode(key)
	} else {
		if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			root.data = root.deleteMinValueNode()
		}
	}
	return root
}
func (bst *tree) Delete(key int) {
	if bst.root == nil {
		return
	}
	bst.root = bst.root.deleteNode(key)
}
func (n *treeNode) deleteMinValueNode() int {
	for n.left.left != nil {
		n = n.left
	}
	minValue := n.left.data
	n.left = n.left.right
	return minValue
}

func (bst *tree) Search(key int) bool {
	current := bst.root
	for current != nil {
		if key < current.data {
			current = current.left
		} else if key > current.data {
			current = current.right
		} else {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func nearestNumber(key, value1, value2 int) int {
	if abs(value1-key) < abs(value2-key) {
		return value1
	} else {
		return value2
	}
}
func (bst *tree) NearestValue(key int) (int, error) {
	if bst.root == nil {
		return 0, fmt.Errorf("tree is empty")
	}
	current := bst.root
	nearest := current.data
	for current != nil {
		nearest = nearestNumber(key, nearest, current.data)
		if current.data > key {
			current = current.left
		} else if current.data < key {
			current = current.right
		} else {
			return key, nil
		}
	}
	return nearest, nil
}

func (bst *tree) IsProperBSTree() bool {
	status, _ := bst.root.checkNGetMax()
	return status
	// return bst.root.checkNode()
}
func (n *treeNode) checkNGetMax() (bool, int) {
	if n.left != nil {
		ok, value := n.left.checkNGetMax()
		if !ok {
			return false, 0
		}
		if value >= n.data {
			return false, 0
		}
	}
	if n.right != nil {
		ok, rightMax := n.right.checkNGetMax()
		if !ok {
			return false, 0
		}
		if rightMax <= n.data {
			return false, 0
		}
		return true, rightMax
	} else {
		return true, n.data
	}
}
