package ds

import (
	"errors"
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
}

func (list *LinkedList) InsertAtTail(value int) {
	newnode := &node{data: value}

	if list.head == nil {
		list.head = newnode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newnode
	}
}

// Print list
func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, ",")
		current = current.next
	}
	fmt.Println()
}

func (list *LinkedList) PrintReverse() {
	current := list.head
	defer fmt.Println()
	for current != nil {
		defer fmt.Print(current.data, ",")
		current = current.next
	}
}

func (list *LinkedList) PrintReverse2() {
	current := list.head
	var dataArr []int
	for current != nil {
		dataArr = append(dataArr, current.data)
		current = current.next
	}
	for i := len(dataArr) - 1; i >= 0; i-- {
		fmt.Print(dataArr[i], ",")
	}
	fmt.Println()
}

// deletes the first node containing the value
func (list *LinkedList) DeleteValue(value int) error {
	if list.head == nil {
		return errors.New("error: empty list")
	}
	if list.head.data == value {
		list.head = list.head.next
	} else {
		current := list.head
		for current.next != nil {
			if current.next.data == value {
				current.next = current.next.next
				return nil
			}
			current = current.next
		}
		return errors.New("value not present")
	}
	return nil
}
func (list *LinkedList) AddAtBeginning(value int) {
	list.head = &node{
		data: value,
		next: list.head,
	}
}
func (list *LinkedList) DeleteMiddle() {
	var length int

	current := list.head
	for current != nil {
		length++
		current = current.next
	}
	current = list.head
	for i := 0; i < length/2-1; i++ {
		current = current.next
	}
	current.next = current.next.next
}
func (list *LinkedList) InsertAfterValue(value int, newValue int) error {
	if list.head == nil {
		return errors.New("error: empty list")
	}
	current := list.head
	for current != nil {
		if current.data == value {
			current.next = &node{
				data: newValue,
				next: current.next,
			}
			return nil
		}
		current = current.next
	}
	return errors.New("value not present")
}

func (list *LinkedList) InsertBeforeValue(value int, newValue int) error {
	if list.head == nil {
		return errors.New("error: empty list")
	}
	if list.head.data == value {
		list.head = &node{
			data: newValue,
			next: list.head,
		}
		return nil
	}
	current := list.head
	for current.next != nil {
		if current.next.data == value {
			current.next = &node{
				data: newValue,
				next: current.next,
			}
			return nil
		}
		current = current.next
	}
	return errors.New("value not present")
}

func (list *LinkedList) RemoveDuplicatesFromSortedList() {
	if list.head == nil {
		return
	}
	current := list.head
	for current.next != nil {
		{
			if current.data == current.next.data {
				current.next = current.next.next
				continue
			}
			current = current.next
		}
	}
}