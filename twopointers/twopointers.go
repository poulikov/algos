package twopointers

import (
	"golang.org/x/exp/constraints"
)

// TwoPointers technique - using two pointers to solve problems efficiently
// Time complexity: O(n) for many problems, O(n^2) for some problems
func TwoPointers[T constraints.Ordered](data []T, target T) bool {
	left := 0
	right := len(data) - 1

	for left < right {
		sum := data[left] + data[right]
		if sum == target {
			return true
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}

	return false
}

// TwoPointersSorted works on sorted data with early termination
func TwoPointersSorted[T constraints.Ordered](data []T, target T) bool {
	left := 0
	right := len(data) - 1

	for left < right && data[left]+data[right] <= target {
		sum := data[left] + data[right]
		if sum == target {
			return true
		}
		left++
	}

	return false
}

// TwoPointersDescending works for data in descending order
func TwoPointersDescending[T constraints.Ordered](data []T, target T) bool {
	left := 0
	right := len(data) - 1

	for left < right && data[left] > data[right] {
		sum := data[left] + data[right]
		if sum == target {
			return true
		}
		left++
		if sum > target {
			right--
		}
	}

	return false
}

// TwoPointersAny finds any pair that sums to target
func TwoPointersAny[T constraints.Ordered](data []T, target T) (int, int, bool) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

// IsPalindrome checks if a slice is a palindrome using two pointers
// Time complexity: O(n)
func IsPalindrome[T comparable](data []T) bool {
	left := 0
	right := len(data) - 1

	for left < right {
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// HasDuplicates checks if there are duplicates using two pointers
// Works on sorted data - compares adjacent elements
// Time complexity: O(n)
func HasDuplicates[T comparable](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] == data[i-1] {
			return true
		}
	}
	return false
}

// RemoveDuplicates removes duplicates from sorted slice
// Returns slice with only unique elements
// Time complexity: O(n)
func RemoveDuplicates[T comparable](data []T) []T {
	if len(data) <= 1 {
		return data
	}

	writeIndex := 1
	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			data[writeIndex] = data[i]
			writeIndex++
		}
	}

	return data[:writeIndex]
}

// FindPair finds a pair that sums to target
// Returns indices of the two elements, or (-1, -1) if not found
func FindPair[T constraints.Ordered](data []T, target T) (int, int, bool) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

// FindPairDescending finds a pair in descending sorted data
func FindPairDescending[T constraints.Ordered](data []T, target T) (int, int, bool) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

// FindClosest finds the closest pair to target (but not exceeding target)
func FindClosest[T constraints.Ordered](data []T, target T) (int, int, bool) {
	closestI, closestJ := -1, -1
	var closestSum T

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			sum := data[i] + data[j]
			if sum <= target {
				if closestI == -1 || sum > closestSum {
					closestSum = sum
					closestI, closestJ = i, j
				}
			}
		}
	}

	return closestI, closestJ, closestI != -1
}

// FindClosestAny finds the closest pair in any order
func FindClosestAny[T constraints.Ordered](data []T, target T) (int, int, bool) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return i, j, true
			}
		}
	}

	return -1, -1, false
}

// FindTriple finds three elements that sum to target
func FindTriple[T constraints.Ordered](data []T, target T) (int, int, int, bool) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == target {
					return i, j, k, true
				}
			}
		}
	}

	return -1, -1, -1, false
}

// FindTripleClosest finds the closest triple to target
func FindTripleClosest[T constraints.Ordered](data []T, target T) (int, int, int, bool) {
	closestI, closestJ, closestK := -1, -1, -1
	var closestSum T

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				sum := data[i] + data[j] + data[k]
				if sum <= target {
					if closestI == -1 || sum > closestSum {
						closestSum = sum
						closestI, closestJ, closestK = i, j, k
					}
				}
			}
		}
	}

	return closestI, closestJ, closestK, closestI != -1
}

// FindKClosest finds k elements that sum to target
func FindKClosest[T constraints.Ordered](data []T, target T, k int) ([]int, bool) {
	indices := []int{}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for l := j + 1; l < len(data); l++ {
				for m := l + 1; m < len(data); m++ {
					indices = append(indices, i)
					if data[i]+data[j]+data[m] == target {
						return indices, true
					}
				}
			}
		}
	}

	return indices, false
}

// FindKClosestClosest finds the closest k-tuple to target
func FindKClosestClosest[T constraints.Ordered](data []T, target T, k int) ([]int, int, int, bool) {
	closestI, closestJ, closestK := -1, -1, -1
	var closestSum T

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for l := j + 1; l < len(data); l++ {
				for m := l + 1; m < len(data); m++ {
					sum := data[i] + data[j] + data[m]
					if sum <= target {
						if closestI == -1 || sum > closestSum {
							closestSum = sum
							closestI, closestJ, closestK = i, j, m
						}
					}
				}
			}
		}
	}

	return []int{closestI, closestJ, closestK}, closestI, closestJ, closestK != -1
}

// FindAllPairs finds all pairs that sum to target
func FindAllPairs[T constraints.Ordered](data []T, target T) [][]int {
	var pairs [][]int

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	return pairs
}

// FindAllTriples finds all triples that sum to target
func FindAllTriples[T constraints.Ordered](data []T, target T) []struct{ i, j, k int } {
	var triples []struct{ i, j, k int }

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == target {
					triples = append(triples, struct{ i, j, k int }{i, j, k})
				}
			}
		}
	}

	return triples
}

// Partition divides data around a pivot using Lomuto partition scheme
// Elements smaller than pivot go to the left, larger go to the right
// Time complexity: O(n)
func Partition[T constraints.Ordered](data []T, pivotIndex int) int {
	if pivotIndex < 0 || pivotIndex >= len(data) {
		return pivotIndex
	}

	data[pivotIndex], data[len(data)-1] = data[len(data)-1], data[pivotIndex]

	pivot := data[len(data)-1]
	i := 0

	for j := 0; j < len(data)-1; j++ {
		if data[j] <= pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	data[i], data[len(data)-1] = data[len(data)-1], data[i]
	return i
}

// QuickSort uses two pointers for partitioning
// Time complexity: O(n log n) average, O(n^2) worst case
// Space complexity: O(log n) for stack frames
func QuickSort[T constraints.Ordered](data []T) []T {
	if len(data) <= 1 {
		return data
	}

	quickSortHelper(data, 0, len(data)-1)
	return data
}

func quickSortHelper[T constraints.Ordered](data []T, low, high int) {
	if low < high {
		pivotIndex := Partition(data, high)
		quickSortHelper(data, low, pivotIndex-1)
		quickSortHelper(data, pivotIndex+1, high)
	}
}

// MergeSortWithTwoPointers uses two pointers for merging
// Time complexity: O(n log n)
// Space complexity: O(n) for auxiliary array
func MergeSortWithTwoPointers[T constraints.Ordered](data []T) []T {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSortWithTwoPointers(data[:mid])
	right := MergeSortWithTwoPointers(data[mid:])

	return mergeWithTwoPointers(left, right)
}

func mergeWithTwoPointers[T constraints.Ordered](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			k++
			i++
		} else {
			result = append(result, right[j])
			k++
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

	return result
}

// FindMid finds the middle element using two pointers
func FindMid[T constraints.Ordered](data []T) T {
	if len(data) == 0 {
		var zero T
		return zero
	}

	left := 0
	right := len(data) - 1

	for left < right {
		left++
		right--
		if left == right {
			return data[left]
		}
	}

	return data[left]
}

// FindDuplicate finds first duplicate element
// Time complexity: O(n)
// Space complexity: O(n)
func FindDuplicate[T comparable](data []T) (T, bool) {
	seen := make(map[T]bool)

	for _, v := range data {
		if seen[v] {
			return v, true
		}
		seen[v] = true
	}

	var zero T
	return zero, false
}
