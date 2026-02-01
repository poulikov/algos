package trees

type FenwickTree struct {
	tree []int
	n    int
}

func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, n+1),
		n:    n,
	}
}

func NewFenwickTreeFromArray(arr []int) *FenwickTree {
	ft := NewFenwickTree(len(arr))
	for i, val := range arr {
		ft.Update(i, val)
	}
	return ft
}

func (ft *FenwickTree) Update(index, delta int) {
	index++
	for index <= ft.n {
		ft.tree[index] += delta
		index += index & (-index)
	}
}

func (ft *FenwickTree) Set(index, value int) {
	index++
	oldValue := ft.PointQuery(index - 1)
	delta := value - oldValue
	for index <= ft.n {
		ft.tree[index] += delta
		index += index & (-index)
	}
}

func (ft *FenwickTree) Query(index int) int {
	index++
	sum := 0
	for index > 0 {
		sum += ft.tree[index]
		index -= index & (-index)
	}
	return sum
}

func (ft *FenwickTree) RangeQuery(left, right int) int {
	if left == 0 {
		return ft.Query(right)
	}
	return ft.Query(right) - ft.Query(left-1)
}

func (ft *FenwickTree) PointQuery(index int) int {
	return ft.RangeQuery(index, index)
}

func (ft *FenwickTree) Size() int {
	return ft.n
}

func (ft *FenwickTree) LowerBound(target int) int {
	index := 0
	bitmap := 1
	for bitmap <= ft.n {
		bitmap <<= 1
	}
	bitmap >>= 1

	for bitmap != 0 {
		temp := index + bitmap
		if temp <= ft.n && ft.tree[temp] < target {
			index = temp
			target -= ft.tree[index]
		}
		bitmap >>= 1
	}
	return index + 1
}

func (ft *FenwickTree) FindKth(k int) int {
	if k <= 0 || k > ft.Query(ft.n-1) {
		return -1
	}
	index := 0
	bitmap := 1
	for bitmap <= ft.n {
		bitmap <<= 1
	}
	bitmap >>= 1

	for bitmap != 0 {
		temp := index + bitmap
		if temp <= ft.n && ft.tree[temp] < k {
			index = temp
			k -= ft.tree[index]
		}
		bitmap >>= 1
	}
	return index
}
