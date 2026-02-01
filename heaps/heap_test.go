package heaps

import (
	"testing"
)

func TestNew(t *testing.T) {
	h := New[int](MinHeap, func(a, b int) bool { return a < b })
	if h == nil {
		t.Fatal("New() returned nil")
	}
	if !h.IsEmpty() {
		t.Fatal("New() heap should be empty")
	}
	if h.Size() != 0 {
		t.Fatalf("New() heap should have size 0, got %d", h.Size())
	}
}

func TestNewMinHeap(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	if h == nil {
		t.Fatal("NewMinHeap() returned nil")
	}
	if h.HeapType() != MinHeap {
		t.Fatal("NewMinHeap() should create MinHeap")
	}
	if !h.IsEmpty() {
		t.Fatal("NewMinHeap() heap should be empty")
	}
}

func TestNewMaxHeap(t *testing.T) {
	h := NewMaxHeap[int](func(a, b int) bool { return a < b })
	if h == nil {
		t.Fatal("NewMaxHeap() returned nil")
	}
	if h.HeapType() != MaxHeap {
		t.Fatal("NewMaxHeap() should create MaxHeap")
	}
	if !h.IsEmpty() {
		t.Fatal("NewMaxHeap() heap should be empty")
	}
}

func TestInsert(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	if h.Size() != 4 {
		t.Fatalf("Expected size 4, got %d", h.Size())
	}
}

func TestExtractMinHeap(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	item, err := h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to extract 1, got %d", item)
	}
	if h.Size() != 3 {
		t.Fatalf("Expected size 3 after extract, got %d", h.Size())
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 3 {
		t.Fatalf("Expected to extract 3, got %d", item)
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 5 {
		t.Fatalf("Expected to extract 5, got %d", item)
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 8 {
		t.Fatalf("Expected to extract 8, got %d", item)
	}
}

func TestExtractMaxHeap(t *testing.T) {
	h := NewMaxHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	item, err := h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 8 {
		t.Fatalf("Expected to extract 8, got %d", item)
	}
	if h.Size() != 3 {
		t.Fatalf("Expected size 3 after extract, got %d", h.Size())
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 5 {
		t.Fatalf("Expected to extract 5, got %d", item)
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 3 {
		t.Fatalf("Expected to extract 3, got %d", item)
	}

	item, err = h.Extract()
	if err != nil {
		t.Fatalf("Extract() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to extract 1, got %d", item)
	}
}

func TestExtractEmptyHeap(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	_, err := h.Extract()
	if err != ErrHeapEmpty {
		t.Fatalf("Expected ErrHeapEmpty, got %v", err)
	}
}

func TestPeek(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	item, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to peek 1, got %d", item)
	}
	if h.Size() != 4 {
		t.Fatalf("Peek() should not change size, got %d", h.Size())
	}

	// Peek multiple times
	item, _ = h.Peek()
	if item != 1 {
		t.Fatalf("Expected to peek 1 again, got %d", item)
	}
}

func TestPeekMaxHeap(t *testing.T) {
	h := NewMaxHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	item, err := h.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}
	if item != 8 {
		t.Fatalf("Expected to peek 8, got %d", item)
	}
}

func TestPeekEmptyHeap(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	_, err := h.Peek()
	if err != ErrHeapEmpty {
		t.Fatalf("Expected ErrHeapEmpty, got %v", err)
	}
}

func TestIsEmpty(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	if !h.IsEmpty() {
		t.Fatal("New heap should be empty")
	}

	h.Insert(1)
	if h.IsEmpty() {
		t.Fatal("Heap with items should not be empty")
	}

	h.Extract()
	if !h.IsEmpty() {
		t.Fatal("Heap after extracting all items should be empty")
	}
}

func TestSize(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	if h.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", h.Size())
	}

	for i := 1; i <= 5; i++ {
		h.Insert(i)
		if h.Size() != i {
			t.Fatalf("Expected size %d, got %d", i, h.Size())
		}
	}
}

func TestClear(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	h.Clear()

	if !h.IsEmpty() {
		t.Fatal("Clear() should make heap empty")
	}
	if h.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", h.Size())
	}
}

func TestCopy(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	hCopy := h.Copy()

	if hCopy.Size() != h.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", hCopy.Size(), h.Size())
	}

	h.Extract()
	h.Insert(4)

	hCopyItem, _ := hCopy.Peek()
	hItem, _ := h.Peek()
	if hCopyItem == hItem {
		t.Fatalf("Copy should be independent, peek returned same value: %d", hCopyItem)
	}
}

func TestToSlice(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(1)
	h.Insert(2)
	h.Insert(3)

	slice := h.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("Expected slice length 3, got %d", len(slice))
	}
	// Note: ToSlice doesn't guarantee order
}

func TestToSortedSlice(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	sorted := h.ToSortedSlice()
	if len(sorted) != 4 {
		t.Fatalf("Expected sorted slice length 4, got %d", len(sorted))
	}
	if sorted[0] != 1 || sorted[1] != 3 || sorted[2] != 5 || sorted[3] != 8 {
		t.Fatalf("Expected [1 3 5 8], got %v", sorted)
	}
}

func TestToSortedSliceMaxHeap(t *testing.T) {
	h := NewMaxHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(8)
	h.Insert(1)

	sorted := h.ToSortedSlice()
	if len(sorted) != 4 {
		t.Fatalf("Expected sorted slice length 4, got %d", len(sorted))
	}
	if sorted[0] != 8 || sorted[1] != 5 || sorted[2] != 3 || sorted[3] != 1 {
		t.Fatalf("Expected [8 5 3 1], got %v", sorted)
	}
}

func TestString(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	str := h.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	h.Insert(1)
	h.Insert(2)
	h.Insert(3)
	str = h.String()
	// Just check it's not empty, exact format depends on slice representation
	if str == "" {
		t.Fatal("String() should return non-empty string")
	}
}

func TestHeapType(t *testing.T) {
	minHeap := NewMinHeap[int](func(a, b int) bool { return a < b })
	if minHeap.HeapType() != MinHeap {
		t.Fatal("MinHeap should return MinHeap type")
	}

	maxHeap := NewMaxHeap[int](func(a, b int) bool { return a < b })
	if maxHeap.HeapType() != MaxHeap {
		t.Fatal("MaxHeap should return MaxHeap type")
	}
}

func TestHeapPropertyMinHeap(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })

	// Insert in random order
	h.Insert(10)
	h.Insert(5)
	h.Insert(15)
	h.Insert(3)
	h.Insert(7)
	h.Insert(12)
	h.Insert(1)
	h.Insert(9)

	// Extract should always give minimum
	prev := -1
	for i := 0; i < 8; i++ {
		item, err := h.Extract()
		if err != nil {
			t.Fatalf("Extract() returned error: %v", err)
		}
		if i > 0 && item < prev {
			t.Fatalf("Heap property violated: extracted %d after %d", item, prev)
		}
		prev = item
	}
}

func TestHeapPropertyMaxHeap(t *testing.T) {
	h := NewMaxHeap[int](func(a, b int) bool { return a < b })

	// Insert in random order
	h.Insert(10)
	h.Insert(5)
	h.Insert(15)
	h.Insert(3)
	h.Insert(7)
	h.Insert(12)
	h.Insert(1)
	h.Insert(9)

	// Extract should always give maximum
	prev := 100
	for i := 0; i < 8; i++ {
		item, err := h.Extract()
		if err != nil {
			t.Fatalf("Extract() returned error: %v", err)
		}
		if i > 0 && item > prev {
			t.Fatalf("Heap property violated: extracted %d after %d", item, prev)
		}
		prev = item
	}
}

func TestDuplicates(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(5)
	h.Insert(3)
	h.Insert(5)
	h.Insert(3)
	h.Insert(5)

	items := []int{}
	for i := 0; i < 5; i++ {
		item, _ := h.Extract()
		items = append(items, item)
	}

	expected := []int{3, 3, 5, 5, 5}
	if len(items) != len(expected) {
		t.Fatalf("Expected %d items, got %d", len(expected), len(items))
	}
	for i := range items {
		if items[i] != expected[i] {
			t.Fatalf("Expected item %d at position %d, got %d", expected[i], i, items[i])
		}
	}
}

func TestStringType(t *testing.T) {
	h := NewMinHeap[string](func(a, b string) bool { return a < b })
	h.Insert("zebra")
	h.Insert("apple")
	h.Insert("banana")

	item, _ := h.Peek()
	if item != "apple" {
		t.Fatalf("Expected 'apple', got '%s'", item)
	}

	item, _ = h.Extract()
	if item != "apple" {
		t.Fatalf("Expected 'apple', got '%s'", item)
	}

	item, _ = h.Extract()
	if item != "banana" {
		t.Fatalf("Expected 'banana', got '%s'", item)
	}

	item, _ = h.Extract()
	if item != "zebra" {
		t.Fatalf("Expected 'zebra', got '%s'", item)
	}
}

func TestFloatType(t *testing.T) {
	h := NewMinHeap[float64](func(a, b float64) bool { return a < b })
	h.Insert(3.14)
	h.Insert(2.71)
	h.Insert(1.41)

	item, _ := h.Extract()
	if item != 1.41 {
		t.Fatalf("Expected 1.41, got %f", item)
	}

	item, _ = h.Extract()
	if item != 2.71 {
		t.Fatalf("Expected 2.71, got %f", item)
	}

	item, _ = h.Extract()
	if item != 3.14 {
		t.Fatalf("Expected 3.14, got %f", item)
	}
}

func TestSingleElement(t *testing.T) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	h.Insert(42)

	if h.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", h.Size())
	}

	item, _ := h.Peek()
	if item != 42 {
		t.Fatalf("Expected peek 42, got %d", item)
	}

	item, _ = h.Extract()
	if item != 42 {
		t.Fatalf("Expected extract 42, got %d", item)
	}

	if !h.IsEmpty() {
		t.Fatal("Heap should be empty after extracting single element")
	}
}

func BenchmarkInsert(b *testing.B) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
}

func BenchmarkExtract(b *testing.B) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Extract()
	}
}

func BenchmarkPeek(b *testing.B) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Peek()
	}
}

func BenchmarkInsertExtract(b *testing.B) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		h.Insert(i)
		h.Extract()
	}
}

func BenchmarkToSortedSlice(b *testing.B) {
	h := NewMinHeap[int](func(a, b int) bool { return a < b })
	for i := 0; i < b.N; i++ {
		h.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		_ = h.ToSortedSlice()
	}
}
