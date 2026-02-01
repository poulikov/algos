package sorting

import (
	"math/rand"
	"testing"
)

func TestRadixSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	RadixSort(slice)

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestRadixSortEmpty(t *testing.T) {
	slice := []int{}
	RadixSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestRadixSortSingle(t *testing.T) {
	slice := []int{5}
	RadixSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestRadixSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	RadixSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestRadixSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	RadixSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestRadixSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	RadixSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestRadixSortLarge(t *testing.T) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = 1000 - i
	}
	RadixSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != i+1 {
			t.Errorf("Large slice should be sorted: expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestRadixSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := RadixSortCopy(original)

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

func TestRadixSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := RadixSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestRadixSortCopySingle(t *testing.T) {
	original := []int{5}
	sorted := RadixSortCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	if original[0] != 5 {
		t.Error("Original should not be modified")
	}
}

func TestRadixSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	RadixSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestRadixSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	RadixSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestRadixSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	RadixSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestRadixSortDescendingSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	RadixSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestRadixSortDescendingCopy(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := RadixSortDescendingCopy(original)

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

func TestRadixSortDescendingCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := RadixSortDescendingCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestRadixSortDescendingCopySingle(t *testing.T) {
	original := []int{5}
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)

	sorted := RadixSortDescendingCopy(original)

	if len(sorted) != 1 || sorted[0] != 5 {
		t.Errorf("Expected [5], got %v", sorted)
	}

	for i := range original {
		if original[i] != originalCopy[i] {
			t.Error("Original should not be modified")
		}
	}
}

func TestRadixSortNegativePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSort should panic on negative numbers")
		}
	}()

	slice := []int{-1, 2, 3}
	RadixSort(slice)
}

func TestRadixSortWithBase(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	RadixSortWithBase(slice, 10)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != i+1 {
			t.Errorf("Expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestRadixSortWithBase2(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	RadixSortWithBase(slice, 2)

	for i := 0; i < len(slice); i++ {
		if slice[i] != i {
			t.Errorf("Expected %d at index %d, got %d", i, i, slice[i])
		}
	}
}

func TestRadixSortWithBase16(t *testing.T) {
	slice := []int{255, 16, 0, 128, 64}
	RadixSortWithBase(slice, 16)

	expected := []int{0, 16, 64, 128, 255}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestRadixSortWithBase32(t *testing.T) {
	slice := []int{31, 0, 1, 30, 15}
	RadixSortWithBase(slice, 32)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted with base 32: %v", slice)
		}
	}
}

func TestRadixSortWithBasePanicInvalidLow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSortWithBase should panic on base < 2")
		}
	}()

	slice := []int{1, 2, 3}
	RadixSortWithBase(slice, 1)
}

func TestRadixSortWithBasePanicInvalidHigh(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSortWithBase should panic on base > 36")
		}
	}()

	slice := []int{1, 2, 3}
	RadixSortWithBase(slice, 37)
}

func TestRadixSortWithBasePanicNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("RadixSortWithBase should panic on negative numbers even with custom base")
		}
	}()

	slice := []int{-1, 2, 3}
	RadixSortWithBase(slice, 10)
}

func TestRadixSortUint32(t *testing.T) {
	slice := []uint32{5, 2, 8, 1, 9, 3, 7, 4, 6}
	RadixSortUint32(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != uint32(i+1) {
			t.Errorf("Expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestRadixSortUint32Empty(t *testing.T) {
	slice := []uint32{}
	RadixSortUint32(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestRadixSortUint32Single(t *testing.T) {
	slice := []uint32{5}
	RadixSortUint32(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestRadixSortUint32Large(t *testing.T) {
	slice := make([]uint32, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = uint32(1000 - i)
	}
	RadixSortUint32(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != uint32(i+1) {
			t.Errorf("Large uint32 slice should be sorted: expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestRadixSortUint32MaxValue(t *testing.T) {
	slice := []uint32{4294967295, 0, 2147483647}
	RadixSortUint32(slice)

	if slice[0] != 0 || slice[1] != 2147483647 || slice[2] != 4294967295 {
		t.Errorf("Uint32 slice with max value should be sorted: %v", slice)
	}
}

func TestRadixSortUint32Duplicates(t *testing.T) {
	slice := []uint32{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	RadixSortUint32(slice)

	if !IsSorted(slice) {
		t.Errorf("Uint32 slice with duplicates should be sorted: %v", slice)
	}
}

func TestRadixSortZeros(t *testing.T) {
	slice := []int{0, 0, 0, 0, 0}
	RadixSort(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			t.Errorf("All elements should be 0: %v", slice)
		}
	}
}

func TestRadixSortLargeNumbers(t *testing.T) {
	slice := []int{1000000, 999999, 1000001, 999998}
	RadixSort(slice)

	expected := []int{999998, 999999, 1000000, 1000001}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestRadixSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	RadixSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestRadixSortDescendingTwoElements(t *testing.T) {
	slice := []int{1, 9}
	RadixSortDescending(slice)

	if slice[0] != 9 || slice[1] != 1 {
		t.Errorf("Expected [9, 1], got %v", slice)
	}
}

func TestRadixSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	RadixSort(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestRadixSortDescendingAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	RadixSortDescending(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestRadixSortUint32AllSame(t *testing.T) {
	slice := []uint32{5, 5, 5, 5, 5}
	RadixSortUint32(slice)

	for i := 0; i < len(slice); i++ {
		if slice[i] != 5 {
			t.Errorf("All elements should be 5: %v", slice)
		}
	}
}

func TestRadixSortPreservesLength(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	originalLength := len(slice)
	RadixSort(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestRadixSortDescendingPreservesLength(t *testing.T) {
	slice := []int{1, 3, 5, 4, 2}
	originalLength := len(slice)
	RadixSortDescending(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestRadixSortUint32PreservesLength(t *testing.T) {
	slice := []uint32{5, 3, 1, 4, 2}
	originalLength := len(slice)
	RadixSortUint32(slice)

	if len(slice) != originalLength {
		t.Errorf("Length should be preserved: expected %d, got %d", originalLength, len(slice))
	}
}

func TestRadixSortRandom(t *testing.T) {
	slice := make([]int, 100)
	for i := 0; i < 100; i++ {
		slice[i] = rand.Intn(10000)
	}

	RadixSort(slice)

	if !IsSorted(slice) {
		t.Error("Random slice should be sorted")
	}
}

func TestRadixSortUint32Random(t *testing.T) {
	slice := make([]uint32, 100)
	for i := 0; i < 100; i++ {
		slice[i] = uint32(rand.Intn(10000))
	}

	RadixSortUint32(slice)

	if !IsSorted(slice) {
		t.Error("Random uint32 slice should be sorted")
	}
}

func TestRadixSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	RadixSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestRadixSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	RadixSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func BenchmarkRadixSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		RadixSort(slice)
	}
}

func BenchmarkRadixSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		RadixSortDescending(slice)
	}
}

func BenchmarkRadixSortCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		RadixSortCopy(slice)
	}
}

func BenchmarkRadixSortDescendingCopy(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		RadixSortDescendingCopy(slice)
	}
}

func BenchmarkRadixSortWithBase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		RadixSortWithBase(slice, 16)
	}
}

func BenchmarkRadixSortUint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]uint32, 1000)
		for j := range slice {
			slice[j] = uint32(rand.Intn(100000))
		}
		RadixSortUint32(slice)
	}
}

func BenchmarkRadixSortLarge(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 5000)
		for j := range slice {
			slice[j] = rand.Intn(1000000)
		}
		RadixSort(slice)
	}
}
