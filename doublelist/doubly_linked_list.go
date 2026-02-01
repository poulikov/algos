package doublelist

import (
	"fmt"
	"reflect"
)

// Node represents a node in the doubly linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// New creates a new empty doubly linked list
func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Append adds a new node with the given value to the end of the list
func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Prev = dll.Tail
		dll.Tail.Next = newNode
		dll.Tail = newNode
	}

	dll.Size++
}

// Prepend adds a new node with the given value to the beginning of the list
func (dll *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}

	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		newNode.Next = dll.Head
		dll.Head.Prev = newNode
		dll.Head = newNode
	}

	dll.Size++
}

// InsertAt inserts a new node with the given value at the specified index
func (dll *DoublyLinkedList[T]) InsertAt(index int, value T) error {
	if index < 0 || index > dll.Size {
		return fmt.Errorf("index out of bounds")
	}

	if index == 0 {
		dll.Prepend(value)
		return nil
	}

	if index == dll.Size {
		dll.Append(value)
		return nil
	}

	newNode := &Node[T]{Value: value}
	current := dll.getNodeAt(index)

	newNode.Next = current
	newNode.Prev = current.Prev
	current.Prev.Next = newNode
	current.Prev = newNode

	dll.Size++
	return nil
}

// RemoveAt removes the node at the specified index and returns its value
func (dll *DoublyLinkedList[T]) RemoveAt(index int) (T, error) {
	var zero T
	if index < 0 || index >= dll.Size {
		return zero, fmt.Errorf("index out of bounds")
	}

	node := dll.getNodeAt(index)
	value := node.Value

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		dll.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		dll.Tail = node.Prev
	}

	dll.Size--
	return value, nil
}

// RemoveFirst removes the first node and returns its value
func (dll *DoublyLinkedList[T]) RemoveFirst() (T, error) {
	var zero T
	if dll.Head == nil {
		return zero, fmt.Errorf("list is empty")
	}

	value := dll.Head.Value

	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Head = dll.Head.Next
		dll.Head.Prev = nil
	}

	dll.Size--
	return value, nil
}

// RemoveLast removes the last node and returns its value
func (dll *DoublyLinkedList[T]) RemoveLast() (T, error) {
	var zero T
	if dll.Tail == nil {
		return zero, fmt.Errorf("list is empty")
	}

	value := dll.Tail.Value

	if dll.Head == dll.Tail {
		dll.Head = nil
		dll.Tail = nil
	} else {
		dll.Tail = dll.Tail.Prev
		dll.Tail.Next = nil
	}

	dll.Size--
	return value, nil
}

// Get returns the value at the specified index
func (dll *DoublyLinkedList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= dll.Size {
		return zero, fmt.Errorf("index out of bounds")
	}

	node := dll.getNodeAt(index)
	return node.Value, nil
}

// Set updates the value at the specified index
func (dll *DoublyLinkedList[T]) Set(index int, value T) error {
	if index < 0 || index >= dll.Size {
		return fmt.Errorf("index out of bounds")
	}

	node := dll.getNodeAt(index)
	node.Value = value
	return nil
}

// IndexOf returns the index of the first occurrence of the value in the list
func (dll *DoublyLinkedList[T]) IndexOf(value T) int {
	current := dll.Head
	index := 0

	for current != nil {
		if reflect.DeepEqual(current.Value, value) {
			return index
		}
		current = current.Next
		index++
	}

	return -1
}

// Contains checks if the list contains the specified value
func (dll *DoublyLinkedList[T]) Contains(value T) bool {
	return dll.IndexOf(value) != -1
}

// Size returns the number of elements in the list
func (dll *DoublyLinkedList[T]) Len() int {
	return dll.Size
}

// IsEmpty checks if the list is empty
func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	return dll.Size == 0
}

// Clear removes all elements from the list
func (dll *DoublyLinkedList[T]) Clear() {
	dll.Head = nil
	dll.Tail = nil
	dll.Size = 0
}

// ToSlice returns a slice containing all elements in the list from head to tail
func (dll *DoublyLinkedList[T]) ToSlice() []T {
	result := make([]T, dll.Size)
	current := dll.Head
	index := 0

	for current != nil {
		result[index] = current.Value
		current = current.Next
		index++
	}

	return result
}

// ToSliceReverse returns a slice containing all elements in the list from tail to head
func (dll *DoublyLinkedList[T]) ToSliceReverse() []T {
	result := make([]T, dll.Size)
	current := dll.Tail
	index := 0

	for current != nil {
		result[index] = current.Value
		current = current.Prev
		index++
	}

	return result
}

// ForEach applies the given function to each element in the list from head to tail
// If the function returns an error, iteration stops and the error is returned
func (dll *DoublyLinkedList[T]) ForEach(fn func(T) error) error {
	current := dll.Head
	for current != nil {
		if err := fn(current.Value); err != nil {
			return err
		}
		current = current.Next
	}
	return nil
}

// ForEachReverse applies the given function to each element in the list from tail to head
// If the function returns an error, iteration stops and the error is returned
func (dll *DoublyLinkedList[T]) ForEachReverse(fn func(T) error) error {
	current := dll.Tail
	for current != nil {
		if err := fn(current.Value); err != nil {
			return err
		}
		current = current.Prev
	}
	return nil
}

// getNodeAt returns the node at the specified index
func (dll *DoublyLinkedList[T]) getNodeAt(index int) *Node[T] {
	// Optimize by choosing direction based on index position
	if index < dll.Size/2 {
		// Traverse from head
		current := dll.Head
		for i := 0; i < index; i++ {
			current = current.Next
		}
		return current
	} else {
		// Traverse from tail
		current := dll.Tail
		for i := dll.Size - 1; i > index; i-- {
			current = current.Prev
		}
		return current
	}
}
