package dp

import (
	"testing"
)

func TestEditDistance(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"sunday", "saturday", 3},
		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"same", "same", 0},
		{"book", "back", 2},
		{"intention", "execution", 5},
	}

	for _, test := range tests {
		result := EditDistance(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("EditDistance(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestEditDistanceOptimized(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"sunday", "saturday", 3},
		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"same", "same", 0},
		{"book", "back", 2},
		{"intention", "execution", 5},
	}

	for _, test := range tests {
		result := EditDistanceOptimized(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("EditDistanceOptimized(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestEditDistanceWithPath(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"sunday", "saturday", 3},
		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"same", "same", 0},
		{"book", "back", 2},
	}

	for _, test := range tests {
		result, path := EditDistanceWithPath(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("EditDistanceWithPath(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}

		editOps := 0
		for _, op := range path {
			if op.Op != Match {
				editOps++
			}
		}

		if editOps != result {
			t.Errorf("EditDistanceWithPath(%q, %q): edit operations %d != distance %d", test.s1, test.s2, editOps, result)
		}
	}
}

func TestEditDistanceConsistency(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
	}{
		{"hello", "world"},
		{"algorithm", "rhythm"},
		{"test", "testing"},
	}

	for _, test := range tests {
		dist1 := EditDistance(test.s1, test.s2)
		dist2 := EditDistanceOptimized(test.s1, test.s2)
		dist3, _ := EditDistanceWithPath(test.s1, test.s2)

		if dist1 != dist2 || dist1 != dist3 {
			t.Errorf("EditDistance methods return different values for %q, %q: %d, %d, %d",
				test.s1, test.s2, dist1, dist2, dist3)
		}
	}
}

func TestEditDistanceSymmetric(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
	}{
		{"kitten", "sitting"},
		{"abc", "def"},
		{"", "test"},
	}

	for _, test := range tests {
		dist1 := EditDistance(test.s1, test.s2)
		dist2 := EditDistance(test.s2, test.s1)

		if dist1 != dist2 {
			t.Errorf("EditDistance is not symmetric: EditDistance(%q, %q) = %d, EditDistance(%q, %q) = %d",
				test.s1, test.s2, dist1, test.s2, test.s1, dist2)
		}
	}
}
