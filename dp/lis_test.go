package dp

import (
	"reflect"
	"testing"
)

func TestLISLength(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		result := LISLength(test.nums)
		if result != test.expected {
			t.Errorf("LISLength(%v) = %d, expected %d", test.nums, result, test.expected)
		}
	}
}

func TestLISLengthOptimized(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		result := LISLengthOptimized(test.nums)
		if result != test.expected {
			t.Errorf("LISLengthOptimized(%v) = %d, expected %d", test.nums, result, test.expected)
		}
	}
}

func TestLIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, []int{2, 5, 7, 101}},
		{[]int{0, 1, 0, 3, 2, 3}, []int{0, 1, 2, 3}},
		{[]int{7, 7, 7, 7, 7, 7, 7}, []int{7}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1}},
	}

	for _, test := range tests {
		result := LIS(test.nums)
		if len(result) != len(test.expected) {
			t.Errorf("LIS(%v) = %v, length %d expected %d", test.nums, result, len(result), len(test.expected))
			continue
		}
		if !reflect.DeepEqual(result, test.expected) && !isLIS(result, test.nums) {
			t.Errorf("LIS(%v) = %v, expected %v or any valid LIS", test.nums, result, test.expected)
		}
	}
}

func TestLISOptimized(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, []int{2, 3, 7, 101}},
		{[]int{0, 1, 0, 3, 2, 3}, []int{0, 1, 2, 3}},
		{[]int{7, 7, 7, 7, 7, 7, 7}, []int{7}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1}},
	}

	for _, test := range tests {
		result := LISOptimized(test.nums)
		if len(result) != len(test.expected) {
			t.Errorf("LISOptimized(%v) = %v, length %d expected %d", test.nums, result, len(result), len(test.expected))
			continue
		}
		if !reflect.DeepEqual(result, test.expected) && !isLIS(result, test.nums) {
			t.Errorf("LISOptimized(%v) = %v, expected %v or any valid LIS", test.nums, result, test.expected)
		}
	}
}

func TestLNDSLength(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 7},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 1},
	}

	for _, test := range tests {
		result := LNDSLength(test.nums)
		if result != test.expected {
			t.Errorf("LNDSLength(%v) = %d, expected %d", test.nums, result, test.expected)
		}
	}
}

func isLIS(subseq, nums []int) bool {
	if len(subseq) == 0 {
		return len(nums) == 0
	}

	for i := 1; i < len(subseq); i++ {
		if subseq[i] <= subseq[i-1] {
			return false
		}
	}

	return true
}
