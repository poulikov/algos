package searching

import (
	"testing"
)

func TestExponentialSearch(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 4},
		{6, 5},
		{7, 6},
		{8, 7},
		{9, 8},
		{10, 9},
		{0, -1},
		{11, -1},
	}

	for _, test := range tests {
		result := ExponentialSearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("ExponentialSearch(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestExponentialSearchEmpty(t *testing.T) {
	slice := []int{}
	result := ExponentialSearch(slice, 5)
	if result != -1 {
		t.Fatalf("ExponentialSearch on empty slice should return -1, got %d", result)
	}
}

func TestExponentialSearchSingle(t *testing.T) {
	slice := []int{5}
	result := ExponentialSearch(slice, 5)
	if result != 0 {
		t.Fatalf("ExponentialSearch(%d) = %d, expected 0", 5, result)
	}

	result = ExponentialSearch(slice, 3)
	if result != -1 {
		t.Fatalf("ExponentialSearch(%d) should return -1, got %d", 3, result)
	}
}

func TestExponentialSearchLarge(t *testing.T) {
	slice := make([]int, 1000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 2
	}

	for i := 0; i < len(slice); i++ {
		result := ExponentialSearch(slice, slice[i])
		if result != i {
			t.Fatalf("ExponentialSearch(%d) = %d, expected %d", slice[i], result, i)
		}
	}
}

func TestExponentialSearchRecursive(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{3, 2},
		{5, 4},
		{7, 6},
		{9, 8},
		{0, -1},
		{10, -1},
	}

	for _, test := range tests {
		result := ExponentialSearchRecursive(slice, test.target)
		if result != test.expected {
			t.Fatalf("ExponentialSearchRecursive(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestExponentialSearchLowerBound(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 4},
		{9, 4},
		{10, 5},
	}

	for _, test := range tests {
		result := ExponentialSearchLowerBound(slice, test.target)
		if result != test.expected {
			t.Fatalf("ExponentialSearchLowerBound(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestExponentialSearchUpperBound(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 4},
		{8, 4},
		{9, 5},
		{10, 5},
	}

	for _, test := range tests {
		result := ExponentialSearchUpperBound(slice, test.target)
		if result != test.expected {
			t.Fatalf("ExponentialSearchUpperBound(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestExponentialSearchRange(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13}

	lower, upper := ExponentialSearchRange(slice, 5, 11)

	if lower != 2 {
		t.Fatalf("Expected lower bound 2, got %d", lower)
	}
	if upper != 6 {
		t.Fatalf("Expected upper bound 6, got %d", upper)
	}
}

func TestExponentialSearchEmptyRange(t *testing.T) {
	slice := []int{}
	lower, upper := ExponentialSearchRange(slice, 1, 10)

	if lower != 0 || upper != 0 {
		t.Fatalf("Expected (0, 0), got (%d, %d)", lower, upper)
	}
}

func BenchmarkExponentialSearch(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExponentialSearch(slice, i*2)
	}
}

func BenchmarkExponentialSearchLarge(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ExponentialSearch(slice, i)
	}
}
