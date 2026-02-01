package stacks

import (
	"errors"
	"fmt"
)

var (
	// ErrStackEmpty returned when trying to pop from an empty stack
	ErrStackEmpty = errors.New("stack is empty")
)

// Stack represents a LIFO (Last In, First Out) data structure
type Stack[T any] struct {
	items []T
}

// New creates a new empty stack
func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an element to the top of the stack
// Time complexity: O(1)
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the element at the top of the stack
// Time complexity: O(1)
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrStackEmpty
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

// Peek returns the element at the top of the stack without removing it
// Time complexity: O(1)
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrStackEmpty
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack
func (s *Stack[T]) Clear() {
	s.items = make([]T, 0)
}

// Copy creates a new stack with the same elements
// Time complexity: O(n)
func (s *Stack[T]) Copy() *Stack[T] {
	newItems := make([]T, len(s.items))
	copy(newItems, s.items)
	return &Stack[T]{items: newItems}
}

// ToSlice returns a slice representation of the stack
// The first element is the bottom of the stack, the last element is the top
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, len(s.items))
	copy(result, s.items)
	return result
}

// String returns a string representation of the stack
func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", s.items)
}
