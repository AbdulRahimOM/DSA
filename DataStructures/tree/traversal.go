package tree

import "fmt"

// InorderTraversal
func (bst *tree) InorderTraversal() {
	bst.root.inorderTraversalThroughNodes()
	fmt.Println()
}
func (n *treeNode) inorderTraversalThroughNodes() {
	if n == nil {
		return
	}
	n.left.inorderTraversalThroughNodes()
	fmt.Print(n.data, ",")
	n.right.inorderTraversalThroughNodes()
}

// PreOrderTraversal
func (bst *tree) PreOrderTraversal() {
	bst.root.preOrderTraversalThroughNodes()
	fmt.Println()
}
func (n *treeNode) preOrderTraversalThroughNodes() {
	if n == nil {
		return
	}

	fmt.Print(n.data, ",")
	n.left.preOrderTraversalThroughNodes()
	n.right.preOrderTraversalThroughNodes()
}

// PostOrderTraversal
func (bst *tree) PostOrderTraversal() {
	bst.root.postOrderTraversalThroughNodes()
	fmt.Println()
}
func (n *treeNode) postOrderTraversalThroughNodes() {
	if n == nil {
		return
	}

	n.left.postOrderTraversalThroughNodes()
	n.right.postOrderTraversalThroughNodes()
	fmt.Print(n.data, ",")
}
