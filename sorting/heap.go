package sorting

import (
	"golang.org/x/exp/constraints"
)

// HeapSort performs heap sort on a slice in ascending order
// Time complexity: O(n log n)
// Space complexity: O(1)
func HeapSort[T constraints.Ordered](slice []T) {
	n := len(slice)
	if n <= 1 {
		return
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(slice, n, i)
	}

	for i := n - 1; i > 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapify(slice, i, 0)
	}
}

// heapify maintains max-heap property
func heapify[T constraints.Ordered](slice []T, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && slice[left] > slice[largest] {
		largest = left
	}

	if right < n && slice[right] > slice[largest] {
		largest = right
	}

	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		heapify(slice, n, largest)
	}
}

// HeapSortDescending performs heap sort on a slice in descending order
// Time complexity: O(n log n)
// Space complexity: O(1)
func HeapSortDescending[T constraints.Ordered](slice []T) {
	n := len(slice)
	if n <= 1 {
		return
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapifyMin(slice, n, i)
	}

	for i := n - 1; i > 0; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		heapifyMin(slice, i, 0)
	}
}

// heapifyMin maintains min-heap property
func heapifyMin[T constraints.Ordered](slice []T, n, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && slice[left] < slice[smallest] {
		smallest = left
	}

	if right < n && slice[right] < slice[smallest] {
		smallest = right
	}

	if smallest != i {
		slice[i], slice[smallest] = slice[smallest], slice[i]
		heapifyMin(slice, n, smallest)
	}
}

// HeapSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n log n)
// Space complexity: O(n)
func HeapSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	HeapSort(result)
	return result
}

// HeapSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n log n)
// Space complexity: O(n)
func HeapSortDescendingCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	HeapSortDescending(result)
	return result
}
