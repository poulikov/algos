package searching

import (
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{3, 1},
		{5, 2},
		{7, 3},
		{9, 4},
		{11, 5},
		{13, 6},
		{15, 7},
		{17, 8},
		{19, 9},
		{2, -1},
		{8, -1},
		{0, -1},
		{20, -1},
	}

	for _, test := range tests {
		result := InterpolationSearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("InterpolationSearch(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestInterpolationSearchEmpty(t *testing.T) {
	slice := []int{}
	result := InterpolationSearch(slice, 5)
	if result != -1 {
		t.Fatalf("InterpolationSearch on empty slice should return -1, got %d", result)
	}
}

func TestInterpolationSearchSingle(t *testing.T) {
	slice := []int{5}
	result := InterpolationSearch(slice, 5)
	if result != 0 {
		t.Fatalf("InterpolationSearch(%d) = %d, expected 0", 5, result)
	}

	result = InterpolationSearch(slice, 3)
	if result != -1 {
		t.Fatalf("InterpolationSearch(%d) should return -1, got %d", 3, result)
	}
}

func TestInterpolationSearchUniform(t *testing.T) {
	slice := []int{0, 10, 20, 30, 40, 50}

	tests := []struct {
		target   int
		expected int
	}{
		{0, 0},
		{10, 1},
		{20, 2},
		{30, 3},
		{40, 4},
		{50, 5},
		{5, -1},
		{45, -1},
		{55, -1},
	}

	for _, test := range tests {
		result := InterpolationSearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("InterpolationSearch(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestInterpolationSearchNonUniform(t *testing.T) {
	slice := []int{1, 2, 4, 8, 16, 32, 64, 128}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{2, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{32, 5},
		{64, 6},
		{128, 7},
		{3, -1},
		{5, -1},
	}

	for _, test := range tests {
		result := InterpolationSearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("InterpolationSearch(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestInterpolationSearchDuplicates(t *testing.T) {
	slice := []int{1, 1, 1, 2, 2, 3, 3}

	for _, test := range []int{1, 2, 3} {
		result := InterpolationSearch(slice, test)
		if result == -1 {
			t.Fatalf("InterpolationSearch(%d) should find a match, got -1", test)
		}
		if slice[result] != test {
			t.Fatalf("InterpolationSearch(%d) = %d, expected value %d at index %d", test, result, test, result)
		}
	}
}

func TestInterpolationSearchFloat(t *testing.T) {
	slice := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	tests := []struct {
		target   float64
		expected int
	}{
		{1.0, 0},
		{2.0, 1},
		{3.0, 2},
		{4.0, 3},
		{5.0, 4},
		{2.5, -1},
		{6.0, -1},
	}

	for _, test := range tests {
		result := InterpolationSearch(slice, test.target)
		if result != test.expected {
			t.Fatalf("InterpolationSearch(%.1f) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func TestInterpolationSearchRecursive(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9, 11}

	tests := []struct {
		target   int
		expected int
	}{
		{1, 0},
		{3, 1},
		{5, 2},
		{7, 3},
		{9, 4},
		{11, 5},
		{2, -1},
		{12, -1},
	}

	for _, test := range tests {
		result := InterpolationSearchRecursive(slice, test.target)
		if result != test.expected {
			t.Fatalf("InterpolationSearchRecursive(%d) = %d, expected %d", test.target, result, test.expected)
		}
	}
}

func BenchmarkInterpolationSearch(b *testing.B) {
	slice := make([]int, 100000)
	for i := 0; i < len(slice); i++ {
		slice[i] = i * 2
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InterpolationSearch(slice, i*2)
	}
}

func BenchmarkInterpolationSearchNonUniform(b *testing.B) {
	slice := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InterpolationSearch(slice, i%10+1)
	}
}
