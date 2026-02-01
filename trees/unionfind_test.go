package trees

import (
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	uf := NewUnionFind[int]()
	if uf == nil {
		t.Fatal("NewUnionFind() returned nil")
	}
	if !uf.IsEmpty() {
		t.Fatal("New UnionFind should be empty")
	}
	if uf.Count() != 0 {
		t.Fatalf("New UnionFind should have 0 sets, got %d", uf.Count())
	}
}

func TestNewUnionFindFromSlice(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	uf := NewUnionFindFromSlice(elements)

	if uf.Count() != 5 {
		t.Fatalf("Expected 5 sets, got %d", uf.Count())
	}

	for _, element := range elements {
		if !uf.Contains(element) {
			t.Fatalf("Element %d not found in UnionFind", element)
		}
	}
}

func TestMakeSet(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)

	if uf.Count() != 2 {
		t.Fatalf("Expected 2 sets, got %d", uf.Count())
	}

	if !uf.Contains(1) || !uf.Contains(2) {
		t.Fatal("Elements not found after MakeSet")
	}

	size1, _ := uf.Size(1)
	if size1 != 1 {
		t.Fatalf("Expected size 1, got %d", size1)
	}
}

func TestMakeSetDuplicate(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(1)

	if uf.Count() != 1 {
		t.Fatalf("Expected 1 set after duplicate MakeSet, got %d", uf.Count())
	}
}

func TestFind(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	root1, err := uf.Find(1)
	if err != nil {
		t.Fatalf("Find() returned error: %v", err)
	}
	if root1 != 1 {
		t.Fatalf("Expected root 1, got %d", root1)
	}

	root2, err := uf.Find(2)
	if err != nil {
		t.Fatalf("Find() returned error: %v", err)
	}
	if root2 != 2 {
		t.Fatalf("Expected root 2, got %d", root2)
	}
}

func TestFindNonExistent(t *testing.T) {
	uf := NewUnionFind[int]()
	_, err := uf.Find(1)
	if err == nil {
		t.Fatal("Find() should return error for non-existent element")
	}
}

func TestUnion(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	err := uf.Union(1, 2)
	if err != nil {
		t.Fatalf("Union() returned error: %v", err)
	}

	root1, _ := uf.Find(1)
	root2, _ := uf.Find(2)
	if root1 != root2 {
		t.Fatal("Elements 1 and 2 should be in same set after Union")
	}

	if uf.Count() != 2 {
		t.Fatalf("Expected 2 sets after Union, got %d", uf.Count())
	}
}

func TestUnionSameSet(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)

	uf.Union(1, 2)
	err := uf.Union(1, 2)

	if err != nil {
		t.Fatalf("Union() of same set should not error: %v", err)
	}
}

func TestUnionNonExistent(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)

	err := uf.Union(1, 2)
	if err == nil {
		t.Fatal("Union() should return error for non-existent element")
	}
}

func TestConnected(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	connected, err := uf.Connected(1, 2)
	if err != nil {
		t.Fatalf("Connected() returned error: %v", err)
	}
	if connected {
		t.Fatal("Elements 1 and 2 should not be connected initially")
	}

	uf.Union(1, 2)
	connected, err = uf.Connected(1, 2)
	if err != nil {
		t.Fatalf("Connected() returned error: %v", err)
	}
	if !connected {
		t.Fatal("Elements 1 and 2 should be connected after Union")
	}

	connected, _ = uf.Connected(1, 3)
	if connected {
		t.Fatal("Element 3 should not be connected to 1 and 2")
	}
}

func TestUFSize(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.MakeSet(4)

	uf.Union(1, 2)
	uf.Union(2, 3)

	size, err := uf.Size(1)
	if err != nil {
		t.Fatalf("Size() returned error: %v", err)
	}
	if size != 3 {
		t.Fatalf("Expected size 3, got %d", size)
	}

	size, _ = uf.Size(4)
	if size != 1 {
		t.Fatalf("Expected size 1, got %d", size)
	}
}

func TestCount(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	if uf.Count() != 3 {
		t.Fatalf("Expected 3 sets, got %d", uf.Count())
	}

	uf.Union(1, 2)
	if uf.Count() != 2 {
		t.Fatalf("Expected 2 sets after Union, got %d", uf.Count())
	}

	uf.Union(1, 3)
	if uf.Count() != 1 {
		t.Fatalf("Expected 1 set after second Union, got %d", uf.Count())
	}
}

func TestContains(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)

	if !uf.Contains(1) {
		t.Fatal("Contains(1) should return true")
	}

	if uf.Contains(2) {
		t.Fatal("Contains(2) should return false")
	}
}

func TestAdd(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.Add(1)
	uf.Add(2)

	if uf.Count() != 2 {
		t.Fatalf("Expected 2 sets, got %d", uf.Count())
	}
}

func TestUFClear(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	uf.Clear()

	if !uf.IsEmpty() {
		t.Fatal("Clear() should make UnionFind empty")
	}
	if uf.Count() != 0 {
		t.Fatalf("Clear() should result in 0 sets, got %d", uf.Count())
	}
}

func TestUFCopy(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)

	ufCopy := uf.Copy()

	if ufCopy.Count() != uf.Count() {
		t.Fatalf("Copy should have same count, got %d vs %d", ufCopy.Count(), uf.Count())
	}

	connected1, _ := uf.Connected(1, 2)
	connected2, _ := ufCopy.Connected(1, 2)
	if connected1 != connected2 {
		t.Fatal("Copy should preserve connections")
	}

	uf.Union(1, 3)
	connected1, _ = uf.Connected(1, 3)
	connected2, _ = ufCopy.Connected(1, 3)
	if connected1 == connected2 {
		t.Fatal("Copy should be independent")
	}
}

func TestToMap(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)

	elementMap := uf.ToMap()

	if len(elementMap) != 3 {
		t.Fatalf("Expected 3 elements in map, got %d", len(elementMap))
	}

	root1 := elementMap[1]
	root2 := elementMap[2]
	if root1 != root2 {
		t.Fatal("Elements 1 and 2 should have same root in map")
	}
}

func TestToGroups(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.MakeSet(4)
	uf.MakeSet(5)

	uf.Union(1, 2)
	uf.Union(2, 3)
	uf.Union(4, 5)

	groups := uf.ToGroups()

	if len(groups) != 2 {
		t.Fatalf("Expected 2 groups, got %d", len(groups))
	}

	for _, group := range groups {
		if len(group) == 3 {
			expected := map[int]bool{1: true, 2: true, 3: true}
			for _, elem := range group {
				if !expected[elem] {
					t.Fatalf("Unexpected element %d in group", elem)
				}
			}
		} else if len(group) == 2 {
			expected := map[int]bool{4: true, 5: true}
			for _, elem := range group {
				if !expected[elem] {
					t.Fatalf("Unexpected element %d in group", elem)
				}
			}
		} else {
			t.Fatalf("Unexpected group size: %d", len(group))
		}
	}
}

func TestElements(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	elements := uf.Elements()

	if len(elements) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(elements))
	}

	elemSet := map[int]bool{1: true, 2: true, 3: true}
	for _, elem := range elements {
		if !elemSet[elem] {
			t.Fatalf("Unexpected element %d", elem)
		}
	}
}

func TestUFString(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)

	str := uf.String()
	if str == "" {
		t.Fatal("String() should not return empty string")
	}
}

func TestGetParent(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)

	parent, err := uf.GetParent(1)
	if err != nil {
		t.Fatalf("GetParent() returned error: %v", err)
	}
	if parent != 1 {
		t.Fatalf("Expected parent 1, got %d", parent)
	}

	_, err = uf.GetParent(3)
	if err == nil {
		t.Fatal("GetParent() should return error for non-existent element")
	}
}

func TestGetRank(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)

	rank, err := uf.GetRank(1)
	if err != nil {
		t.Fatalf("GetRank() returned error: %v", err)
	}
	if rank != 0 {
		t.Fatalf("Expected rank 0, got %d", rank)
	}
}

func TestSizeOf(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)

	if uf.SizeOf() != 2 {
		t.Fatalf("Expected size 2, got %d", uf.SizeOf())
	}
}

func TestPathCompression(t *testing.T) {
	uf := NewUnionFind[int]()

	for i := 0; i < 100; i++ {
		uf.MakeSet(i)
	}

	for i := 0; i < 99; i++ {
		uf.Union(i, i+1)
	}

	root, err := uf.Find(0)
	if err != nil {
		t.Fatalf("Find() returned error: %v", err)
	}

	for i := 1; i < 100; i++ {
		currentRoot, _ := uf.Find(i)
		if currentRoot != root {
			t.Fatalf("Element %d has different root than element 0", i)
		}
	}

	parent, _ := uf.GetParent(50)
	if parent != root {
		t.Fatal("Path compression should have flattened the tree")
	}
}

func TestUFStringType(t *testing.T) {
	uf := NewUnionFind[string]()
	uf.MakeSet("a")
	uf.MakeSet("b")
	uf.MakeSet("c")
	uf.Union("a", "b")

	root1, _ := uf.Find("a")
	root2, _ := uf.Find("b")
	if root1 != root2 {
		t.Fatal("Strings 'a' and 'b' should be in same set")
	}
}

func TestFindSafe(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.Union(1, 2)

	root, err := uf.FindSafe(1)
	if err != nil {
		t.Fatalf("FindSafe() returned error: %v", err)
	}
	if root != 1 && root != 2 {
		t.Fatalf("Unexpected root: %d", root)
	}
}

func TestUnionSafe(t *testing.T) {
	uf := NewUnionFind[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)

	err := uf.UnionSafe(1, 2)
	if err != nil {
		t.Fatalf("UnionSafe() returned error: %v", err)
	}

	connected, _ := uf.Connected(1, 2)
	if !connected {
		t.Fatal("Elements 1 and 2 should be connected after UnionSafe")
	}
}

func BenchmarkMakeSet(b *testing.B) {
	uf := NewUnionFind[int]()
	for i := 0; i < b.N; i++ {
		uf.MakeSet(i)
	}
}

func BenchmarkFind(b *testing.B) {
	uf := NewUnionFind[int]()
	for i := 0; i < 1000; i++ {
		uf.MakeSet(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Find(i % 1000)
	}
}

func BenchmarkUnion(b *testing.B) {
	uf := NewUnionFind[int]()
	for i := 0; i < 1000; i++ {
		uf.MakeSet(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Union(i, (i+1)%1000)
	}
}

func BenchmarkConnected(b *testing.B) {
	uf := NewUnionFind[int]()
	for i := 0; i < 1000; i++ {
		uf.MakeSet(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uf.Connected(i, (i+1)%1000)
	}
}
