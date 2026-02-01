package heaps

import (
	"golang.org/x/exp/constraints"
)

// PriorityType determines whether higher or lower priority values are served first
type PriorityType int

const (
	// MaxPriority serves elements with higher priority values first
	MaxPriority PriorityType = iota
	// MinPriority serves elements with lower priority values first
	MinPriority
)

// PriorityQueue represents a priority queue that serves elements based on their priority
// It can be configured as MaxPriority (higher values first) or MinPriority (lower values first)
type PriorityQueue[T any, P constraints.Ordered] struct {
	heap         *Heap[priorityItem[T, P]]
	priorityType PriorityType
}

type priorityItem[T any, P constraints.Ordered] struct {
	value    T
	priority P
}

// NewPriorityQueue creates a new priority queue
// Time complexity: O(1)
func NewPriorityQueue[T any, P constraints.Ordered](priorityType PriorityType) *PriorityQueue[T, P] {
	less := func(a, b priorityItem[T, P]) bool {
		return a.priority < b.priority
	}

	var heapType HeapType
	if priorityType == MaxPriority {
		heapType = MaxHeap
	} else {
		heapType = MinHeap
	}

	return &PriorityQueue[T, P]{
		heap:         New(heapType, less),
		priorityType: priorityType,
	}
}

// NewMaxPriorityQueue creates a new max priority queue (higher values served first)
// Time complexity: O(1)
func NewMaxPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T, P](MaxPriority)
}

// NewMinPriorityQueue creates a new min priority queue (lower values served first)
// Time complexity: O(1)
func NewMinPriorityQueue[T any, P constraints.Ordered]() *PriorityQueue[T, P] {
	return NewPriorityQueue[T, P](MinPriority)
}

// Enqueue adds an element to the priority queue with the given priority
// Time complexity: O(log n)
func (pq *PriorityQueue[T, P]) Enqueue(value T, priority P) {
	item := priorityItem[T, P]{value: value, priority: priority}
	pq.heap.Insert(item)
}

// Dequeue removes and returns the element with the highest priority
// For MaxPriority: returns element with highest priority value
// For MinPriority: returns element with lowest priority value
// Time complexity: O(log n)
func (pq *PriorityQueue[T, P]) Dequeue() (T, P, error) {
	item, err := pq.heap.Extract()
	if err != nil {
		var zero T
		var zeroPriority P
		return zero, zeroPriority, err
	}
	return item.value, item.priority, nil
}

// Peek returns the element with the highest priority without removing it
// Time complexity: O(1)
func (pq *PriorityQueue[T, P]) Peek() (T, P, error) {
	item, err := pq.heap.Peek()
	if err != nil {
		var zero T
		var zeroPriority P
		return zero, zeroPriority, err
	}
	return item.value, item.priority, nil
}

// IsEmpty returns true if the priority queue is empty
// Time complexity: O(1)
func (pq *PriorityQueue[T, P]) IsEmpty() bool {
	return pq.heap.IsEmpty()
}

// Size returns the number of elements in the priority queue
// Time complexity: O(1)
func (pq *PriorityQueue[T, P]) Size() int {
	return pq.heap.Size()
}

// Clear removes all elements from the priority queue
// Time complexity: O(1)
func (pq *PriorityQueue[T, P]) Clear() {
	pq.heap.Clear()
}

// Copy creates a new priority queue with the same elements
// Time complexity: O(n)
func (pq *PriorityQueue[T, P]) Copy() *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		heap:         pq.heap.Copy(),
		priorityType: pq.priorityType,
	}
}

// ToSlice returns a slice of all elements in priority order
// Time complexity: O(n log n)
func (pq *PriorityQueue[T, P]) ToSlice() []priorityItem[T, P] {
	slice := pq.heap.ToSlice()
	result := make([]priorityItem[T, P], len(slice))
	copy(result, slice)
	return result
}

// Values returns a slice of all values in priority order
// Time complexity: O(n log n)
func (pq *PriorityQueue[T, P]) Values() []T {
	items := pq.heap.ToSortedSlice()
	result := make([]T, len(items))
	for i, item := range items {
		result[i] = item.value
	}
	return result
}

// UpdatePriority updates the priority of an existing value
// This operation is O(n) as it needs to find the element
// Time complexity: O(n)
func (pq *PriorityQueue[T, P]) UpdatePriority(value T, newPriority P, equals func(T, T) bool) bool {
	items := pq.heap.ToSlice()

	for i, item := range items {
		if equals(item.value, value) {
			pq.heap.Clear()

			for j, existingItem := range items {
				if i != j {
					pq.heap.Insert(existingItem)
				} else {
					priorityItem := priorityItem[T, P]{value: value, priority: newPriority}
					pq.heap.Insert(priorityItem)
				}
			}

			return true
		}
	}

	return false
}

// Contains checks if a value exists in the priority queue
// Time complexity: O(n)
func (pq *PriorityQueue[T, P]) Contains(value T, equals func(T, T) bool) bool {
	slice := pq.heap.ToSlice()

	for _, item := range slice {
		if equals(item.value, value) {
			return true
		}
	}

	return false
}

// Remove removes a specific value from the priority queue
// Time complexity: O(n)
func (pq *PriorityQueue[T, P]) Remove(value T, equals func(T, T) bool) bool {
	items := pq.heap.ToSlice()

	for i, item := range items {
		if equals(item.value, value) {
			pq.heap.Clear()

			for j, existingItem := range items {
				if i != j {
					pq.heap.Insert(existingItem)
				}
			}

			return true
		}
	}

	return false
}

// PriorityType returns the priority type (MaxPriority or MinPriority)
// Time complexity: O(1)
func (pq *PriorityQueue[T, P]) PriorityType() PriorityType {
	return pq.priorityType
}

// DequeueAll removes and returns all elements in priority order
// Time complexity: O(n log n)
func (pq *PriorityQueue[T, P]) DequeueAll() []priorityItem[T, P] {
	items := make([]priorityItem[T, P], 0, pq.heap.Size())

	for !pq.IsEmpty() {
		item, _ := pq.heap.Extract()
		items = append(items, item)
	}

	return items
}

// Merge merges another priority queue into this one
// Time complexity: O(m log n) where m is the size of the other queue
func (pq *PriorityQueue[T, P]) Merge(other *PriorityQueue[T, P]) {
	otherItems := other.heap.ToSlice()

	for _, item := range otherItems {
		pq.heap.Insert(item)
	}
}

// ForEach applies a function to each element in priority order
// Time complexity: O(n log n)
func (pq *PriorityQueue[T, P]) ForEach(action func(T, P)) {
	slice := pq.heap.ToSlice()

	for _, item := range slice {
		action(item.value, item.priority)
	}
}

// Filter creates a new priority queue with elements that satisfy the predicate
// Time complexity: O(n log n)
func (pq *PriorityQueue[T, P]) Filter(predicate func(T, P) bool) *PriorityQueue[T, P] {
	newPQ := NewPriorityQueue[T, P](pq.priorityType)

	slice := pq.heap.ToSlice()

	for _, item := range slice {
		if predicate(item.value, item.priority) {
			newPQ.Enqueue(item.value, item.priority)
		}
	}

	return newPQ
}
