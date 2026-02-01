package trees

import (
	"testing"
)

func TestSegmentTree(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 0, 1},
		{1, 3, 15},
		{0, 5, 36},
		{2, 4, 21},
	}

	for _, test := range tests {
		result := st.Query(test.left, test.right)
		if result != test.expected {
			t.Errorf("Query(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}
	}
}

func TestSegmentTreeUpdate(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)

	st.Update(2, 10)

	if st.Query(0, 5) != 41 {
		t.Errorf("After update, Query(0, 5) = %d, expected 41", st.Query(0, 5))
	}
	if st.Query(2, 2) != 10 {
		t.Errorf("After update, Query(2, 2) = %d, expected 10", st.Query(2, 2))
	}
}

func TestSegmentTreeMin(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3}
	st := NewSegmentTreeMin(arr)

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 0, 5},
		{1, 4, 1},
		{0, 5, 1},
		{2, 3, 1},
	}

	for _, test := range tests {
		result := st.Query(test.left, test.right)
		if result != test.expected {
			t.Errorf("Min Query(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}
	}
}

func TestSegmentTreeMax(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3}
	st := NewSegmentTreeMax(arr)

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 0, 5},
		{1, 4, 9},
		{0, 5, 9},
		{2, 3, 8},
	}

	for _, test := range tests {
		result := st.Query(test.left, test.right)
		if result != test.expected {
			t.Errorf("Max Query(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}
	}
}

func TestSegmentTreeLazy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	st := NewSegmentTreeLazy(arr)

	st.UpdateRange(1, 3, 2)

	if st.Query(0, 4) != 21 {
		t.Errorf("Lazy: Query(0, 4) = %d, expected 21", st.Query(0, 4))
	}
	if st.Query(1, 3) != 15 {
		t.Errorf("Lazy: Query(1, 3) = %d, expected 15", st.Query(1, 3))
	}

	st.UpdateRange(0, 4, 3)
	if st.Query(0, 4) != 36 {
		t.Errorf("Lazy: Query(0, 4) = %d, expected 36", st.Query(0, 4))
	}
}

func TestSegmentTreeSize(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	st := NewSegmentTree(arr)

	if st.Size() != 5 {
		t.Errorf("Size() = %d, expected 5", st.Size())
	}

	stMin := NewSegmentTreeMin(arr)
	if stMin.Size() != 5 {
		t.Errorf("Min Size() = %d, expected 5", stMin.Size())
	}

	stMax := NewSegmentTreeMax(arr)
	if stMax.Size() != 5 {
		t.Errorf("Max Size() = %d, expected 5", stMax.Size())
	}

	stLazy := NewSegmentTreeLazy(arr)
	if stLazy.Size() != 5 {
		t.Errorf("Lazy Size() = %d, expected 5", stLazy.Size())
	}
}

func TestSegmentTreeConsistency(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	st := NewSegmentTree(arr)
	ft := NewFenwickTreeFromArray(arr)

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			stSum := st.Query(i, j)
			ftSum := ft.RangeQuery(i, j)
			if stSum != ftSum {
				t.Errorf("Inconsistent: SegmentTree.Query(%d, %d) = %d, FenwickTree.RangeQuery(%d, %d) = %d",
					i, j, stSum, i, j, ftSum)
			}
		}
	}
}

func TestSegmentTreeEdgeCases(t *testing.T) {
	arr := []int{5}
	st := NewSegmentTree(arr)

	if st.Query(0, 0) != 5 {
		t.Errorf("Single element: Query(0, 0) = %d, expected 5", st.Query(0, 0))
	}

	st.Update(0, 10)
	if st.Query(0, 0) != 10 {
		t.Errorf("After update: Query(0, 0) = %d, expected 10", st.Query(0, 0))
	}
}

func TestSegmentTreeMultipleUpdates(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0}
	st := NewSegmentTree(arr)

	for i := 0; i < 5; i++ {
		st.Update(i, i+1)
	}

	if st.Query(0, 4) != 15 {
		t.Errorf("After multiple updates: Query(0, 4) = %d, expected 15", st.Query(0, 4))
	}
}
