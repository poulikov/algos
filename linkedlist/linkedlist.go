package linkedlist

// Node represents a node in the singly linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a singly linked list
type LinkedList[T any] struct {
	Head *Node[T]
}

// NewLinkedList creates a new empty linked list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds a new node with the given value to the end of the list
func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Prepend adds a new node with the given value to the beginning of the list
func (ll *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: ll.Head}
	ll.Head = newNode
}

// ForEach iterates through the list and applies the given function to each element
func (ll *LinkedList[T]) ForEach(processFunc func(T)) {
	current := ll.Head
	for current != nil {
		processFunc(current.Value)
		current = current.Next
	}
}

// Reverse reverses the linked list in-place
func (ll *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := ll.Head

	for current != nil {
		next := current.Next // Store the next node
		current.Next = prev  // Reverse the link
		prev = current       // Move prev to current
		current = next       // Move to the next node
	}

	ll.Head = prev // Update head to the new first node
}

// ToSlice converts the linked list to a slice for easier testing
func (ll *LinkedList[T]) ToSlice() []T {
	var result []T
	current := ll.Head

	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}

	return result
}
