package sorting

import (
	"golang.org/x/exp/constraints"
)

// MergeSort performs merge sort on a slice in ascending order
// Time complexity: O(n log n) average, O(n log n) worst case
// Space complexity: O(n) for auxiliary array
func MergeSort[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}

	aux := make([]T, len(slice))
	mergeSortHelper(slice, aux, 0, len(slice)-1)
}

// MergeSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n log n)
// Space complexity: O(n)
func MergeSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	MergeSort(result)
	return result
}

// mergeSortHelper implements merge sort algorithm in-place using auxiliary array
func mergeSortHelper[T constraints.Ordered](slice, aux []T, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2
	mergeSortHelper(slice, aux, low, mid)
	mergeSortHelper(slice, aux, mid+1, high)
	merge(slice, aux, low, mid, high)
}

// merge merges two sorted halves of slice using auxiliary array
func merge[T constraints.Ordered](slice, aux []T, low, mid, high int) {
	for k := low; k <= high; k++ {
		aux[k] = slice[k]
	}

	i, j := low, mid+1
	for k := low; k <= high; k++ {
		if i > mid {
			slice[k] = aux[j]
			j++
		} else if j > high {
			slice[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			slice[k] = aux[i]
			i++
		} else {
			slice[k] = aux[j]
			j++
		}
	}
}

// MergeSortDescending performs merge sort in descending order
// Time complexity: O(n log n)
func MergeSortDescending[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}
	result := make([]T, len(slice))
	copy(result, slice)
	mergeSortDescendingHelper(result, 0, len(result)-1)
	copy(slice, result)
}

func mergeSortDescendingHelper[T constraints.Ordered](source []T, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSortDescendingHelper(source, start, mid)
	mergeSortDescendingHelper(source, mid+1, end)

	mergeDescending(source, start, mid, end)
}

func mergeDescending[T constraints.Ordered](source []T, start, mid, end int) {
	left := source[start : mid+1]
	right := source[mid+1 : end+1]

	result := make([]T, 0, end-start+1)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] >= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	for k := 0; k < len(result); k++ {
		source[start+k] = result[k]
	}
}

// MergeSortWithComparator performs merge sort with custom comparator
// Time complexity: O(n log n)
func MergeSortWithComparator[T any](slice []T, compare func(T, T) int) {
	if len(slice) <= 1 {
		return
	}
	mergeSortWithComparatorHelper(slice, compare)
}

func mergeSortWithComparatorHelper[T any](slice []T, compare func(T, T) int) {
	n := len(slice)
	if n <= 1 {
		return
	}

	mid := n / 2
	mergeSortWithComparatorHelper(slice[:mid], compare)
	mergeSortWithComparatorHelper(slice[mid:], compare)

	mergeWithComparator(slice, mid, compare)
}

func mergeWithComparator[T any](slice []T, mid int, compare func(T, T) int) {
	left := slice[:mid]
	right := slice[mid:]

	result := make([]T, len(slice))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if compare(left[i], right[j]) <= 0 {
			result[k] = left[i]
			i++
			k++
		} else {
			result[k] = right[j]
			j++
			k++
		}
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	for idx := 0; idx < len(result); idx++ {
		slice[idx] = result[idx]
	}
}
