package test

import (
	"DSA/DataStructures/heap"
	"DSA/tools"
	"fmt"
)

func TestHeapSort() {
	// arr := []int{12, 11, 13, 5, 6, 7}

	fmt.Println("Ascending order:")
	arr := tools.GenerateRandomIntSlice(10, 20)
	fmt.Println("Original array:", arr)

	heap.HeapSortAsc(arr)

	fmt.Println("Sorted array:", arr)

	fmt.Println("______________\nDescending order:")
	arr2 := tools.GenerateRandomIntSlice(10, 20)
	fmt.Println("Original array:", arr2)

	heap.HeapSortDesc(arr2)

	fmt.Println("Sorted array:", arr2)
}
