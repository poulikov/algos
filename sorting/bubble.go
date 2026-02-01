package sorting

import (
	"golang.org/x/exp/constraints"
)

// BubbleSort performs bubble sort on a slice in ascending order
// Time complexity: O(n^2) worst and average case, O(n) best case (already sorted)
// Space complexity: O(1)
func BubbleSort[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortDescending performs bubble sort on a slice in descending order
// Time complexity: O(n^2) worst and average case, O(n) best case (already sorted)
// Space complexity: O(1)
func BubbleSortDescending[T constraints.Ordered](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// BubbleSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n^2)
// Space complexity: O(n)
func BubbleSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	BubbleSort(result)
	return result
}

// BubbleSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n^2)
// Space complexity: O(n)
func BubbleSortDescendingCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	BubbleSortDescending(result)
	return result
}
