package trees

import (
	"sort"
	"testing"
)

func TestNewBTree(t *testing.T) {
	bt := NewBTree(3)

	if !bt.IsEmpty() {
		t.Errorf("New B-Tree should be empty")
	}

	if bt.Size() != 0 {
		t.Errorf("New B-Tree size should be 0, got %d", bt.Size())
	}
}

func TestBTreeInsert(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(5)

	if bt.Size() != 3 {
		t.Errorf("Expected size 3, got %d", bt.Size())
	}

	if !bt.Search(10) {
		t.Errorf("Should contain 10")
	}

	if !bt.Search(20) {
		t.Errorf("Should contain 20")
	}

	if !bt.Search(5) {
		t.Errorf("Should contain 5")
	}
}

func TestBTreeSearch(t *testing.T) {
	bt := NewBTree(3)

	values := []int{10, 20, 5, 15, 25, 30}

	for _, v := range values {
		bt.Insert(v)
	}

	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Should contain %d", v)
		}
	}

	if bt.Search(100) {
		t.Errorf("Should not contain 100")
	}
}

func TestBTreeDelete(t *testing.T) {
	bt := NewBTree(3)

	values := []int{10, 20, 5, 15, 25, 30}

	for _, v := range values {
		bt.Insert(v)
	}

	bt.Delete(10)

	if bt.Search(10) {
		t.Errorf("Should not contain 10 after deletion")
	}

	if bt.Size() != 5 {
		t.Errorf("Expected size 5 after deletion, got %d", bt.Size())
	}
}

func TestBTreeDeleteNonExistent(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)
	bt.Insert(20)

	deleted := bt.Delete(30)

	if deleted {
		t.Errorf("Delete of non-existent key should return false")
	}

	if bt.Size() != 2 {
		t.Errorf("Size should remain 2, got %d", bt.Size())
	}
}

func TestBTreeInOrderTraversal(t *testing.T) {
	bt := NewBTree(3)

	values := []int{10, 20, 5, 15, 25, 30}

	for _, v := range values {
		bt.Insert(v)
	}

	traversal := bt.InOrderTraversal()

	if len(traversal) != len(values) {
		t.Errorf("Traversal length %d != values length %d", len(traversal), len(values))
	}

	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)

	for i := range traversal {
		if traversal[i] != sorted[i] {
			t.Errorf("Traversal not sorted: %v, expected %v", traversal, sorted)
		}
	}
}

func TestBTreeMinMax(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(5)
	bt.Insert(15)
	bt.Insert(25)

	if bt.Min() != 5 {
		t.Errorf("Min should be 5, got %d", bt.Min())
	}

	if bt.Max() != 25 {
		t.Errorf("Max should be 25, got %d", bt.Max())
	}
}

func TestBTreeClear(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)
	bt.Insert(20)
	bt.Insert(5)

	bt.Clear()

	if !bt.IsEmpty() {
		t.Errorf("Should be empty after clear")
	}

	if bt.Size() != 0 {
		t.Errorf("Size should be 0 after clear, got %d", bt.Size())
	}

	if bt.Search(10) {
		t.Errorf("Should not contain 10 after clear")
	}
}

func TestBTreeLarge(t *testing.T) {
	bt := NewBTree(4)

	n := 100
	for i := 0; i < n; i++ {
		bt.Insert(i)
	}

	if bt.Size() != n {
		t.Errorf("Expected size %d, got %d", n, bt.Size())
	}

	traversal := bt.InOrderTraversal()
	if len(traversal) != n {
		t.Errorf("Traversal length %d != n %d", len(traversal), n)
	}

	for i := 0; i < n; i++ {
		if !bt.Search(i) {
			t.Errorf("Should contain %d", i)
		}
	}
}

func TestBTreeDeleteAll(t *testing.T) {
	bt := NewBTree(3)

	values := []int{10, 20, 5, 15, 25, 30}

	for _, v := range values {
		bt.Insert(v)
	}

	for _, v := range values {
		bt.Delete(v)
	}

	if !bt.IsEmpty() {
		t.Errorf("Should be empty after deleting all")
	}

	if bt.Size() != 0 {
		t.Errorf("Size should be 0 after deleting all, got %d", bt.Size())
	}
}

func TestBTreeDuplicateInsert(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)
	bt.Insert(10)
	bt.Insert(10)

	if bt.Size() != 1 {
		t.Errorf("Size should be 1 (no duplicates), got %d", bt.Size())
	}
}

func TestBTreeConsistency(t *testing.T) {
	bt := NewBTree(4)

	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, v := range values {
		bt.Insert(v)
	}

	for _, v := range values {
		if !bt.Search(v) {
			t.Errorf("Should contain %d", v)
		}
	}

	traversal := bt.InOrderTraversal()
	sorted := make([]int, len(values))
	copy(sorted, values)
	sort.Ints(sorted)

	for i := range traversal {
		if traversal[i] != sorted[i] {
			t.Errorf("Traversal mismatch at %d: got %d, expected %d", i, traversal[i], sorted[i])
		}
	}

	if bt.Min() != sorted[0] {
		t.Errorf("Min mismatch: got %d, expected %d", bt.Min(), sorted[0])
	}

	if bt.Max() != sorted[len(sorted)-1] {
		t.Errorf("Max mismatch: got %d, expected %d", bt.Max(), sorted[len(sorted)-1])
	}
}

func TestBTreeMinMaxEmpty(t *testing.T) {
	bt := NewBTree(3)

	if bt.Min() != 0 {
		t.Errorf("Min of empty tree should be 0, got %d", bt.Min())
	}

	if bt.Max() != 0 {
		t.Errorf("Max of empty tree should be 0, got %d", bt.Max())
	}
}

func TestBTreeSingleNode(t *testing.T) {
	bt := NewBTree(3)

	bt.Insert(10)

	if !bt.Search(10) {
		t.Errorf("Should contain 10")
	}

	if bt.Min() != 10 {
		t.Errorf("Min should be 10, got %d", bt.Min())
	}

	if bt.Max() != 10 {
		t.Errorf("Max should be 10, got %d", bt.Max())
	}

	bt.Delete(10)

	if !bt.IsEmpty() {
		t.Errorf("Should be empty after deleting only node")
	}
}
