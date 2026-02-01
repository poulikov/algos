package structures

import (
	"sort"
	"testing"
)

func TestSkipListNew(t *testing.T) {
	sl := NewSkipList()

	if sl.Size() != 0 {
		t.Errorf("Expected size 0, got %d", sl.Size())
	}
	if !sl.IsEmpty() {
		t.Errorf("Expected empty skip list")
	}
}

func TestSkipListInsert(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	if sl.Size() != 4 {
		t.Errorf("Expected size 4, got %d", sl.Size())
	}

	if sl.IsEmpty() {
		t.Errorf("Expected non-empty skip list")
	}
}

func TestSkipListSearch(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	tests := []struct {
		value    int
		expected bool
	}{
		{1, true},
		{3, true},
		{5, true},
		{7, true},
		{2, false},
		{4, false},
		{6, false},
		{8, false},
	}

	for _, test := range tests {
		result := sl.Search(test.value)
		if result != test.expected {
			t.Errorf("Search(%d) = %v, expected %v", test.value, result, test.expected)
		}
	}
}

func TestSkipListDelete(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	if !sl.Delete(3) {
		t.Errorf("Delete(3) should return true")
	}
	if sl.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sl.Size())
	}
	if sl.Search(3) {
		t.Errorf("Element 3 should not exist after deletion")
	}

	if sl.Delete(2) {
		t.Errorf("Delete(2) should return false")
	}
}

func TestSkipListToSlice(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	slice := sl.ToSlice()

	if len(slice) != 4 {
		t.Errorf("Expected slice length 4, got %d", len(slice))
	}

	for i := 1; i < len(slice); i++ {
		if slice[i] <= slice[i-1] {
			t.Errorf("Slice should be sorted: %v", slice)
		}
	}
}

func TestSkipListMinMax(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	if sl.Min() != 1 {
		t.Errorf("Min() = %d, expected 1", sl.Min())
	}
	if sl.Max() != 7 {
		t.Errorf("Max() = %d, expected 7", sl.Max())
	}
}

func TestSkipListClear(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	sl.Clear()

	if !sl.IsEmpty() {
		t.Errorf("Expected empty skip list after clear")
	}
	if sl.Size() != 0 {
		t.Errorf("Expected size 0, got %d", sl.Size())
	}
}

func TestSkipListDuplicate(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(5)
	sl.Insert(7)

	if sl.Size() != 3 {
		t.Errorf("Expected size 3 (no duplicates), got %d", sl.Size())
	}
}

func TestSkipListOrder(t *testing.T) {
	sl := NewSkipList()

	values := []int{5, 3, 7, 1, 9, 2, 8, 4, 6, 0}

	for _, v := range values {
		sl.Insert(v)
	}

	slice := sl.ToSlice()
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)

	for i := range slice {
		if slice[i] != sorted[i] {
			t.Errorf("SkipList order mismatch: got %v, expected %v", slice, sorted)
			return
		}
	}
}

func TestSkipListContains(t *testing.T) {
	sl := NewSkipList()

	sl.Insert(5)
	sl.Insert(3)
	sl.Insert(7)
	sl.Insert(1)

	if !sl.Contains(5) {
		t.Errorf("Contains(5) should return true")
	}
	if sl.Contains(2) {
		t.Errorf("Contains(2) should return false")
	}
}

func TestSkipListLarge(t *testing.T) {
	sl := NewSkipList()

	n := 1000
	for i := 0; i < n; i++ {
		sl.Insert(i)
	}

	if sl.Size() != n {
		t.Errorf("Expected size %d, got %d", n, sl.Size())
	}

	for i := 0; i < n; i++ {
		if !sl.Search(i) {
			t.Errorf("Search(%d) should return true", i)
		}
	}
}

func TestSkipListConsistency(t *testing.T) {
	sl := NewSkipList()

	values := []int{5, 3, 7, 1, 9, 2, 8, 4, 6, 0}

	for _, v := range values {
		sl.Insert(v)
	}

	slice := sl.ToSlice()

	if len(slice) != len(values) {
		t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
	}

	if sl.Min() != 0 {
		t.Errorf("Min() = %d, expected 0", sl.Min())
	}

	if sl.Max() != 9 {
		t.Errorf("Max() = %d, expected 9", sl.Max())
	}

	for _, v := range values {
		if !sl.Search(v) {
			t.Errorf("Search(%d) should return true", v)
		}
	}
}
