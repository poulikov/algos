package sorting

import (
	"math/rand"
	"testing"
)

func TestHeapSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestHeapSortEmpty(t *testing.T) {
	slice := []int{}
	HeapSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestHeapSortSingle(t *testing.T) {
	slice := []int{5}
	HeapSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestHeapSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	HeapSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestHeapSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	HeapSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestHeapSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	HeapSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestHeapSortStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry", "date"}
	HeapSort(slice)

	expected := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestHeapSortFloats(t *testing.T) {
	slice := []float64{3.14, 1.41, 2.71, 0.577, 1.618}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestHeapSortLarge(t *testing.T) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = 1000 - i
	}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Large slice should be sorted in ascending order")
		}
	}
}

func TestHeapSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := HeapSortCopy(original)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("Sorted slice should be in ascending order: %v", sorted)
		}
	}

	originalExpected := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	for i, v := range originalExpected {
		if original[i] != v {
			t.Errorf("Original slice should not be modified: expected %d at index %d, got %d", v, i, original[i])
		}
	}
}

func TestHeapSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := HeapSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestHeapSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	HeapSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestHeapSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	HeapSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestHeapSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	HeapSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestHeapSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	HeapSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	HeapSortDescending(slice)

	if !IsSortedDesc(slice) {
		t.Errorf("Already descending sorted slice should remain sorted: %v", slice)
	}
}

func TestHeapSortDescendingStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	HeapSortDescending(slice)

	expected := []string{"cherry", "banana", "apple"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestHeapSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := HeapSortDescendingCopy(original)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] < sorted[i+1] {
			t.Errorf("Sorted slice should be in descending order: %v", sorted)
		}
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Errorf("Original slice should not be modified: expected %d at index %d, got %d", originalCopy[i], i, original[i])
		}
	}
}

func TestHeapSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := HeapSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestHeapSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	HeapSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	HeapSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	HeapSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestHeapSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	HeapSort(slice)

	expected := []int{-5, -2, -1, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortDescendingNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	HeapSortDescending(slice)

	expected := []int{8, 3, -1, -2, -5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != 5 || slice[i] != slice[i+1] {
			t.Errorf("All elements should be 5, got %v", slice)
		}
	}
}

func TestHeapSortMaxInt(t *testing.T) {
	slice := []int{2147483647, -2147483648, 0, 1, -1}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice with max/min int should be sorted: %v", slice)
		}
	}
}

func TestHeapSortRandom(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = rand.Intn(1000)
	}

	HeapSort(slice)

	if !IsSorted(slice) {
		t.Error("Random slice should be sorted")
	}
}

func TestHeapSortUint(t *testing.T) {
	slice := []uint{5, 2, 8, 1, 9}
	HeapSort(slice)

	expected := []uint{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortDescendingUint(t *testing.T) {
	slice := []uint{1, 2, 5, 8, 9}
	HeapSortDescending(slice)

	expected := []uint{9, 8, 5, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortInt8(t *testing.T) {
	slice := []int8{5, -2, 8, 1, -9}
	HeapSort(slice)

	expected := []int8{-9, -2, 1, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortInt64(t *testing.T) {
	slice := []int64{5000000000, 2000000000, 8000000000, 1000000000}
	HeapSort(slice)

	expected := []int64{1000000000, 2000000000, 5000000000, 8000000000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestHeapSortFloat32(t *testing.T) {
	slice := []float32{3.14, 1.41, 2.71, 0.577}
	HeapSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float32 slice should be sorted: %v", slice)
		}
	}
}

func TestHeapSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	HeapSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestHeapSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	HeapSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestHeapSortCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := HeapSortCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		HeapSort(slice)
	}
}

func BenchmarkHeapSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		HeapSortDescending(slice)
	}
}

func BenchmarkHeapSortBestCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = j
		}
		HeapSort(slice)
	}
}

func BenchmarkHeapSortWorstCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := 0; j < 1000; j++ {
			slice[j] = 1000 - j
		}
		HeapSort(slice)
	}
}

func BenchmarkHeapSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		HeapSortCopy(slice)
	}
}

func BenchmarkHeapSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		HeapSortDescendingCopy(slice)
	}
}
