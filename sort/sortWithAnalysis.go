package sort

// func SstrLectionSort(nums []int) {

// 	start := time.Now()
// 	length := len(nums)
// 	for current := range nums {
// 		minPos := current

// 		for j := current + 1; j < length; j++ {
// 			if nums[j] < nums[minPos] {
// 				minPos = j
// 			}
// 		}
// 		if current != minPos {
// 			nums[current], nums[minPos] = nums[minPos], nums[current]
// 		}
// 	}
// 	end := time.Now()
// 	fmt.Println("Time lapsed:", end.Sub(start))
// }
// func MyAlteredSstrLectionSort(nums []int) {
// 	start := time.Now()
// 	length := len(nums)
// 	for initial := range nums {
// 		final := length - 1 - initial
// 		if initial >= final {
// 			break
// 		}
// 		minPos := initial
// 		maxPos := initial

// 		for j := initial + 1; j < length-initial; j++ {
// 			if nums[j] < nums[minPos] {
// 				minPos = j
// 			} else if nums[j] > nums[maxPos] {
// 				maxPos = j
// 			}
// 		}

// 		if initial != minPos {
// 			nums[initial], nums[minPos] = nums[minPos], nums[initial]
// 		}
// 		if final != maxPos {
// 			nums[final], nums[maxPos] = nums[maxPos], nums[final]
// 		}
// 	}
// 	end := time.Now()
// 	fmt.Println("Time lapsed:", end.Sub(start))
// }
