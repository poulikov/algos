package sorting

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	MergeSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestMergeSortEmpty(t *testing.T) {
	slice := []int{}
	MergeSort(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestMergeSortSingle(t *testing.T) {
	slice := []int{5}
	MergeSort(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestMergeSortSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	MergeSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Already sorted slice should remain sorted: %v", slice)
	}
}

func TestMergeSortReverseSorted(t *testing.T) {
	slice := []int{5, 4, 3, 2, 1}
	MergeSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Reverse sorted slice should be sorted: %v", slice)
	}
}

func TestMergeSortDuplicates(t *testing.T) {
	slice := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	MergeSort(slice)

	if !IsSorted(slice) {
		t.Errorf("Slice with duplicates should be sorted: %v", slice)
	}
}

func TestMergeSortStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry", "date"}
	MergeSort(slice)

	expected := []string{"apple", "banana", "cherry", "date"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestMergeSortFloats(t *testing.T) {
	slice := []float64{3.14, 1.41, 2.71, 0.577, 1.618}
	MergeSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Float slice should be sorted in ascending order: %v", slice)
		}
	}
}

func TestMergeSortLarge(t *testing.T) {
	slice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		slice[i] = 1000 - i
	}
	MergeSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Error("Large slice should be sorted in ascending order")
		}
	}
}

func TestMergeSortCopy(t *testing.T) {
	original := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	sorted := MergeSortCopy(original)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("Sorted slice should be in ascending order: %v", sorted)
		}
	}

	if IsSorted(original) {
		t.Error("Original slice should not be modified")
	}
}

func TestMergeSortCopyEmpty(t *testing.T) {
	original := []int{}
	sorted := MergeSortCopy(original)

	if len(sorted) != 0 {
		t.Error("Empty copy should be empty")
	}
}

func TestMergeSortDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	MergeSortDescending(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestMergeSortDescendingEmpty(t *testing.T) {
	slice := []int{}
	MergeSortDescending(slice)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestMergeSortDescendingSingle(t *testing.T) {
	slice := []int{5}
	MergeSortDescending(slice)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestMergeSortDescendingReverseSorted(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	MergeSortDescending(slice)

	expected := []int{5, 4, 3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestMergeSortDescendingStrings(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	MergeSortDescending(slice)

	expected := []string{"cherry", "banana", "apple"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestMergeSortWithComparator(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}

	compare := func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	MergeSortWithComparator(slice, compare)

	if !IsSorted(slice) {
		t.Errorf("Slice should be sorted: %v", slice)
	}
}

func TestMergeSortWithComparatorDescending(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}

	compare := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		}
		return 0
	}

	MergeSortWithComparator(slice, compare)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			t.Errorf("Slice should be sorted in descending order: %v", slice)
		}
	}
}

func TestMergeSortWithComparatorStrings(t *testing.T) {
	slice := []string{"banana", "apple", "cherry"}

	compare := func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	MergeSortWithComparator(slice, compare)

	expected := []string{"apple", "banana", "cherry"}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, slice[i])
		}
	}
}

func TestMergeSortWithComparatorEmpty(t *testing.T) {
	slice := []int{}
	compare := func(a, b int) int { return 0 }

	MergeSortWithComparator(slice, compare)

	if len(slice) != 0 {
		t.Error("Empty slice should remain empty")
	}
}

func TestMergeSortWithComparatorSingle(t *testing.T) {
	slice := []int{5}
	compare := func(a, b int) int { return 0 }

	MergeSortWithComparator(slice, compare)

	if len(slice) != 1 || slice[0] != 5 {
		t.Error("Single element should remain unchanged")
	}
}

func TestMergeSortEvenLength(t *testing.T) {
	slice := []int{5, 2, 8, 1}
	MergeSort(slice)

	expected := []int{1, 2, 5, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestMergeSortOddLength(t *testing.T) {
	slice := []int{5, 2, 8, 1, 9}
	MergeSort(slice)

	expected := []int{1, 2, 5, 8, 9}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestMergeSortTwoElements(t *testing.T) {
	slice := []int{9, 1}
	MergeSort(slice)

	if slice[0] != 1 || slice[1] != 9 {
		t.Errorf("Expected [1, 9], got %v", slice)
	}
}

func TestMergeSortNegativeNumbers(t *testing.T) {
	slice := []int{-5, 3, -2, 8, -1}
	MergeSort(slice)

	expected := []int{-5, -2, -1, 3, 8}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestMergeSortAllSame(t *testing.T) {
	slice := []int{5, 5, 5, 5, 5}
	MergeSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] != 5 || slice[i] != slice[i+1] {
			t.Errorf("All elements should be 5, got %v", slice)
		}
	}
}

func TestMergeSortMaxInt(t *testing.T) {
	slice := []int{2147483647, -2147483648, 0, 1, -1}
	MergeSort(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("Slice with max/min int should be sorted: %v", slice)
		}
	}
}

func TestMergeSortStability(t *testing.T) {
	type pair struct {
		first  int
		second int
	}

	slice := []pair{
		{2, 1},
		{1, 2},
		{2, 3},
		{1, 4},
	}

	compare := func(a, b pair) int {
		if a.first < b.first {
			return -1
		} else if a.first > b.first {
			return 1
		}
		return 0
	}

	MergeSortWithComparator(slice, compare)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i].first > slice[i+1].first {
			t.Errorf("Slice should be sorted by first element: %v", slice)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		MergeSort(slice)
	}
}

func BenchmarkMergeSortDescending(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		MergeSortDescending(slice)
	}
}

func BenchmarkMergeSortWithComparator(b *testing.B) {
	compare := func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 1000)
		for j := range slice {
			slice[j] = rand.Intn(10000)
		}
		MergeSortWithComparator(slice, compare)
	}
}
