package heaps

import (
	"errors"
	"fmt"
)

var (
	// ErrHeapEmpty returned when trying to extract from an empty heap
	ErrHeapEmpty = errors.New("heap is empty")
)

// HeapType defines whether the heap is a MinHeap or MaxHeap
type HeapType int

const (
	// MinHeap represents a minimum heap where the root is the smallest element
	MinHeap HeapType = iota
	// MaxHeap represents a maximum heap where the root is the largest element
	MaxHeap
)

// Heap represents a binary heap data structure
// It can be configured as either a MinHeap or MaxHeap
type Heap[T any] struct {
	items    []T
	heapType HeapType
	less     func(T, T) bool
}

// New creates a new empty heap with the specified type
// For MinHeap, the root will always be the minimum element
// For MaxHeap, the root will always be the maximum element
// Time complexity: O(1)
func New[T any](heapType HeapType, less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		items:    make([]T, 0),
		heapType: heapType,
		less:     less,
	}
}

// NewMinHeap creates a new minimum heap
// Time complexity: O(1)
func NewMinHeap[T any](less func(a, b T) bool) *Heap[T] {
	return New(MinHeap, less)
}

// NewMaxHeap creates a new maximum heap
// Time complexity: O(1)
func NewMaxHeap[T any](less func(a, b T) bool) *Heap[T] {
	return New(MaxHeap, less)
}

// Insert adds an element to the heap
// Time complexity: O(log n)
func (h *Heap[T]) Insert(item T) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

// Extract removes and returns the root element of the heap
// For MinHeap: returns the minimum element
// For MaxHeap: returns the maximum element
// Time complexity: O(log n)
func (h *Heap[T]) Extract() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, ErrHeapEmpty
	}

	root := h.items[0]
	last := len(h.items) - 1

	h.items[0] = h.items[last]
	h.items = h.items[:last]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return root, nil
}

// Peek returns the root element without removing it
// For MinHeap: returns the minimum element
// For MaxHeap: returns the maximum element
// Time complexity: O(1)
func (h *Heap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, ErrHeapEmpty
	}
	return h.items[0], nil
}

// IsEmpty returns true if the heap contains no elements
// Time complexity: O(1)
func (h *Heap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

// Size returns the number of elements in the heap
// Time complexity: O(1)
func (h *Heap[T]) Size() int {
	return len(h.items)
}

// Clear removes all elements from the heap
// Time complexity: O(1)
func (h *Heap[T]) Clear() {
	h.items = make([]T, 0)
}

// Copy creates a new heap with the same elements
// Time complexity: O(n)
func (h *Heap[T]) Copy() *Heap[T] {
	newItems := make([]T, len(h.items))
	copy(newItems, h.items)
	return &Heap[T]{
		items:    newItems,
		heapType: h.heapType,
		less:     h.less,
	}
}

// ToSlice returns a slice representation of the heap
// Note: The slice is not guaranteed to be sorted
// Time complexity: O(n)
func (h *Heap[T]) ToSlice() []T {
	result := make([]T, len(h.items))
	copy(result, h.items)
	return result
}

// ToSortedSlice returns a sorted slice by extracting all elements
// Time complexity: O(n log n)
func (h *Heap[T]) ToSortedSlice() []T {
	heapCopy := h.Copy()
	result := make([]T, heapCopy.Size())
	for i := 0; i < len(result); i++ {
		result[i], _ = heapCopy.Extract()
	}
	return result
}

// String returns a string representation of the heap
func (h *Heap[T]) String() string {
	if h.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", h.items)
}

// HeapType returns the type of the heap (MinHeap or MaxHeap)
func (h *Heap[T]) HeapType() HeapType {
	return h.heapType
}

// heapifyUp moves the element at the given index up the heap
// to maintain the heap property
// Time complexity: O(log n)
func (h *Heap[T]) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.shouldSwap(index, parent) {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

// heapifyDown moves the element at the given index down the heap
// to maintain the heap property
// Time complexity: O(log n)
func (h *Heap[T]) heapifyDown(index int) {
	n := len(h.items)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.shouldSwap(left, smallest) {
			smallest = left
		}
		if right < n && h.shouldSwap(right, smallest) {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.swap(index, smallest)
		index = smallest
	}
}

// shouldSwap determines if two elements should be swapped based on heap type
func (h *Heap[T]) shouldSwap(i, j int) bool {
	if h.heapType == MinHeap {
		return h.less(h.items[i], h.items[j])
	}
	return h.less(h.items[j], h.items[i])
}

// swap swaps two elements in the heap
func (h *Heap[T]) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
