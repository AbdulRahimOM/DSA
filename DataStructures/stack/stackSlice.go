package stack

import (
	types "DSA/types"
	"fmt"
	"testing"
)

// Example implementation of a stack using slice
type StackSlice struct {
	data      []interface{}
	lastIndex int
}

func (s *StackSlice) Push(value interface{}) {
	s.data = append(s.data, value)
	s.lastIndex++
}

func (s *StackSlice) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	value := s.data[s.lastIndex]
	s.data = s.data[:s.lastIndex]
	s.lastIndex--
	return value, true
}

func (s *StackSlice) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.data[s.lastIndex], true
}

func (s *StackSlice) IsEmpty() bool {
	return s.lastIndex == -1
}

func (s *StackSlice) Size() int {
	return s.lastIndex
}

func (s *StackSlice) PrintItems() {
	fmt.Println(s.data...)
}

func TestSingle(t *testing.T) {
	var stack types.Stack = &StackSlice{}
	// var stack types.Stack = &stack.StackList{}

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

// // Example implementation of a stack using slice
// type StackSlice struct {
// 	data []interface{}
// }

// func (s *StackSlice) Push(value interface{}) {
// 	s.data = append(s.data, value)
// }

// func (s *StackSlice) Pop() (interface{}, bool) {
// 	if s.IsEmpty() {
// 		return nil, false
// 	}
// 	lastIndex := len(s.data) - 1
// 	value := s.data[lastIndex]
// 	s.data = s.data[:lastIndex]
// 	return value, true
// }

// func (s *StackSlice) Peek() (interface{}, bool) {
// 	if s.IsEmpty() {
// 		return nil, false
// 	}
// 	return s.data[len(s.data)-1], true
// }

// func (s *StackSlice) IsEmpty() bool {
// 	return s.data[0]==nil
// }

// func (s *StackSlice) Size() int {
// 	return len(s.data)
// }

// func (s *StackSlice) PrintItems() {
// 	fmt.Println(s.data...)
// }

// func TestSingle(t *testing.T) {
// 	var stack types.Stack = &StackSlice{}
// 	// var stack types.Stack = &stack.StackList{}

// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)

// 	//print elements
// 	fmt.Print("Elements are:  ")
// 	stack.PrintItems()

// 	//size
// 	fmt.Println("Stack size:", stack.Size())

// 	//pop
// 	if output, ok := stack.Pop(); !ok {
// 		fmt.Println("Stack is Empty!!!")
// 	} else {
// 		fmt.Println("Pop:", output)
// 	}

// 	fmt.Print("Elements are:  ")
// 	stack.PrintItems()

// 	if output, ok := stack.Peek(); !ok {
// 		fmt.Println("Stack is Empty!!!")
// 	} else {
// 		fmt.Println("Peek:", output)
// 	}

// 	fmt.Print("Elements are:  ")
// 	stack.PrintItems()

// 	fmt.Println("Stack empty status:", stack.IsEmpty())
// }
