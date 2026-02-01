package trees

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AVLTree represents a self-balancing binary search tree
// with guaranteed height of O(log n)
type AVLTree[T constraints.Ordered] struct {
	root *avlNode[T]
	size int
}

type avlNode[T constraints.Ordered] struct {
	value  T
	left   *avlNode[T]
	right  *avlNode[T]
	height int
}

// NewAVLTree creates a new empty AVL tree
// Time complexity: O(1)
func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{}
}

// Insert adds a value to the AVL tree
// Time complexity: O(log n)
func (avl *AVLTree[T]) Insert(value T) {
	avl.root = avl.insert(avl.root, value)
}

func (avl *AVLTree[T]) insert(node *avlNode[T], value T) *avlNode[T] {
	if node == nil {
		avl.size++
		return &avlNode[T]{value: value, height: 1}
	}

	if value < node.value {
		node.left = avl.insert(node.left, value)
	} else if value > node.value {
		node.right = avl.insert(node.right, value)
	} else {
		return node
	}

	node.height = 1 + max(avl.height(node.left), avl.height(node.right))

	return avl.balance(node)
}

// Delete removes a value from the AVL tree
// Time complexity: O(log n)
func (avl *AVLTree[T]) Delete(value T) bool {
	if !avl.Search(value) {
		return false
	}
	avl.root = avl.delete(avl.root, value)
	avl.size--
	return true
}

func (avl *AVLTree[T]) delete(node *avlNode[T], value T) *avlNode[T] {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = avl.delete(node.left, value)
	} else if value > node.value {
		node.right = avl.delete(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}

		minNode := avl.findMin(node.right)
		node.value = minNode.value
		node.right = avl.delete(node.right, minNode.value)
	}

	node.height = 1 + max(avl.height(node.left), avl.height(node.right))

	return avl.balance(node)
}

// Search checks if a value exists in the AVL tree
// Time complexity: O(log n)
func (avl *AVLTree[T]) Search(value T) bool {
	return avl.search(avl.root, value)
}

func (avl *AVLTree[T]) search(node *avlNode[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.value {
		return true
	}

	if value < node.value {
		return avl.search(node.left, value)
	}
	return avl.search(node.right, value)
}

// FindMin returns the minimum value in the tree
// Time complexity: O(log n)
func (avl *AVLTree[T]) FindMin() (T, error) {
	if avl.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	return avl.findMin(avl.root).value, nil
}

func (avl *AVLTree[T]) findMin(node *avlNode[T]) *avlNode[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// FindMax returns the maximum value in the tree
// Time complexity: O(log n)
func (avl *AVLTree[T]) FindMax() (T, error) {
	if avl.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}
	return avl.findMax(avl.root).value, nil
}

func (avl *AVLTree[T]) findMax(node *avlNode[T]) *avlNode[T] {
	for node.right != nil {
		node = node.right
	}
	return node
}

// Inorder returns a slice of values in inorder traversal (left, root, right)
// Time complexity: O(n)
func (avl *AVLTree[T]) Inorder() []T {
	result := make([]T, 0, avl.size)
	avl.inorder(avl.root, &result)
	return result
}

func (avl *AVLTree[T]) inorder(node *avlNode[T], result *[]T) {
	if node == nil {
		return
	}
	avl.inorder(node.left, result)
	*result = append(*result, node.value)
	avl.inorder(node.right, result)
}

// Preorder returns a slice of values in preorder traversal (root, left, right)
// Time complexity: O(n)
func (avl *AVLTree[T]) Preorder() []T {
	result := make([]T, 0, avl.size)
	avl.preorder(avl.root, &result)
	return result
}

func (avl *AVLTree[T]) preorder(node *avlNode[T], result *[]T) {
	if node == nil {
		return
	}
	*result = append(*result, node.value)
	avl.preorder(node.left, result)
	avl.preorder(node.right, result)
}

// Postorder returns a slice of values in postorder traversal (left, right, root)
// Time complexity: O(n)
func (avl *AVLTree[T]) Postorder() []T {
	result := make([]T, 0, avl.size)
	avl.postorder(avl.root, &result)
	return result
}

func (avl *AVLTree[T]) postorder(node *avlNode[T], result *[]T) {
	if node == nil {
		return
	}
	avl.postorder(node.left, result)
	avl.postorder(node.right, result)
	*result = append(*result, node.value)
}

// Height returns the height of the tree
// Time complexity: O(1)
func (avl *AVLTree[T]) Height() int {
	return avl.height(avl.root)
}

func (avl *AVLTree[T]) height(node *avlNode[T]) int {
	if node == nil {
		return 0
	}
	return node.height
}

// IsEmpty returns true if the tree is empty
// Time complexity: O(1)
func (avl *AVLTree[T]) IsEmpty() bool {
	return avl.root == nil
}

// Size returns the number of nodes in the tree
// Time complexity: O(1)
func (avl *AVLTree[T]) Size() int {
	return avl.size
}

// Clear removes all nodes from the tree
// Time complexity: O(1)
func (avl *AVLTree[T]) Clear() {
	avl.root = nil
	avl.size = 0
}

// ToSlice returns a slice of all values in sorted order
// Time complexity: O(n)
func (avl *AVLTree[T]) ToSlice() []T {
	return avl.Inorder()
}

// String returns a string representation of the tree
func (avl *AVLTree[T]) String() string {
	if avl.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", avl.Inorder())
}

// GetBalanceFactor returns the balance factor of a node
// Time complexity: O(1)
func (avl *AVLTree[T]) GetBalanceFactor(node *avlNode[T]) int {
	if node == nil {
		return 0
	}
	return avl.height(node.left) - avl.height(node.right)
}

// balance balances the subtree rooted at the given node
func (avl *AVLTree[T]) balance(node *avlNode[T]) *avlNode[T] {
	balanceFactor := avl.GetBalanceFactor(node)

	if balanceFactor > 1 {
		if avl.GetBalanceFactor(node.left) < 0 {
			node.left = avl.rotateLeft(node.left)
		}
		return avl.rotateRight(node)
	}

	if balanceFactor < -1 {
		if avl.GetBalanceFactor(node.right) > 0 {
			node.right = avl.rotateRight(node.right)
		}
		return avl.rotateLeft(node)
	}

	return node
}

// rotateRight performs a right rotation on the node
func (avl *AVLTree[T]) rotateRight(y *avlNode[T]) *avlNode[T] {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = 1 + max(avl.height(y.left), avl.height(y.right))
	x.height = 1 + max(avl.height(x.left), avl.height(x.right))

	return x
}

// rotateLeft performs a left rotation on the node
func (avl *AVLTree[T]) rotateLeft(x *avlNode[T]) *avlNode[T] {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = 1 + max(avl.height(x.left), avl.height(x.right))
	y.height = 1 + max(avl.height(y.left), avl.height(y.right))

	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Copy creates a new AVL tree with the same values
// Time complexity: O(n log n)
func (avl *AVLTree[T]) Copy() *AVLTree[T] {
	newAVL := NewAVLTree[T]()
	for _, value := range avl.Inorder() {
		newAVL.Insert(value)
	}
	return newAVL
}

// LevelOrder returns a slice of values in level order (BFS traversal)
// Time complexity: O(n)
func (avl *AVLTree[T]) LevelOrder() []T {
	if avl.IsEmpty() {
		return []T{}
	}

	result := make([]T, 0, avl.size)
	queue := []*avlNode[T]{avl.root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return result
}

// Contains checks if a value exists (alias for Search)
// Time complexity: O(log n)
func (avl *AVLTree[T]) Contains(value T) bool {
	return avl.Search(value)
}
