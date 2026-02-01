package searching

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// BinarySearch performs binary search on a sorted slice
// Returns index of target element, or -1 if not found
// Time complexity: O(log n)
func BinarySearch[T constraints.Ordered](slice []T, target T) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		}

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs binary search recursively
// Returns index of target element, or -1 if not found
// Time complexity: O(log n)
func BinarySearchRecursive[T constraints.Ordered](slice []T, target T) int {
	return binarySearchRecursiveHelper(slice, target, 0, len(slice)-1)
}

func binarySearchRecursiveHelper[T constraints.Ordered](slice []T, target T, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if slice[mid] == target {
		return mid
	}

	if slice[mid] < target {
		return binarySearchRecursiveHelper(slice, target, mid+1, right)
	}
	return binarySearchRecursiveHelper(slice, target, left, mid-1)
}

// BinarySearchInsertionPoint finds index where target should be inserted
// Returns index where target should be inserted to maintain sorted order
// Time complexity: O(log n)
func BinarySearchInsertionPoint[T constraints.Ordered](slice []T, target T) int {
	left := 0
	right := len(slice)

	for left < right {
		mid := left + (right-left)/2

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// BinarySearchFirstOccurrence finds first occurrence of target in sorted slice
// Returns index of first occurrence, or -1 if not found
// Time complexity: O(log n) average, O(n) worst case for duplicates
func BinarySearchFirstOccurrence[T constraints.Ordered](slice []T, target T) int {
	index := BinarySearch[T](slice, target)
	if index == -1 {
		return -1
	}

	// Move left to find first occurrence
	for index > 0 && slice[index-1] == target {
		index--
	}

	return index
}

// BinarySearchLastOccurrence finds last occurrence of target in sorted slice
// Returns index of last occurrence, or -1 if not found
// Time complexity: O(log n) average, O(n) worst case for duplicates
func BinarySearchLastOccurrence[T constraints.Ordered](slice []T, target T) int {
	index := BinarySearch[T](slice, target)
	if index == -1 {
		return -1
	}

	// Move right to find last occurrence
	for index < len(slice)-1 && slice[index+1] == target {
		index++
	}

	return index
}

// BinarySearchCountOccurrences counts how many times target appears in sorted slice
// Time complexity: O(log n) average, O(n) worst case for duplicates
func BinarySearchCountOccurrences[T constraints.Ordered](slice []T, target T) int {
	first := BinarySearchFirstOccurrence[T](slice, target)
	if first == -1 {
		return 0
	}

	last := BinarySearchLastOccurrence[T](slice, target)
	return last - first + 1
}

// BinarySearchLowerBound finds index of first element >= target
// Time complexity: O(log n)
func BinarySearchLowerBound[T constraints.Ordered](slice []T, target T) int {
	left := 0
	right := len(slice)

	for left < right {
		mid := left + (right-left)/2

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// BinarySearchUpperBound finds index of first element > target
// Time complexity: O(log n)
func BinarySearchUpperBound[T constraints.Ordered](slice []T, target T) int {
	left := 0
	right := len(slice)

	for left < right {
		mid := left + (right-left)/2

		if slice[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// BinarySearchRange finds range of elements in [lower, upper]
// Returns [lowerBound, upperBound) indices
// Time complexity: O(log n)
func BinarySearchRange[T constraints.Ordered](slice []T, lower, upper T) (int, int) {
	lowerBound := BinarySearchLowerBound[T](slice, lower)
	upperBound := BinarySearchUpperBound[T](slice, upper)
	return lowerBound, upperBound
}

// BinarySearchNearest finds element nearest to target in sorted slice
// Only works for numeric types (int, float, etc.)
// Returns index of nearest element
// Time complexity: O(log n)
func BinarySearchNearest[T constraints.Integer | constraints.Float](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := left + (right-left)/2

		if slice[mid] == target {
			return mid
		}

		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Find nearest without exact match
	nearestIdx := -1
	minDist := float64(0)

	if left < len(slice) {
		dist := abs(float64(slice[left]) - float64(target))
		nearestIdx = left
		minDist = dist
	}

	if right >= 0 {
		dist := abs(float64(slice[right]) - float64(target))
		if nearestIdx == -1 || dist < minDist {
			nearestIdx = right
			minDist = dist
		}
	}

	return nearestIdx
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// BinarySearchCustom performs binary search with custom comparison function
// Comparison function should return -1 if a < b, 0 if a == b, 1 if a > b
// Time complexity: O(log n)
func BinarySearchCustom[T any](slice []T, target T, compare func(T, T) int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := left + (right-left)/2
		cmp := compare(slice[mid], target)

		if cmp == 0 {
			return mid
		}

		if cmp < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// IsSorted checks if a slice is sorted in ascending order
// Time complexity: O(n)
func IsSorted[T constraints.Ordered](slice []T) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

// RotateLeft rotates slice left by k positions
// Time complexity: O(n)
func RotateLeft[T any](slice []T, k int) []T {
	if len(slice) == 0 {
		return slice
	}

	k = k % len(slice)
	if k < 0 {
		k += len(slice)
	}

	result := make([]T, len(slice))
	for i := 0; i < len(slice); i++ {
		result[i] = slice[(i+k)%len(slice)]
	}

	return result
}

// RotateRight rotates slice right by k positions
// Time complexity: O(n)
func RotateRight[T any](slice []T, k int) []T {
	if len(slice) == 0 {
		return slice
	}

	k = k % len(slice)
	if k < 0 {
		k += len(slice)
	}

	result := make([]T, len(slice))
	for i := 0; i < len(slice); i++ {
		result[(i+k)%len(slice)] = slice[i]
	}

	return result
}

// Reverse reverses a slice in place
// Time complexity: O(n)
func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// FindKthLargest finds k-th largest element in a sorted slice
// Time complexity: O(1)
func FindKthLargest[T constraints.Ordered](slice []T, k int) (T, error) {
	if k < 1 || k > len(slice) {
		var zero T
		return zero, fmt.Errorf("k out of range")
	}
	return slice[len(slice)-k], nil
}

// FindKthSmallest finds k-th smallest element in a sorted slice
// Time complexity: O(1)
func FindKthSmallest[T constraints.Ordered](slice []T, k int) (T, error) {
	if k < 1 || k > len(slice) {
		var zero T
		return zero, fmt.Errorf("k out of range")
	}
	return slice[k-1], nil
}

// Median finds median of a sorted slice
// Time complexity: O(1)
func Median[T constraints.Ordered](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, fmt.Errorf("slice is empty")
	}

	n := len(slice)
	if n%2 == 1 {
		return slice[n/2], nil
	}

	// For even length, return lower middle element
	return slice[n/2-1], nil
}

// Floor finds largest element <= target in a sorted slice
// Time complexity: O(log n)
func Floor[T constraints.Ordered](slice []T, target T) (T, bool) {
	index := BinarySearchLowerBound[T](slice, target)

	if index == len(slice) {
		return slice[index-1], true
	}

	if slice[index] == target {
		return target, true
	}

	if index == 0 {
		var zero T
		return zero, false
	}

	return slice[index-1], true
}

// Ceiling finds smallest element >= target in a sorted slice
// Time complexity: O(log n)
func Ceiling[T constraints.Ordered](slice []T, target T) (T, bool) {
	index := BinarySearchLowerBound[T](slice, target)

	if index == len(slice) {
		var zero T
		return zero, false
	}

	return slice[index], true
}
