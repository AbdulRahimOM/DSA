package test

import (
	"DSA/DataStructures/trie"
	"fmt"
)

func TestTrie() {
	trie := trie.NewTrie()
	words := []string{"apple", "banana", "app", "orange", "grape", "aplicant", "applicant", "ap"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Search("apple"))  // Output: true
	fmt.Println(trie.Search("orange")) // Output: true
	fmt.Println(trie.Search("apples")) // Output: false

	predictions := trie.Predict("app")
	fmt.Println("Predictions=", predictions)

}
