package trees

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestNewRBTree(t *testing.T) {
	rbt := NewRBTree[int]()
	if rbt == nil {
		t.Fatal("NewRBTree() returned nil")
	}
	if !rbt.IsEmpty() {
		t.Fatal("NewRBTree() tree should be empty")
	}
	if rbt.Size() != 0 {
		t.Fatalf("NewRBTree() tree should have size 0, got %d", rbt.Size())
	}
}

func TestRBInsert(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(5)
	rbt.Insert(3)
	rbt.Insert(7)

	if rbt.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", rbt.Size())
	}
	if !rbt.Search(5) || !rbt.Search(3) || !rbt.Search(7) {
		t.Fatal("Inserted values not found")
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after insert")
	}
}

func TestRBInsertDuplicate(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(5)
	rbt.Insert(5)

	if rbt.Size() != 1 {
		t.Fatalf("Expected size 1 after duplicate insert, got %d", rbt.Size())
	}
}

func TestRBSearch(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)
	rbt.Insert(7)

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
		if result := rbt.Search(test.value); result != test.expected {
			t.Fatalf("Search(%d) = %v, expected %v", test.value, result, test.expected)
		}
	}
}

func TestRBDelete(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	if !rbt.Delete(5) {
		t.Fatal("Delete(5) returned false")
	}
	if rbt.Search(5) {
		t.Fatal("Value 5 still exists after deletion")
	}
	if rbt.Size() != 2 {
		t.Fatalf("Expected size 2 after deletion, got %d", rbt.Size())
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after delete")
	}
}

func TestRBDeleteNonExistent(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)

	if rbt.Delete(15) {
		t.Fatal("Delete(15) should return false for non-existent value")
	}
	if rbt.Size() != 2 {
		t.Fatalf("Size should remain 2, got %d", rbt.Size())
	}
}

func TestRBDeleteLeaf(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)

	rbt.Delete(3)
	if rbt.Search(3) {
		t.Fatal("Leaf node 3 still exists")
	}
	if rbt.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", rbt.Size())
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after leaf deletion")
	}
}

func TestRBDeleteNodeWithOneChild(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(3)

	rbt.Delete(5)
	if rbt.Search(5) {
		t.Fatal("Node 5 with one child still exists")
	}
	if !rbt.Search(3) {
		t.Fatal("Child node 3 was deleted")
	}
	if rbt.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", rbt.Size())
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after deletion")
	}
}

func TestRBDeleteNodeWithTwoChildren(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)
	rbt.Insert(7)

	rbt.Delete(5)
	if rbt.Search(5) {
		t.Fatal("Node 5 with two children still exists")
	}
	if !rbt.Search(3) || !rbt.Search(7) {
		t.Fatal("Children nodes were deleted")
	}
	if rbt.Size() != 4 {
		t.Fatalf("Expected size 4, got %d", rbt.Size())
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after deletion")
	}
}

func TestRBDeleteRoot(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	rbt.Delete(10)
	if rbt.Search(10) {
		t.Fatal("Root still exists")
	}
	if rbt.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", rbt.Size())
	}
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after root deletion")
	}
}

func TestRBFindMin(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)
	rbt.Insert(7)

	min, err := rbt.FindMin()
	if err != nil {
		t.Fatalf("FindMin() returned error: %v", err)
	}
	if min != 3 {
		t.Fatalf("Expected min 3, got %d", min)
	}
}

func TestRBFindMinEmpty(t *testing.T) {
	rbt := NewRBTree[int]()
	_, err := rbt.FindMin()
	if err == nil {
		t.Fatal("FindMin() should return error for empty tree")
	}
}

func TestRBFindMax(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)
	rbt.Insert(7)

	max, err := rbt.FindMax()
	if err != nil {
		t.Fatalf("FindMax() returned error: %v", err)
	}
	if max != 15 {
		t.Fatalf("Expected max 15, got %d", max)
	}
}

func TestRBFindMaxEmpty(t *testing.T) {
	rbt := NewRBTree[int]()
	_, err := rbt.FindMax()
	if err == nil {
		t.Fatal("FindMax() should return error for empty tree")
	}
}

func TestRBInorder(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	rbt.Insert(3)
	rbt.Insert(7)

	inorder := rbt.Inorder()
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

func TestRBPreorder(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	preorder := rbt.Preorder()
	if len(preorder) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(preorder))
	}
}

func TestRBPostorder(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	postorder := rbt.Postorder()
	if len(postorder) != 3 {
		t.Fatalf("Expected 3 elements, got %d", len(postorder))
	}
}

func TestRBIsEmpty(t *testing.T) {
	rbt := NewRBTree[int]()
	if !rbt.IsEmpty() {
		t.Fatal("New tree should be empty")
	}

	rbt.Insert(5)
	if rbt.IsEmpty() {
		t.Fatal("Tree with elements should not be empty")
	}
}

func TestRBSize(t *testing.T) {
	rbt := NewRBTree[int]()
	if rbt.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", rbt.Size())
	}

	for i := 1; i <= 5; i++ {
		rbt.Insert(i)
		if rbt.Size() != i {
			t.Fatalf("Expected size %d, got %d", i, rbt.Size())
		}
	}
}

func TestRBClear(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(1)
	rbt.Insert(2)
	rbt.Insert(3)

	rbt.Clear()

	if !rbt.IsEmpty() {
		t.Fatal("Clear() should make tree empty")
	}
	if rbt.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", rbt.Size())
	}
}

func TestRBCopy(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	rbtCopy := rbt.Copy()

	if rbtCopy.Size() != rbt.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", rbtCopy.Size(), rbt.Size())
	}

	rbt.Insert(20)
	rbt.Delete(10)

	if rbtCopy.Search(20) {
		t.Fatal("Copy should be independent, but found new value in copy")
	}
	if !rbtCopy.Search(10) {
		t.Fatal("Copy should be independent, but missing original value")
	}
}

func TestRBToSlice(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)

	slice := rbt.ToSlice()
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

func TestRBHeight(t *testing.T) {
	rbt := NewRBTree[int]()
	if rbt.Height() != 0 {
		t.Fatalf("Empty tree should have height 0, got %d", rbt.Height())
	}

	rbt.Insert(10)
	if rbt.Height() != 1 {
		t.Fatalf("Tree with root should have height 1, got %d", rbt.Height())
	}

	rbt.Insert(5)
	rbt.Insert(15)
	if rbt.Height() != 2 {
		t.Fatalf("Tree with 3 nodes should have height 2, got %d", rbt.Height())
	}
}

func TestRBString(t *testing.T) {
	rbt := NewRBTree[int]()
	str := rbt.String()
	if str != "[]" {
		t.Fatalf("Expected '[]', got '%s'", str)
	}

	rbt.Insert(10)
	rbt.Insert(5)
	rbt.Insert(15)
	str = rbt.String()
	if str != "[5 10 15]" {
		t.Fatalf("Expected '[5 10 15]', got '%s'", str)
	}
}

func TestRBContains(t *testing.T) {
	rbt := NewRBTree[int]()
	rbt.Insert(10)
	rbt.Insert(5)

	if !rbt.Contains(10) {
		t.Fatal("Contains() should return true for existing value")
	}
	if rbt.Contains(15) {
		t.Fatal("Contains() should return false for non-existing value")
	}
}

func TestRBStringType(t *testing.T) {
	rbt := NewRBTree[string]()
	rbt.Insert("banana")
	rbt.Insert("apple")
	rbt.Insert("cherry")

	inorder := rbt.Inorder()
	if inorder[0] != "apple" || inorder[1] != "banana" || inorder[2] != "cherry" {
		t.Fatalf("Expected alphabetical order, got %v", inorder)
	}
}

func TestRBTreeBalancing(t *testing.T) {
	rbt := NewRBTree[int]()

	for i := 1; i <= 100; i++ {
		rbt.Insert(i)
	}

	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after sequential insert")
	}

	height := rbt.Height()
	if height > 20 {
		t.Fatalf("Tree height %d is too large for 100 elements (unbalanced)", height)
	}
}

func TestRBTreeComplexOperations(t *testing.T) {
	rbt := NewRBTree[int]()

	values := []int{7, 3, 18, 10, 22, 8, 11, 26}
	for _, v := range values {
		rbt.Insert(v)
	}

	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after initial insertions")
	}

	rbt.Delete(18)
	if !validateRBTree(rbt) {
		t.Fatal("Tree violates red-black properties after deleting node with two children")
	}

	inorder := rbt.Inorder()
	expected := []int{3, 7, 8, 10, 11, 22, 26}
	if len(inorder) != len(expected) {
		t.Fatalf("Expected %d elements, got %d", len(expected), len(inorder))
	}
	for i, v := range expected {
		if inorder[i] != v {
			t.Fatalf("Inorder[%d] = %d, expected %d", i, inorder[i], v)
		}
	}
}

func BenchmarkRBInsert(b *testing.B) {
	rbt := NewRBTree[int]()
	for i := 0; i < b.N; i++ {
		rbt.Insert(i)
	}
}

func BenchmarkRBSearch(b *testing.B) {
	rbt := NewRBTree[int]()
	for i := 0; i < 1000; i++ {
		rbt.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Search(i % 1000)
	}
}

func BenchmarkRBDelete(b *testing.B) {
	rbt := NewRBTree[int]()
	for i := 0; i < 1000; i++ {
		rbt.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Insert(i + 1000)
		rbt.Delete(i)
	}
}

func BenchmarkRBInorder(b *testing.B) {
	rbt := NewRBTree[int]()
	for i := 0; i < 1000; i++ {
		rbt.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rbt.Inorder()
	}
}

func validateRBTree[T constraints.Ordered](rbt *RBTree[T]) bool {
	if rbt.IsEmpty() {
		return true
	}

	if rbt.root.color != BLACK {
		return false
	}

	_, valid := validateNodeProperties(rbt.root, 0)
	return valid
}

func validateNodeProperties[T constraints.Ordered](node *rbNode[T], blackCount int) (int, bool) {
	if node == nil {
		return blackCount, true
	}

	if node.color == RED {
		if (node.left != nil && node.left.color == RED) || (node.right != nil && node.right.color == RED) {
			return 0, false
		}
	} else {
		blackCount++
	}

	leftBlackCount, leftValid := validateNodeProperties(node.left, blackCount)
	if !leftValid {
		return 0, false
	}

	rightBlackCount, rightValid := validateNodeProperties(node.right, blackCount)
	if !rightValid {
		return 0, false
	}

	if leftBlackCount != rightBlackCount {
		return 0, false
	}

	return leftBlackCount, true
}
