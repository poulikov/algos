package trees

import (
	"testing"
)

func TestNewBST(t *testing.T) {
	bst := NewBST[int]()
	if bst == nil {
		t.Fatal("NewBST() returned nil")
	}
	if !bst.IsEmpty() {
		t.Fatal("NewBST() tree should be empty")
	}
	if bst.Size() != 0 {
		t.Fatalf("NewBST() tree should have size 0, got %d", bst.Size())
	}
}

func TestInsert(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)

	if bst.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", bst.Size())
	}
	if !bst.Search(5) || !bst.Search(3) || !bst.Search(7) {
		t.Fatal("Inserted values not found")
	}
}

func TestInsertDuplicate(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(5)
	bst.Insert(5)

	if bst.Size() != 1 {
		t.Fatalf("Expected size 1 after duplicate insert, got %d", bst.Size())
	}
}

func TestSearch(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	tests := []struct {
		value    int
		expected bool
	}{
		{10, true},
		{5, true},
		{15, true},
		{3, true},
		{7, true},
		{1, false},
		{20, false},
	}

	for _, test := range tests {
		if result := bst.Search(test.value); result != test.expected {
			t.Fatalf("Search(%d) = %v, expected %v", test.value, result, test.expected)
		}
	}
}

func TestDelete(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	if !bst.Delete(5) {
		t.Fatal("Delete(5) returned false")
	}
	if bst.Search(5) {
		t.Fatal("Value 5 still exists after deletion")
	}
	if bst.Size() != 2 {
		t.Fatalf("Expected size 2 after deletion, got %d", bst.Size())
	}
}

func TestDeleteNonExistent(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)

	if bst.Delete(15) {
		t.Fatal("Delete(15) should return false for non-existent value")
	}
	if bst.Size() != 2 {
		t.Fatalf("Size should remain 2, got %d", bst.Size())
	}
}

func TestDeleteLeaf(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)

	bst.Delete(3)
	if bst.Search(3) {
		t.Fatal("Leaf node 3 still exists")
	}
	if bst.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", bst.Size())
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(3)

	bst.Delete(5)
	if bst.Search(5) {
		t.Fatal("Node 5 with one child still exists")
	}
	if !bst.Search(3) {
		t.Fatal("Child node 3 was deleted")
	}
	if bst.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", bst.Size())
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	bst.Delete(5)
	if bst.Search(5) {
		t.Fatal("Node 5 with two children still exists")
	}
	if !bst.Search(3) || !bst.Search(7) {
		t.Fatal("Children nodes were deleted")
	}
	if bst.Size() != 4 {
		t.Fatalf("Expected size 4, got %d", bst.Size())
	}
}

func TestDeleteRoot(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	bst.Delete(10)
	if bst.Search(10) {
		t.Fatal("Root still exists")
	}
	if bst.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", bst.Size())
	}
}

func TestFindMin(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	min, err := bst.FindMin()
	if err != nil {
		t.Fatalf("FindMin() returned error: %v", err)
	}
	if min != 3 {
		t.Fatalf("Expected min 3, got %d", min)
	}
}

func TestFindMinEmpty(t *testing.T) {
	bst := NewBST[int]()
	_, err := bst.FindMin()
	if err == nil {
		t.Fatal("FindMin() should return error for empty tree")
	}
}

func TestFindMax(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	max, err := bst.FindMax()
	if err != nil {
		t.Fatalf("FindMax() returned error: %v", err)
	}
	if max != 15 {
		t.Fatalf("Expected max 15, got %d", max)
	}
}

func TestFindMaxEmpty(t *testing.T) {
	bst := NewBST[int]()
	_, err := bst.FindMax()
	if err == nil {
		t.Fatal("FindMax() should return error for empty tree")
	}
}

func TestInorder(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(7)

	inorder := bst.Inorder()
	expected := []int{3, 5, 7, 10, 15}
	if len(inorder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(inorder))
	}
	for i, v := range expected {
		if inorder[i] != v {
			t.Fatalf("Inorder[%d] = %d, expected %d", i, inorder[i], v)
		}
	}
}

func TestPreorder(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	preorder := bst.Preorder()
	expected := []int{10, 5, 15}
	if len(preorder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(preorder))
	}
	for i, v := range expected {
		if preorder[i] != v {
			t.Fatalf("Preorder[%d] = %d, expected %d", i, preorder[i], v)
		}
	}
}

func TestPostorder(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	postorder := bst.Postorder()
	expected := []int{5, 15, 10}
	if len(postorder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(postorder))
	}
	for i, v := range expected {
		if postorder[i] != v {
			t.Fatalf("Postorder[%d] = %d, expected %d", i, postorder[i], v)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	bst := NewBST[int]()
	if !bst.IsEmpty() {
		t.Fatal("New tree should be empty")
	}

	bst.Insert(5)
	if bst.IsEmpty() {
		t.Fatal("Tree with elements should not be empty")
	}
}

func TestSize(t *testing.T) {
	bst := NewBST[int]()
	if bst.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", bst.Size())
	}

	for i := 1; i <= 5; i++ {
		bst.Insert(i)
		if bst.Size() != i {
			t.Fatalf("Expected size %d, got %d", i, bst.Size())
		}
	}
}

func TestClear(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)

	bst.Clear()

	if !bst.IsEmpty() {
		t.Fatal("Clear() should make tree empty")
	}
	if bst.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", bst.Size())
	}
}

func TestCopy(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	bstCopy := bst.Copy()

	if bstCopy.Size() != bst.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", bstCopy.Size(), bst.Size())
	}

	bst.Insert(20)
	bst.Delete(10)

	if bstCopy.Search(20) {
		t.Fatal("Copy should be independent, but found new value in copy")
	}
	if !bstCopy.Search(10) {
		t.Fatal("Copy should be independent, but missing original value")
	}
}

func TestToSlice(t *testing.T) {
	bst := NewBST[int]()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)

	slice := bst.ToSlice()
	expected := []int{5, 10, 15}
	if len(slice) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(slice))
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Fatalf("ToSlice()[%d] = %d, expected %d", i, slice[i], v)
		}
	}
}

func TestHeight(t *testing.T) {
	bst := NewBST[int]()
	if bst.Height() != 0 {
		t.Fatalf("Empty tree should have height 0, got %d", bst.Height())
	}

	bst.Insert(10)
	if bst.Height() != 1 {
		t.Fatalf("Tree with root should have height 1, got %d", bst.Height())
	}

	bst.Insert(5)
	bst.Insert(15)
	if bst.Height() != 2 {
		t.Fatalf("Tree with 3 nodes should have height 2, got %d", bst.Height())
	}
}

func TestString(t *testing.T) {
	bst := NewBST[int]()
	str := bst.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	str = bst.String()
	if str != "[5 10 15]" {
		t.Fatalf("Expected '[5 10 15]', got '%s'", str)
	}
}

func TestBalance(t *testing.T) {
	bst := NewBST[int]()
	for i := 1; i <= 7; i++ {
		bst.Insert(i)
	}

	originalHeight := bst.Height()
	bst.Balance()
	balancedHeight := bst.Height()

	if balancedHeight >= originalHeight {
		t.Fatalf("Balanced tree should have less height, got %d vs %d", balancedHeight, originalHeight)
	}

	inorder := bst.Inorder()
	for i, v := range inorder {
		if v != i+1 {
			t.Fatalf("Inorder should preserve order, got %d at index %d", v, i)
		}
	}
}

func TestStringType(t *testing.T) {
	bst := NewBST[string]()
	bst.Insert("banana")
	bst.Insert("apple")
	bst.Insert("cherry")

	inorder := bst.Inorder()
	if inorder[0] != "apple" || inorder[1] != "banana" || inorder[2] != "cherry" {
		t.Fatalf("Expected alphabetical order, got %v", inorder)
	}
}

func BenchmarkInsert(b *testing.B) {
	bst := NewBST[int]()
	for i := 0; i < b.N; i++ {
		bst.Insert(i)
	}
}

func BenchmarkSearch(b *testing.B) {
	bst := NewBST[int]()
	for i := 0; i < 1000; i++ {
		bst.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Search(i % 1000)
	}
}

func BenchmarkDelete(b *testing.B) {
	bst := NewBST[int]()
	for i := 0; i < 1000; i++ {
		bst.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Insert(i + 1000)
		bst.Delete(i)
	}
}

func BenchmarkInorder(b *testing.B) {
	bst := NewBST[int]()
	for i := 0; i < 1000; i++ {
		bst.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bst.Inorder()
	}
}
