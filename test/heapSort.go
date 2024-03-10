package test

import (
	"DSA/sort"
	"DSA/tools"
	"fmt"
)

func TestHeapSort() {
	// arr := []int{12, 11, 13, 5, 6, 7}

	fmt.Println("Ascending order:")
	arr := tools.GenerateRandomIntSlice(10, 20)
	fmt.Println("Original array:", arr)

	sort.HeapSortAsc(arr)

	fmt.Println("Sorted array:", arr)

	fmt.Println("______________\nDescending order:")
	arr2 := tools.GenerateRandomIntSlice(10, 20)
	fmt.Println("Original array:", arr2)

	sort.HeapSortDesc(arr2)

	fmt.Println("Sorted array:", arr2)
}
