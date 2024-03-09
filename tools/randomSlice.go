package tools

import (
	"math/rand"
	"time"
)

func GenerateRandomIntSlice(size int, max int) []int {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice with 1000 elements
	numbers := make([]int, size)
	// Fill the slice with random numbers
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(max + 1) // Adjust the upper bound as needed
	}

	return numbers
}
