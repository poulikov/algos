package queues

import (
	"testing"
)

func TestNewDeque(t *testing.T) {
	d := NewDeque[int]()
	if d == nil {
		t.Fatal("NewDeque() returned nil")
	}
	if !d.IsEmpty() {
		t.Fatal("NewDeque() deque should be empty")
	}
	if d.Size() != 0 {
		t.Fatalf("NewDeque() deque should have size 0, got %d", d.Size())
	}
}

func TestPushFront(t *testing.T) {
	d := NewDeque[int]()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	if d.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", d.Size())
	}

	item, _ := d.PeekFront()
	if item != 3 {
		t.Fatalf("Expected front to be 3, got %d", item)
	}
}

func TestPushBack(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	if d.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", d.Size())
	}

	item, _ := d.PeekBack()
	if item != 3 {
		t.Fatalf("Expected back to be 3, got %d", item)
	}
}

func TestPopFront(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item, err := d.PopFront()
	if err != nil {
		t.Fatalf("PopFront() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to pop front 1, got %d", item)
	}
	if d.Size() != 2 {
		t.Fatalf("Expected size 2 after pop front, got %d", d.Size())
	}

	item, err = d.PopFront()
	if err != nil {
		t.Fatalf("PopFront() returned error: %v", err)
	}
	if item != 2 {
		t.Fatalf("Expected to pop front 2, got %d", item)
	}
}

func TestPopFrontEmptyDeque(t *testing.T) {
	d := NewDeque[int]()
	_, err := d.PopFront()
	if err != ErrDequeEmpty {
		t.Fatalf("Expected ErrDequeEmpty, got %v", err)
	}
}

func TestPopBack(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item, err := d.PopBack()
	if err != nil {
		t.Fatalf("PopBack() returned error: %v", err)
	}
	if item != 3 {
		t.Fatalf("Expected to pop back 3, got %d", item)
	}
	if d.Size() != 2 {
		t.Fatalf("Expected size 2 after pop back, got %d", d.Size())
	}

	item, err = d.PopBack()
	if err != nil {
		t.Fatalf("PopBack() returned error: %v", err)
	}
	if item != 2 {
		t.Fatalf("Expected to pop back 2, got %d", item)
	}
}

func TestPopBackEmptyDeque(t *testing.T) {
	d := NewDeque[int]()
	_, err := d.PopBack()
	if err != ErrDequeEmpty {
		t.Fatalf("Expected ErrDequeEmpty, got %v", err)
	}
}

func TestPeekFront(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item, err := d.PeekFront()
	if err != nil {
		t.Fatalf("PeekFront() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to peek front 1, got %d", item)
	}
	if d.Size() != 3 {
		t.Fatalf("PeekFront() should not change size, got %d", d.Size())
	}
}

func TestPeekFrontEmptyDeque(t *testing.T) {
	d := NewDeque[int]()
	_, err := d.PeekFront()
	if err != ErrDequeEmpty {
		t.Fatalf("Expected ErrDequeEmpty, got %v", err)
	}
}

func TestPeekBack(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item, err := d.PeekBack()
	if err != nil {
		t.Fatalf("PeekBack() returned error: %v", err)
	}
	if item != 3 {
		t.Fatalf("Expected to peek back 3, got %d", item)
	}
	if d.Size() != 3 {
		t.Fatalf("PeekBack() should not change size, got %d", d.Size())
	}
}

func TestPeekBackEmptyDeque(t *testing.T) {
	d := NewDeque[int]()
	_, err := d.PeekBack()
	if err != ErrDequeEmpty {
		t.Fatalf("Expected ErrDequeEmpty, got %v", err)
	}
}

func TestDequeIsEmpty(t *testing.T) {
	d := NewDeque[int]()
	if !d.IsEmpty() {
		t.Fatal("New deque should be empty")
	}

	d.PushBack(1)
	if d.IsEmpty() {
		t.Fatal("Deque with items should not be empty")
	}

	d.PopBack()
	if !d.IsEmpty() {
		t.Fatal("Deque after popping all items should be empty")
	}
}

func TestDequeSize(t *testing.T) {
	d := NewDeque[int]()
	if d.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", d.Size())
	}

	for i := 1; i <= 5; i++ {
		d.PushBack(i)
		if d.Size() != i {
			t.Fatalf("Expected size %d, got %d", i, d.Size())
		}
	}
}

func TestDequeClear(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	d.Clear()

	if !d.IsEmpty() {
		t.Fatal("Clear() should make deque empty")
	}
	if d.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", d.Size())
	}
}

func TestDequeCopy(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	dCopy := d.Copy()

	if dCopy.Size() != d.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", dCopy.Size(), d.Size())
	}

	d.PopFront()
	d.PushBack(4)

	dCopyFront, _ := dCopy.PeekFront()
	dFront, _ := d.PeekFront()
	if dCopyFront == dFront {
		t.Fatalf("Copy should be independent, peek returned same value: %d", dCopyFront)
	}
}

func TestDequeToSlice(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	slice := d.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("Expected slice length 3, got %d", len(slice))
	}
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Fatalf("Expected [1 2 3], got %v", slice)
	}
}

func TestDequeString(t *testing.T) {
	d := NewDeque[int]()
	str := d.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)
	str = d.String()
	if str != "[1 2 3]" {
		t.Fatalf("Expected '[1 2 3]', got '%s'", str)
	}
}

func TestDequeOrder(t *testing.T) {
	d := NewDeque[int]()

	// Push from both ends
	d.PushBack(20)
	d.PushFront(10)
	d.PushBack(30)
	d.PushFront(5)

	// Expected: [5, 10, 20, 30]

	item, _ := d.PeekFront()
	if item != 5 {
		t.Fatalf("Expected front 5, got %d", item)
	}

	item, _ = d.PeekBack()
	if item != 30 {
		t.Fatalf("Expected back 30, got %d", item)
	}

	// Pop from front
	item, _ = d.PopFront()
	if item != 5 {
		t.Fatalf("Expected to pop front 5, got %d", item)
	}

	// Pop from back
	item, _ = d.PopBack()
	if item != 30 {
		t.Fatalf("Expected to pop back 30, got %d", item)
	}

	if d.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", d.Size())
	}
}

func TestDequeStringType(t *testing.T) {
	d := NewDeque[string]()
	d.PushBack("first")
	d.PushFront("zero")

	item, _ := d.PopFront()
	if item != "zero" {
		t.Fatalf("Expected 'zero', got '%s'", item)
	}
}

func BenchmarkPushFront(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}

func BenchmarkPushBack(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
}

func BenchmarkPopFront(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopFront()
	}
}

func BenchmarkPopBack(b *testing.B) {
	d := NewDeque[int]()
	for i := 0; i < b.N; i++ {
		d.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PopBack()
	}
}
