package stack

import "fmt"

// Example implementation of a stack using linked list
type StackList struct {
	top *node
}
type node struct {
	data interface{}
	next *node
}

func (s *StackList) Push(value interface{}) {
	s.top = &node{
		data: value,
		next: s.top,
	}
}

func (s *StackList) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	returnValue := s.top.data
	s.top = s.top.next
	return returnValue, true
}

func (s *StackList) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	returnValue := s.top.data
	return returnValue, true
}

func (s *StackList) IsEmpty() bool {
	return s.top == nil
}

func (s *StackList) Size() int {
	current := s.top
	var size int
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (s *StackList) PrintItems() {
	defer fmt.Println()
	current := s.top
	for current != nil {
		defer fmt.Printf("%v, ", current.data)
		current = current.next
	}
}
