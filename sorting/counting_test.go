package sorting

import (
	"testing"
)

func TestCountingSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	CountingSort(slice)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortEmpty(t *testing.T) {
	slice := []int{}
	CountingSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestCountingSortSingle(t *testing.T) {
	slice := []int{5}
	CountingSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestCountingSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	CountingSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestCountingSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	CountingSort(slice)

	expected := []int{1, 2, 3, 4, 5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	CountingSort(slice)

	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1, 0}
	CountingSort(slice)

	expected := []int{-5, -2, -1, 0, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortAllNegative(t *testing.T) {
	slice := []int{-5, -3, -1, -4, -2}
	CountingSort(slice)

	expected := []int{-5, -4, -3, -2, -1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortLarge(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = 100 - i
	}
	CountingSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != i+1 {
			t.Errorf("Large slice should be sorted: expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestCountingSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := CountingSortCopy(original)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range expected {
		if sorted[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, sorted[i])
		}
	}

	originalExpected := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	for i, v := range originalExpected {
		if original[i] != v {
			t.Errorf("Original slice should not be modified: expected %d at index %d, got %d", v, i, original[i])
		}
	}
}

func TestCountingSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := CountingSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestCountingSortCopySingle(t *testing.T) {
	original := []int{5}
	sorted := CountingSortCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	if original[0] != 5 {
		t.Error("Original should not be modified")
	}
}

func TestCountingSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	CountingSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	CountingSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestCountingSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	CountingSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestCountingSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	CountingSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	CountingSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDescendingNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1, 0}
	CountingSortDescending(slice)

	expected := []int{8, 3, 0, -1, -2, -5}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := CountingSortDescendingCopy(original)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if sorted[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, sorted[i])
		}
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Errorf("Original slice should not be modified: expected %d at index %d, got %d", originalCopy[i], i, original[i])
		}
	}
}

func TestCountingSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := CountingSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestCountingSortDescendingCopySingle(t *testing.T) {
	original := []int{5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := CountingSortDescendingCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Errorf("Original should not be modified")
		}
	}
}

func TestCountingSortUint8(t *testing.T) {
	slice := []uint8{5, 2, 8, 1, 9, 3, 7, 4, 6}
	CountingSortUint8(slice)

	expected := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortUint8Empty(t *testing.T) {
	slice := []uint8{}
	CountingSortUint8(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestCountingSortUint8Single(t *testing.T) {
	slice := []uint8{5}
	CountingSortUint8(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestCountingSortUint8Duplicates(t *testing.T) {
	slice := []uint8{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	CountingSortUint8(slice)

	expected := []uint8{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortUint8FullRange(t *testing.T) {
	slice := []uint8{255, 0, 128, 64, 192}
	CountingSortUint8(slice)

	expected := []uint8{0, 64, 128, 192, 255}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortUint8Large(t *testing.T) {
	slice := make([]uint8, 255)
	for i := 0; i < 255; i++ {
		slice[i] = uint8(254 - i)
	}
	CountingSortUint8(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != uint8(i) {
			t.Errorf("Large uint8 slice should be sorted: expected %d at index %d, got %d", i, i, slice[i])
		}
	}
}

func TestCountingSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	CountingSort(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestCountingSortZeroAndNegative(t *testing.T) {
	slice := []int{0, -1, 1, -2, 2}
	CountingSort(slice)

	expected := []int{-2, -1, 0, 1, 2}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortDescendingAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	CountingSortDescending(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestCountingSortUint8AllSame(t *testing.T) {
	slice := []uint8{5, 5, 5, 5, 5}
	CountingSortUint8(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestCountingSortLargeRange(t *testing.T) {
	slice := []int{1000, -1000, 500, -500, 0, 999, -999}
	CountingSort(slice)

	expected := []int{-1000, -999, -500, 0, 500, 999, 1000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}

	if !IsSorted(slice) {
		t.Errorf("Slice should be sorted: %v", slice)
	}
}

func TestCountingSortDescendingLargeRange(t *testing.T) {
	slice := []int{1000, -1000, 500, -500, 0, 999, -999}
	CountingSortDescending(slice)

	expected := []int{1000, 999, 500, 0, -500, -999, -1000}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}

	if !IsSortedDesc(slice) {
		t.Errorf("Slice should be sorted in descending order: %v", slice)
	}
}

func TestCountingSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	CountingSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestCountingSortDescendingTwoElements(t *testing.T) {
	slice := []int{1, 9}
	CountingSortDescending(slice)

	if slice[0] != 9 || slice[1] != 1 {
		t.Errorf("Expected [9, 1], got %v", slice)
	}
}

func TestCountingSortUint8TwoElements(t *testing.T) {
	slice := []uint8{9, 1}
	CountingSortUint8(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestCountingSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	CountingSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	CountingSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestCountingSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	CountingSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestCountingSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	CountingSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestCountingSortUint8PreservesLength(t *testing.T) {
	slice := []uint8{5, 3, 1, 4, 2}
	originalLength := len(slice)
	CountingSortUint8(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func BenchmarkCountingSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = (j * 7) % 100
		}
		CountingSort(slice)
	}
}

func BenchmarkCountingSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = (j * 7) % 100
		}
		CountingSortDescending(slice)
	}
}

func BenchmarkCountingSortUint8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]uint8, 100)
		for j := range slice {
			slice[j] = uint8((j * 7) % 100)
		}
		CountingSortUint8(slice)
	}
}

func BenchmarkCountingSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = (j * 7) % 100
		}
		CountingSortCopy(slice)
	}
}

func BenchmarkCountingSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = (j * 7) % 100
		}
		CountingSortDescendingCopy(slice)
	}
}

func BenchmarkCountingSortUint8FullRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]uint8, 256)
		for j := range slice {
			slice[j] = uint8(255 - j)
		}
		CountingSortUint8(slice)
	}
}

func BenchmarkCountingSortLargeRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = j * 1000
		}
		CountingSort(slice)
	}
}
