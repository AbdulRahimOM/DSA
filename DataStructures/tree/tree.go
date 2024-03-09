package tree

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

type tree struct {
	root *treeNode
}
func NewTree() *tree {
	return &tree{}
}
func InitiateTreeWithRootValue(value int) *tree {
	return &tree{
		root: &treeNode{
			data: value,
		},
	}
}
func InitiateTreeWithRootNode(root *treeNode) *tree {
	return &tree{
		root: root,
	}
}
func (tree *tree) PrintTree() {
	if tree.root.left==nil{
		fmt.Println("tree.root.left==nil")
	}else{
		fmt.Println("tree.root.left=",tree.root.left)
	}
	if tree.root != nil {
		tree.root.printThrough("")
	}
}

func (n *treeNode) printThrough(preText string) {
	strBlank := "   "
	strT := "├──"
	strL := "└──"
	strLine := "│   "

	fmt.Println(n.data)
	if n.left != nil && n.right != nil {

		fmt.Print(preText, strT)
		n.right.printThrough(preText + strLine)

		fmt.Print(preText, strL)
		n.left.printThrough(preText + strBlank)

	} else if n.right != nil {
		fmt.Print(preText, strL)
		n.right.printThrough(preText + strBlank)
	} else if n.left != nil {
		fmt.Print(preText, strL)
		n.left.printThrough(preText + strBlank)
	}
}
