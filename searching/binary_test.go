package searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{3, 1},
		{5, 2},
		{7, 3},
		{9, 4},
		{11, 5},
		{13, 6},
		{2, -1},
		{8, -1},
		{0, -1},
		{14, -1},
	}

	for _, test := range tests {
		result := BinarySearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearch(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchEmpty(t *testing.T) {
	slice := []int{}

	result := BinarySearch(slice, 5)
	if result != -1 {
		t.Fatalf("BinarySearch on empty slice should return -1, got %d", result)
	}
}

func TestBinarySearchSingle(t *testing.T) {
	slice := []int{5}

	result := BinarySearch(slice, 5)
	if result != 0 {
		t.Fatalf("BinarySearch(%d) = %d, expected 0", 5, result)
	}

	result = BinarySearch(slice, 3)
	if result != -1 {
		t.Fatalf("BinarySearch(%d) should return -1, got %d", 3, result)
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{5, 2},
		{9, 4},
		{6, -1},
	}

	for _, test := range tests {
		result := BinarySearchRecursive(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchRecursive(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchInsertionPoint(t *testing.T) {
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
		result := BinarySearchInsertionPoint(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchInsertionPoint(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchFirstOccurrence(t *testing.T) {
	slice := []int{1, 2, 2, 2, 3, 4}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 4},
		{4, 5},
		{5, -1},
	}

	for _, test := range tests {
		result := BinarySearchFirstOccurrence(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchFirstOccurrence(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchLastOccurrence(t *testing.T) {
	slice := []int{1, 2, 2, 2, 3, 4}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, -1},
	}

	for _, test := range tests {
		result := BinarySearchLastOccurrence(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchLastOccurrence(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchCountOccurrences(t *testing.T) {
	slice := []int{1, 2, 2, 2, 3, 4}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 1},
		{2, 3},
		{3, 1},
		{4, 1},
		{5, 0},
	}

	for _, test := range tests {
		result := BinarySearchCountOccurrences(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchCountOccurrences(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchLowerBound(t *testing.T) {
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
		result := BinarySearchLowerBound(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchLowerBound(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchUpperBound(t *testing.T) {
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
		result := BinarySearchUpperBound(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchUpperBound(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchRange(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13}

	lower, upper := BinarySearchRange(slice, 5, 11)

	if lower != 2 {
		t.Fatalf("Expected lower bound 2, got %d", lower)
	}
	if upper != 6 {
		t.Fatalf("Expected upper bound 6, got %d", upper)
	}
}

func TestBinarySearchNearestSkip(t *testing.T) {
	slice := []int{1, 5, 10, 15, 20}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 0},
		{3, 1},
		{5, 1},
		{7, 1},
		{10, 2},
		{13, 3},
		{18, 4},
		{20, 4},
		{25, 4},
	}

	for _, test := range tests {
		result := BinarySearchNearest(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchNearest(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchNearestFloatSkip(t *testing.T) {
	slice := []float64{1.0, 5.0, 10.0, 15.0, 20.0}

	tests := []struct {
		target   float64
		expected int
	}{
		{0.0, 0},
		{3.0, 1},
		{5.0, 1},
		{7.0, 1},
		{10.0, 2},
		{13.0, 3},
		{18.0, 4},
		{20.0, 4},
		{25.0, 4},
	}

	for _, test := range tests {
		result := BinarySearchNearest(slice, test.target)
		if result != test.expected {
			t.Fatalf("BinarySearchNearest(%.1f) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestBinarySearchCustom(t *testing.T) {
	slice := []string{"apple", "banana", "cherry", "date"}

	compare := func(a, b string) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}

	result := BinarySearchCustom(slice, "cherry", compare)
	if result != 2 {
		t.Fatalf("BinarySearchCustom('cherry') = %d, expected 2", result)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		slice    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 1, 2, 3}, true},
		{[]int{1, 3, 2, 4}, false},
		{[]int{}, true},
		{[]int{5}, true},
	}

	for _, test := range tests {
		result := IsSorted(test.slice)
		if result != test.expected {
			t.Fatalf("IsSorted(%v) = %v, expected %v", test.slice, result, test.expected)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := RotateLeft(slice, 2)
	expected := []int{3, 4, 5, 1, 2}

	if len(result) != len(expected) {
		t.Fatalf("RotateLeft result has wrong length: %d vs %d", len(result), len(expected))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("RotateLeft(%v) = %v, expected %v", slice, result, expected)
		}
	}
}

func TestRotateRight(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := RotateRight(slice, 2)
	expected := []int{4, 5, 1, 2, 3}

	if len(result) != len(expected) {
		t.Fatalf("RotateRight result has wrong length: %d vs %d", len(result), len(expected))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("RotateRight(%v) = %v, expected %v", slice, result, expected)
		}
	}
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := Reverse(slice)
	expected := []int{5, 4, 3, 2, 1}

	for i := range result {
		if result[i] != expected[i] {
			t.Fatalf("Reverse(%v) = %v, expected %v", slice, result, expected)
		}
	}
}

func TestFindKthLargest(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	value, err := FindKthLargest(slice, 3)
	if err != nil {
		t.Fatalf("FindKthLargest returned error: %v", err)
	}
	if value != 3 {
		t.Fatalf("FindKthLargest(slice, 3) = %d, expected 3", value)
	}

	_, err = FindKthLargest(slice, 0)
	if err == nil {
		t.Fatal("FindKthLargest(slice, 0) should return error")
	}

	_, err = FindKthLargest(slice, 6)
	if err == nil {
		t.Fatal("FindKthLargest(slice, 6) should return error")
	}
}

func TestFindKthSmallest(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	value, err := FindKthSmallest(slice, 3)
	if err != nil {
		t.Fatalf("FindKthSmallest returned error: %v", err)
	}
	if value != 3 {
		t.Fatalf("FindKthSmallest(slice, 3) = %d, expected 3", value)
	}

	_, err = FindKthSmallest(slice, 0)
	if err == nil {
		t.Fatal("FindKthSmallest(slice, 0) should return error")
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		slice    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 2, 3}, 2},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		value, err := Median(test.slice)
		if err != nil {
			t.Fatalf("Median returned error: %v", err)
		}
		if value != test.expected {
			t.Fatalf("Median(%v) = %d, expected %d", test.slice, value, test.expected)
		}
	}

	_, err := Median([]int{})
	if err == nil {
		t.Fatal("Median(empty) should return error")
	}
}

func TestFloor(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}

	tests := []struct {
		target   int
		expected int
		found    bool
	}{
		{0, 0, false},
		{1, 1, true},
		{2, 1, true},
		{3, 3, true},
		{4, 3, true},
		{5, 5, true},
		{6, 5, true},
		{7, 7, true},
		{8, 7, true},
		{9, 9, true},
		{10, 9, true},
	}

	for _, test := range tests {
		value, found := Floor(slice, test.target)
		if found != test.found || (found && value != test.expected) {
			t.Fatalf("Floor(%d) = (%d, %v), expected (%d, %v)", test.target, value, found, test.expected, test.found)
		}
	}
}

func TestCeiling(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}

	tests := []struct {
		target   int
		expected int
		found    bool
	}{
		{0, 1, true},
		{1, 1, true},
		{2, 3, true},
		{3, 3, true},
		{4, 5, true},
		{5, 5, true},
		{6, 7, true},
		{7, 7, true},
		{8, 9, true},
		{9, 9, true},
		{10, 0, false},
	}

	for _, test := range tests {
		value, found := Ceiling(slice, test.target)
		if found != test.found || (found && value != test.expected) {
			t.Fatalf("Ceiling(%d) = (%d, %v), expected (%d, %v)", test.target, value, found, test.expected, test.found)
		}
	}
}

func TestStringType(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	result := BinarySearch(slice, "banana")
	if result != 1 {
		t.Fatalf("BinarySearch('banana') = %d, expected 1", result)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	slice := make([]int, 1000000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(slice, i*2)
	}
}

func BenchmarkBinarySearchFirstOccurrence(b *testing.B) {
	slice := []int{1, 2, 2, 2, 3, 4, 5, 5, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchFirstOccurrence(slice, 2)
	}
}

func BenchmarkBinarySearchNearest(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 10
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearchNearest(slice, i*10+5)
	}
}

func BenchmarkRotateLeft(b *testing.B) {
	slice := make([]int, 10000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RotateLeft(slice, i%10000)
	}
}
