package trees

type SegmentTree struct {
	tree []int
	n    int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	st := &SegmentTree{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.build(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTree) build(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	st.build(arr, 2*node+1, start, mid)
	st.build(arr, 2*node+2, mid+1, end)
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegmentTree) Update(index, value int) {
	st.update(0, 0, st.n-1, index, value)
}

func (st *SegmentTree) update(node, start, end, index, value int) {
	if start == end {
		st.tree[node] = value
		return
	}
	mid := (start + end) / 2
	if index <= mid {
		st.update(2*node+1, start, mid, index, value)
	} else {
		st.update(2*node+2, mid+1, end, index, value)
	}
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegmentTree) Query(left, right int) int {
	return st.query(0, 0, st.n-1, left, right)
}

func (st *SegmentTree) query(node, start, end, left, right int) int {
	if right < start || end < left {
		return 0
	}
	if left <= start && end <= right {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftSum := st.query(2*node+1, start, mid, left, right)
	rightSum := st.query(2*node+2, mid+1, end, left, right)
	return leftSum + rightSum
}

func (st *SegmentTree) Size() int {
	return st.n
}

type SegmentTreeMin struct {
	tree []int
	n    int
}

func NewSegmentTreeMin(arr []int) *SegmentTreeMin {
	n := len(arr)
	st := &SegmentTreeMin{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.buildMin(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTreeMin) buildMin(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	st.buildMin(arr, 2*node+1, start, mid)
	st.buildMin(arr, 2*node+2, mid+1, end)
	st.tree[node] = minInt(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegmentTreeMin) Update(index, value int) {
	st.updateMin(0, 0, st.n-1, index, value)
}

func (st *SegmentTreeMin) updateMin(node, start, end, index, value int) {
	if start == end {
		st.tree[node] = value
		return
	}
	mid := (start + end) / 2
	if index <= mid {
		st.updateMin(2*node+1, start, mid, index, value)
	} else {
		st.updateMin(2*node+2, mid+1, end, index, value)
	}
	st.tree[node] = minInt(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegmentTreeMin) Query(left, right int) int {
	return st.queryMin(0, 0, st.n-1, left, right)
}

func (st *SegmentTreeMin) queryMin(node, start, end, left, right int) int {
	if right < start || end < left {
		return 1<<31 - 1
	}
	if left <= start && end <= right {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftMin := st.queryMin(2*node+1, start, mid, left, right)
	rightMin := st.queryMin(2*node+2, mid+1, end, left, right)
	return minInt(leftMin, rightMin)
}

func (st *SegmentTreeMin) Size() int {
	return st.n
}

type SegmentTreeMax struct {
	tree []int
	n    int
}

func NewSegmentTreeMax(arr []int) *SegmentTreeMax {
	n := len(arr)
	st := &SegmentTreeMax{
		tree: make([]int, 4*n),
		n:    n,
	}
	st.buildMax(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTreeMax) buildMax(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	st.buildMax(arr, 2*node+1, start, mid)
	st.buildMax(arr, 2*node+2, mid+1, end)
	st.tree[node] = maxInt(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegmentTreeMax) Update(index, value int) {
	st.updateMax(0, 0, st.n-1, index, value)
}

func (st *SegmentTreeMax) updateMax(node, start, end, index, value int) {
	if start == end {
		st.tree[node] = value
		return
	}
	mid := (start + end) / 2
	if index <= mid {
		st.updateMax(2*node+1, start, mid, index, value)
	} else {
		st.updateMax(2*node+2, mid+1, end, index, value)
	}
	st.tree[node] = maxInt(st.tree[2*node+1], st.tree[2*node+2])
}

func (st *SegmentTreeMax) Query(left, right int) int {
	return st.queryMax(0, 0, st.n-1, left, right)
}

func (st *SegmentTreeMax) queryMax(node, start, end, left, right int) int {
	if right < start || end < left {
		return ^int(^uint(0) >> 1)
	}
	if left <= start && end <= right {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftMax := st.queryMax(2*node+1, start, mid, left, right)
	rightMax := st.queryMax(2*node+2, mid+1, end, left, right)
	return maxInt(leftMax, rightMax)
}

func (st *SegmentTreeMax) Size() int {
	return st.n
}

type SegmentTreeLazy struct {
	tree []int
	lazy []int
	n    int
}

func NewSegmentTreeLazy(arr []int) *SegmentTreeLazy {
	n := len(arr)
	st := &SegmentTreeLazy{
		tree: make([]int, 4*n),
		lazy: make([]int, 4*n),
		n:    n,
	}
	st.buildLazy(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTreeLazy) buildLazy(arr []int, node, start, end int) {
	if start == end {
		st.tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	st.buildLazy(arr, 2*node+1, start, mid)
	st.buildLazy(arr, 2*node+2, mid+1, end)
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegmentTreeLazy) UpdateRange(left, right, value int) {
	st.updateRange(0, 0, st.n-1, left, right, value)
}

func (st *SegmentTreeLazy) updateRange(node, start, end, left, right, value int) {
	if st.lazy[node] != 0 {
		st.tree[node] += st.lazy[node] * (end - start + 1)
		if start != end {
			st.lazy[2*node+1] += st.lazy[node]
			st.lazy[2*node+2] += st.lazy[node]
		}
		st.lazy[node] = 0
	}

	if right < start || end < left {
		return
	}

	if left <= start && end <= right {
		st.tree[node] += value * (end - start + 1)
		if start != end {
			st.lazy[2*node+1] += value
			st.lazy[2*node+2] += value
		}
		return
	}

	mid := (start + end) / 2
	st.updateRange(2*node+1, start, mid, left, right, value)
	st.updateRange(2*node+2, mid+1, end, left, right, value)
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func (st *SegmentTreeLazy) Query(left, right int) int {
	return st.queryLazy(0, 0, st.n-1, left, right)
}

func (st *SegmentTreeLazy) queryLazy(node, start, end, left, right int) int {
	if st.lazy[node] != 0 {
		st.tree[node] += st.lazy[node] * (end - start + 1)
		if start != end {
			st.lazy[2*node+1] += st.lazy[node]
			st.lazy[2*node+2] += st.lazy[node]
		}
		st.lazy[node] = 0
	}

	if right < start || end < left {
		return 0
	}

	if left <= start && end <= right {
		return st.tree[node]
	}

	mid := (start + end) / 2
	leftSum := st.queryLazy(2*node+1, start, mid, left, right)
	rightSum := st.queryLazy(2*node+2, mid+1, end, left, right)
	return leftSum + rightSum
}

func (st *SegmentTreeLazy) Size() int {
	return st.n
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
