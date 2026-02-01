package trees

import (
	"fmt"
)

// UnionFind represents a disjoint set union data structure with path compression and union by rank
type UnionFind[T comparable] struct {
	parent map[T]T
	rank   map[T]int
	size   map[T]int
	count  int
}

// NewUnionFind creates a new empty UnionFind data structure
// Time complexity: O(1)
func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
		size:   make(map[T]int),
	}
}

// NewUnionFindFromSlice creates a new UnionFind with elements initialized as separate sets
// Time complexity: O(n)
func NewUnionFindFromSlice[T comparable](elements []T) *UnionFind[T] {
	uf := NewUnionFind[T]()
	for _, element := range elements {
		uf.MakeSet(element)
	}
	return uf
}

// MakeSet creates a new set containing only the given element
// Time complexity: O(1)
func (uf *UnionFind[T]) MakeSet(element T) {
	if _, exists := uf.parent[element]; exists {
		return
	}
	uf.parent[element] = element
	uf.rank[element] = 0
	uf.size[element] = 1
	uf.count++
}

// Find finds the representative (root) of the set containing the given element
// Uses path compression for optimal performance
// Time complexity: O(α(n)) where α is the inverse Ackermann function (practically constant)
func (uf *UnionFind[T]) Find(element T) (T, error) {
	if _, exists := uf.parent[element]; !exists {
		var zero T
		return zero, fmt.Errorf("element not found in UnionFind")
	}

	root := element
	for uf.parent[root] != root {
		root = uf.parent[root]
	}

	uf.parent[element] = root
	return root, nil
}

// FindSafe finds the representative without modifying the structure
// Time complexity: O(log n) worst case
func (uf *UnionFind[T]) FindSafe(element T) (T, error) {
	if _, exists := uf.parent[element]; !exists {
		var zero T
		return zero, fmt.Errorf("element not found in UnionFind")
	}

	root := element
	for uf.parent[root] != root {
		root = uf.parent[root]
	}

	return root, nil
}

// Union merges the sets containing element1 and element2
// Uses union by rank for optimal performance
// Time complexity: O(α(n)) where α is the inverse Ackermann function
func (uf *UnionFind[T]) Union(element1, element2 T) error {
	root1, err := uf.Find(element1)
	if err != nil {
		return err
	}

	root2, err := uf.Find(element2)
	if err != nil {
		return err
	}

	if root1 == root2 {
		return nil
	}

	if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
		uf.size[root2] += uf.size[root1]
	} else if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += uf.size[root2]
	} else {
		uf.parent[root2] = root1
		uf.rank[root1]++
		uf.size[root1] += uf.size[root2]
	}

	uf.count--
	return nil
}

// UnionSafe merges sets without path compression (for debugging/testing)
// Time complexity: O(log n) worst case
func (uf *UnionFind[T]) UnionSafe(element1, element2 T) error {
	root1, err := uf.FindSafe(element1)
	if err != nil {
		return err
	}

	root2, err := uf.FindSafe(element2)
	if err != nil {
		return err
	}

	if root1 == root2 {
		return nil
	}

	if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
		uf.size[root2] += uf.size[root1]
	} else if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += uf.size[root2]
	} else {
		uf.parent[root2] = root1
		uf.rank[root1]++
		uf.size[root1] += uf.size[root2]
	}

	uf.count--
	return nil
}

// Connected checks if two elements are in the same set
// Time complexity: O(α(n))
func (uf *UnionFind[T]) Connected(element1, element2 T) (bool, error) {
	root1, err := uf.Find(element1)
	if err != nil {
		return false, err
	}

	root2, err := uf.Find(element2)
	if err != nil {
		return false, err
	}

	return root1 == root2, nil
}

// Size returns the size of the set containing the given element
// Time complexity: O(α(n))
func (uf *UnionFind[T]) Size(element T) (int, error) {
	root, err := uf.Find(element)
	if err != nil {
		return 0, err
	}
	return uf.size[root], nil
}

// Count returns the number of disjoint sets
// Time complexity: O(1)
func (uf *UnionFind[T]) Count() int {
	return uf.count
}

// Contains checks if an element exists in the UnionFind
// Time complexity: O(1)
func (uf *UnionFind[T]) Contains(element T) bool {
	_, exists := uf.parent[element]
	return exists
}

// Add is an alias for MakeSet
// Time complexity: O(1)
func (uf *UnionFind[T]) Add(element T) {
	uf.MakeSet(element)
}

// Clear removes all elements from the UnionFind
// Time complexity: O(1)
func (uf *UnionFind[T]) Clear() {
	uf.parent = make(map[T]T)
	uf.rank = make(map[T]int)
	uf.size = make(map[T]int)
	uf.count = 0
}

// Copy creates a new UnionFind with the same elements and sets
// Time complexity: O(n)
func (uf *UnionFind[T]) Copy() *UnionFind[T] {
	newUF := NewUnionFind[T]()

	for k, v := range uf.parent {
		newUF.parent[k] = v
	}

	for k, v := range uf.rank {
		newUF.rank[k] = v
	}

	for k, v := range uf.size {
		newUF.size[k] = v
	}

	newUF.count = uf.count

	return newUF
}

// ToMap returns a map of elements to their representative roots
// Time complexity: O(n * α(n))
func (uf *UnionFind[T]) ToMap() map[T]T {
	result := make(map[T]T)
	for element := range uf.parent {
		root, _ := uf.Find(element)
		result[element] = root
	}
	return result
}

// ToGroups returns a map grouping elements by their representative roots
// Time complexity: O(n * α(n))
func (uf *UnionFind[T]) ToGroups() map[T][]T {
	groups := make(map[T][]T)

	for element := range uf.parent {
		root, _ := uf.Find(element)
		groups[root] = append(groups[root], element)
	}

	return groups
}

// Elements returns all elements in the UnionFind
// Time complexity: O(n)
func (uf *UnionFind[T]) Elements() []T {
	elements := make([]T, 0, len(uf.parent))
	for element := range uf.parent {
		elements = append(elements, element)
	}
	return elements
}

// String returns a string representation of the UnionFind
func (uf *UnionFind[T]) String() string {
	groups := uf.ToGroups()
	return fmt.Sprintf("%v", groups)
}

// GetParent returns the parent of an element (for debugging/testing)
// Time complexity: O(1)
func (uf *UnionFind[T]) GetParent(element T) (T, error) {
	if parent, exists := uf.parent[element]; exists {
		return parent, nil
	}
	var zero T
	return zero, fmt.Errorf("element not found")
}

// GetRank returns the rank of an element (for debugging/testing)
// Time complexity: O(1)
func (uf *UnionFind[T]) GetRank(element T) (int, error) {
	if rank, exists := uf.rank[element]; exists {
		return rank, nil
	}
	return 0, fmt.Errorf("element not found")
}

// SizeOf returns the total number of elements in all sets
// Time complexity: O(1)
func (uf *UnionFind[T]) SizeOf() int {
	return len(uf.parent)
}

// IsEmpty returns true if the UnionFind contains no elements
// Time complexity: O(1)
func (uf *UnionFind[T]) IsEmpty() bool {
	return len(uf.parent) == 0
}
