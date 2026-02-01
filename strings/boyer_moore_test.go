package strings

import (
	"testing"
)

func TestBoyerMooreSearch(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected []int
	}{
		{"hello world", "world", []int{6}},
		{"abracadabra", "abra", []int{0, 7}},
		{"aaaaa", "aa", []int{0, 1, 2, 3}},
		{"mississippi", "issi", []int{1, 4}},
		{"", "pattern", []int{}},
		{"text", "", []int{}},
		{"abc", "xyz", []int{}},
		{"a", "a", []int{0}},
	}

	for _, test := range tests {
		bm := NewBoyerMoore()
		result := bm.Search(test.text, test.pattern)

		if len(result) != len(test.expected) {
			t.Errorf("Search(%q, %q) = %v, expected %v", test.text, test.pattern, result, test.expected)
		}

		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Search(%q, %q)[%d] = %d, expected %d", test.text, test.pattern, i, result[i], test.expected[i])
			}
		}
	}
}

func TestBoyerMooreSearchFirst(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected int
	}{
		{"hello world", "world", 6},
		{"abracadabra", "abra", 0},
		{"mississippi", "issi", 1},
		{"abc", "xyz", -1},
		{"", "pattern", -1},
		{"text", "", -1},
	}

	for _, test := range tests {
		bm := NewBoyerMoore()
		result := bm.SearchFirst(test.text, test.pattern)

		if result != test.expected {
			t.Errorf("SearchFirst(%q, %q) = %d, expected %d", test.text, test.pattern, result, test.expected)
		}
	}
}

func TestBoyerMooreCount(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected int
	}{
		{"abracadabra", "abra", 2},
		{"aaaaa", "aa", 4},
		{"hello world", "world", 1},
		{"abc", "xyz", 0},
	}

	for _, test := range tests {
		bm := NewBoyerMoore()
		result := bm.Count(test.text, test.pattern)

		if result != test.expected {
			t.Errorf("Count(%q, %q) = %d, expected %d", test.text, test.pattern, result, test.expected)
		}
	}
}

func TestBoyerMooreContains(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", true},
		{"abc", "xyz", false},
		{"", "pattern", false},
	}

	for _, test := range tests {
		bm := NewBoyerMoore()
		result := bm.Contains(test.text, test.pattern)

		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %v, expected %v", test.text, test.pattern, result, test.expected)
		}
	}
}
