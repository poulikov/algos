package sorting

import (
	"golang.org/x/exp/constraints"
)

// SelectionSort performs selection sort on a slice in ascending order
// Time complexity: O(n^2) worst, average, and best case
// Space complexity: O(1)
func SelectionSort[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}
}

// SelectionSortDescending performs selection sort on a slice in descending order
// Time complexity: O(n^2) worst, average, and best case
// Space complexity: O(1)
func SelectionSortDescending[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if slice[j] > slice[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i {
			slice[i], slice[maxIdx] = slice[maxIdx], slice[i]
		}
	}
}

// SelectionSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n^2)
// Space complexity: O(n)
func SelectionSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	SelectionSort(result)
	return result
}

// SelectionSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n^2)
// Space complexity: O(n)
func SelectionSortDescendingCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	SelectionSortDescending(result)
	return result
}

// SelectionSortRange performs selection sort on a slice range [low, high] in ascending order
// Time complexity: O(n^2)
// Space complexity: O(1)
func SelectionSortRange[T constraints.Ordered](slice []T, low, high int) {
	for i := low; i < high; i++ {
		minIdx := i
		for j := i + 1; j <= high; j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}
}
