package sorting

import (
	"golang.org/x/exp/constraints"
)

// InsertionSort performs insertion sort on a slice in ascending order
// Time complexity: O(n^2) worst and average case, O(n) best case (already sorted)
// Space complexity: O(1)
func InsertionSort[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 1; i < n; i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

// InsertionSortDescending performs insertion sort on a slice in descending order
// Time complexity: O(n^2) worst and average case, O(n) best case (already sorted)
// Space complexity: O(1)
func InsertionSortDescending[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 1; i < n; i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] < key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}

// InsertionSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n^2)
// Space complexity: O(n)
func InsertionSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	InsertionSort(result)
	return result
}

// InsertionSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n^2)
// Space complexity: O(n)
func InsertionSortDescendingCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	InsertionSortDescending(result)
	return result
}

// InsertionSortRange performs insertion sort on a slice range [low, high] in ascending order
// Time complexity: O(n^2)
// Space complexity: O(1)
func InsertionSortRange[T constraints.Ordered](slice []T, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := slice[i]
		j := i - 1
		for j >= low && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
}
