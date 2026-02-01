package twopointers

import (
	"testing"
)

func TestTwoPointers(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	target := 6

	result := TwoPointers(data, target)

	if !result {
		t.Error("Should find pair that sums to 6 (1+5)")
	}
}

func TestTwoPointersNoPair(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	target := 10

	result := TwoPointers(data, target)

	if result {
		t.Error("Should not find pair that sums to 10")
	}
}

func TestTwoPointersEmpty(t *testing.T) {
	data := []int{}
	target := 5

	result := TwoPointers(data, target)

	if result {
		t.Error("Empty data should return false")
	}
}

func TestTwoPointersSorted(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	target := 6

	result := TwoPointersSorted(data, target)

	if !result {
		t.Error("Should find pair that sums to 6 (1+5)")
	}
}

func TestTwoPointersSortedNoPair(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	target := 20

	result := TwoPointersSorted(data, target)

	if result {
		t.Error("Should not find pair for large target")
	}
}

func TestTwoPointersDescending(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	target := 6

	result := TwoPointersDescending(data, target)

	if !result {
		t.Error("Should find pair that sums to 6 (4+2)")
	}
}

func TestTwoPointersAny(t *testing.T) {
	data := []int{3, 5, 2, -4, 8, 11, -7}
	target := 7

	i, j, found := TwoPointersAny(data, target)

	if !found {
		t.Error("Should find pair that sums to 7")
	}

	if data[i]+data[j] != target {
		t.Errorf("Found pair should sum to %d, got %d+%d", target, data[i], data[j])
	}
}

func TestTwoPointersAnyNoPair(t *testing.T) {
	data := []int{1, 2, 3}
	target := 10

	_, _, found := TwoPointersAny(data, target)

	if found {
		t.Error("Should not find pair")
	}
}

func TestIsPalindrome(t *testing.T) {
	data := []int{1, 2, 3, 2, 1}

	result := IsPalindrome(data)

	if !result {
		t.Error("Should be palindrome")
	}
}

func TestIsPalindromeNot(t *testing.T) {
	data := []int{1, 2, 3, 4}

	result := IsPalindrome(data)

	if result {
		t.Error("Should not be palindrome")
	}
}

func TestIsPalindromeEmpty(t *testing.T) {
	data := []int{}

	result := IsPalindrome(data)

	if !result {
		t.Error("Empty slice should be palindrome")
	}
}

func TestIsPalindromeSingle(t *testing.T) {
	data := []int{5}

	result := IsPalindrome(data)

	if !result {
		t.Error("Single element should be palindrome")
	}
}

func TestIsPalindromeStrings(t *testing.T) {
	data := []string{"hello", "world", "hello"}

	result := IsPalindrome(data)

	if !result {
		t.Error("Should be palindrome")
	}
}

func TestHasDuplicates(t *testing.T) {
	data := []int{1, 2, 3, 3, 5}

	result := HasDuplicates(data)

	if !result {
		t.Error("Should find duplicates in array with duplicates")
	}
}

func TestHasDuplicatesNone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	result := HasDuplicates(data)

	if result {
		t.Error("Should not detect duplicates in unique array")
	}
}

func TestHasDuplicatesEmpty(t *testing.T) {
	data := []int{}

	result := HasDuplicates(data)

	if result {
		t.Error("Empty array should have no duplicates")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	data := []int{1, 2, 2, 3, 3, 3, 4}

	result := RemoveDuplicates(data)

	expected := []int{1, 2, 3, 4}
	if len(result) != len(expected) {
		t.Errorf("Expected %d unique elements, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestRemoveDuplicatesEmpty(t *testing.T) {
	data := []int{}

	result := RemoveDuplicates(data)

	if len(result) != 0 {
		t.Error("Empty array should remain empty")
	}
}

func TestRemoveDuplicatesSingle(t *testing.T) {
	data := []int{5}

	result := RemoveDuplicates(data)

	if len(result) != 1 {
		t.Error("Single element should remain")
	}
}

func TestFindPair(t *testing.T) {
	data := []int{2, 7, 11, 15}
	target := 9

	i, j, found := FindPair(data, target)

	if !found {
		t.Error("Should find pair 2+7=9")
	}

	if data[i]+data[j] != target {
		t.Errorf("Found pair should sum to %d", target)
	}
}

func TestFindPairDescending(t *testing.T) {
	data := []int{15, 11, 7, 2}
	target := 9

	i, j, found := FindPairDescending(data, target)

	if !found {
		t.Error("Should find pair 7+2=9")
	}

	if data[i]+data[j] != target {
		t.Errorf("Found pair should sum to %d", target)
	}
}

func TestFindClosest(t *testing.T) {
	data := []int{1, 4, 5, 6, 8, 9}
	target := 11

	i, j, found := FindClosest(data, target)

	if !found {
		t.Error("Should find closest pair")
	}

	sum := data[i] + data[j]
	if sum > target {
		t.Errorf("Sum %d should not exceed target %d", sum, target)
	}
}

func TestFindClosestNone(t *testing.T) {
	data := []int{10, 20, 30}
	target := 5

	_, _, found := FindClosest(data, target)

	if found {
		t.Error("Should not find pair where sum <= 5")
	}
}

func TestFindClosestAny(t *testing.T) {
	data := []int{1, 4, 5, 6}
	target := 9

	i, j, found := FindClosestAny(data, target)

	if !found {
		t.Error("Should find pair 4+5=9")
	}

	if data[i]+data[j] != target {
		t.Errorf("Found pair should sum to %d", target)
	}
}

func TestFindTriple(t *testing.T) {
	data := []int{1, 4, 5, 6, 7}
	target := 10

	i, j, k, found := FindTriple(data, target)

	if !found {
		t.Error("Should find triple 1+4+5=10")
	}

	sum := data[i] + data[j] + data[k]
	if sum != target {
		t.Errorf("Triple should sum to %d, got %d", target, sum)
	}
}

func TestFindTripleNoMatch(t *testing.T) {
	data := []int{1, 2, 3}
	target := 100

	_, _, _, found := FindTriple(data, target)

	if found {
		t.Error("Should not find triple")
	}
}

func TestFindTripleClosest(t *testing.T) {
	data := []int{1, 4, 5, 6, 7}
	target := 13

	i, j, k, found := FindTripleClosest(data, target)

	if !found {
		t.Error("Should find triple")
	}

	sum := data[i] + data[j] + data[k]
	if sum > target {
		t.Errorf("Triple sum %d should not exceed target %d", sum, target)
	}
}

func TestFindKClosest(t *testing.T) {
	data := []int{1, 4, 5, 6}
	target := 11
	k := 4

	indices, found := FindKClosest(data, target, k)

	if !found {
		t.Error("Should find k-tuple")
	}

	if len(indices) > k {
		t.Errorf("Should not return more than %d indices", k)
	}
}

func TestFindKClosestClosest(t *testing.T) {
	data := []int{1, 4, 5, 6}
	target := 12
	k := 4

	_, _, _, found := FindKClosestClosest(data, target, k)

	if !found {
		t.Error("Should find k-tuple")
	}
}

func TestFindAllPairs(t *testing.T) {
	data := []int{1, 5, 7, -1, 5}
	target := 6

	pairs := FindAllPairs(data, target)

	if len(pairs) != 3 {
		t.Errorf("Expected 3 pairs, got %d", len(pairs))
	}

	for _, pair := range pairs {
		if data[pair[0]]+data[pair[1]] != target {
			t.Errorf("Pair %v doesn't sum to %d", pair, target)
		}
	}
}

func TestFindAllPairsNone(t *testing.T) {
	data := []int{1, 2, 3}
	target := 100

	pairs := FindAllPairs(data, target)

	if len(pairs) != 0 {
		t.Error("Should return empty slice for no pairs")
	}
}

func TestFindAllTriples(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	target := 9

	triples := FindAllTriples(data, target)

	if len(triples) != 2 {
		t.Errorf("Expected 2 triples, got %d", len(triples))
	}
}

func TestPartition(t *testing.T) {
	data := []int{3, 6, 8, 10, 1, 2, 1}
	pivotIndex := len(data) - 1

	result := Partition(data, pivotIndex)

	expectedPivotValue := data[result]
	leftAllSmaller := true
	for i := 0; i < result; i++ {
		if data[i] > expectedPivotValue {
			leftAllSmaller = false
			break
		}
	}

	rightAllLarger := true
	for i := result + 1; i < len(data); i++ {
		if data[i] < expectedPivotValue {
			rightAllLarger = false
			break
		}
	}

	if !leftAllSmaller || !rightAllLarger {
		t.Error("Partition should place all smaller elements left of pivot and larger right of pivot")
	}
}

func TestQuickSort(t *testing.T) {
	data := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	sorted := QuickSort(data)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", sorted)
		}
	}
}

func TestQuickSortEmpty(t *testing.T) {
	data := []int{}

	sorted := QuickSort(data)

	if len(sorted) != 0 {
		t.Error("Empty array should remain empty")
	}
}

func TestMergeSortWithTwoPointers(t *testing.T) {
	data := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	sorted := MergeSortWithTwoPointers(data)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Errorf("Slice should be sorted in ascending order: %v", sorted)
		}
	}
}

func TestMergeSortWithTwoPointersEmpty(t *testing.T) {
	data := []int{}

	sorted := MergeSortWithTwoPointers(data)

	if len(sorted) != 0 {
		t.Error("Empty array should remain empty")
	}
}

func TestFindMid(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	mid := FindMid(data)

	if mid != 3 {
		t.Errorf("Expected middle 3, got %d", mid)
	}
}

func TestFindMidEmpty(t *testing.T) {
	data := []int{}

	mid := FindMid(data)

	if mid != 0 {
		t.Error("Empty array should return 0")
	}
}

func TestFindDuplicate(t *testing.T) {
	data := []int{1, 2, 3, 4, 2, 5}

	duplicate, found := FindDuplicate(data)

	if !found {
		t.Error("Should find duplicate in array with duplicates")
	}

	if duplicate != 2 {
		t.Errorf("Expected duplicate 2, got %d", duplicate)
	}
}

func TestFindDuplicateNone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	_, found := FindDuplicate(data)

	if found {
		t.Error("Should not find duplicate in unique array")
	}
}

func TestFindDuplicateEmpty(t *testing.T) {
	data := []int{}

	_, found := FindDuplicate(data)

	if found {
		t.Error("Empty array should have no duplicates")
	}
}

func TestTwoPointersStrings(t *testing.T) {
	data := []string{"a", "b", "c", "d", "e"}
	target := "ae"

	result := TwoPointers(data, target)

	if !result {
		t.Error("Should find pair that sums to 'ae' ('a'+'e')")
	}
}

func TestIsPalindromeEven(t *testing.T) {
	data := []int{1, 2, 2, 1}

	result := IsPalindrome(data)

	if !result {
		t.Error("Even length palindrome should be detected")
	}
}

func TestIsPalindromeOdd(t *testing.T) {
	data := []int{1, 2, 3, 2, 1}

	result := IsPalindrome(data)

	if !result {
		t.Error("Odd length palindrome should be detected")
	}
}

func TestTwoPointersNegative(t *testing.T) {
	data := []int{-5, -3, -1, 0, 2, 4, 6}
	target := 1

	result := TwoPointers(data, target)

	if !result {
		t.Error("Should find pair -1+2=1")
	}
}

func TestFindPairFloats(t *testing.T) {
	data := []float64{1.5, 2.5, 3.5}
	target := 4.0

	i, j, found := FindPair(data, target)

	if !found {
		t.Error("Should find pair 1.5+2.5=4.0")
	}

	if data[i]+data[j] != target {
		t.Errorf("Found pair should sum to %f", target)
	}
}

func TestLargeArray(t *testing.T) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}

	target := 1997
	i, j, found := TwoPointersAny(data, target)

	if !found {
		t.Error("Should find pair in large array")
	}

	if data[i]+data[j] != target {
		t.Errorf("Pair should sum to %d", target)
	}
}
