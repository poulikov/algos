package searching

import (
	"golang.org/x/exp/constraints"
)

// ExponentialSearch performs exponential search on a sorted slice
// Useful for searching in infinite or very large arrays
// Returns index of target element, or -1 if not found
// Time complexity: O(log n)
func ExponentialSearch[T constraints.Ordered](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	if slice[0] == target {
		return 0
	}

	bound := 1
	n := len(slice)

	for bound < n && slice[bound] < target {
		bound *= 2
	}

	left := bound / 2
	right := bound + 1
	if right > n {
		right = n
	}

	result := BinarySearch(slice[left:right], target)
	if result == -1 {
		return -1
	}
	return result + left
}

// ExponentialSearchRecursive performs exponential search recursively
// Returns index of target element, or -1 if not found
// Time complexity: O(log n)
func ExponentialSearchRecursive[T constraints.Ordered](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	if slice[0] == target {
		return 0
	}

	bound := 1
	n := len(slice)

	for bound < n && slice[bound] < target {
		bound *= 2
	}

	left := bound / 2
	right := bound + 1
	if right > n {
		right = n
	}

	result := BinarySearchRecursive(slice[left:right], target)
	if result == -1 {
		return -1
	}
	return result + left
}

// ExponentialSearchLowerBound finds exponential search and returns lower bound
// Returns index of first element >= target
// Time complexity: O(log n)
func ExponentialSearchLowerBound[T constraints.Ordered](slice []T, target T) int {
	if len(slice) == 0 {
		return 0
	}

	bound := 1
	n := len(slice)

	for bound < n && slice[bound] < target {
		bound *= 2
	}

	left := bound / 2
	right := bound + 1
	if right > n {
		right = n
	}

	return BinarySearchLowerBound(slice[left:right], target) + left
}

// ExponentialSearchUpperBound finds exponential search and returns upper bound
// Returns index of first element > target
// Time complexity: O(log n)
func ExponentialSearchUpperBound[T constraints.Ordered](slice []T, target T) int {
	if len(slice) == 0 {
		return 0
	}

	if slice[0] > target {
		return 0
	}

	bound := 1
	n := len(slice)

	for bound < n && slice[bound] <= target {
		bound *= 2
	}

	left := bound / 2
	right := bound + 1
	if right > n {
		right = n
	}

	return BinarySearchUpperBound(slice[left:right], target) + left
}

// ExponentialSearchRange performs exponential search and returns range
// Returns [lowerBound, upperBound) indices
// Time complexity: O(log n)
func ExponentialSearchRange[T constraints.Ordered](slice []T, lower, upper T) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}

	lowerBound := ExponentialSearchLowerBound[T](slice, lower)
	upperBound := ExponentialSearchUpperBound[T](slice, upper)
	return lowerBound, upperBound
}
