package tools

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func PrintArrayAsBinaryTree(arr []int) {
	length := len(arr)
	if length == 0 {
		return
	}

	newArr := make([]*node, length)

	newArr[0] = &node{
		data: arr[0],
	}
	for i := 1; i < length; i++ {
		newArr[i] = &node{
			data: arr[i],
		}
		if i%2 == 1 { //isOdd
			newArr[i/2].left = newArr[i]
		} else {
			newArr[(i-1)/2].right = newArr[i]
		}

	}
	newArr[0].printThrough("")
}

func (n *node) printThrough(preText string) {
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
