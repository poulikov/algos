package queues

import (
	"errors"
	"fmt"
)

var (
	// ErrDequeEmpty returned when trying to pop from an empty deque
	ErrDequeEmpty = errors.New("deque is empty")
)

// Deque represents a double-ended queue (deque) that allows insertion and deletion at both ends
type Deque[T any] struct {
	items []T
}

// NewDeque creates a new empty deque
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		items: make([]T, 0),
	}
}

// PushFront adds an element to the front of the deque
// Time complexity: O(n) - due to slice shifting
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// PushBack adds an element to the back of the deque
// Time complexity: O(1) amortized
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// PopFront removes and returns the element from the front of the deque
// Time complexity: O(n) - due to slice shifting
func (d *Deque[T]) PopFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrDequeEmpty
	}

	item := d.items[0]
	d.items = d.items[1:]
	return item, nil
}

// PopBack removes and returns the element from the back of the deque
// Time complexity: O(1)
func (d *Deque[T]) PopBack() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrDequeEmpty
	}

	index := len(d.items) - 1
	item := d.items[index]
	d.items = d.items[:index]
	return item, nil
}

// PeekFront returns the element at the front without removing it
// Time complexity: O(1)
func (d *Deque[T]) PeekFront() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrDequeEmpty
	}
	return d.items[0], nil
}

// PeekBack returns the element at the back without removing it
// Time complexity: O(1)
func (d *Deque[T]) PeekBack() (T, error) {
	if d.IsEmpty() {
		var zero T
		return zero, ErrDequeEmpty
	}
	return d.items[len(d.items)-1], nil
}

// IsEmpty returns true if the deque contains no elements
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

// Size returns the number of elements in the deque
func (d *Deque[T]) Size() int {
	return len(d.items)
}

// Clear removes all elements from the deque
func (d *Deque[T]) Clear() {
	d.items = make([]T, 0)
}

// Copy creates a new deque with the same elements
// Time complexity: O(n)
func (d *Deque[T]) Copy() *Deque[T] {
	newItems := make([]T, len(d.items))
	copy(newItems, d.items)
	return &Deque[T]{items: newItems}
}

// ToSlice returns a slice representation of the deque
func (d *Deque[T]) ToSlice() []T {
	result := make([]T, len(d.items))
	copy(result, d.items)
	return result
}

// String returns a string representation of the deque
func (d *Deque[T]) String() string {
	if d.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", d.items)
}
