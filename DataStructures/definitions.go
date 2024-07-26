package types

/*	These define the methods that the data structures in this package has implemented.
 */

// Stack interface
type Stack interface {
	Push(value interface{})    // Pushes a value onto the stack
	Pop() (interface{}, bool)  // Pops a value from the stack
	Peek() (interface{}, bool) // Peeks at the top value of the stack without removing it
	IsEmpty() bool             // Checks if the stack is empty
	Size() int                 // Returns the size of the stack
	PrintItems()               // Prints  all elements in the stack
}

// Queue interface
type Queue interface {
	Enqueue(int)      // Adds an element to the queue
	Dequeue() int     // Removes an element from the queue
	IsEmpty() bool    // Checks if the queue is empty
	DisplayElements() // Displays all elements in the queue
}

// Graph interface
type Graph interface {
	LookUpByBFS(int) // Look up by Breadth First Search
	LookUpByDFS(int) // Look up by Depth First Search
	AddVertex(int) error
	AddEdge(int, int) error
	PrintGraph()
}

// Heap interface
type Heap interface {
	Insert(int)               // Inserts a value into the heap
	ExtractMax() (int, error) // Extracts the maximum value from the heap (the value will be removed as well)
	Delete(int) bool          // Deletes a value from the heap
	BuildHeap([]int)          // Builds a heap from an array
	HeapifyDown(int, int)     // Heapify down may be used for sorting
	PrintAsTree()
}

// LinkedList interface
type LinkedList interface {
	AddAtBeginning(int)               // Inserts a value at the head of the linked list
	InsertAtTail(int)                 // Inserts a value at the tail of the linked list
	DeleteValue(int)                  // Deletes a value from the linked list
	DeleteMiddle()                    // Deletes the middle element of the linked list
	InsertAfterValue(int, int) error  // Inserts a value after a given value in the linked list
	InsertBeforeValue(int, int) error // Inserts a value before a given value in the linked list
	Search(int) bool                  // Searches for a value in the linked list
	PrintList()                       // Prints the linked list
	PrintReverse()                    // Prints the linked list in reverse
	RemoveDuplicatesFromSortedList()  // Removes duplicates from a sorted linked list
}

// Tree interface
type Tree interface {
	PrintTree()          // Prints the tree in a visualisable way
	PreOrderTraversal()  // Pre-order traversal of the binary search tree
	InOrderTraversal()   // In-order traversal of the binary search tree
	PostOrderTraversal() // Post-order traversal of the binary search tree
}

// BinarySearchTree interface
type BinarySearchTree interface {
	Tree                  // All methods of Tree interface should be implemented
	InsertIntoBST(int)    // Inserts a value into the binary search tree
	Search(int) bool      // Searches for a value in the binary search tree
	Delete(int) bool      // Deletes a value from the binary search tree
	NearestValue(int) int // Finds the nearest value to the given value in the binary search tree
	IsProperBSTree() bool // Checks if the binary search tree is a proper binary search tree
}

// Trie interface
type Trie interface {
	Insert(string)           // Inserts a string into the trie
	Search(string) bool      // Searches for a string in the trie
	Predict(string) []string // Predicts the words that start with the given prefix
}
