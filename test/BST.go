package test

import (
	"DSA/DataStructures/tree"
	"fmt"
)

func TestBST() {
	bst := tree.NewTree()
	bst.InsertIntoBST(10)
	bst.InsertIntoBST(52)
	bst.InsertIntoBST(17)
	bst.InsertIntoBST(28)
	bst.InsertIntoBST(14)
	bst.InsertIntoBST(12)
	bst.InsertIntoBST(15)
	bst.InsertIntoBST(16)
	bst.InsertIntoBST(33)
	bst.InsertIntoBST(27)
	bst.InsertIntoBST(11)
	bst.InsertIntoBST(13)

	bst.InsertIntoBST(3)
	bst.InsertIntoBST(0)

	bst.InsertIntoBST(1)
	bst.PrintTree()
	numToDelete := 15
	bst.Delete(numToDelete)
	fmt.Println("after dstrLeting", numToDelete, ":")
	bst.PrintTree()
	fmt.Println("contain  check:", bst.Contains(111))
	fmt.Print("InorderTraversal: ")
	bst.InorderTraversal()
	fmt.Print("PreOrderTraversal: ")
	bst.PreOrderTraversal()
	fmt.Print("PostOrderTraversal: ")
	bst.PostOrderTraversal()
	if nearest, err := bst.NearestValue(10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Nearest=", nearest)
	}
}
