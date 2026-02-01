package slidingwindow

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// SlidingWindowResult represents the result of max/min sliding window operations
type SlidingWindowResult[T constraints.Ordered] struct {
	Max    T
	MaxIdx int
	Min    T
	MinIdx int
}

// SlidingWindowSum represents a window with its sum and boundaries
type SlidingWindowSum[T any] struct {
	Sum      T
	StartIdx int
	EndIdx   int
	Window   []T
}

// MaxSlidingWindow finds the maximum value in any window of size k
// Returns the maximum value and its index
// Time complexity: O(n * k) where n is array length, k is window size
func MaxSlidingWindow[T constraints.Ordered](data []T, windowSize int) SlidingWindowResult[T] {
	if len(data) == 0 || windowSize <= 0 || windowSize > len(data) {
		return SlidingWindowResult[T]{}
	}

	max := data[0]
	maxIdx := 0

	for i := 0; i <= len(data)-windowSize; i++ {
		for j := i; j < i+windowSize; j++ {
			if data[j] > max {
				max = data[j]
				maxIdx = j
			}
		}
	}

	return SlidingWindowResult[T]{
		Max:    max,
		MaxIdx: maxIdx,
	}
}

// MinSlidingWindow finds the minimum value in any window of size k
// Returns the minimum value and its index
// Time complexity: O(n * k)
func MinSlidingWindow[T constraints.Ordered](data []T, windowSize int) SlidingWindowResult[T] {
	if len(data) == 0 || windowSize <= 0 || windowSize > len(data) {
		return SlidingWindowResult[T]{}
	}

	min := data[0]
	minIdx := 0

	for i := 0; i <= len(data)-windowSize; i++ {
		for j := i; j < i+windowSize; j++ {
			if data[j] < min {
				min = data[j]
				minIdx = j
			}
		}
	}

	return SlidingWindowResult[T]{
		Min:    min,
		MinIdx: minIdx,
	}
}

// SumSlidingWindow calculates sum of elements in window [start, end]
// Time complexity: O(n) where n is window size
func SumSlidingWindow[T constraints.Integer](data []T, start, end int) T {
	var sum T

	if start < 0 || end >= len(data) || start > end {
		return sum
	}

	for i := start; i <= end; i++ {
		sum += data[i]
	}

	return sum
}

// AverageSlidingWindow calculates average in window [start, end]
// Time complexity: O(n)
func AverageSlidingWindow[T constraints.Integer](data []T, start, end int) float64 {
	if start < 0 || end >= len(data) || start > end || len(data) == 0 {
		return 0
	}

	sum := SumSlidingWindow(data, start, end)
	count := end - start + 1

	return float64(sum) / float64(count)
}

// ContainsInWindow checks if a value exists in window [start, end]
// Time complexity: O(n) where n is window size
func ContainsInWindow[T comparable](data []T, value T, start, end int) bool {
	if start < 0 || end >= len(data) || start > end {
		return false
	}

	for i := start; i <= end; i++ {
		if data[i] == value {
			return true
		}
	}

	return false
}

// CountInWindow counts occurrences of a value in window [start, end]
// Time complexity: O(n)
func CountInWindow[T comparable](data []T, value T, start, end int) int {
	if start < 0 || end >= len(data) || start > end {
		return 0
	}

	count := 0
	for i := start; i <= end; i++ {
		if data[i] == value {
			count++
		}
	}

	return count
}

// FindInWindow finds indices of a value in window [start, end]
// Time complexity: O(n)
func FindInWindow[T comparable](data []T, value T, start, end int) []int {
	if start < 0 || end >= len(data) || start > end {
		return []int{}
	}

	indices := []int{}

	for i := start; i <= end; i++ {
		if data[i] == value {
			indices = append(indices, i)
		}
	}

	return indices
}

// FirstOccurrence finds first occurrence index of value in window [start, end]
// Time complexity: O(n)
func FirstOccurrence[T comparable](data []T, value T, start, end int) (int, bool) {
	indices := FindInWindow(data, value, start, end)

	if len(indices) == 0 {
		return 0, false
	}

	return indices[0], true
}

// AllOccurrences returns all indices of value in window [start, end]
// Time complexity: O(n)
func AllOccurrences[T comparable](data []T, value T, start, end int) []int {
	return FindInWindow(data, value, start, end)
}

// MaxSlidingWindowSum finds window with maximum sum
// Returns the window with maximum sum and its boundaries
// Time complexity: O(n * k) where n is array length, k is window size
func MaxSlidingWindowSum[T constraints.Integer](data []T, windowSize int) *SlidingWindowSum[T] {
	if len(data) == 0 || windowSize <= 0 || windowSize > len(data) {
		return nil
	}

	maxSum := SumSlidingWindow(data, 0, windowSize-1)
	maxIdx := 0

	for i := 1; i <= len(data)-windowSize; i++ {
		sum := SumSlidingWindow(data, i, i+windowSize-1)
		if sum > maxSum {
			maxSum = sum
			maxIdx = i
		}
	}

	return &SlidingWindowSum[T]{
		Sum:      maxSum,
		StartIdx: maxIdx,
		EndIdx:   maxIdx + windowSize - 1,
		Window:   data[maxIdx : maxIdx+windowSize],
	}
}

// MinSlidingWindowSum finds window with minimum sum
// Returns the window with minimum sum and its boundaries
// Time complexity: O(n * k)
func MinSlidingWindowSum[T constraints.Integer](data []T, windowSize int) *SlidingWindowSum[T] {
	if len(data) == 0 || windowSize <= 0 || windowSize > len(data) {
		return nil
	}

	minSum := SumSlidingWindow(data, 0, windowSize-1)
	minIdx := 0

	for i := 1; i <= len(data)-windowSize; i++ {
		sum := SumSlidingWindow(data, i, i+windowSize-1)
		if sum < minSum {
			minSum = sum
			minIdx = i
		}
	}

	return &SlidingWindowSum[T]{
		Sum:      minSum,
		StartIdx: minIdx,
		EndIdx:   minIdx + windowSize - 1,
		Window:   data[minIdx : minIdx+windowSize],
	}
}

// FixedSizeSlidingWindow finds the window with maximum sum
// Returns the window information
// Time complexity: O(n * k)
func FixedSizeSlidingWindow[T constraints.Integer](data []T, windowSize int) *SlidingWindowSum[T] {
	return MaxSlidingWindowSum(data, windowSize)
}

// AllMaxSlidingWindow returns all windows of given size with their sums
// Time complexity: O(n * k)
func AllMaxSlidingWindow[T constraints.Integer](data []T, windowSize int) []*SlidingWindowSum[T] {
	if len(data) == 0 || windowSize <= 0 || windowSize > len(data) {
		return []*SlidingWindowSum[T]{}
	}

	results := make([]*SlidingWindowSum[T], 0)

	for i := 0; i <= len(data)-windowSize; i++ {
		sum := SumSlidingWindow(data, i, i+windowSize-1)
		results = append(results, &SlidingWindowSum[T]{
			Sum:      sum,
			StartIdx: i,
			EndIdx:   i + windowSize - 1,
			Window:   data[i : i+windowSize],
		})
	}

	return results
}

// AllMinSlidingWindow returns all windows of given size with their sums
// Time complexity: O(n * k)
func AllMinSlidingWindow[T constraints.Integer](data []T, windowSize int) []*SlidingWindowSum[T] {
	return AllMaxSlidingWindow(data, windowSize)
}

// String returns a string representation of SlidingWindowResult
func (r *SlidingWindowResult[T]) String() string {
	return fmt.Sprintf("Max: %v (index: %d), Min: %v (index: %d)",
		r.Max, r.MaxIdx, r.Min, r.MinIdx)
}

// String returns a string representation of SlidingWindowSum
func (r *SlidingWindowSum[T]) String() string {
	if r == nil {
		return "nil"
	}
	return fmt.Sprintf("Sum: %v (idx: %d-%d)", r.Sum, r.StartIdx, r.EndIdx)
}
