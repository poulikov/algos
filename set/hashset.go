package set

import (
	"sync/atomic"
)

// HashSet represents a hash set implementation
// It's built on top of HashMap for efficient operations
type HashSet[T comparable] struct {
	data *HashMap[T, struct{}]
}

// NewHashSet creates a new empty HashSet
// Time complexity: O(1)
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{
		data: NewHashMap[T, struct{}](),
	}
}

// NewHashSetWithCapacity creates a new empty HashSet with specified capacity
// Time complexity: O(1)
func NewHashSetWithCapacity[T comparable](capacity int) *HashSet[T] {
	return &HashSet[T]{
		data: NewHashMapWithCapacity[T, struct{}](capacity),
	}
}

// NewHashSetWithLoadFactor creates a new empty HashSet with specified capacity and load factor
// Time complexity: O(1)
func NewHashSetWithLoadFactor[T comparable](capacity int, loadFactor float64) *HashSet[T] {
	return &HashSet[T]{
		data: NewHashMapWithLoadFactor[T, struct{}](capacity, loadFactor),
	}
}

// NewHashSetFromSlice creates a new HashSet from a slice
// Time complexity: O(n)
func NewHashSetFromSlice[T comparable](elements []T) *HashSet[T] {
	hs := NewHashSet[T]()
	hs.AddAll(elements)
	return hs
}

// Add adds an element to the set
// Time complexity: O(1) average, O(n) worst case
func (hs *HashSet[T]) Add(element T) bool {
	return hs.data.PutIfAbsent(element, struct{}{})
}

// AddAll adds all elements from a slice to the set
// Time complexity: O(n)
func (hs *HashSet[T]) AddAll(elements []T) {
	for _, element := range elements {
		hs.Add(element)
	}
}

// Contains checks if an element exists in the set
// Time complexity: O(1) average, O(n) worst case
func (hs *HashSet[T]) Contains(element T) bool {
	return hs.data.ContainsKey(element)
}

// Remove removes an element from the set
// Time complexity: O(1) average, O(n) worst case
func (hs *HashSet[T]) Remove(element T) bool {
	return hs.data.Remove(element)
}

// RemoveAll removes all elements from a slice from the set
// Time complexity: O(n)
func (hs *HashSet[T]) RemoveAll(elements []T) {
	for _, element := range elements {
		hs.Remove(element)
	}
}

// Size returns the number of elements in the set
// Time complexity: O(1)
func (hs *HashSet[T]) Size() int {
	return hs.data.Size()
}

// IsEmpty returns true if the set is empty
// Time complexity: O(1)
func (hs *HashSet[T]) IsEmpty() bool {
	return hs.data.IsEmpty()
}

// Capacity returns the current capacity of the set
// Time complexity: O(1)
func (hs *HashSet[T]) Capacity() int {
	return hs.data.Capacity()
}

// Clear removes all elements from the set
// Time complexity: O(n)
func (hs *HashSet[T]) Clear() {
	hs.data.Clear()
}

// ToSlice returns a slice of all elements in the set
// Time complexity: O(n)
func (hs *HashSet[T]) ToSlice() []T {
	return hs.data.Keys()
}

// ContainsAll checks if the set contains all elements from another set
// Time complexity: O(n) where n is the size of the other set
func (hs *HashSet[T]) ContainsAll(other *HashSet[T]) bool {
	for _, element := range other.ToSlice() {
		if !hs.Contains(element) {
			return false
		}
	}
	return true
}

// ContainsAny checks if the set contains any element from another set
// Time complexity: O(n) where n is the size of the other set
func (hs *HashSet[T]) ContainsAny(other *HashSet[T]) bool {
	for _, element := range other.ToSlice() {
		if hs.Contains(element) {
			return true
		}
	}
	return false
}

// Union creates a new set that is the union of this set and another
// Time complexity: O(n + m) where n and m are sizes of the two sets
func (hs *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	newSet := hs.Copy()
	newSet.AddAll(other.ToSlice())
	return newSet
}

// Intersection creates a new set that is the intersection of this set and another
// Time complexity: O(n + m) where n and m are sizes of the two sets
func (hs *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for _, element := range hs.ToSlice() {
		if other.Contains(element) {
			result.Add(element)
		}
	}

	return result
}

// Difference creates a new set that is the difference of this set and another
// Time complexity: O(n + m) where n and m are sizes of the two sets
func (hs *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for _, element := range hs.ToSlice() {
		if !other.Contains(element) {
			result.Add(element)
		}
	}

	return result
}

// SymmetricDifference creates a new set that is the symmetric difference of this set and another
// Time complexity: O(n + m) where n and m are sizes of the two sets
func (hs *HashSet[T]) SymmetricDifference(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()

	for _, element := range hs.ToSlice() {
		if !other.Contains(element) {
			result.Add(element)
		}
	}

	for _, element := range other.ToSlice() {
		if !hs.Contains(element) {
			result.Add(element)
		}
	}

	return result
}

// Subset checks if this set is a subset of another set
// Time complexity: O(n) where n is the size of this set
func (hs *HashSet[T]) Subset(other *HashSet[T]) bool {
	for _, element := range hs.ToSlice() {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

// Superset checks if this set is a superset of another set
// Time complexity: O(n) where n is the size of the other set
func (hs *HashSet[T]) Superset(other *HashSet[T]) bool {
	return other.Subset(hs)
}

// Equals checks if this set equals another set
// Time complexity: O(n) where n is the size of this set
func (hs *HashSet[T]) Equals(other *HashSet[T]) bool {
	if hs.Size() != other.Size() {
		return false
	}
	return hs.Subset(other)
}

// Copy creates a shallow copy of the set
// Time complexity: O(n)
func (hs *HashSet[T]) Copy() *HashSet[T] {
	return &HashSet[T]{
		data: hs.data.Copy(),
	}
}

// ForEach applies a function to each element in the set
// Time complexity: O(n)
func (hs *HashSet[T]) ForEach(action func(T)) {
	hs.data.ForEach(func(key T, _ struct{}) {
		action(key)
	})
}

// Filter creates a new set with elements that satisfy the predicate
// Time complexity: O(n)
func (hs *HashSet[T]) Filter(predicate func(T) bool) *HashSet[T] {
	result := NewHashSet[T]()

	hs.ForEach(func(element T) {
		if predicate(element) {
			result.Add(element)
		}
	})

	return result
}

// Map creates a new set by applying a function to each element
// Time complexity: O(n)
func (hs *HashSet[T]) Map(mapper func(T) T) *HashSet[T] {
	result := NewHashSet[T]()

	hs.ForEach(func(element T) {
		result.Add(mapper(element))
	})

	return result
}

// Reduce reduces the set to a single value using the reduce function
// Time complexity: O(n)
func (hs *HashSet[T]) Reduce(initial T, reducer func(T, T) T) T {
	result := initial

	hs.ForEach(func(element T) {
		result = reducer(result, element)
	})

	return result
}

// Any checks if any element satisfies the predicate
// Time complexity: O(n)
func (hs *HashSet[T]) Any(predicate func(T) bool) bool {
	result := atomic.Bool{}

	hs.ForEach(func(element T) {
		if predicate(element) {
			result.Store(true)
		}
	})

	return result.Load()
}

// All checks if all elements satisfy the predicate
// Time complexity: O(n)
func (hs *HashSet[T]) All(predicate func(T) bool) bool {
	result := atomic.Bool{}
	result.Store(true)

	hs.ForEach(func(element T) {
		if !predicate(element) {
			result.Store(false)
		}
	})

	return result.Load()
}

// Find returns the first element that satisfies the predicate
// Time complexity: O(n)
func (hs *HashSet[T]) Find(predicate func(T) bool) (T, bool) {
	var zero T
	var found atomic.Bool
	var result atomic.Value

	hs.ForEach(func(element T) {
		if !found.Load() && predicate(element) {
			found.Store(true)
			result.Store(element)
		}
	})

	if found.Load() {
		return result.Load().(T), true
	}
	return zero, false
}

// RetainOnly retains only the elements in the given slice
// Time complexity: O(n) where n is the size of the set
func (hs *HashSet[T]) RetainOnly(elements []T) {
	hs.data = hs.data.Filter(func(key T, _ struct{}) bool {
		for _, element := range elements {
			if key == element {
				return true
			}
		}
		return false
	})
}

// RemoveIf removes elements that satisfy the predicate
// Time complexity: O(n)
func (hs *HashSet[T]) RemoveIf(predicate func(T) bool) {
	hs.data = hs.data.Filter(func(key T, _ struct{}) bool {
		return !predicate(key)
	})
}

// AddIfAbsent adds an element only if it's not already present
// Time complexity: O(1) average, O(n) worst case
func (hs *HashSet[T]) AddIfAbsent(element T) bool {
	return hs.Add(element)
}

// ComputeIfAbsent computes an element only if it's not already present
// Time complexity: O(1) average, O(n) worst case
func (hs *HashSet[T]) ComputeIfAbsent(element T, computeFunc func(T) T) T {
	if hs.Contains(element) {
		return element
	}
	result := computeFunc(element)
	hs.Add(result)
	return result
}
