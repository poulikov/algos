package trees

import (
	"fmt"
)

type BTreeNode struct {
	keys     []int
	children []*BTreeNode
	isLeaf   bool
}

type BTree struct {
	root *BTreeNode
	t    int
}

func NewBTree(t int) *BTree {
	return &BTree{
		t: t,
	}
}

func (bt *BTree) Insert(key int) {
	if bt.Search(key) {
		return
	}

	if bt.root == nil {
		bt.root = &BTreeNode{
			keys:     []int{key},
			children: nil,
			isLeaf:   true,
		}
		return
	}

	if len(bt.root.keys) == 2*bt.t-1 {
		newRoot := &BTreeNode{
			keys:     []int{},
			children: []*BTreeNode{bt.root},
			isLeaf:   false,
		}
		bt.splitChild(newRoot, 0)
		bt.root = newRoot
		bt.insertNonFull(newRoot, key)
	} else {
		bt.insertNonFull(bt.root, key)
	}
}

func (bt *BTree) insertNonFull(node *BTreeNode, key int) {
	if node.isLeaf {
		i := len(node.keys) - 1
		node.keys = append(node.keys, 0)

		for i >= 0 && key < node.keys[i] {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = key
	} else {
		i := len(node.keys) - 1
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		if len(node.children[i].keys) == 2*bt.t-1 {
			bt.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		bt.insertNonFull(node.children[i], key)
	}
}

func (bt *BTree) splitChild(parent *BTreeNode, i int) {
	t := bt.t
	y := parent.children[i]
	z := &BTreeNode{
		keys:     make([]int, t-1),
		children: make([]*BTreeNode, t),
		isLeaf:   y.isLeaf,
	}

	for j := 0; j < t-1; j++ {
		z.keys[j] = y.keys[j+t]
	}

	if !y.isLeaf {
		for j := 0; j < t; j++ {
			z.children[j] = y.children[j+t]
		}
	}

	midKey := y.keys[t-1]

	y.keys = y.keys[:t-1]

	parent.keys = append(parent.keys, 0)
	parent.children = append(parent.children, nil)

	for j := len(parent.keys) - 1; j > i; j-- {
		parent.keys[j] = parent.keys[j-1]
		parent.children[j+1] = parent.children[j]
	}

	parent.keys[i] = midKey
	parent.children[i+1] = z
}

func (bt *BTree) Search(key int) bool {
	return bt.search(bt.root, key)
}

func (bt *BTree) search(node *BTreeNode, key int) bool {
	if node == nil {
		return false
	}

	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	if i < len(node.keys) && key == node.keys[i] {
		return true
	}

	if node.isLeaf {
		return false
	}

	return bt.search(node.children[i], key)
}

func (bt *BTree) Delete(key int) bool {
	if !bt.Search(key) {
		return false
	}
	bt.delete(bt.root, key)

	if bt.root != nil && len(bt.root.keys) == 0 {
		if bt.root.isLeaf {
			bt.root = nil
		} else {
			bt.root = bt.root.children[0]
		}
	}
	return true
}

func (bt *BTree) delete(node *BTreeNode, key int) {
	idx := bt.findKey(node, key)

	if idx < len(node.keys) && node.keys[idx] == key {
		if node.isLeaf {
			bt.removeFromLeaf(node, idx)
		} else {
			bt.removeFromNonLeaf(node, idx)
		}
	} else {
		if node.isLeaf {
			return
		}

		flag := (idx == len(node.keys))

		if len(node.children[idx].keys) < bt.t {
			bt.fill(node, idx)
		}

		if flag && idx > len(node.keys) {
			bt.delete(node.children[idx-1], key)
		} else {
			bt.delete(node.children[idx], key)
		}
	}
}

func (bt *BTree) removeFromLeaf(node *BTreeNode, idx int) {
	node.keys = append(node.keys[:idx], node.keys[idx+1:]...)
}

func (bt *BTree) removeFromNonLeaf(node *BTreeNode, idx int) {
	key := node.keys[idx]

	if len(node.children[idx].keys) >= bt.t {
		pred := bt.getPredecessor(node.children[idx])
		node.keys[idx] = pred
		bt.delete(node.children[idx], pred)
	} else if len(node.children[idx+1].keys) >= bt.t {
		succ := bt.getSuccessor(node.children[idx+1])
		node.keys[idx] = succ
		bt.delete(node.children[idx+1], succ)
	} else {
		bt.merge(node, idx)
		bt.delete(node.children[idx], key)
	}
}

func (bt *BTree) getPredecessor(node *BTreeNode) int {
	for !node.isLeaf {
		node = node.children[len(node.keys)]
	}
	return node.keys[len(node.keys)-1]
}

func (bt *BTree) getSuccessor(node *BTreeNode) int {
	for !node.isLeaf {
		node = node.children[0]
	}
	return node.keys[0]
}

func (bt *BTree) fill(node *BTreeNode, idx int) {
	if idx != 0 && len(node.children[idx-1].keys) >= bt.t {
		bt.borrowFromPrev(node, idx)
	} else if idx != len(node.keys) && len(node.children[idx+1].keys) >= bt.t {
		bt.borrowFromNext(node, idx)
	} else {
		if idx != len(node.keys) {
			bt.merge(node, idx)
		} else {
			bt.merge(node, idx-1)
		}
	}
}

func (bt *BTree) borrowFromPrev(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx-1]

	for i := len(child.keys) - 1; i >= 0; i-- {
		child.keys = append(child.keys[:i+1], child.keys[i:]...)
	}

	child.keys = append([]int{node.keys[idx-1]}, child.keys...)

	if !child.isLeaf {
		for i := len(child.children) - 1; i >= 0; i-- {
			child.children = append(child.children[:i+1], child.children[i:]...)
		}
		child.children = append([]*BTreeNode{sibling.children[len(sibling.children)-1]}, child.children...)
	}

	node.keys[idx-1] = sibling.keys[len(sibling.keys)-1]

	sibling.keys = sibling.keys[:len(sibling.keys)-1]
	if !sibling.isLeaf {
		sibling.children = sibling.children[:len(sibling.children)-1]
	}
}

func (bt *BTree) borrowFromNext(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx+1]

	child.keys = append(child.keys, node.keys[idx])

	if !child.isLeaf {
		child.children = append(child.children, sibling.children[0])
	}

	node.keys[idx] = sibling.keys[0]

	sibling.keys = sibling.keys[1:]
	if !sibling.isLeaf {
		sibling.children = sibling.children[1:]
	}
}

func (bt *BTree) merge(node *BTreeNode, idx int) {
	child := node.children[idx]
	sibling := node.children[idx+1]

	child.keys = append(child.keys, node.keys[idx])
	child.keys = append(child.keys, sibling.keys...)

	if !child.isLeaf {
		child.children = append(child.children, sibling.children...)
	}

	node.keys = append(node.keys[:idx], node.keys[idx+1:]...)
	node.children = append(node.children[:idx+1], node.children[idx+2:]...)
}

func (bt *BTree) findKey(node *BTreeNode, key int) int {
	idx := 0
	for idx < len(node.keys) && node.keys[idx] < key {
		idx++
	}
	return idx
}

func (bt *BTree) InOrderTraversal() []int {
	var result []int
	bt.inOrder(bt.root, &result)
	return result
}

func (bt *BTree) inOrder(node *BTreeNode, result *[]int) {
	if node == nil {
		return
	}

	for i := 0; i < len(node.keys); i++ {
		if !node.isLeaf {
			bt.inOrder(node.children[i], result)
		}
		*result = append(*result, node.keys[i])
	}

	if !node.isLeaf {
		bt.inOrder(node.children[len(node.keys)], result)
	}
}

func (bt *BTree) Min() int {
	if bt.root == nil {
		return 0
	}

	current := bt.root
	for !current.isLeaf {
		current = current.children[0]
	}

	return current.keys[0]
}

func (bt *BTree) Max() int {
	if bt.root == nil {
		return 0
	}

	current := bt.root
	for !current.isLeaf {
		current = current.children[len(current.keys)]
	}

	return current.keys[len(current.keys)-1]
}

func (bt *BTree) Size() int {
	result := 0
	bt.size(bt.root, &result)
	return result
}

func (bt *BTree) size(node *BTreeNode, result *int) {
	if node == nil {
		return
	}

	for i := 0; i < len(node.keys); i++ {
		if !node.isLeaf {
			bt.size(node.children[i], result)
		}
		(*result)++
	}

	if !node.isLeaf {
		bt.size(node.children[len(node.keys)], result)
	}
}

func (bt *BTree) IsEmpty() bool {
	return bt.root == nil
}

func (bt *BTree) Clear() {
	bt.root = nil
}

func (bt *BTree) Print() {
	bt.printNode(bt.root, 0)
}

func (bt *BTree) printNode(node *BTreeNode, level int) {
	if node == nil {
		return
	}

	fmt.Printf("Level %d: %v\n", level, node.keys)

	if !node.isLeaf {
		for i := range node.children {
			bt.printNode(node.children[i], level+1)
		}
	}
}
