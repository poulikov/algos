package sorting

import (
	"time"

	"golang.org/x/exp/constraints"
)

// QuickSort performs quicksort on a slice in ascending order
// Time complexity: O(n log n) average, O(n^2) worst case
// Space complexity: O(log n) for stack frames
func QuickSort[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}
	quickSort(slice, 0, len(slice)-1)
}

// QuickSortDescending performs quicksort in descending order
// Time complexity: O(n log n) average, O(n^2) worst case
func QuickSortDescending[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}
	quickSortDescending(slice, 0, len(slice)-1)
}

func quickSort[T constraints.Ordered](slice []T, low, high int) {
	if low < high {
		p := partition(slice, low, high)
		quickSort(slice, low, p-1)
		quickSort(slice, p+1, high)
	}
}

func quickSortDescending[T constraints.Ordered](slice []T, low, high int) {
	if low < high {
		p := partitionDescending(slice, low, high)
		quickSortDescending(slice, low, p-1)
		quickSortDescending(slice, p+1, high)
	}
}

func partition[T constraints.Ordered](slice []T, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	i++
	slice[i], slice[high] = slice[high], slice[i]
	return i
}

func partitionDescending[T constraints.Ordered](slice []T, low, high int) int {
	pivot := slice[high]
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] >= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	i++
	slice[i], slice[high] = slice[high], slice[i]
	return i
}

// QuickSortWithTimeLimit runs quicksort with a time limit
// Returns true if sorting completed within time limit
func QuickSortWithTimeLimit[T constraints.Ordered](slice []T, duration time.Duration) bool {
	done := make(chan bool)
	timeout := time.After(duration)

	go func() {
		QuickSort(slice)
		done <- true
	}()

	select {
	case <-done:
		return true
	case <-timeout:
		return false
	}
}

// QuickSortWithTime measures how long quicksort takes
// Returns duration of sorting operation
func QuickSortWithTime[T constraints.Ordered](slice []T) time.Duration {
	start := time.Now()
	QuickSort(slice)
	return time.Since(start)
}

// MergeSortInPlace performs merge sort with O(1) space complexity in ascending order
// Time complexity: O(n)
func IsSorted[T constraints.Ordered](slice []T) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}
	return true
}

// QuickSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n log n)
// Space complexity: O(n)
func QuickSortCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	QuickSort(result)
	return result
}

// QuickSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n log n)
// Space complexity: O(n)
func QuickSortDescendingCopy[T constraints.Ordered](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	QuickSortDescending(result)
	return result
}

// QuickSortInPlace sorts a slice in place and returns it
// Time complexity: O(n log n)
func QuickSortInPlace[T constraints.Ordered](slice []T) []T {
	QuickSort(slice)
	return slice
}

// insertionSort implements insertion sort for small partitions
// Time complexity: O(n^2)
func insertionSort[T constraints.Ordered](slice []T, low, high int) {
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
