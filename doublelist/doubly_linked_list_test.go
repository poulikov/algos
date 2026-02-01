package doublelist

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	dll := New[int]()

	if dll == nil {
		t.Fatal("New() returned nil")
	}

	if dll.Head != nil || dll.Tail != nil {
		t.Error("New list should have nil head and tail")
	}

	if dll.Size != 0 {
		t.Errorf("New list should have size 0, got %d", dll.Size)
	}
}

func TestAppend(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	if dll.Size != 3 {
		t.Errorf("Expected size 3, got %d", dll.Size)
	}

	if dll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", dll.Head.Value)
	}

	if dll.Tail.Value != 3 {
		t.Errorf("Expected tail value 3, got %d", dll.Tail.Value)
	}
}

func TestAppendSingle(t *testing.T) {
	dll := New[int]()
	dll.Append(5)

	if dll.Size != 1 {
		t.Errorf("Expected size 1, got %d", dll.Size)
	}

	if dll.Head == nil || dll.Tail == nil {
		t.Error("Head and tail should not be nil")
	}

	if dll.Head != dll.Tail {
		t.Error("Head and tail should be the same node")
	}

	if dll.Head.Value != 5 {
		t.Errorf("Expected value 5, got %d", dll.Head.Value)
	}
}

func TestPrepend(t *testing.T) {
	dll := New[int]()
	dll.Prepend(3)
	dll.Prepend(2)
	dll.Prepend(1)

	if dll.Size != 3 {
		t.Errorf("Expected size 3, got %d", dll.Size)
	}

	if dll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", dll.Head.Value)
	}

	if dll.Tail.Value != 3 {
		t.Errorf("Expected tail value 3, got %d", dll.Tail.Value)
	}
}

func TestPrependSingle(t *testing.T) {
	dll := New[int]()
	dll.Prepend(5)

	if dll.Size != 1 {
		t.Errorf("Expected size 1, got %d", dll.Size)
	}

	if dll.Head.Value != 5 {
		t.Errorf("Expected value 5, got %d", dll.Head.Value)
	}
}

func TestInsertAt(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(3)

	err := dll.InsertAt(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	if dll.Size != 3 {
		t.Errorf("Expected size 3, got %d", dll.Size)
	}

	val, _ := dll.Get(1)
	if val != 2 {
		t.Errorf("Expected value 2 at index 1, got %d", val)
	}
}

func TestInsertAtBeginning(t *testing.T) {
	dll := New[int]()
	dll.Append(2)
	dll.Append(3)

	err := dll.InsertAt(0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if dll.Head.Value != 1 {
		t.Errorf("Expected head value 1, got %d", dll.Head.Value)
	}
}

func TestInsertAtEnd(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)

	err := dll.InsertAt(2, 3)
	if err != nil {
		t.Fatal(err)
	}

	if dll.Tail.Value != 3 {
		t.Errorf("Expected tail value 3, got %d", dll.Tail.Value)
	}
}

func TestInsertAtInvalidIndex(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)

	err := dll.InsertAt(5, 3)
	if err == nil {
		t.Error("Expected error for invalid index")
	}

	err = dll.InsertAt(-1, 3)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestInsertAtEmpty(t *testing.T) {
	dll := New[int]()

	err := dll.InsertAt(0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if dll.Size != 1 {
		t.Errorf("Expected size 1, got %d", dll.Size)
	}
}

func TestRemoveAt(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.RemoveAt(1)
	if err != nil {
		t.Fatal(err)
	}

	if val != 2 {
		t.Errorf("Expected removed value 2, got %d", val)
	}

	if dll.Size != 2 {
		t.Errorf("Expected size 2 after removal, got %d", dll.Size)
	}
}

func TestRemoveAtHead(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.RemoveAt(0)
	if err != nil {
		t.Fatal(err)
	}

	if val != 1 {
		t.Errorf("Expected removed value 1, got %d", val)
	}

	if dll.Head.Value != 2 {
		t.Errorf("Expected new head value 2, got %d", dll.Head.Value)
	}
}

func TestRemoveAtTail(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.RemoveAt(2)
	if err != nil {
		t.Fatal(err)
	}

	if val != 3 {
		t.Errorf("Expected removed value 3, got %d", val)
	}

	if dll.Tail.Value != 2 {
		t.Errorf("Expected new tail value 2, got %d", dll.Tail.Value)
	}
}

func TestRemoveAtInvalidIndex(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)

	_, err := dll.RemoveAt(5)
	if err == nil {
		t.Error("Expected error for invalid index")
	}

	_, err = dll.RemoveAt(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestRemoveFirst(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.RemoveFirst()
	if err != nil {
		t.Fatal(err)
	}

	if val != 1 {
		t.Errorf("Expected removed value 1, got %d", val)
	}

	if dll.Head.Value != 2 {
		t.Errorf("Expected new head value 2, got %d", dll.Head.Value)
	}
}

func TestRemoveFirstEmpty(t *testing.T) {
	dll := New[int]()

	_, err := dll.RemoveFirst()
	if err == nil {
		t.Error("Expected error when removing from empty list")
	}
}

func TestRemoveFirstSingleElement(t *testing.T) {
	dll := New[int]()
	dll.Append(1)

	val, err := dll.RemoveFirst()
	if err != nil {
		t.Fatal(err)
	}

	if val != 1 {
		t.Errorf("Expected removed value 1, got %d", val)
	}

	if dll.Head != nil || dll.Tail != nil {
		t.Error("Head and tail should be nil after removing single element")
	}
}

func TestRemoveLast(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.RemoveLast()
	if err != nil {
		t.Fatal(err)
	}

	if val != 3 {
		t.Errorf("Expected removed value 3, got %d", val)
	}

	if dll.Tail.Value != 2 {
		t.Errorf("Expected new tail value 2, got %d", dll.Tail.Value)
	}
}

func TestRemoveLastEmpty(t *testing.T) {
	dll := New[int]()

	_, err := dll.RemoveLast()
	if err == nil {
		t.Error("Expected error when removing from empty list")
	}
}

func TestRemoveLastSingleElement(t *testing.T) {
	dll := New[int]()
	dll.Append(1)

	val, err := dll.RemoveLast()
	if err != nil {
		t.Fatal(err)
	}

	if val != 1 {
		t.Errorf("Expected removed value 1, got %d", val)
	}

	if dll.Head != nil || dll.Tail != nil {
		t.Error("Head and tail should be nil after removing single element")
	}
}

func TestGet(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	val, err := dll.Get(1)
	if err != nil {
		t.Fatal(err)
	}

	if val != 2 {
		t.Errorf("Expected value 2 at index 1, got %d", val)
	}
}

func TestGetInvalidIndex(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)

	_, err := dll.Get(5)
	if err == nil {
		t.Error("Expected error for invalid index")
	}

	_, err = dll.Get(-1)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestSet(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	err := dll.Set(1, 10)
	if err != nil {
		t.Fatal(err)
	}

	val, _ := dll.Get(1)
	if val != 10 {
		t.Errorf("Expected value 10 at index 1, got %d", val)
	}
}

func TestSetInvalidIndex(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)

	err := dll.Set(5, 10)
	if err == nil {
		t.Error("Expected error for invalid index")
	}

	err = dll.Set(-1, 10)
	if err == nil {
		t.Error("Expected error for negative index")
	}
}

func TestIndexOf(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	index := dll.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index 1 for value 2, got %d", index)
	}

	index = dll.IndexOf(10)
	if index != -1 {
		t.Errorf("Expected index -1 for non-existent value, got %d", index)
	}
}

func TestIndexOfDuplicates(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(2)
	dll.Append(3)

	index := dll.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected first occurrence index 1 for value 2, got %d", index)
	}
}

func TestContains(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	if !dll.Contains(2) {
		t.Error("List should contain value 2")
	}

	if dll.Contains(10) {
		t.Error("List should not contain value 10")
	}
}

func TestLen(t *testing.T) {
	dll := New[int]()

	if dll.Len() != 0 {
		t.Errorf("Expected length 0, got %d", dll.Len())
	}

	dll.Append(1)
	dll.Append(2)

	if dll.Len() != 2 {
		t.Errorf("Expected length 2, got %d", dll.Len())
	}
}

func TestIsEmpty(t *testing.T) {
	dll := New[int]()

	if !dll.IsEmpty() {
		t.Error("New list should be empty")
	}

	dll.Append(1)

	if dll.IsEmpty() {
		t.Error("List with element should not be empty")
	}
}

func TestClear(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.Clear()

	if !dll.IsEmpty() {
		t.Error("List should be empty after Clear")
	}

	if dll.Head != nil || dll.Tail != nil {
		t.Error("Head and tail should be nil after Clear")
	}
}

func TestToSlice(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	slice := dll.ToSlice()

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
	dll := New[int]()

	slice := dll.ToSlice()

	if len(slice) != 0 {
		t.Errorf("Expected empty slice, got length %d", len(slice))
	}
}

func TestToSliceReverse(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	slice := dll.ToSliceReverse()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	expected := []int{3, 2, 1}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, slice[i])
		}
	}
}

func TestForEach(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	sum := 0
	err := dll.ForEach(func(val int) error {
		sum += val
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestForEachWithError(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	err := dll.ForEach(func(val int) error {
		if val == 2 {
			return errors.New("found 2")
		}
		return nil
	})

	if err == nil {
		t.Error("Expected error from ForEach")
	}
}

func TestForEachReverse(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	sum := 0
	err := dll.ForEachReverse(func(val int) error {
		sum += val
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if sum != 6 {
		t.Errorf("Expected sum 6, got %d", sum)
	}
}

func TestForEachReverseWithError(t *testing.T) {
	dll := New[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	err := dll.ForEachReverse(func(val int) error {
		if val == 2 {
			return errors.New("found 2")
		}
		return nil
	})

	if err == nil {
		t.Error("Expected error from ForEachReverse")
	}
}

func TestDoublyLinkedListStrings(t *testing.T) {
	dll := New[string]()
	dll.Append("hello")
	dll.Append("world")

	val, _ := dll.Get(0)
	if val != "hello" {
		t.Errorf("Expected 'hello', got %s", val)
	}

	val, _ = dll.Get(1)
	if val != "world" {
		t.Errorf("Expected 'world', got %s", val)
	}
}

func TestDoublyLinkedListCustomType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	dll := New[Person]()
	dll.Append(Person{Name: "Alice", Age: 30})
	dll.Append(Person{Name: "Bob", Age: 25})

	person, _ := dll.Get(0)
	if person.Name != "Alice" || person.Age != 30 {
		t.Errorf("Expected Alice age 30, got %s age %d", person.Name, person.Age)
	}
}

func TestMixedOperations(t *testing.T) {
	dll := New[int]()

	dll.Append(1)
	dll.Append(2)
	dll.Prepend(0)

	if dll.Len() != 3 {
		t.Errorf("Expected length 3, got %d", dll.Len())
	}

	dll.RemoveAt(1)
	if dll.Len() != 2 {
		t.Errorf("Expected length 2, got %d", dll.Len())
	}

	dll.InsertAt(1, 5)
	val, _ := dll.Get(1)
	if val != 5 {
		t.Errorf("Expected value 5, got %d", val)
	}
}

func TestLargeList(t *testing.T) {
	dll := New[int]()

	for i := 0; i < 1000; i++ {
		dll.Append(i)
	}

	if dll.Len() != 1000 {
		t.Errorf("Expected length 1000, got %d", dll.Len())
	}

	val, _ := dll.Get(500)
	if val != 500 {
		t.Errorf("Expected value 500 at index 500, got %d", val)
	}

	val, _ = dll.Get(999)
	if val != 999 {
		t.Errorf("Expected value 999 at index 999, got %d", val)
	}
}
