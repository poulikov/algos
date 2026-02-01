package stacks

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()

	if s == nil {
		t.Fatal("New() returned nil")
	}

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	if s.Size() != 0 {
		t.Errorf("New stack should have size 0, got %d", s.Size())
	}
}

func TestPush(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}

	if s.IsEmpty() {
		t.Error("Stack should not be empty after pushes")
	}
}

func TestPop(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Pop()
	if err != nil {
		t.Fatal(err)
	}

	if item != 3 {
		t.Errorf("Expected popped value 3, got %d", item)
	}

	if s.Size() != 2 {
		t.Errorf("Expected size 2 after pop, got %d", s.Size())
	}
}

func TestPopEmpty(t *testing.T) {
	s := New[int]()

	_, err := s.Pop()
	if err != ErrStackEmpty {
		t.Errorf("Expected ErrStackEmpty, got %v", err)
	}
}

func TestPopAll(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 3; i >= 1; i-- {
		item, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if item != i {
			t.Errorf("Expected %d, got %d", i, item)
		}
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, err := s.Peek()
	if err != nil {
		t.Fatal(err)
	}

	if item != 3 {
		t.Errorf("Expected peek value 3, got %d", item)
	}

	if s.Size() != 3 {
		t.Errorf("Peek should not change size, got %d", s.Size())
	}
}

func TestPeekEmpty(t *testing.T) {
	s := New[int]()

	_, err := s.Peek()
	if err != ErrStackEmpty {
		t.Errorf("Expected ErrStackEmpty, got %v", err)
	}
}

func TestLIFO(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	item1, _ := s.Pop()
	item2, _ := s.Pop()
	item3, _ := s.Pop()

	if item1 != 3 || item2 != 2 || item3 != 1 {
		t.Errorf("LIFO violation: got %d, %d, %d", item1, item2, item3)
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Error("Stack with element should not be empty")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

func TestSize(t *testing.T) {
	s := New[int]()

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	s.Push(1)
	s.Push(2)

	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}

	s.Pop()

	if s.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Clear()

	if !s.IsEmpty() {
		t.Error("Stack should be empty after Clear")
	}

	if s.Size() != 0 {
		t.Errorf("Size should be 0 after Clear, got %d", s.Size())
	}
}

func TestClearEmpty(t *testing.T) {
	s := New[int]()

	s.Clear()

	if !s.IsEmpty() {
		t.Error("Clearing empty stack should keep it empty")
	}
}

func TestCopy(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	copied := s.Copy()

	if copied.Size() != s.Size() {
		t.Errorf("Copied stack should have same size, got %d vs %d", copied.Size(), s.Size())
	}

	s.Pop()

	if copied.Size() != 3 {
		t.Errorf("Modifying original should not affect copy, got size %d", copied.Size())
	}
}

func TestCopyEmpty(t *testing.T) {
	s := New[int]()

	copied := s.Copy()

	if copied.Size() != 0 {
		t.Errorf("Copied empty stack should have size 0, got %d", copied.Size())
	}
}

func TestToSlice(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	slice := s.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestToSliceEmpty(t *testing.T) {
	s := New[int]()

	slice := s.ToSlice()

	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(slice))
	}
}

func TestString(t *testing.T) {
	s := New[int]()

	str := s.String()
	if str != "[]" {
		t.Errorf("Expected '[]', got %s", str)
	}

	s.Push(1)
	s.Push(2)

	str = s.String()
	if str == "[]" {
		t.Error("Non-empty stack should not return '[]'")
	}
}

func TestStackStrings(t *testing.T) {
	s := New[string]()
	s.Push("hello")
	s.Push("world")

	item, _ := s.Pop()
	if item != "world" {
		t.Errorf("Expected 'world', got %s", item)
	}

	item, _ = s.Pop()
	if item != "hello" {
		t.Errorf("Expected 'hello', got %s", item)
	}
}

func TestStackFloats(t *testing.T) {
	s := New[float64]()
	s.Push(3.14)
	s.Push(2.71)

	item, _ := s.Pop()
	if item != 2.71 {
		t.Errorf("Expected 2.71, got %f", item)
	}
}

func TestStackCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	s := New[Person]()
	s.Push(Person{Name: "Alice", Age: 30})
	s.Push(Person{Name: "Bob", Age: 25})

	person, _ := s.Pop()
	if person.Name != "Bob" || person.Age != 25 {
		t.Errorf("Expected Bob age 25, got %s age %d", person.Name, person.Age)
	}
}

func TestLargeStack(t *testing.T) {
	s := New[int]()

	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	if s.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", s.Size())
	}

	item, _ := s.Peek()
	if item != 999 {
		t.Errorf("Expected top value 999, got %d", item)
	}

	for i := 999; i >= 0; i-- {
		item, err := s.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if item != i {
			t.Errorf("Expected %d, got %d", i, item)
		}
	}
}

func TestPushPopAlternating(t *testing.T) {
	s := New[int]()

	s.Push(1)
	item, _ := s.Pop()
	if item != 1 {
		t.Errorf("Expected 1, got %d", item)
	}

	s.Push(2)
	item, _ = s.Pop()
	if item != 2 {
		t.Errorf("Expected 2, got %d", item)
	}

	s.Push(3)
	item, _ = s.Pop()
	if item != 3 {
		t.Errorf("Expected 3, got %d", item)
	}
}

func TestMultiplePeeks(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 0; i < 5; i++ {
		item, err := s.Peek()
		if err != nil {
			t.Fatal(err)
		}
		if item != 3 {
			t.Errorf("Peek should always return 3, got %d", item)
		}
	}
}

func TestStackAfterClear(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)

	s.Clear()

	s.Push(3)
	s.Push(4)

	item, _ := s.Pop()
	if item != 4 {
		t.Errorf("Expected 4, got %d", item)
	}

	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
}

func TestCopyPreservesOrder(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	copied := s.Copy()

	slice1 := s.ToSlice()
	slice2 := copied.ToSlice()

	if len(slice1) != len(slice2) {
		t.Errorf("Copied slice should have same length")
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			t.Errorf("Copy should preserve order: %d != %d", slice1[i], slice2[i])
		}
	}
}

func BenchmarkPush(b *testing.B) {
	s := New[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func BenchmarkPop(b *testing.B) {
	s := New[int]()

	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if s.Size() == 0 {
			for j := 0; j < 1000; j++ {
				s.Push(j)
			}
		}
		s.Pop()
	}
}

func BenchmarkPeek(b *testing.B) {
	s := New[int]()

	for i := 0; i < 1000; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}
