package linkedlist

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()

	if ll == nil {
		t.Fatal("NewLinkedList() returned nil")
	}

	if ll.Head != nil {
		t.Error("New list should have nil head")
	}

	slice := ll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("New list should be empty, got %v", slice)
	}
}

func TestAppend(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	slice := ll.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}

	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestAppendSingle(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(42)

	slice := ll.ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected length 1, got %d", len(slice))
	}

	if slice[0] != 42 {
		t.Errorf("Expected 42, got %d", slice[0])
	}
}

func TestAppendEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	slice := ll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty list, got %v", slice)
	}
}

func TestPrepend(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)

	slice := ll.ToSlice()
	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}

	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestPrependSingle(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Prepend(99)

	slice := ll.ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected length 1, got %d", len(slice))
	}

	if slice[0] != 99 {
		t.Errorf("Expected 99, got %d", slice[0])
	}
}

func TestPrependAfterAppend(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(2)
	ll.Append(3)
	ll.Prepend(1)

	slice := ll.ToSlice()
	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestAppendAfterPrepend(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Prepend(1)
	ll.Append(2)
	ll.Append(3)

	slice := ll.ToSlice()
	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestForEach(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	sum := 0
	ll.ForEach(func(value int) {
		sum += value
	})

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestForEachEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	called := false
	ll.ForEach(func(value int) {
		called = true
	})

	if called {
		t.Error("ForEach should not be called on empty list")
	}
}

func TestForEachCollect(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	var result []int
	ll.ForEach(func(value int) {
		result = append(result, value)
	})

	expected := []int{10, 20, 30}
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Reverse()

	slice := ll.ToSlice()
	expected := []int{3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestReverseSingle(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(42)

	ll.Reverse()

	slice := ll.ToSlice()
	if len(slice) != 1 {
		t.Errorf("Expected length 1, got %d", len(slice))
	}

	if slice[0] != 42 {
		t.Errorf("Expected 42, got %d", slice[0])
	}
}

func TestReverseEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.Reverse()

	slice := ll.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Expected empty list after reverse, got %v", slice)
	}
}

func TestReverseTwice(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Reverse()
	ll.Reverse()

	slice := ll.ToSlice()
	expected := []int{1, 2, 3}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestToSlice(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	slice := ll.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}

	expected := []int{10, 20, 30}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestToSliceEmpty(t *testing.T) {
	ll := NewLinkedList[int]()

	slice := ll.ToSlice()

	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(slice))
	}
}

func TestLinkedListStrings(t *testing.T) {
	ll := NewLinkedList[string]()
	ll.Append("hello")
	ll.Append("world")

	slice := ll.ToSlice()

	if len(slice) != 2 {
		t.Errorf("Expected length 2, got %d", len(slice))
	}

	if slice[0] != "hello" || slice[1] != "world" {
		t.Errorf("Expected [hello world], got %v", slice)
	}
}

func TestLinkedListFloats(t *testing.T) {
	ll := NewLinkedList[float64]()
	ll.Append(3.14)
	ll.Append(2.71)
	ll.Append(1.41)

	slice := ll.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected length 3, got %d", len(slice))
	}

	expected := []float64{3.14, 2.71, 1.41}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %f at index %d, got %f", v, i, slice[i])
		}
	}
}

func TestLinkedListCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	ll := NewLinkedList[Person]()
	ll.Append(Person{Name: "Alice", Age: 30})
	ll.Append(Person{Name: "Bob", Age: 25})

	slice := ll.ToSlice()

	if len(slice) != 2 {
		t.Errorf("Expected length 2, got %d", len(slice))
	}

	if slice[0].Name != "Alice" || slice[0].Age != 30 {
		t.Errorf("Expected Alice age 30, got %s age %d", slice[0].Name, slice[0].Age)
	}
}

func TestLinkedListManyElements(t *testing.T) {
	ll := NewLinkedList[int]()

	for i := 1; i <= 100; i++ {
		ll.Append(i)
	}

	slice := ll.ToSlice()

	if len(slice) != 100 {
		t.Errorf("Expected length 100, got %d", len(slice))
	}

	for i := 0; i < 100; i++ {
		if slice[i] != i+1 {
			t.Errorf("Expected %d at index %d, got %d", i+1, i, slice[i])
		}
	}
}

func TestLinkedListReverseLarge(t *testing.T) {
	ll := NewLinkedList[int]()

	for i := 1; i <= 1000; i++ {
		ll.Append(i)
	}

	ll.Reverse()

	slice := ll.ToSlice()

	if len(slice) != 1000 {
		t.Errorf("Expected length 1000, got %d", len(slice))
	}

	for i := 0; i < 1000; i++ {
		expected := 1000 - i
		if slice[i] != expected {
			t.Errorf("Expected %d at index %d, got %d", expected, i, slice[i])
		}
	}
}

func TestLinkedListMixedOperations(t *testing.T) {
	ll := NewLinkedList[int]()

	ll.Append(2)
	ll.Prepend(1)
	ll.Append(4)
	ll.Prepend(0)
	ll.Append(5)

	slice := ll.ToSlice()
	expected := []int{0, 1, 2, 4, 5}

	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestLinkedListForEachModify(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	appended := false
	ll.ForEach(func(value int) {
		if value%2 == 0 && !appended {
			ll.Append(value * 10)
			appended = true
		}
	})

	slice := ll.ToSlice()

	if len(slice) != 4 {
		t.Errorf("Expected length 4, got %d", len(slice))
	}

	expected := []int{1, 2, 3, 20}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	ll := NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Append(i)
	}
}

func BenchmarkPrepend(b *testing.B) {
	ll := NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Prepend(i)
	}
}

func BenchmarkReverse(b *testing.B) {
	ll := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Reverse()
		ll.Reverse()
	}
}

func BenchmarkToSlice(b *testing.B) {
	ll := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.ToSlice()
	}
}

func BenchmarkForEach(b *testing.B) {
	ll := NewLinkedList[int]()
	for i := 0; i < 1000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum := 0
		ll.ForEach(func(value int) {
			sum += value
		})
	}
}
