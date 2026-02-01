package sorting

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	slice := []int{5, 3, 1, 4, 2}
	QuickSort(slice)
	if !IsSorted(slice) {
		t.Fatal("QuickSort failed")
	}
}

func TestQuickSortEmpty(t *testing.T) {
	slice := []int{}
	QuickSort(slice)
	if len(slice) != 0 {
		t.Fatal("Empty slice should remain empty")
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice := make([]int, 100)
		for j := range slice {
			slice[j] = rand.Intn(100000)
		}
		QuickSort(slice)
	}
}
