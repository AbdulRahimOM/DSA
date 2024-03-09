package types

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
