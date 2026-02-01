package sorting

import (
	"math/rand"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestSelectionSortEmpty(t *testing.T) {
	slice := []int{}
	SelectionSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestSelectionSortSingle(t *testing.T) {
	slice := []int{5}
	SelectionSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestSelectionSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	SelectionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestSelectionSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	SelectionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestSelectionSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	SelectionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestSelectionSortStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry", "date"}
	SelectionSort(slice)

	expected := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestSelectionSortFloats(t *testing.T) {
	slice := []float64{3.14, 1.41, 2.71, 0.577, 1.618}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestSelectionSortLarge(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = 100 - i
	}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Large slice should be sorted in ascending order")
		}
	}
}

func TestSelectionSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := SelectionSortCopy(original)

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

func TestSelectionSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := SelectionSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestSelectionSortCopySingle(t *testing.T) {
	original := []int{5}
	sorted := SelectionSortCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	if original[0] != 5 {
		t.Error("Original should not be modified")
	}
}

func TestSelectionSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	SelectionSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestSelectionSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	SelectionSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestSelectionSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	SelectionSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestSelectionSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	SelectionSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	SelectionSortDescending(slice)

	if !IsSortedDesc(slice) {
		t.Errorf("Already descending sorted slice should remain sorted: %v", slice)
	}
}

func TestSelectionSortDescendingStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	SelectionSortDescending(slice)

	expected := []string{"cherry", "banana", "apple"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestSelectionSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := SelectionSortDescendingCopy(original)

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

func TestSelectionSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := SelectionSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestSelectionSortDescendingCopySingle(t *testing.T) {
	original := []int{5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := SelectionSortDescendingCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Error("Original should not be modified")
		}
	}
}

func TestSelectionSortRange(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	SelectionSortRange(slice, 2, 6)

	for i := 2; i < 6; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Range [2,6] should be sorted: %v", slice[2:7])
		}
	}
}

func TestSelectionSortRangeFull(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	SelectionSortRange(slice, 0, len(slice)-1)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Full range should be sorted in ascending order")
		}
	}
}

func TestSelectionSortRangeSingle(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	SelectionSortRange(slice, 2, 2)

	for i := 0; i < len(slice); i++ {
		if i == 2 {
			if slice[i] != 8 {
				t.Errorf("Single element range should not change: expected 8, got %d", slice[i])
			}
		}
	}
}

func TestSelectionSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	SelectionSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	SelectionSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	SelectionSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestSelectionSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	SelectionSort(slice)

	expected := []int{-5, -2, -1, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortDescendingNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	SelectionSortDescending(slice)

	expected := []int{8, 3, -1, -2, -5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != 5 || slice[i] != slice[i+1] {
			t.Errorf("All elements should be 5, got %v", slice)
		}
	}
}

func TestSelectionSortMaxInt(t *testing.T) {
	slice := []int{2147483647, -2147483648, 0, 1, -1}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice with max/min int should be sorted: %v", slice)
		}
	}
}

func TestSelectionSortRandom(t *testing.T) {
	slice := make([]int, 50)
	for i := 0; i < 50; i++ {
		slice[i] = rand.Intn(100)
	}

	SelectionSort(slice)

	if !IsSorted(slice) {
		t.Error("Random slice should be sorted")
	}
}

func TestSelectionSortUint(t *testing.T) {
	slice := []uint{5, 2, 8, 1, 9}
	SelectionSort(slice)

	expected := []uint{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortDescendingUint(t *testing.T) {
	slice := []uint{1, 2, 5, 8, 9}
	SelectionSortDescending(slice)

	expected := []uint{9, 8, 5, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortInt8(t *testing.T) {
	slice := []int8{5, -2, 8, 1, -9}
	SelectionSort(slice)

	expected := []int8{-9, -2, 1, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortInt64(t *testing.T) {
	slice := []int64{5000000000, 2000000000, 8000000000, 1000000000}
	SelectionSort(slice)

	expected := []int64{1000000000, 2000000000, 5000000000, 8000000000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestSelectionSortFloat32(t *testing.T) {
	slice := []float32{3.14, 1.41, 2.71, 0.577}
	SelectionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float32 slice should be sorted: %v", slice)
		}
	}
}

func TestSelectionSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	SelectionSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestSelectionSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	SelectionSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestSelectionSortRangePreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2, 7, 6}
	originalLength := len(slice)
	SelectionSortRange(slice, 1, 4)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestSelectionSortCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := SelectionSortCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func TestSelectionSortDescendingCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := SelectionSortDescendingCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		SelectionSort(slice)
	}
}

func BenchmarkSelectionSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		SelectionSortDescending(slice)
	}
}

func BenchmarkSelectionSortRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		SelectionSortRange(slice, 20, 80)
	}
}

func BenchmarkSelectionSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		SelectionSortCopy(slice)
	}
}

func BenchmarkSelectionSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		SelectionSortDescendingCopy(slice)
	}
}
