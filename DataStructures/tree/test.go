package tree

import "fmt"

var newBST tree = tree{
	root: &treeNode{
		data: 10,
		left: &treeNode{
			data: 5,
			left: &treeNode{
				data: 0,
				left: nil,
				right: &treeNode{
					data:  3,
					left:  nil,
					right: &treeNode{4, nil, nil},
				},
			},
			right: &treeNode{8, nil, nil},
		},
		right: &treeNode{25, nil, nil},
	},
}

func InnerPkgTest() {
	newBST.PrintTree()
	fmt.Println("test result (if BST or not):", newBST.IsProperBSTree())
}
