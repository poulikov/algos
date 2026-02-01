package searching

import (
	"golang.org/x/exp/constraints"
)

// InterpolationSearch performs interpolation search on a sorted slice
// Works best for uniformly distributed numeric data
// Returns index of target element, or -1 if not found
// Time complexity: O(log log n) average, O(n) worst case
func InterpolationSearch[T constraints.Integer | constraints.Float](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	low := 0
	high := len(slice) - 1

	for low <= high && target >= slice[low] && target <= slice[high] {
		if low == high {
			if slice[low] == target {
				return low
			}
			return -1
		}

		pos := low + int(float64(high-low)*float64(target-slice[low])/float64(slice[high]-slice[low]))

		if pos < low || pos > high {
			break
		}

		if slice[pos] == target {
			return pos
		}

		if slice[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}

// InterpolationSearchRecursive performs interpolation search recursively
// Returns index of target element, or -1 if not found
// Time complexity: O(log log n) average
func InterpolationSearchRecursive[T constraints.Integer | constraints.Float](slice []T, target T) int {
	if len(slice) == 0 {
		return -1
	}

	return interpolationSearchHelper(slice, target, 0, len(slice)-1)
}

func interpolationSearchHelper[T constraints.Integer | constraints.Float](slice []T, target T, low, high int) int {
	if low > high {
		return -1
	}

	if target < slice[low] || target > slice[high] {
		return -1
	}

	if low == high {
		if slice[low] == target {
			return low
		}
		return -1
	}

	pos := low + int(float64(high-low)*float64(target-slice[low])/float64(slice[high]-slice[low]))

	if slice[pos] == target {
		return pos
	}

	if slice[pos] < target {
		return interpolationSearchHelper(slice, target, pos+1, high)
	}
	return interpolationSearchHelper(slice, target, low, pos-1)
}
