package queues

import (
	"errors"
	"fmt"
)

var (
	// ErrQueueEmpty returned when trying to dequeue from an empty queue
	ErrQueueEmpty = errors.New("queue is empty")
)

// Queue represents a FIFO (First In, First Out) data structure
type Queue[T any] struct {
	items []T
}

// New creates a new empty queue
func New[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an element to the back of the queue
// Time complexity: O(1) amortized
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element at the front of the queue
// Time complexity: O(n) - due to slice shifting
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrQueueEmpty
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// DequeueFast removes and returns the element at the front using circular buffer approach
// Time complexity: O(1) amortized
func (q *Queue[T]) DequeueFast() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrQueueEmpty
	}

	item := q.items[0]
	q.items[0] = *new(T) // Help GC
	q.items = q.items[1:]
	return item, nil
}

// Peek returns the element at the front of the queue without removing it
// Time complexity: O(1)
func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, ErrQueueEmpty
	}
	return q.items[0], nil
}

// IsEmpty returns true if the queue contains no elements
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue
func (q *Queue[T]) Clear() {
	q.items = make([]T, 0)
}

// Copy creates a new queue with the same elements
// Time complexity: O(n)
func (q *Queue[T]) Copy() *Queue[T] {
	newItems := make([]T, len(q.items))
	copy(newItems, q.items)
	return &Queue[T]{items: newItems}
}

// ToSlice returns a slice representation of the queue
// The first element is the front of the queue, the last element is the back
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, len(q.items))
	copy(result, q.items)
	return result
}

// String returns a string representation of the queue
func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", q.items)
}
