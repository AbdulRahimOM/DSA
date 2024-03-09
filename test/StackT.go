package test

import (
	"DSA/DataStructures/stack"
	"DSA/types"
	"fmt"
)

func TestStack() {
	// var stack types.Stack = &stack.StackSlice{}
	var stack types.Stack = &stack.StackList{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	//print elements
	fmt.Print("Elements are:  ")
	stack.PrintItems()

	//size
	fmt.Println("Stack size:", stack.Size())

	//pop
	if output, ok := stack.Pop(); !ok {
		fmt.Println("Stack is Empty!!!")
	} else {
		fmt.Println("Pop:", output)
	}

	fmt.Print("Elements are:  ")
	stack.PrintItems()

	if output, ok := stack.Peek(); !ok {
		fmt.Println("Stack is Empty!!!")
	} else {
		fmt.Println("Peek:", output)
	}

	fmt.Print("Elements are:  ")
	stack.PrintItems()

	fmt.Println("Stack empty status:", stack.IsEmpty())
}
