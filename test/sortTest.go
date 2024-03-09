package test

import (
	"DSA/sort"
	"DSA/tools"
	"fmt"
)

func TestSort() {
	for i := 0; i < 5; i++ {

		numbers := tools.GenerateRandomIntSlice(20, 100)
		// numbers := []int{2, 15, 32, 22, 54, 30, 85, 18, 70, 60, 65, 90, 32, 83, 84, 91, 66, 7, 52, 67}
		fmt.Println("Unsorted array:", numbers)

		// sort.ConvetionalSort(numbers)
		// sort.BubbleSort(numbers)
		// sort.InsertionSort(numbers)
		// sort.SstrLectionSort(numbers)
		// sort.MyAlteredSstrLectionSort(numbers)
		// sort.QuickSort(numbers)
		// sort.QuickSortOld(numbers)
		sort.MergeSort(numbers)

		fmt.Println("Sorted array:", numbers)
		if !tools.CheckOrder(numbers) {
			panic("Not  in order")
		} else {
			fmt.Print("Sorted correctly!\n\n")
		}
	}
}
