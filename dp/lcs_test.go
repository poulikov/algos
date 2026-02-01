package dp

import (
	"testing"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"", "abc", 0},
		{"abc", "", 0},
		{"a", "a", 1},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"ABCD", "ACBAD", 3},
		{"AGGTAB", "GXTXAYB", 4},
		{"stone", "longest", 3},
		{"ABCBDAB", "BDCABA", 4},
		{"XMJYAUZ", "MZJAWXU", 4},
		{"aaaa", "aa", 2},
		{"abcdefg", "aceg", 4},
	}

	for _, test := range tests {
		result := LCS(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("LCS(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestLCSString(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected []string
	}{
		{"", "", []string{""}},
		{"", "abc", []string{""}},
		{"abc", "", []string{""}},
		{"a", "a", []string{"a"}},
		{"abc", "abc", []string{"abc"}},
		{"abc", "def", []string{""}},
		{"ABCD", "ACBAD", []string{"ABD", "ACD"}},
		{"AGGTAB", "GXTXAYB", []string{"GTAB"}},
		{"stone", "longest", []string{"one"}},
	}

	for _, test := range tests {
		result := LCSString(test.s1, test.s2)
		found := false
		for _, exp := range test.expected {
			if result == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("LCSString(%q, %q) = %q, expected one of %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestLCSOptimized(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"", "abc", 0},
		{"abc", "", 0},
		{"a", "a", 1},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"ABCD", "ACBAD", 3},
		{"AGGTAB", "GXTXAYB", 4},
		{"stone", "longest", 3},
		{"ABCBDAB", "BDCABA", 4},
		{"XMJYAUZ", "MZJAWXU", 4},
		{"abcdefg", "aceg", 4},
	}

	for _, test := range tests {
		result := LCSOptimized(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("LCSOptimized(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestLCSOptimizedMatchesLCS(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
	}{
		{"AGGTAB", "GXTXAYB"},
		{"ABCBDAB", "BDCABA"},
		{"XMJYAUZ", "MZJAWXU"},
		{"abcdefg", "aceg"},
		{"aaaa", "aa"},
	}

	for _, test := range tests {
		normalResult := LCS(test.s1, test.s2)
		optimizedResult := LCSOptimized(test.s1, test.s2)
		if normalResult != optimizedResult {
			t.Errorf("LCS(%q, %q) = %d, but LCSOptimized = %d", test.s1, test.s2, normalResult, optimizedResult)
		}
	}
}

func TestLCSAll(t *testing.T) {
	tests := []struct {
		s1            string
		s2            string
		expectedCount int
		lcsLength     int
		checkFunc     func(t *testing.T, results []string)
	}{
		{
			"", "", 1, 0,
			func(t *testing.T, results []string) {
				if len(results) != 1 || results[0] != "" {
					t.Errorf("Expected [\"\"], got %v", results)
				}
			},
		},
		{
			"abc", "", 1, 0,
			func(t *testing.T, results []string) {
				if len(results) != 1 || results[0] != "" {
					t.Errorf("Expected [\"\"], got %v", results)
				}
			},
		},
		{
			"aaa", "aaa", 1, 3,
			func(t *testing.T, results []string) {
				if len(results) != 1 || results[0] != "aaa" {
					t.Errorf("Expected [\"aaa\"], got %v", results)
				}
			},
		},
		{
			"ABCD", "ACBAD", 2, 3,
			func(t *testing.T, results []string) {
				for _, r := range results {
					if r != "ABD" && r != "ACD" {
						t.Errorf("Expected \"ABD\" or \"ACD\", got %q", r)
					}
					if len(r) != 3 {
						t.Errorf("Expected length 3, got %d", len(r))
					}
				}
			},
		},
		{
			"ABCBDAB", "BDCABA", 3, 4,
			func(t *testing.T, results []string) {
				if len(results) == 0 {
					t.Error("Expected at least one result")
				}
				for _, r := range results {
					if len(r) != 4 {
						t.Errorf("Expected length 4, got %d", len(r))
					}
				}
			},
		},
		{
			"AGGTAB", "GXTXAYB", 1, 4,
			func(t *testing.T, results []string) {
				if len(results) == 0 {
					t.Error("Expected at least one result")
				}
				for _, r := range results {
					if len(r) != 4 {
						t.Errorf("Expected length 4, got %d", len(r))
					}
				}
			},
		},
	}

	for _, test := range tests {
		results := LCSAll(test.s1, test.s2)
		if test.expectedCount > 0 && len(results) != test.expectedCount {
			t.Logf("Note: LCSAll(%q, %q) returned %d results, expected %d (may vary by implementation)",
				test.s1, test.s2, len(results), test.expectedCount)
		}
		for _, r := range results {
			if len(r) != test.lcsLength {
				t.Errorf("LCS %q has length %d, expected %d", r, len(r), test.lcsLength)
			}
		}
		if test.checkFunc != nil {
			test.checkFunc(t, results)
		}
	}
}

func TestLCSDistance(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"", "abc", 3},
		{"abc", "", 3},
		{"a", "a", 0},
		{"abc", "abc", 0},
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"intention", "execution", 5},
		{"algorithm", "altruistic", 6},
		{"sunday", "saturday", 3},
		{"", "", 0},
		{"abc", "def", 3},
		{"abcd", "abef", 2},
	}

	for _, test := range tests {
		result := LCSDistance(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("LCSDistance(%q, %q) = %d, expected %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestLCSDistanceWithOperations(t *testing.T) {
	tests := []struct {
		s1               string
		s2               string
		expectedDistance int
		checkOperations  bool
	}{
		{"", "", 0, true},
		{"", "abc", 3, true},
		{"abc", "", 3, true},
		{"a", "a", 0, true},
		{"abc", "abc", 0, true},
		{"kitten", "sitting", 3, true},
		{"flaw", "lawn", 2, true},
	}

	for _, test := range tests {
		distance, ops := LCSDistanceWithOperations(test.s1, test.s2)
		if distance != test.expectedDistance {
			t.Errorf("LCSDistanceWithOperations(%q, %q) distance = %d, expected %d",
				test.s1, test.s2, distance, test.expectedDistance)
		}

		if test.checkOperations {
			transformed := applyOperations(test.s1, ops)
			if transformed != test.s2 {
				t.Errorf("Operations don't transform %q to %q. Got: %q, Ops: %v",
					test.s1, test.s2, transformed, ops)
			}
		}
	}
}

func TestLCSDistanceMatchesLCSDistanceWithOperations(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
	}{
		{"kitten", "sitting"},
		{"flaw", "lawn"},
		{"intention", "execution"},
		{"algorithm", "altruistic"},
		{"sunday", "saturday"},
	}

	for _, test := range tests {
		distance1 := LCSDistance(test.s1, test.s2)
		distance2, _ := LCSDistanceWithOperations(test.s1, test.s2)
		if distance1 != distance2 {
			t.Errorf("LCSDistance(%q, %q) = %d, but LCSDistanceWithOperations = %d",
				test.s1, test.s2, distance1, distance2)
		}
	}
}

func TestLCSSimilarity(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected float64
	}{
		{"", "", 0.0},
		{"", "abc", 0.0},
		{"abc", "", 0.0},
		{"a", "a", 1.0},
		{"abc", "abc", 1.0},
		{"abc", "def", 0.0},
		{"abc", "abd", 0.6666666666666666},
		{"kitten", "sitting", 0.5714285714285714},
		{"flaw", "lawn", 0.75},
		{"ABCD", "ACBAD", 0.6},
		{"AGGTAB", "GXTXAYB", 0.5714285714285714},
	}

	for _, test := range tests {
		result := LCSSimilarity(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("LCSSimilarity(%q, %q) = %f, expected %f", test.s1, test.s2, result, test.expected)
		}
		if result < 0 || result > 1 {
			t.Errorf("LCSSimilarity(%q, %q) = %f, should be between 0 and 1", test.s1, test.s2, result)
		}
	}
}

func TestLCSSimilaritySymmetric(t *testing.T) {
	tests := []struct {
		s1 string
		s2 string
	}{
		{"abc", "def"},
		{"AGGTAB", "GXTXAYB"},
		{"kitten", "sitting"},
		{"ABCD", "ACBAD"},
	}

	for _, test := range tests {
		result1 := LCSSimilarity(test.s1, test.s2)
		result2 := LCSSimilarity(test.s2, test.s1)
		if result1 != result2 {
			t.Errorf("LCSSimilarity(%q, %q) = %f, but LCSSimilarity(%q, %q) = %f (should be symmetric)",
				test.s1, test.s2, result1, test.s2, test.s1, result2)
		}
	}
}

func TestLCSMultiple(t *testing.T) {
	tests := []struct {
		strings   []string
		minLength int
		checkLCS  bool
	}{
		{[]string{}, 0, false},
		{[]string{"abc"}, 3, false},
		{[]string{"abc", "abc"}, 3, true},
		{[]string{"abc", "def"}, 0, true},
		{[]string{"ABCD", "ACBAD", "ABD"}, 1, true},
		{[]string{"AGGTAB", "GXTXAYB", "GTAB"}, 3, true},
		{[]string{"abc", "abc", "abc"}, 3, true},
		{[]string{"ABCBDAB", "BDCABA", "BCBA"}, 0, true},
	}

	for _, test := range tests {
		result := LCSMultiple(test.strings)
		if len(result) < test.minLength {
			t.Errorf("LCSMultiple(%v) = %q, expected at least length %d", test.strings, result, test.minLength)
		}

		if test.checkLCS && len(result) > 0 {
			for _, s := range test.strings {
				if LCS(s, result) != len(result) {
					t.Errorf("Result %q is not a subsequence of %q", result, s)
				}
			}
		}
	}
}

func TestLCSMultipleEmpty(t *testing.T) {
	result := LCSMultiple([]string{})
	if result != "" {
		t.Errorf("LCSMultiple([]) = %q, expected \"\"", result)
	}
}

func TestLCSMultipleSingle(t *testing.T) {
	result := LCSMultiple([]string{"abc"})
	if result != "abc" {
		t.Errorf("LCSMultiple([\"abc\"]) = %q, expected \"abc\"", result)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{[]int{5}, 5},
		{[]int{5, 3, 7, 1, 9}, 1},
		{[]int{-1, -5, -3}, -5},
		{[]int{0, 0, 0}, 0},
		{[]int{100, 50, 75}, 50},
	}

	for _, test := range tests {
		result := min(test.values...)
		if result != test.expected {
			t.Errorf("min(%v) = %d, expected %d", test.values, result, test.expected)
		}
	}
}

func applyOperations(s1 string, ops []EditOperation) string {
	s2 := ""
	i := 0

	for _, op := range ops {
		switch op.Operation {
		case "match":
			if i < len(s1) {
				s2 += string(s1[i])
				i++
			}
		case "delete":
			if i < len(s1) && s1[i] == op.Char {
				i++
			}
		case "insert":
			s2 += string(op.Char)
		}
	}

	return s2
}

func BenchmarkLCS(b *testing.B) {
	s1 := "AGGTABAGGTABAGGTABAGGTABAGGTAB"
	s2 := "GXTXAYBGXTXAYBGXTXAYBGXTXAYBGXTXAYB"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCS(s1, s2)
	}
}

func BenchmarkLCSString(b *testing.B) {
	s1 := "AGGTABAGGTABAGGTABAGGTABAGGTAB"
	s2 := "GXTXAYBGXTXAYBGXTXAYBGXTXAYBGXTXAYB"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSString(s1, s2)
	}
}

func BenchmarkLCSOptimized(b *testing.B) {
	s1 := "AGGTABAGGTABAGGTABAGGTABAGGTAB"
	s2 := "GXTXAYBGXTXAYBGXTXAYBGXTXAYBGXTXAYB"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSOptimized(s1, s2)
	}
}

func BenchmarkLCSDistance(b *testing.B) {
	s1 := "kittenkittenkittenkittenkitten"
	s2 := "sittingsittingsittingsittingsitting"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSDistance(s1, s2)
	}
}

func BenchmarkLCSSimilarity(b *testing.B) {
	s1 := "AGGTABAGGTABAGGTABAGGTABAGGTAB"
	s2 := "GXTXAYBGXTXAYBGXTXAYBGXTXAYBGXTXAYB"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSSimilarity(s1, s2)
	}
}

func BenchmarkLCSMultiple(b *testing.B) {
	strings := []string{"AGGTABAGGTAB", "GXTXAYBGXTXAYB", "GTABGTABGTAB"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LCSMultiple(strings)
	}
}
