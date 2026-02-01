package sorting

import (
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestInsertionSortEmpty(t *testing.T) {
	slice := []int{}
	InsertionSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestInsertionSortSingle(t *testing.T) {
	slice := []int{5}
	InsertionSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestInsertionSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	InsertionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestInsertionSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	InsertionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestInsertionSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	InsertionSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestInsertionSortStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry", "date"}
	InsertionSort(slice)

	expected := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestInsertionSortFloats(t *testing.T) {
	slice := []float64{3.14, 1.41, 2.71, 0.577, 1.618}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestInsertionSortLarge(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = 100 - i
	}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Large slice should be sorted in ascending order")
		}
	}
}

func TestInsertionSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := InsertionSortCopy(original)

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

func TestInsertionSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := InsertionSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestInsertionSortCopySingle(t *testing.T) {
	original := []int{5}
	sorted := InsertionSortCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	if original[0] != 5 {
		t.Error("Original should not be modified")
	}
}

func TestInsertionSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	InsertionSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestInsertionSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	InsertionSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestInsertionSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	InsertionSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestInsertionSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	InsertionSortDescending(slice)

	if !IsSortedDesc(slice) {
		t.Errorf("Already descending sorted slice should remain sorted: %v", slice)
	}
}

func TestInsertionSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	InsertionSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortDescendingStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	InsertionSortDescending(slice)

	expected := []string{"cherry", "banana", "apple"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestInsertionSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := InsertionSortDescendingCopy(original)

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

func TestInsertionSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := InsertionSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestInsertionSortDescendingCopySingle(t *testing.T) {
	original := []int{5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := InsertionSortDescendingCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Error("Original should not be modified")
		}
	}
}

func TestInsertionSortRange(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	InsertionSortRange(slice, 2, 6)

	expected := []int{5, 2, 1, 3, 7, 8, 9, 4, 6}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}

	for i := 2; i < 6; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Range [2,6] should be sorted: %v", slice[2:7])
		}
	}
}

func TestInsertionSortRangeFull(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	InsertionSortRange(slice, 0, len(slice)-1)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Full range should be sorted in ascending order")
		}
	}
}

func TestInsertionSortRangeSingle(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	InsertionSortRange(slice, 2, 2)

	for i := 0; i < len(slice); i++ {
		if i == 2 {
			if slice[i] != 8 {
				t.Errorf("Single element range should not change: expected 8, got %d", slice[i])
			}
		}
	}
}

func TestInsertionSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	InsertionSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	InsertionSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	InsertionSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestInsertionSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	InsertionSort(slice)

	expected := []int{-5, -2, -1, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortDescendingNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	InsertionSortDescending(slice)

	expected := []int{8, 3, -1, -2, -5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != 5 || slice[i] != slice[i+1] {
			t.Errorf("All elements should be 5, got %v", slice)
		}
	}
}

func TestInsertionSortMaxInt(t *testing.T) {
	slice := []int{2147483647, -2147483648, 0, 1, -1}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice with max/min int should be sorted: %v", slice)
		}
	}
}

func TestInsertionSortBestCase(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	InsertionSort(slice)

	if !IsSorted(slice) {
		t.Error("Already sorted slice should remain sorted (best case)")
	}

	if len(slice) != 10 {
		t.Error("Length should not change")
	}
}

func TestInsertionSortWorstCase(t *testing.T) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	InsertionSort(slice)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Worst case should be sorted: expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortRandom(t *testing.T) {
	slice := make([]int, 50)
	for i := 0; i < 50; i++ {
		slice[i] = rand.Intn(100)
	}

	InsertionSort(slice)

	if !IsSorted(slice) {
		t.Error("Random slice should be sorted")
	}
}

func TestInsertionSortUint(t *testing.T) {
	slice := []uint{5, 2, 8, 1, 9}
	InsertionSort(slice)

	expected := []uint{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortDescendingUint(t *testing.T) {
	slice := []uint{1, 2, 5, 8, 9}
	InsertionSortDescending(slice)

	expected := []uint{9, 8, 5, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortInt8(t *testing.T) {
	slice := []int8{5, -2, 8, 1, -9}
	InsertionSort(slice)

	expected := []int8{-9, -2, 1, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortInt64(t *testing.T) {
	slice := []int64{5000000000, 2000000000, 8000000000, 1000000000}
	InsertionSort(slice)

	expected := []int64{1000000000, 2000000000, 5000000000, 8000000000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestInsertionSortFloat32(t *testing.T) {
	slice := []float32{3.14, 1.41, 2.71, 0.577}
	InsertionSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float32 slice should be sorted: %v", slice)
		}
	}
}

func TestInsertionSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	InsertionSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestInsertionSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	InsertionSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestInsertionSortRangePreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2, 7, 6}
	originalLength := len(slice)
	InsertionSortRange(slice, 1, 4)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestInsertionSortCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := InsertionSortCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func TestInsertionSortDescendingCopyPreservesLength(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	sorted := InsertionSortDescendingCopy(original)

	if len(sorted) != len(original) {
		t.Errorf("Copy should have same length: expected %d, got %d", len(original), len(sorted))
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		InsertionSort(slice)
	}
}

func BenchmarkInsertionSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		InsertionSortDescending(slice)
	}
}

func BenchmarkInsertionSortBestCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = j
		}
		InsertionSort(slice)
	}
}

func BenchmarkInsertionSortWorstCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := 0; j < 100; j++ {
			slice[j] = 100 - j
		}
		InsertionSort(slice)
	}
}

func BenchmarkInsertionSortRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		InsertionSortRange(slice, 20, 80)
	}
}

func BenchmarkInsertionSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		InsertionSortCopy(slice)
	}
}

func BenchmarkInsertionSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		InsertionSortDescendingCopy(slice)
	}
}
