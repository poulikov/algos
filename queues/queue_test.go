package queues

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := New[int]()
	if q == nil {
		t.Fatal("New() returned nil")
	}
	if !q.IsEmpty() {
		t.Fatal("New() queue should be empty")
	}
	if q.Size() != 0 {
		t.Fatalf("New() queue should have size 0, got %d", q.Size())
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", q.Size())
	}
}

func TestDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to dequeue 1, got %d", item)
	}
	if q.Size() != 2 {
		t.Fatalf("Expected size 2 after dequeue, got %d", q.Size())
	}

	item, err = q.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() returned error: %v", err)
	}
	if item != 2 {
		t.Fatalf("Expected to dequeue 2, got %d", item)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := New[int]()
	_, err := q.Dequeue()
	if err != ErrQueueEmpty {
		t.Fatalf("Expected ErrQueueEmpty, got %v", err)
	}
}

func TestDequeueFast(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.DequeueFast()
	if err != nil {
		t.Fatalf("DequeueFast() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to dequeue 1, got %d", item)
	}

	item, err = q.DequeueFast()
	if err != nil {
		t.Fatalf("DequeueFast() returned error: %v", err)
	}
	if item != 2 {
		t.Fatalf("Expected to dequeue 2, got %d", item)
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, err := q.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}
	if item != 1 {
		t.Fatalf("Expected to peek 1, got %d", item)
	}
	if q.Size() != 3 {
		t.Fatalf("Peek() should not change size, got %d", q.Size())
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	q := New[int]()
	_, err := q.Peek()
	if err != ErrQueueEmpty {
		t.Fatalf("Expected ErrQueueEmpty, got %v", err)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := New[int]()
	if !q.IsEmpty() {
		t.Fatal("New queue should be empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Fatal("Queue with items should not be empty")
	}

	q.Dequeue()
	if !q.IsEmpty() {
		t.Fatal("Queue after dequeuing all items should be empty")
	}
}

func TestQueueSize(t *testing.T) {
	q := New[int]()
	if q.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", q.Size())
	}

	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
		if q.Size() != i {
			t.Fatalf("Expected size %d, got %d", i, q.Size())
		}
	}
}

func TestQueueClear(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if !q.IsEmpty() {
		t.Fatal("Clear() should make queue empty")
	}
	if q.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", q.Size())
	}
}

func TestQueueCopy(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	qCopy := q.Copy()

	if qCopy.Size() != q.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", qCopy.Size(), q.Size())
	}

	q.Dequeue()
	q.Enqueue(4)

	qCopyItem, _ := qCopy.Peek()
	qItem, _ := q.Peek()
	if qCopyItem == qItem {
		t.Fatalf("Copy should be independent, peek returned same value: %d", qCopyItem)
	}
}

func TestQueueToSlice(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	slice := q.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("Expected slice length 3, got %d", len(slice))
	}
	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 {
		t.Fatalf("Expected [1 2 3], got %v", slice)
	}
}

func TestQueueString(t *testing.T) {
	q := New[int]()
	str := q.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	str = q.String()
	if str != "[1 2 3]" {
		t.Fatalf("Expected '[1 2 3]', got '%s'", str)
	}
}

func TestFIFOOrder(t *testing.T) {
	q := New[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	item, _ := q.Dequeue()
	if item != 10 {
		t.Fatalf("Expected 10, got %d", item)
	}

	item, _ = q.Dequeue()
	if item != 20 {
		t.Fatalf("Expected 20, got %d", item)
	}

	item, _ = q.Dequeue()
	if item != 30 {
		t.Fatalf("Expected 30, got %d", item)
	}

	if !q.IsEmpty() {
		t.Fatal("Queue should be empty after dequeuing all elements")
	}
}

func TestQueueStringType(t *testing.T) {
	q := New[string]()
	q.Enqueue("first")
	q.Enqueue("second")

	item, _ := q.Dequeue()
	if item != "first" {
		t.Fatalf("Expected 'first', got '%s'", item)
	}
}

func BenchmarkEnqueue(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkDequeueFast(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.DequeueFast()
	}
}

func BenchmarkPeek(b *testing.B) {
	q := New[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Peek()
	}
}
