package structures

import (
	"math/rand"
	"time"
)

type SkipListNode struct {
	value   int
	forward []*SkipListNode
}

type SkipList struct {
	head     *SkipListNode
	maxLevel int
	level    int
	size     int
}

const MAX_LEVEL = 16
const P = 0.5

func NewSkipList() *SkipList {
	rand.Seed(time.Now().UnixNano())

	head := &SkipListNode{
		value:   0,
		forward: make([]*SkipListNode, MAX_LEVEL),
	}

	return &SkipList{
		head:     head,
		maxLevel: MAX_LEVEL,
		level:    1,
		size:     0,
	}
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for rand.Float64() < P && level < sl.maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Search(value int) bool {
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].value < value {
			current = current.forward[i]
		}
	}

	current = current.forward[0]
	return current != nil && current.value == value
}

func (sl *SkipList) Insert(value int) {
	update := make([]*SkipListNode, sl.maxLevel)
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].value < value {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current != nil && current.value == value {
		return
	}

	newLevel := sl.randomLevel()

	if newLevel > sl.level {
		for i := sl.level; i < newLevel; i++ {
			update[i] = sl.head
		}
		sl.level = newLevel
	}

	newNode := &SkipListNode{
		value:   value,
		forward: make([]*SkipListNode, newLevel),
	}

	for i := 0; i < newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	sl.size++
}

func (sl *SkipList) Delete(value int) bool {
	update := make([]*SkipListNode, sl.maxLevel)
	current := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].value < value {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current == nil || current.value != value {
		return false
	}

	for i := 0; i < sl.level; i++ {
		if update[i].forward[i] != current {
			break
		}
		update[i].forward[i] = current.forward[i]
	}

	for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
		sl.level--
	}

	sl.size--
	return true
}

func (sl *SkipList) Size() int {
	return sl.size
}

func (sl *SkipList) Level() int {
	return sl.level
}

func (sl *SkipList) Contains(value int) bool {
	return sl.Search(value)
}

func (sl *SkipList) IsEmpty() bool {
	return sl.size == 0
}

func (sl *SkipList) ToSlice() []int {
	result := []int{}
	current := sl.head.forward[0]

	for current != nil {
		result = append(result, current.value)
		current = current.forward[0]
	}

	return result
}

func (sl *SkipList) Clear() {
	sl.head = &SkipListNode{
		value:   0,
		forward: make([]*SkipListNode, MAX_LEVEL),
	}
	sl.level = 1
	sl.size = 0
}

func (sl *SkipList) Min() int {
	if sl.IsEmpty() {
		return 0
	}
	return sl.head.forward[0].value
}

func (sl *SkipList) Max() int {
	if sl.IsEmpty() {
		return 0
	}
	current := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for current.forward[i] != nil {
			current = current.forward[i]
		}
	}
	return current.value
}
