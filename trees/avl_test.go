package trees

import (
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree[int]()
	if tree == nil {
		t.Fatal("NewAVLTree returned nil")
	}
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
	if tree.Size() != 0 {
		t.Errorf("New tree should have size 0, got %d", tree.Size())
	}
}

func TestAVLInsert(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if tree.Size() != 3 {
		t.Errorf("Expected size 3, got %d", tree.Size())
	}
	if !tree.Search(5) || !tree.Search(3) || !tree.Search(7) {
		t.Error("Inserted values should be searchable")
	}
}

func TestAVLInsertDuplicates(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(5)
	tree.Insert(5)

	if tree.Size() != 1 {
		t.Errorf("Duplicate insertions should not increase size, got %d", tree.Size())
	}
}

func TestAVLDelete(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	deleted := tree.Delete(5)
	if !deleted {
		t.Error("Delete should return true for existing element")
	}
	if tree.Size() != 2 {
		t.Errorf("Expected size 2 after deletion, got %d", tree.Size())
	}
	if tree.Search(5) {
		t.Error("Deleted element should not be searchable")
	}
}

func TestAVLDeleteNonExistent(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)

	deleted := tree.Delete(999)
	if deleted {
		t.Error("Delete should return false for non-existent element")
	}
}

func TestAVLDeleteFromEmpty(t *testing.T) {
	tree := NewAVLTree[int]()
	deleted := tree.Delete(5)
	if deleted {
		t.Error("Delete from empty tree should return false")
	}
}

func TestAVLSearch(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	tests := []struct {
		value    int
		expected bool
	}{
		{5, true},
		{3, true},
		{7, true},
		{1, false},
		{10, false},
	}

	for _, test := range tests {
		result := tree.Search(test.value)
		if result != test.expected {
			t.Errorf("Search(%d): expected %v, got %v", test.value, test.expected, result)
		}
	}
}

func TestAVLContains(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)

	if !tree.Contains(5) {
		t.Error("Contains should return true for existing element")
	}
	if tree.Contains(10) {
		t.Error("Contains should return false for non-existent element")
	}
}

func TestAVLFindMin(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)

	min, err := tree.FindMin()
	if err != nil {
		t.Fatal(err)
	}
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}
}

func TestAVLFindMinEmpty(t *testing.T) {
	tree := NewAVLTree[int]()
	_, err := tree.FindMin()
	if err == nil {
		t.Error("FindMin on empty tree should return error")
	}
}

func TestAVLFindMax(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(10)

	max, err := tree.FindMax()
	if err != nil {
		t.Fatal(err)
	}
	if max != 10 {
		t.Errorf("Expected max 10, got %d", max)
	}
}

func TestAVLFindMaxEmpty(t *testing.T) {
	tree := NewAVLTree[int]()
	_, err := tree.FindMax()
	if err == nil {
		t.Error("FindMax on empty tree should return error")
	}
}

func TestAVLInorder(t *testing.T) {
	tree := NewAVLTree[int]()
	values := []int{5, 3, 7, 2, 4, 6, 8}
	for _, v := range values {
		tree.Insert(v)
	}

	result := tree.Inorder()
	if len(result) != 7 {
		t.Errorf("Expected 7 elements, got %d", len(result))
	}

	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			t.Errorf("Inorder should produce sorted array: %v", result)
		}
	}
}

func TestAVLPreorder(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	result := tree.Preorder()
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	if result[0] != 5 {
		t.Errorf("Preorder should start with root (5), got %d", result[0])
	}
}

func TestAVLPostorder(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	result := tree.Postorder()
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	if result[2] != 5 {
		t.Errorf("Postorder should end with root (5), got %d", result[2])
	}
}

func TestAVLLevelOrder(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	result := tree.LevelOrder()
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}
	if result[0] != 5 {
		t.Errorf("Level order should start with root (5), got %d", result[0])
	}
}

func TestAVLHeight(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	height := tree.Height()
	if height != 2 {
		t.Errorf("Expected height 2, got %d", height)
	}
}

func TestAVLHeightEmpty(t *testing.T) {
	tree := NewAVLTree[int]()
	if tree.Height() != 0 {
		t.Errorf("Empty tree should have height 0, got %d", tree.Height())
	}
}

func TestAVLClear(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after Clear")
	}
	if tree.Size() != 0 {
		t.Errorf("Size should be 0 after Clear, got %d", tree.Size())
	}
}

func TestAVLCopy(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	copied := tree.Copy()
	if copied.Size() != tree.Size() {
		t.Error("Copied tree should have same size")
	}

	tree.Delete(5)
	if !copied.Search(5) {
		t.Error("Modifying original should not affect copy - element 5 should still exist in copy")
	}

	if tree.Size() != 2 {
		t.Errorf("Original tree should have size 2 after deletion, got %d", tree.Size())
	}

	if copied.Size() != 3 {
		t.Errorf("Copied tree should still have size 3, got %d", copied.Size())
	}
}

func TestAVLToSlice(t *testing.T) {
	tree := NewAVLTree[int]()
	values := []int{5, 3, 7}
	for _, v := range values {
		tree.Insert(v)
	}

	result := tree.ToSlice()
	if len(result) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(result))
	}

	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			t.Errorf("ToSlice should return sorted array: %v", result)
		}
	}
}

func TestAVLString(t *testing.T) {
	tree := NewAVLTree[int]()
	str := tree.String()
	if str != "[]" {
		t.Errorf("Empty tree string should be '[]', got %s", str)
	}

	tree.Insert(5)
	str = tree.String()
	if str == "[]" {
		t.Error("Non-empty tree string should not be '[]'")
	}
}

func TestAVLBalancing(t *testing.T) {
	tree := NewAVLTree[int]()

	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		tree.Insert(v)
	}

	height := tree.Height()
	expectedHeight := 3

	if height > expectedHeight {
		t.Errorf("Tree should be balanced. Expected height <= %d, got %d", expectedHeight, height)
	}

	inorder := tree.Inorder()
	for i := 1; i < len(inorder); i++ {
		if inorder[i] <= inorder[i-1] {
			t.Errorf("Tree should maintain BST property after balancing: %v", inorder)
		}
	}
}

func TestAVLLargeInsertion(t *testing.T) {
	tree := NewAVLTree[int]()

	for i := 1; i <= 1000; i++ {
		tree.Insert(i)
	}

	if tree.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", tree.Size())
	}

	height := tree.Height()
	if height > 11 {
		t.Errorf("Tree height should be O(log n). Expected ~10, got %d", height)
	}

	for i := 1; i <= 1000; i++ {
		if !tree.Search(i) {
			t.Errorf("Element %d should be searchable", i)
		}
	}
}

func TestAVLMixedOperations(t *testing.T) {
	tree := NewAVLTree[int]()

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	if tree.Size() != 5 {
		t.Errorf("Expected size 5, got %d", tree.Size())
	}

	tree.Delete(5)
	if tree.Size() != 4 {
		t.Errorf("Expected size 4 after deletion, got %d", tree.Size())
	}

	if tree.Search(5) {
		t.Error("Deleted element should not be found")
	}

	if !tree.Search(3) || !tree.Search(7) {
		t.Error("Remaining elements should still be searchable")
	}
}

func TestAVLDeleteRoot(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	deleted := tree.Delete(5)
	if !deleted {
		t.Error("Delete of root should return true")
	}
	if tree.Search(5) {
		t.Error("Root should be deleted")
	}
	if tree.Size() != 2 {
		t.Errorf("Expected size 2, got %d", tree.Size())
	}
}

func TestAVLDeleteLeaf(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	deleted := tree.Delete(3)
	if !deleted {
		t.Error("Delete of leaf should return true")
	}
	if tree.Size() != 2 {
		t.Errorf("Expected size 2, got %d", tree.Size())
	}
}

func TestAVLDeleteWithRebalancing(t *testing.T) {
	tree := NewAVLTree[int]()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)

	initialHeight := tree.Height()

	tree.Delete(10)

	newHeight := tree.Height()
	if newHeight > initialHeight {
		t.Errorf("Height should not increase after deletion. Before: %d, After: %d", initialHeight, newHeight)
	}
}

func TestAVLStringType(t *testing.T) {
	tree := NewAVLTree[string]()
	tree.Insert("banana")
	tree.Insert("apple")
	tree.Insert("cherry")

	if !tree.Search("banana") {
		t.Error("String element should be searchable")
	}

	min, err := tree.FindMin()
	if err != nil {
		t.Fatal(err)
	}
	if min != "apple" {
		t.Errorf("Expected min 'apple', got %s", min)
	}
}

func TestAVLFloatType(t *testing.T) {
	tree := NewAVLTree[float64]()
	tree.Insert(3.14)
	tree.Insert(2.71)
	tree.Insert(1.41)

	min, err := tree.FindMin()
	if err != nil {
		t.Fatal(err)
	}
	if min != 1.41 {
		t.Errorf("Expected min 1.41, got %f", min)
	}
}

func TestAVLSize(t *testing.T) {
	tree := NewAVLTree[int]()

	for i := 1; i <= 10; i++ {
		tree.Insert(i)
		if tree.Size() != i {
			t.Errorf("Expected size %d, got %d", i, tree.Size())
		}
	}

	for i := 10; i >= 1; i-- {
		tree.Delete(i)
		if tree.Size() != i-1 {
			t.Errorf("Expected size %d after deletion, got %d", i-1, tree.Size())
		}
	}
}

func TestAVLIsEmpty(t *testing.T) {
	tree := NewAVLTree[int]()

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(5)
	if tree.IsEmpty() {
		t.Error("Tree with element should not be empty")
	}

	tree.Delete(5)
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after removing all elements")
	}
}

func TestAVLSingleElement(t *testing.T) {
	tree := NewAVLTree[int]()
	tree.Insert(5)

	if tree.Size() != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size())
	}
	if tree.Height() != 1 {
		t.Errorf("Expected height 1, got %d", tree.Height())
	}
	if !tree.Search(5) {
		t.Error("Single element should be searchable")
	}

	min, _ := tree.FindMin()
	max, _ := tree.FindMax()
	if min != 5 || max != 5 {
		t.Error("Min and max should both be 5 for single element")
	}
}

func TestAVLLeftRotation(t *testing.T) {
	tree := NewAVLTree[int]()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)

	if tree.Height() > 2 {
		t.Errorf("Tree should be balanced after left rotation, height: %d", tree.Height())
	}

	inorder := tree.Inorder()
	expected := []int{10, 20, 30}
	for i, v := range expected {
		if inorder[i] != v {
			t.Errorf("Expected inorder %v, got %v", expected, inorder)
		}
	}
}

func TestAVLRightRotation(t *testing.T) {
	tree := NewAVLTree[int]()

	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(10)

	if tree.Height() > 2 {
		t.Errorf("Tree should be balanced after right rotation, height: %d", tree.Height())
	}

	inorder := tree.Inorder()
	expected := []int{10, 20, 30}
	for i, v := range expected {
		if inorder[i] != v {
			t.Errorf("Expected inorder %v, got %v", expected, inorder)
		}
	}
}
