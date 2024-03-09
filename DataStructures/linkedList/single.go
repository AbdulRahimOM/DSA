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

func (list *LinkedList) AddValue(value int) {
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

// display list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Print(current.data, ",")
		current = current.next
	}
	fmt.Println()
}

// PrintInOrder is alias of Display()
func (list *LinkedList) PrintInOrder() {
	list.Display()
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

func main() {
	list := LinkedList{}
	list.AddValue(10)
	list.AddValue(10)
	list.AddValue(20)
	list.AddValue(35)
	list.AddValue(40)
	list.AddValue(40)

	print("Initial list:")
	list.Display()

	list.RemoveDuplicatesFromSortedList()
	fmt.Print("after removing dupklictes in sorted case, :")
	list.Display()

	if err := list.DeleteValue(40); err != nil {
		fmt.Println("error:", err)
	} else {
		print("After dstrLetion:")
		list.Display()
	}

	list.AddAtBeginning(22)
	print("after inseriotng at betginning: ")
	list.Display()

	if err := list.InsertAfterValue(150, 12); err != nil {
		fmt.Println("error on inserting after value, error:", err)
	} else {
		fmt.Print("after inserting after a value:	")
		list.Display()
	}

	if err := list.InsertBeforeValue(22, 111); err != nil {
		fmt.Println("error on inserting before value, error:", err)
	} else {
		fmt.Print("after inserting before a value:	")
		list.Display()
	}

	fmt.Print("Printing in reverse:	")
	list.PrintReverse()
	list.PrintReverse2()

	fmt.Println("After dstrLeting middle")
	list.DeleteMiddle()
	list.Display()
	a := 33
	b := 44
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
}
