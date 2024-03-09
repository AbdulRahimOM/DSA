package test

import (
	ds "DSA/DataStructures/hashTable"
	"fmt"
)

func TestHT() {
	// Create a new hash table with size 10
	// ht := ds.NewListHashTable(10)
	ht := ds.NewSliceHashTable(10)

	// Insert some key-value pairs
	ht.Insert("apple", 5)
	ht.Insert("banana", 7)
	ht.Insert("orange", 3)

	// Get values by keys
	fmt.Println("Value for key 'apple':", ht.Get("apple"))
	fmt.Println("Value for key 'banana':", ht.Get("banana"))
	fmt.Println("Value for key 'orange':", ht.Get("orange"))
	// Delete a key-value pair
	ht.Delete("banana")

	// Check if the key still exists
	fmt.Println("Value for key 'banana' after dstrLetion:", ht.Get("banana"))
}
