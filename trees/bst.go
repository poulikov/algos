package trees

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// BST represents a Binary Search Tree
type BST[T constraints.Ordered] struct {
	root *bstNode[T]
	size int
}

type bstNode[T constraints.Ordered] struct {
	value T
	left  *bstNode[T]
	right *bstNode[T]
}

// NewBST creates a new empty Binary Search Tree
// Time complexity: O(1)
func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Insert adds a value to the tree
// Time complexity: O(h) where h is the height of the tree (O(log n) average, O(n) worst case)
func (bst *BST[T]) Insert(value T) {
	if bst.root == nil {
		bst.root = &bstNode[T]{value: value}
		bst.size++
		return
	}

	if bst.insertRecursive(bst.root, value) {
		bst.size++
	}
}

func (bst *BST[T]) insertRecursive(node *bstNode[T], value T) bool {
	if value == node.value {
		return false
	}

	if value < node.value {
		if node.left == nil {
			node.left = &bstNode[T]{value: value}
			return true
		}
		return bst.insertRecursive(node.left, value)
	} else {
		if node.right == nil {
			node.right = &bstNode[T]{value: value}
			return true
		}
		return bst.insertRecursive(node.right, value)
	}
}

// Search checks if a value exists in the tree
// Time complexity: O(h) where h is the height of the tree (O(log n) average, O(n) worst case)
func (bst *BST[T]) Search(value T) bool {
	return bst.searchRecursive(bst.root, value)
}

func (bst *BST[T]) searchRecursive(node *bstNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	}

	if value < node.value {
		return bst.searchRecursive(node.left, value)
	}
	return bst.searchRecursive(node.right, value)
}

// Delete removes a value from the tree
// Time complexity: O(h) where h is the height of the tree
func (bst *BST[T]) Delete(value T) bool {
	if !bst.Search(value) {
		return false
	}
	bst.root = bst.deleteRecursive(bst.root, value)
	bst.size--
	return true
}

func (bst *BST[T]) deleteRecursive(node *bstNode[T], value T) *bstNode[T] {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = bst.deleteRecursive(node.left, value)
		return node
	}

	if value > node.value {
		node.right = bst.deleteRecursive(node.right, value)
		return node
	}

	if node.left == nil {
		return node.right
	}

	if node.right == nil {
		return node.left
	}

	minRight := bst.findMinNode(node.right)
	node.value = minRight.value
	node.right = bst.deleteRecursive(node.right, minRight.value)
	return node
}

func (bst *BST[T]) findMinNode(node *bstNode[T]) *bstNode[T] {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// FindMin returns the minimum value in the tree
// Time complexity: O(h)
func (bst *BST[T]) FindMin() (T, error) {
	if bst.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	node := bst.findMinNode(bst.root)
	return node.value, nil
}

// FindMax returns the maximum value in the tree
// Time complexity: O(h)
func (bst *BST[T]) FindMax() (T, error) {
	if bst.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	node := bst.root
	for node.right != nil {
		node = node.right
	}
	return node.value, nil
}

// Inorder returns a slice of values in inorder traversal (left, root, right)
// Time complexity: O(n)
func (bst *BST[T]) Inorder() []T {
	result := make([]T, 0, bst.size)
	bst.inorderRecursive(bst.root, &result)
	return result
}

func (bst *BST[T]) inorderRecursive(node *bstNode[T], result *[]T) {
	if node == nil {
		return
	}
	bst.inorderRecursive(node.left, result)
	*result = append(*result, node.value)
	bst.inorderRecursive(node.right, result)
}

// Preorder returns a slice of values in preorder traversal (root, left, right)
// Time complexity: O(n)
func (bst *BST[T]) Preorder() []T {
	result := make([]T, 0, bst.size)
	bst.preorderRecursive(bst.root, &result)
	return result
}

func (bst *BST[T]) preorderRecursive(node *bstNode[T], result *[]T) {
	if node == nil {
		return
	}
	*result = append(*result, node.value)
	bst.preorderRecursive(node.left, result)
	bst.preorderRecursive(node.right, result)
}

// Postorder returns a slice of values in postorder traversal (left, right, root)
// Time complexity: O(n)
func (bst *BST[T]) Postorder() []T {
	result := make([]T, 0, bst.size)
	bst.postorderRecursive(bst.root, &result)
	return result
}

func (bst *BST[T]) postorderRecursive(node *bstNode[T], result *[]T) {
	if node == nil {
		return
	}
	bst.postorderRecursive(node.left, result)
	bst.postorderRecursive(node.right, result)
	*result = append(*result, node.value)
}

// IsEmpty returns true if the tree is empty
// Time complexity: O(1)
func (bst *BST[T]) IsEmpty() bool {
	return bst.root == nil
}

// Size returns the number of nodes in the tree
// Time complexity: O(1)
func (bst *BST[T]) Size() int {
	return bst.size
}

// Clear removes all nodes from the tree
// Time complexity: O(1)
func (bst *BST[T]) Clear() {
	bst.root = nil
	bst.size = 0
}

// Copy creates a new BST with the same values
// Time complexity: O(n)
func (bst *BST[T]) Copy() *BST[T] {
	newBST := NewBST[T]()
	for _, value := range bst.Inorder() {
		newBST.Insert(value)
	}
	return newBST
}

// ToSlice returns a slice of all values in sorted order
// Time complexity: O(n)
func (bst *BST[T]) ToSlice() []T {
	return bst.Inorder()
}

// Height returns the height of the tree
// Time complexity: O(n)
func (bst *BST[T]) Height() int {
	return bst.heightRecursive(bst.root)
}

func (bst *BST[T]) heightRecursive(node *bstNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := bst.heightRecursive(node.left)
	rightHeight := bst.heightRecursive(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// String returns a string representation of the tree
func (bst *BST[T]) String() string {
	if bst.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", bst.Inorder())
}

// Balance rebalances the tree to achieve optimal height
// Time complexity: O(n log n)
func (bst *BST[T]) Balance() {
	values := bst.Inorder()
	bst.root = bst.buildBalanced(values, 0, len(values)-1)
	bst.size = len(values)
}

func (bst *BST[T]) buildBalanced(values []T, start, end int) *bstNode[T] {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	node := &bstNode[T]{value: values[mid]}

	node.left = bst.buildBalanced(values, start, mid-1)
	node.right = bst.buildBalanced(values, mid+1, end)

	return node
}
