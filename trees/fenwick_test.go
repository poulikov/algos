package trees

import (
	"testing"
)

func TestFenwickTreeNew(t *testing.T) {
	ft := NewFenwickTree(10)
	if ft.Size() != 10 {
		t.Errorf("Expected size 10, got %d", ft.Size())
	}
}

func TestFenwickTreeUpdate(t *testing.T) {
	ft := NewFenwickTree(5)
	ft.Update(0, 1)
	ft.Update(1, 2)
	ft.Update(2, 3)

	if ft.Query(0) != 1 {
		t.Errorf("Query(0) = %d, expected 1", ft.Query(0))
	}
	if ft.Query(1) != 3 {
		t.Errorf("Query(1) = %d, expected 3", ft.Query(1))
	}
	if ft.Query(2) != 6 {
		t.Errorf("Query(2) = %d, expected 6", ft.Query(2))
	}
}

func TestFenwickTreeRangeQuery(t *testing.T) {
	ft := NewFenwickTree(10)
	ft.Update(0, 1)
	ft.Update(1, 2)
	ft.Update(2, 3)
	ft.Update(3, 4)
	ft.Update(4, 5)

	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 0, 1},
		{1, 3, 9},
		{0, 4, 15},
		{2, 4, 12},
	}

	for _, test := range tests {
		result := ft.RangeQuery(test.left, test.right)
		if result != test.expected {
			t.Errorf("RangeQuery(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}
	}
}

func TestFenwickTreePointQuery(t *testing.T) {
	ft := NewFenwickTree(5)
	ft.Update(0, 10)
	ft.Update(1, 20)
	ft.Update(2, 30)

	tests := []struct {
		index    int
		expected int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 0},
		{4, 0},
	}

	for _, test := range tests {
		result := ft.PointQuery(test.index)
		if result != test.expected {
			t.Errorf("PointQuery(%d) = %d, expected %d", test.index, result, test.expected)
		}
	}
}

func TestFenwickTreeSet(t *testing.T) {
	ft := NewFenwickTree(5)
	ft.Set(0, 10)
	ft.Set(1, 20)
	ft.Set(2, 30)

	ft.Set(1, 25)
	ft.Set(2, 35)

	if ft.PointQuery(0) != 10 {
		t.Errorf("PointQuery(0) = %d, expected 10", ft.PointQuery(0))
	}
	if ft.PointQuery(1) != 25 {
		t.Errorf("PointQuery(1) = %d, expected 25", ft.PointQuery(1))
	}
	if ft.PointQuery(2) != 35 {
		t.Errorf("PointQuery(2) = %d, expected 35", ft.PointQuery(2))
	}
}

func TestFenwickTreeFromArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	ft := NewFenwickTreeFromArray(arr)

	if ft.Query(4) != 15 {
		t.Errorf("Query(4) = %d, expected 15", ft.Query(4))
	}
	if ft.RangeQuery(1, 3) != 9 {
		t.Errorf("RangeQuery(1, 3) = %d, expected 9", ft.RangeQuery(1, 3))
	}
}

func TestFenwickTreeFindKth(t *testing.T) {
	ft := NewFenwickTree(10)
	ft.Update(0, 1)
	ft.Update(1, 1)
	ft.Update(2, 1)
	ft.Update(3, 2)
	ft.Update(4, 1)

	tests := []struct {
		k        int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 3},
		{6, 4},
	}

	for _, test := range tests {
		result := ft.FindKth(test.k)
		if result != test.expected {
			t.Errorf("FindKth(%d) = %d, expected %d", test.k, result, test.expected)
		}
	}

	result := ft.FindKth(0)
	if result != -1 {
		t.Errorf("FindKth(0) should return -1, got %d", result)
	}

	result = ft.FindKth(10)
	if result != -1 {
		t.Errorf("FindKth(10) should return -1, got %d", result)
	}
}

func TestFenwickTreeLowerBound(t *testing.T) {
	ft := NewFenwickTree(10)
	ft.Update(0, 1)
	ft.Update(1, 2)
	ft.Update(2, 3)
	ft.Update(3, 4)
	ft.Update(4, 5)

	tests := []struct {
		target   int
		expected int
	}{
		{1, 1},
		{3, 2},
		{6, 3},
		{10, 4},
		{15, 5},
	}

	for _, test := range tests {
		result := ft.LowerBound(test.target)
		if result != test.expected {
			t.Errorf("LowerBound(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestFenwickTreeConsistency(t *testing.T) {
	ft := NewFenwickTree(100)

	ft.Set(0, 5)
	ft.Set(10, 10)
	ft.Set(50, 15)
	ft.Set(99, 20)

	sum := ft.Query(99)
	if sum != 50 {
		t.Errorf("Expected total sum 50, got %d", sum)
	}

	ft.Update(25, 7)
	sum = ft.Query(99)
	if sum != 57 {
		t.Errorf("Expected total sum 57 after update, got %d", sum)
	}

	rangeSum := ft.RangeQuery(10, 50)
	expectedRangeSum := 10 + 7 + 15
	if rangeSum != expectedRangeSum {
		t.Errorf("Expected range sum %d, got %d", expectedRangeSum, rangeSum)
	}
}
