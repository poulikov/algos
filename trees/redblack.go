package trees

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

const (
	RED   = true
	BLACK = false
)

// RBTree represents a Red-Black Tree
// A self-balancing binary search tree with guaranteed O(log n) operations
type RBTree[T constraints.Ordered] struct {
	root *rbNode[T]
	size int
}

type rbNode[T constraints.Ordered] struct {
	value  T
	left   *rbNode[T]
	right  *rbNode[T]
	color  bool
	parent *rbNode[T]
}

// NewRBTree creates a new empty Red-Black Tree
func NewRBTree[T constraints.Ordered]() *RBTree[T] {
	return &RBTree[T]{}
}

// Insert adds a value to the Red-Black Tree
// Time complexity: O(log n)
func (rbt *RBTree[T]) Insert(value T) {
	newNode := &rbNode[T]{
		value: value,
		color: RED,
	}

	if rbt.insertNode(newNode) {
		rbt.size++
	}
}

// insertNode performs the actual insertion and fixes the tree
// Returns true if the node was inserted, false if it was a duplicate
func (rbt *RBTree[T]) insertNode(node *rbNode[T]) bool {
	var current, parent *rbNode[T]

	current = rbt.root
	parent = nil

	for current != nil {
		parent = current
		if node.value < current.value {
			current = current.left
		} else if node.value > current.value {
			current = current.right
		} else {
			return false
		}
	}

	node.parent = parent
	node.left = nil
	node.right = nil

	if parent == nil {
		rbt.root = node
	} else if node.value < parent.value {
		parent.left = node
	} else {
		parent.right = node
	}

	rbt.insertFixup(node)
	return true
}

// insertFixup fixes the Red-Black Tree properties after insertion
func (rbt *RBTree[T]) insertFixup(node *rbNode[T]) {
	for node.parent != nil && node.parent.color == RED {
		if node.parent.parent == nil {
			break
		}

		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right

			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					rbt.rotateLeft(node)
				}
				if node.parent == nil || node.parent.parent == nil {
					break
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbt.rotateRight(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left

			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					rbt.rotateRight(node)
				}
				if node.parent == nil || node.parent.parent == nil {
					break
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rbt.rotateLeft(node.parent.parent)
			}
		}
	}

	rbt.root.color = BLACK
}

// rotateLeft performs a left rotation on a node
func (rbt *RBTree[T]) rotateLeft(node *rbNode[T]) {
	rightChild := node.right

	node.right = rightChild.left
	if rightChild.left != nil {
		rightChild.left.parent = node
	}

	nodeParent := node.parent
	rightChild.parent = nodeParent
	node.parent = rightChild
	rightChild.left = node

	if nodeParent == nil {
		rbt.root = rightChild
	} else if nodeParent.left == node {
		nodeParent.left = rightChild
	} else {
		nodeParent.right = rightChild
	}
}

// rotateRight performs a right rotation on a node
func (rbt *RBTree[T]) rotateRight(node *rbNode[T]) {
	leftChild := node.left

	node.left = leftChild.right
	if leftChild.right != nil {
		leftChild.right.parent = node
	}

	nodeParent := node.parent
	leftChild.parent = nodeParent
	node.parent = leftChild
	leftChild.right = node

	if nodeParent == nil {
		rbt.root = leftChild
	} else if nodeParent.left == node {
		nodeParent.left = leftChild
	} else {
		nodeParent.right = leftChild
	}
}

// Search checks if a value exists in the Red-Black Tree
// Time complexity: O(log n)
func (rbt *RBTree[T]) Search(value T) bool {
	return rbt.Find(value) != nil
}

// Find returns the node containing the value, or nil if not found
func (rbt *RBTree[T]) Find(value T) *rbNode[T] {
	node := rbt.root

	for node != nil {
		if value == node.value {
			return node
		} else if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}

	return nil
}

// FindMin returns the minimum value in the tree
func (rbt *RBTree[T]) FindMin() (T, error) {
	if rbt.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}

	node := rbt.root
	for node.left != nil {
		node = node.left
	}

	return node.value, nil
}

// FindMax returns the maximum value in the tree
func (rbt *RBTree[T]) FindMax() (T, error) {
	if rbt.root == nil {
		var zero T
		return zero, fmt.Errorf("tree is empty")
	}

	node := rbt.root
	for node.right != nil {
		node = node.right
	}

	return node.value, nil
}

// Delete removes a value from the Red-Black Tree
// Time complexity: O(log n)
func (rbt *RBTree[T]) Delete(value T) bool {
	node := rbt.Find(value)
	if node == nil {
		return false
	}

	rbt.deleteNode(node)
	rbt.size--
	return true
}

func (rbt *RBTree[T]) deleteFixupWithParent(parent *rbNode[T], isLeftChild bool) {
	if parent == nil {
		return
	}

	for {
		var sibling *rbNode[T]
		if isLeftChild {
			sibling = parent.right
		} else {
			sibling = parent.left
		}

		if sibling != nil && sibling.color == RED {
			sibling.color = BLACK
			parent.color = RED
			if isLeftChild {
				rbt.rotateLeft(parent)
				sibling = parent.right
			} else {
				rbt.rotateRight(parent)
				sibling = parent.left
			}
		}

		siblingLeftBlack := sibling == nil || sibling.left == nil || sibling.left.color == BLACK
		siblingRightBlack := sibling == nil || sibling.right == nil || sibling.right.color == BLACK

		if siblingLeftBlack && siblingRightBlack {
			if sibling != nil {
				sibling.color = RED
			}
			isLeftChild = parent.parent != nil && parent == parent.parent.left
			parent = parent.parent
			if parent == nil || parent.color == RED {
				break
			}
		} else {
			if isLeftChild {
				if siblingRightBlack {
					if sibling != nil && sibling.left != nil {
						sibling.left.color = BLACK
					}
					if sibling != nil {
						sibling.color = RED
					}
					rbt.rotateRight(sibling)
					sibling = parent.right
				}
				if sibling != nil {
					sibling.color = parent.color
				}
				parent.color = BLACK
				if sibling != nil && sibling.right != nil {
					sibling.right.color = BLACK
				}
				rbt.rotateLeft(parent)
			} else {
				if siblingLeftBlack {
					if sibling != nil && sibling.right != nil {
						sibling.right.color = BLACK
					}
					if sibling != nil {
						sibling.color = RED
					}
					rbt.rotateLeft(sibling)
					sibling = parent.left
				}
				if sibling != nil {
					sibling.color = parent.color
				}
				parent.color = BLACK
				if sibling != nil && sibling.left != nil {
					sibling.left.color = BLACK
				}
				rbt.rotateRight(parent)
			}
			break
		}
	}

	if rbt.root != nil {
		rbt.root.color = BLACK
	}
}

// deleteNode performs the actual deletion and fixes the tree
func (rbt *RBTree[T]) deleteNode(node *rbNode[T]) {
	var y *rbNode[T]
	var x *rbNode[T]
	var yOriginalColor bool
	var yParent *rbNode[T]
	var yIsLeftChild bool

	if node.parent != nil {
		yParent = node.parent
		yIsLeftChild = node == node.parent.left
	}

	y = node
	yOriginalColor = y.color

	if node.left == nil {
		x = node.right
		rbt.transplant(node, node.right)
	} else if node.right == nil {
		x = node.left
		rbt.transplant(node, node.left)
	} else {
		y = rbt.minimumNode(node.right)
		yOriginalColor = y.color
		x = y.right
		yParent = y.parent
		yIsLeftChild = y == y.parent.left

		if y.parent != node {
			rbt.transplant(y, y.right)
			y.right = node.right
			if y.right != nil {
				y.right.parent = y
			}
		}

		rbt.transplant(node, y)
		y.left = node.left
		if y.left != nil {
			y.left.parent = y
		}
		y.color = node.color
	}

	if yOriginalColor == BLACK {
		var parent *rbNode[T]
		var isLeftChild bool

		if x == nil {
			parent = yParent
			isLeftChild = yIsLeftChild
		} else {
			parent = x.parent
			isLeftChild = x.parent != nil && x == x.parent.left
		}

		if parent != nil {
			rbt.deleteFixupWithParent(parent, isLeftChild)
		}
	}
}

// minimumNode returns the minimum node in a subtree
func (rbt *RBTree[T]) minimumNode(node *rbNode[T]) *rbNode[T] {
	for node.left != nil {
		node = node.left
	}
	return node
}

// transplant replaces subtree u with subtree v
func (rbt *RBTree[T]) transplant(u, v *rbNode[T]) {
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

// isRed checks if a node is RED
func (rbt *RBTree[T]) isRed(node *rbNode[T]) bool {
	return node != nil && node.color == RED
}

// Inorder returns a slice of values in inorder traversal
func (rbt *RBTree[T]) Inorder() []T {
	result := make([]T, 0, rbt.size)
	rbt.inorderHelper(rbt.root, &result)
	return result
}

func (rbt *RBTree[T]) inorderHelper(node *rbNode[T], result *[]T) {
	if node == nil {
		return
	}

	rbt.inorderHelper(node.left, result)
	*result = append(*result, node.value)
	rbt.inorderHelper(node.right, result)
}

// Preorder returns a slice of values in preorder traversal
func (rbt *RBTree[T]) Preorder() []T {
	result := make([]T, 0, rbt.size)
	rbt.preorderHelper(rbt.root, &result)
	return result
}

func (rbt *RBTree[T]) preorderHelper(node *rbNode[T], result *[]T) {
	if node == nil {
		return
	}

	*result = append(*result, node.value)
	rbt.preorderHelper(node.left, result)
	rbt.preorderHelper(node.right, result)
}

// Postorder returns a slice of values in postorder traversal
func (rbt *RBTree[T]) Postorder() []T {
	result := make([]T, 0, rbt.size)
	rbt.postorderHelper(rbt.root, &result)
	return result
}

func (rbt *RBTree[T]) postorderHelper(node *rbNode[T], result *[]T) {
	if node == nil {
		return
	}

	rbt.postorderHelper(node.left, result)
	rbt.postorderHelper(node.right, result)
	*result = append(*result, node.value)
}

// Size returns the number of nodes in the tree
func (rbt *RBTree[T]) Size() int {
	return rbt.size
}

// IsEmpty returns true if the tree is empty
func (rbt *RBTree[T]) IsEmpty() bool {
	return rbt.root == nil
}

// Clear removes all nodes from the tree
func (rbt *RBTree[T]) Clear() {
	rbt.root = nil
	rbt.size = 0
}

// ToSlice returns a slice of all values in sorted order
func (rbt *RBTree[T]) ToSlice() []T {
	return rbt.Inorder()
}

// String returns a string representation of the tree
func (rbt *RBTree[T]) String() string {
	if rbt.IsEmpty() {
		return "[]"
	}
	return fmt.Sprintf("%v", rbt.Inorder())
}

// Contains checks if a value exists (alias for Search)
func (rbt *RBTree[T]) Contains(value T) bool {
	return rbt.Search(value)
}

// Height returns the height of the tree
func (rbt *RBTree[T]) Height() int {
	return rbt.heightHelper(rbt.root)
}

func (rbt *RBTree[T]) heightHelper(node *rbNode[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := rbt.heightHelper(node.left)
	rightHeight := rbt.heightHelper(node.right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Copy creates a new Red-Black Tree with the same values
func (rbt *RBTree[T]) Copy() *RBTree[T] {
	newTree := NewRBTree[T]()
	for _, value := range rbt.Inorder() {
		newTree.Insert(value)
	}
	return newTree
}
