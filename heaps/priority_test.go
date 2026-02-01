package heaps

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int, int](MaxPriority)
	if pq == nil {
		t.Fatal("NewPriorityQueue() returned nil")
	}
	if !pq.IsEmpty() {
		t.Fatal("NewPriorityQueue should be empty")
	}
	if pq.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", pq.Size())
	}
}

func TestNewMaxPriorityQueue(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	if pq.PriorityType() != MaxPriority {
		t.Fatal("NewMaxPriorityQueue should create MaxPriority queue")
	}
}

func TestNewMinPriorityQueue(t *testing.T) {
	pq := NewMinPriorityQueue[string, int]()
	if pq.PriorityType() != MinPriority {
		t.Fatal("NewMinPriorityQueue should create MinPriority queue")
	}
}

func TestEnqueueAndDequeueMax(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	value, priority, err := pq.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() returned error: %v", err)
	}
	if value != 2 {
		t.Fatalf("Expected value 2, got %d", value)
	}
	if priority != 30 {
		t.Fatalf("Expected priority 30, got %d", priority)
	}

	value, priority, _ = pq.Dequeue()
	if value != 3 {
		t.Fatalf("Expected value 3, got %d", value)
	}
	if priority != 20 {
		t.Fatalf("Expected priority 20, got %d", priority)
	}
}

func TestEnqueueAndDequeueMin(t *testing.T) {
	pq := NewMinPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	value, priority, err := pq.Dequeue()
	if err != nil {
		t.Fatalf("Dequeue() returned error: %v", err)
	}
	if value != 1 {
		t.Fatalf("Expected value 1, got %d", value)
	}
	if priority != 10 {
		t.Fatalf("Expected priority 10, got %d", priority)
	}

	value, priority, _ = pq.Dequeue()
	if value != 3 {
		t.Fatalf("Expected value 3, got %d", value)
	}
	if priority != 20 {
		t.Fatalf("Expected priority 20, got %d", priority)
	}
}

func TestDequeueEmpty(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	_, _, err := pq.Dequeue()
	if err != ErrHeapEmpty {
		t.Fatalf("Expected ErrHeapEmpty, got %v", err)
	}
}

func TestPQPeek(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	value, priority, err := pq.Peek()
	if err != nil {
		t.Fatalf("Peek() returned error: %v", err)
	}
	if value != 2 {
		t.Fatalf("Expected value 2, got %d", value)
	}
	if priority != 30 {
		t.Fatalf("Expected priority 30, got %d", priority)
	}

	if pq.Size() != 3 {
		t.Fatalf("Peek() should not change size, got %d", pq.Size())
	}
}

func TestPeekEmpty(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	_, _, err := pq.Peek()
	if err != ErrHeapEmpty {
		t.Fatalf("Expected ErrHeapEmpty, got %v", err)
	}
}

func TestPQSize(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()

	for i := 0; i < 5; i++ {
		pq.Enqueue("value", i)
		if pq.Size() != i+1 {
			t.Fatalf("Expected size %d, got %d", i+1, pq.Size())
		}
	}
}

func TestPQIsEmpty(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	if !pq.IsEmpty() {
		t.Fatal("New PriorityQueue should be empty")
	}

	pq.Enqueue("value", 10)
	if pq.IsEmpty() {
		t.Fatal("PriorityQueue with elements should not be empty")
	}
}

func TestPQClear(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	pq.Enqueue("value1", 10)
	pq.Enqueue("value2", 20)

	pq.Clear()

	if !pq.IsEmpty() {
		t.Fatal("Clear() should make PriorityQueue empty")
	}
	if pq.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", pq.Size())
	}
}

func TestPQCopy(t *testing.T) {
	pq := NewMaxPriorityQueue[string, int]()
	pq.Enqueue("value1", 10)
	pq.Enqueue("value2", 20)

	pqCopy := pq.Copy()

	if pqCopy.Size() != pq.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", pqCopy.Size(), pq.Size())
	}

	pq.Enqueue("value3", 30)

	if pqCopy.Size() != 2 {
		t.Fatal("Copy should be independent")
	}
}

func TestPQToSlice(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	slice := pq.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("Expected 3 items, got %d", len(slice))
	}
}

func TestValues(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	values := pq.Values()
	if len(values) != 3 {
		t.Fatalf("Expected 3 values, got %d", len(values))
	}

	if values[0] != 2 {
		t.Fatalf("First value should be 2, got %d", values[0])
	}
}

func TestContains(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 20)

	equals := func(a, b int) bool { return a == b }

	if !pq.Contains(1, equals) {
		t.Fatal("Contains should return true for existing value")
	}

	if pq.Contains(3, equals) {
		t.Fatal("Contains should return false for non-existent value")
	}
}

func TestRemove(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 20)
	pq.Enqueue(3, 30)

	equals := func(a, b int) bool { return a == b }

	if !pq.Remove(2, equals) {
		t.Fatal("Remove should return true for existing value")
	}

	if pq.Contains(2, equals) {
		t.Fatal("Value should be removed")
	}

	if pq.Size() != 2 {
		t.Fatalf("Expected size 2 after removal, got %d", pq.Size())
	}
}

func TestUpdatePriority(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 20)
	pq.Enqueue(3, 30)

	equals := func(a, b int) bool { return a == b }

	if !pq.UpdatePriority(2, 35, equals) {
		t.Fatal("UpdatePriority should return true for existing value")
	}

	value, priority, _ := pq.Dequeue()
	if value != 2 {
		t.Fatalf("Expected value 2 to be dequeued first after priority update, got %d", value)
	}
	if priority != 35 {
		t.Fatalf("Expected updated priority 35, got %d", priority)
	}
}

func TestDequeueAll(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	items := pq.DequeueAll()

	if len(items) != 3 {
		t.Fatalf("Expected 3 items, got %d", len(items))
	}

	if items[0].value != 2 {
		t.Fatalf("First item should have value 2, got %d", items[0].value)
	}

	if !pq.IsEmpty() {
		t.Fatal("PriorityQueue should be empty after DequeueAll")
	}
}

func TestMerge(t *testing.T) {
	pq1 := NewMaxPriorityQueue[int, int]()
	pq1.Enqueue(1, 10)
	pq1.Enqueue(2, 20)

	pq2 := NewMaxPriorityQueue[int, int]()
	pq2.Enqueue(3, 30)
	pq2.Enqueue(4, 40)

	pq1.Merge(pq2)

	if pq1.Size() != 4 {
		t.Fatalf("Expected size 4 after merge, got %d", pq1.Size())
	}

	if !pq1.Contains(3, func(a, b int) bool { return a == b }) {
		t.Fatal("Merge should add elements from other queue")
	}
}

func TestForEach(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 30)
	pq.Enqueue(3, 20)

	count := 0
	pq.ForEach(func(value int, priority int) {
		count++
	})

	if count != 3 {
		t.Fatalf("ForEach should visit all elements, got %d", count)
	}
}

func TestFilter(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 20)
	pq.Enqueue(3, 30)

	filtered := pq.Filter(func(value int, priority int) bool {
		return priority >= 20
	})

	if filtered.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", filtered.Size())
	}

	if !filtered.Contains(2, func(a, b int) bool { return a == b }) {
		t.Fatal("Filter should keep element with priority >= 20")
	}
}

func TestEqualPriorities(t *testing.T) {
	pq := NewMaxPriorityQueue[int, int]()
	pq.Enqueue(1, 10)
	pq.Enqueue(2, 10)
	pq.Enqueue(3, 10)

	for i := 0; i < 3; i++ {
		_, _, err := pq.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue() failed: %v", err)
		}
	}

	if !pq.IsEmpty() {
		t.Fatal("All elements should be dequeued")
	}
}

func TestStringTypePriority(t *testing.T) {
	pq := NewMinPriorityQueue[string, string]()
	pq.Enqueue("low", "low")
	pq.Enqueue("high", "high")
	pq.Enqueue("medium", "medium")

	value, _, _ := pq.Dequeue()
	if value != "high" {
		t.Fatalf("Expected 'high', got '%s'", value)
	}
}

func BenchmarkEnqueue(b *testing.B) {
	pq := NewMaxPriorityQueue[int, int]()
	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	pq := NewMaxPriorityQueue[int, int]()
	for i := 0; i < 1000; i++ {
		pq.Enqueue(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Enqueue(i+1000, i+1000)
		pq.Dequeue()
	}
}

func BenchmarkPQPeek(b *testing.B) {
	pq := NewMaxPriorityQueue[int, int]()
	for i := 0; i < 1000; i++ {
		pq.Enqueue(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Peek()
	}
}

func BenchmarkContains(b *testing.B) {
	pq := NewMaxPriorityQueue[int, int]()
	for i := 0; i < 1000; i++ {
		pq.Enqueue(i, i)
	}
	equals := func(a, b int) bool { return a == b }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Contains(i%1000, equals)
	}
}
