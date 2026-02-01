package set

import (
	"testing"
)

func TestNewHashSet(t *testing.T) {
	hs := NewHashSet[int]()
	if hs == nil {
		t.Fatal("NewHashSet() returned nil")
	}
	if !hs.IsEmpty() {
		t.Fatal("NewHashSet should be empty")
	}
	if hs.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", hs.Size())
	}
}

func TestNewHashSetWithCapacity(t *testing.T) {
	hs := NewHashSetWithCapacity[int](32)
	if hs.Capacity() != 32 {
		t.Fatalf("Expected capacity 32, got %d", hs.Capacity())
	}
}

func TestNewHashSetWithLoadFactor(t *testing.T) {
	hs := NewHashSetWithLoadFactor[int](16, 0.5)
	if hs.Capacity() != 16 {
		t.Fatalf("Expected capacity 16, got %d", hs.Capacity())
	}
}

func TestNewHashSetFromSlice(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	hs := NewHashSetFromSlice(elements)

	if hs.Size() != 5 {
		t.Fatalf("Expected size 5, got %d", hs.Size())
	}

	for _, element := range elements {
		if !hs.Contains(element) {
			t.Fatalf("Element %d not found in HashSet", element)
		}
	}
}

func TestAdd(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	if hs.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", hs.Size())
	}
}

func TestAddDuplicate(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(1)

	if hs.Size() != 1 {
		t.Fatalf("Expected size 1 after duplicate add, got %d", hs.Size())
	}
}

func TestAddAll(t *testing.T) {
	hs := NewHashSet[int]()
	hs.AddAll([]int{1, 2, 3})

	if hs.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", hs.Size())
	}
}

func TestContains(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	if !hs.Contains(1) {
		t.Fatal("Contains(1) should return true")
	}

	if hs.Contains(3) {
		t.Fatal("Contains(3) should return false")
	}
}

func TestHSRemove(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	if !hs.Remove(1) {
		t.Fatal("Remove(1) should return true")
	}

	if hs.Contains(1) {
		t.Fatal("Element 1 should not exist after removal")
	}

	if hs.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", hs.Size())
	}
}

func TestHSRemoveNonExistent(t *testing.T) {
	hs := NewHashSet[int]()
	if hs.Remove(1) {
		t.Fatal("Remove(1) should return false for non-existent element")
	}
}

func TestRemoveAll(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	hs.RemoveAll([]int{1, 2})

	if hs.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", hs.Size())
	}

	if hs.Contains(1) || hs.Contains(2) {
		t.Fatal("Elements 1 and 2 should be removed")
	}
}

func TestHSClear(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	hs.Clear()

	if !hs.IsEmpty() {
		t.Fatal("Clear() should make set empty")
	}
	if hs.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", hs.Size())
	}
}

func TestToSlice(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	slice := hs.ToSlice()
	if len(slice) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(slice))
	}

	elementSet := map[int]bool{1: true, 2: true, 3: true}
	for _, element := range slice {
		if !elementSet[element] {
			t.Fatalf("Unexpected element: %d", element)
		}
	}
}

func TestContainsAll(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)
	hs1.Add(3)

	hs2 := NewHashSet[int]()
	hs2.Add(1)
	hs2.Add(2)

	if !hs1.ContainsAll(hs2) {
		t.Fatal("hs1 should contain all elements from hs2")
	}

	hs2.Add(4)
	if hs1.ContainsAll(hs2) {
		t.Fatal("hs1 should not contain all elements from hs2 with extra element")
	}
}

func TestContainsAny(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)
	hs1.Add(3)

	hs2 := NewHashSet[int]()
	hs2.Add(4)
	hs2.Add(5)

	if hs1.ContainsAny(hs2) {
		t.Fatal("hs1 should not contain any elements from hs2")
	}

	hs2.Add(1)
	if !hs1.ContainsAny(hs2) {
		t.Fatal("hs1 should contain at least one element from hs2")
	}
}

func TestUnion(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)

	hs2 := NewHashSet[int]()
	hs2.Add(2)
	hs2.Add(3)

	union := hs1.Union(hs2)

	if union.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", union.Size())
	}

	for _, element := range []int{1, 2, 3} {
		if !union.Contains(element) {
			t.Fatalf("Union should contain element %d", element)
		}
	}
}

func TestIntersection(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)
	hs1.Add(3)

	hs2 := NewHashSet[int]()
	hs2.Add(2)
	hs2.Add(3)
	hs2.Add(4)

	intersection := hs1.Intersection(hs2)

	if intersection.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", intersection.Size())
	}

	for _, element := range []int{2, 3} {
		if !intersection.Contains(element) {
			t.Fatalf("Intersection should contain element %d", element)
		}
	}

	if intersection.Contains(1) || intersection.Contains(4) {
		t.Fatal("Intersection should not contain non-shared elements")
	}
}

func TestDifference(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)
	hs1.Add(3)

	hs2 := NewHashSet[int]()
	hs2.Add(2)
	hs2.Add(4)

	difference := hs1.Difference(hs2)

	if difference.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", difference.Size())
	}

	for _, element := range []int{1, 3} {
		if !difference.Contains(element) {
			t.Fatalf("Difference should contain element %d", element)
		}
	}

	if difference.Contains(2) {
		t.Fatal("Difference should not contain shared elements")
	}
}

func TestSymmetricDifference(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)

	hs2 := NewHashSet[int]()
	hs2.Add(2)
	hs2.Add(3)

	symDiff := hs1.SymmetricDifference(hs2)

	if symDiff.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", symDiff.Size())
	}

	for _, element := range []int{1, 3} {
		if !symDiff.Contains(element) {
			t.Fatalf("SymmetricDifference should contain element %d", element)
		}
	}

	if symDiff.Contains(2) {
		t.Fatal("SymmetricDifference should not contain shared elements")
	}
}

func TestSubset(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)

	hs2 := NewHashSet[int]()
	hs2.Add(1)
	hs2.Add(2)
	hs2.Add(3)

	if !hs1.Subset(hs2) {
		t.Fatal("hs1 should be a subset of hs2")
	}

	if hs2.Subset(hs1) {
		t.Fatal("hs2 should not be a subset of hs1")
	}
}

func TestSuperset(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)

	hs2 := NewHashSet[int]()
	hs2.Add(1)
	hs2.Add(2)
	hs2.Add(3)

	if !hs2.Superset(hs1) {
		t.Fatal("hs2 should be a superset of hs1")
	}

	if hs1.Superset(hs2) {
		t.Fatal("hs1 should not be a superset of hs2")
	}
}

func TestEquals(t *testing.T) {
	hs1 := NewHashSet[int]()
	hs1.Add(1)
	hs1.Add(2)

	hs2 := NewHashSet[int]()
	hs2.Add(1)
	hs2.Add(2)

	hs3 := NewHashSet[int]()
	hs3.Add(1)
	hs3.Add(2)
	hs3.Add(3)

	if !hs1.Equals(hs2) {
		t.Fatal("hs1 should equal hs2")
	}

	if hs1.Equals(hs3) {
		t.Fatal("hs1 should not equal hs3")
	}
}

func TestHSCopy(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	hsCopy := hs.Copy()

	if hsCopy.Size() != hs.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", hsCopy.Size(), hs.Size())
	}

	hs.Add(3)
	hs.Remove(1)

	if hsCopy.Contains(3) {
		t.Fatal("Copy should be independent")
	}

	if !hsCopy.Contains(1) {
		t.Fatal("Copy should have original elements")
	}
}

func TestHSForEach(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	count := 0
	hs.ForEach(func(element int) {
		count++
	})

	if count != 3 {
		t.Fatalf("ForEach should visit all elements, got %d", count)
	}
}

func TestHSFilter(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	hs.Add(4)

	filtered := hs.Filter(func(element int) bool {
		return element%2 == 0
	})

	if filtered.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", filtered.Size())
	}

	for _, element := range []int{2, 4} {
		if !filtered.Contains(element) {
			t.Fatalf("Filter should keep element %d", element)
		}
	}
}

func TestHSMap(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)

	mapped := hs.Map(func(element int) int {
		return element * 2
	})

	if mapped.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", mapped.Size())
	}

	for _, element := range []int{2, 4} {
		if !mapped.Contains(element) {
			t.Fatalf("Map should contain element %d", element)
		}
	}
}

func TestReduce(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	result := hs.Reduce(0, func(acc, element int) int {
		return acc + element
	})

	if result != 6 {
		t.Fatalf("Expected sum 6, got %d", result)
	}
}

func TestAny(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	if !hs.Any(func(element int) bool {
		return element%2 == 1
	}) {
		t.Fatal("Any should find odd element")
	}

	if hs.Any(func(element int) bool {
		return element > 10
	}) {
		t.Fatal("Any should not find element > 10")
	}
}

func TestAll(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(2)
	hs.Add(4)
	hs.Add(6)

	if !hs.All(func(element int) bool {
		return element%2 == 0
	}) {
		t.Fatal("All should find all elements are even")
	}

	hs.Add(3)
	if hs.All(func(element int) bool {
		return element%2 == 0
	}) {
		t.Fatal("All should not find all elements are even after adding odd")
	}
}

func TestFind(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)

	element, found := hs.Find(func(e int) bool {
		return e%2 == 0
	})

	if !found {
		t.Fatal("Find should find even element")
	}

	if element != 2 {
		t.Fatalf("Expected element 2, got %d", element)
	}
}

func TestRetainOnly(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	hs.Add(4)

	hs.RetainOnly([]int{1, 2})

	if hs.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", hs.Size())
	}

	if !hs.Contains(1) || !hs.Contains(2) {
		t.Fatal("RetainOnly should keep specified elements")
	}
}

func TestRemoveIf(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	hs.Add(4)

	hs.RemoveIf(func(element int) bool {
		return element%2 == 0
	})

	if hs.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", hs.Size())
	}

	if hs.Contains(2) || hs.Contains(4) {
		t.Fatal("RemoveIf should remove even elements")
	}
}

func TestAddIfAbsent(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)

	if hs.AddIfAbsent(1) {
		t.Fatal("AddIfAbsent should return false for existing element")
	}

	if !hs.AddIfAbsent(2) {
		t.Fatal("AddIfAbsent should return true for new element")
	}
}

func TestHSComputeIfAbsent(t *testing.T) {
	hs := NewHashSet[int]()
	hs.Add(1)

	result := hs.ComputeIfAbsent(1, func(e int) int {
		return e * 2
	})

	if result != 1 {
		t.Fatalf("Expected 1, got %d", result)
	}

	result = hs.ComputeIfAbsent(2, func(e int) int {
		return e * 2
	})

	if result != 4 {
		t.Fatalf("Expected 4, got %d", result)
	}

	if !hs.Contains(4) {
		t.Fatal("ComputeIfAbsent should add computed element")
	}
}

func TestStringType(t *testing.T) {
	hs := NewHashSet[string]()
	hs.Add("one")
	hs.Add("two")

	if !hs.Contains("one") {
		t.Fatal("HashSet should work with string type")
	}
}

func BenchmarkAdd(b *testing.B) {
	hs := NewHashSet[int]()
	for i := 0; i < b.N; i++ {
		hs.Add(i)
	}
}

func BenchmarkContains(b *testing.B) {
	hs := NewHashSet[int]()
	for i := 0; i < 1000; i++ {
		hs.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.Contains(i % 1000)
	}
}

func BenchmarkHSRemove(b *testing.B) {
	hs := NewHashSet[int]()
	for i := 0; i < 1000; i++ {
		hs.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.Remove(i % 1000)
		hs.Add(i % 1000)
	}
}

func BenchmarkHSForEach(b *testing.B) {
	hs := NewHashSet[int]()
	for i := 0; i < 1000; i++ {
		hs.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs.ForEach(func(element int) {})
	}
}

func BenchmarkUnion(b *testing.B) {
	hs1 := NewHashSet[int]()
	hs2 := NewHashSet[int]()
	for i := 0; i < 100; i++ {
		hs1.Add(i)
		hs2.Add(i + 100)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs1.Union(hs2)
	}
}

func BenchmarkIntersection(b *testing.B) {
	hs1 := NewHashSet[int]()
	hs2 := NewHashSet[int]()
	for i := 0; i < 100; i++ {
		hs1.Add(i)
		hs2.Add(i + 50)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hs1.Intersection(hs2)
	}
}
