package sorting

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
)

func IsSortedDesc[T constraints.Ordered](slice []T) bool {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestBubbleSortEmpty(t *testing.T) {
	slice := []int{}
	BubbleSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestBubbleSortSingle(t *testing.T) {
	slice := []int{5}
	BubbleSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestBubbleSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	BubbleSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestBubbleSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	BubbleSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestBubbleSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	BubbleSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestBubbleSortStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry", "date"}
	BubbleSort(slice)

	expected := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestBubbleSortFloats(t *testing.T) {
	slice := []float64{3.14, 1.41, 2.71, 0.577, 1.618}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestBubbleSortLarge(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = 100 - i
	}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Large slice should be sorted in ascending order")
		}
	}
}

func TestBubbleSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := BubbleSortCopy(original)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("Sorted slice should be in ascending order: %v", sorted)
		}
	}

	if IsSorted(original) {
		t.Error("Original slice should not be modified")
	}
}

func TestBubbleSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := BubbleSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestBubbleSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	BubbleSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestBubbleSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	BubbleSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestBubbleSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	BubbleSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestBubbleSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	BubbleSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	BubbleSortDescending(slice)

	if !IsSortedDesc(slice) {
		t.Errorf("Already descending sorted slice should remain sorted: %v", slice)
	}
}

func TestBubbleSortDescendingStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	BubbleSortDescending(slice)

	expected := []string{"cherry", "banana", "apple"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestBubbleSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := BubbleSortDescendingCopy(original)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] < sorted[i+1] {
			t.Errorf("Sorted slice should be in descending order: %v", sorted)
		}
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Errorf("Original slice should not be modified: expected %v, got %v", originalCopy, original)
		}
	}
}

func TestBubbleSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := BubbleSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestBubbleSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	BubbleSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	BubbleSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	BubbleSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestBubbleSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	BubbleSort(slice)

	expected := []int{-5, -2, -1, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortDescendingNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	BubbleSortDescending(slice)

	expected := []int{8, 3, -1, -2, -5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != 5 || slice[i] != slice[i+1] {
			t.Errorf("All elements should be 5, got %v", slice)
		}
	}
}

func TestBubbleSortMaxInt(t *testing.T) {
	slice := []int{2147483647, -2147483648, 0, 1, -1}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice with max/min int should be sorted: %v", slice)
		}
	}
}

func TestBubbleSortBestCase(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	BubbleSort(slice)

	if !IsSorted(slice) {
		t.Error("Already sorted slice should remain sorted (best case)")
	}

	if len(slice) != 10 {
		t.Error("Length should not change")
	}
}

func TestBubbleSortWorstCase(t *testing.T) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(slice)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Worst case should be sorted: expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortRandom(t *testing.T) {
	slice := make([]int, 50)
	for i := 0; i < 50; i++ {
		slice[i] = rand.Intn(100)
	}

	BubbleSort(slice)

	if !IsSorted(slice) {
		t.Error("Random slice should be sorted")
	}
}

func TestBubbleSortUint(t *testing.T) {
	slice := []uint{5, 2, 8, 1, 9}
	BubbleSort(slice)

	expected := []uint{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortDescendingUint(t *testing.T) {
	slice := []uint{1, 2, 5, 8, 9}
	BubbleSortDescending(slice)

	expected := []uint{9, 8, 5, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortInt8(t *testing.T) {
	slice := []int8{5, -2, 8, 1, -9}
	BubbleSort(slice)

	expected := []int8{-9, -2, 1, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortInt64(t *testing.T) {
	slice := []int64{5000000000, 2000000000, 8000000000, 1000000000}
	BubbleSort(slice)

	expected := []int64{1000000000, 2000000000, 5000000000, 8000000000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestBubbleSortFloat32(t *testing.T) {
	slice := []float32{3.14, 1.41, 2.71, 0.577}
	BubbleSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float32 slice should be sorted: %v", slice)
		}
	}
}

func TestBubbleSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	BubbleSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestBubbleSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	BubbleSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestBubbleSortCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := BubbleSortCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		BubbleSort(slice)
	}
}

func BenchmarkBubbleSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		BubbleSortDescending(slice)
	}
}

func BenchmarkBubbleSortBestCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = j
		}
		BubbleSort(slice)
	}
}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := 0; j < 100; j++ {
			slice[j] = 100 - j
		}
		BubbleSort(slice)
	}
}

func BenchmarkBubbleSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		BubbleSortCopy(slice)
	}
}

func BenchmarkBubbleSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		BubbleSortDescendingCopy(slice)
	}
}
