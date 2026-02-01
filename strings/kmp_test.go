package strings

import (
	"testing"
)

func TestKMP(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		pattern  string
		expected []int
	}{
		{
			name:     "Pattern found",
			text:     "ABABDABACDABABCABAB",
			pattern:  "ABABCABAB",
			expected: []int{10},
		},
		{
			name:     "Pattern not found",
			text:     "Hello World",
			pattern:  "Python",
			expected: []int{},
		},
		{
			name:     "Pattern longer than text",
			text:     "Hi",
			pattern:  "Hello World",
			expected: []int{},
		},
		{
			name:     "Pattern equal to text",
			text:     "Test",
			pattern:  "Test",
			expected: []int{0},
		},
		{
			name:     "Multiple occurrences",
			text:     "ABABABA",
			pattern:  "ABA",
			expected: []int{0, 2, 4},
		},
		{
			name:     "Empty pattern",
			text:     "Hello",
			pattern:  "",
			expected: []int{},
		},
		{
			name:     "Single character",
			text:     "ABACADABRAC",
			pattern:  "A",
			expected: []int{0, 2, 4, 6, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KMP(tt.text, tt.pattern)

			if len(result) != len(tt.expected) {
				t.Errorf("KMP() returned %d matches, expected %d", len(result), len(tt.expected))
			}

			for i, exp := range tt.expected {
				if i >= len(result) {
					t.Errorf("KMP()[%d] missing, expected %d", i, exp)
				} else if result[i] != exp {
					t.Errorf("KMP()[%d] = %d, expected %d", i, result[i], exp)
				}
			}
		})
	}
}

func TestKMPMultiplePatterns(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		patterns []string
		expected map[string][]int
	}{
		{
			name:     "Multiple patterns found",
			text:     "Hello World Hello",
			patterns: []string{"Hello", "World", "lo"},
			expected: map[string][]int{
				"Hello": {0, 12},
				"World": {6},
				"lo":    {3, 15},
			},
		},
		{
			name:     "Some patterns found",
			text:     "ABABABA",
			patterns: []string{"ABA", "BAB", "ABC"},
			expected: map[string][]int{
				"ABA": {0, 2, 4},
				"BAB": {1, 3},
			},
		},
		{
			name:     "No patterns found",
			text:     "Hello World",
			patterns: []string{"Python", "Java"},
			expected: map[string][]int{},
		},
		{
			name:     "Empty patterns list",
			text:     "Hello",
			patterns: []string{},
			expected: map[string][]int{},
		},
		{
			name:     "Empty text",
			text:     "",
			patterns: []string{"Hello", "World"},
			expected: map[string][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := KMPMultiplePatterns(tt.text, tt.patterns)

			if len(result) != len(tt.expected) {
				t.Errorf("KMPMultiplePatterns() returned %d patterns, expected %d", len(result), len(tt.expected))
			}

			for pattern, expectedIndices := range tt.expected {
				resultIndices, ok := result[pattern]
				if !ok {
					t.Errorf("Pattern %q not found in result", pattern)
				} else if len(resultIndices) != len(expectedIndices) {
					t.Errorf("Pattern %q: got %d matches, expected %d", pattern, len(resultIndices), len(expectedIndices))
				}

				for i, exp := range expectedIndices {
					if i >= len(resultIndices) {
						t.Errorf("Pattern %q: missing match at index %d", pattern, exp)
					} else if resultIndices[i] != exp {
						t.Errorf("Pattern %q: match at index %d = %d, expected %d", pattern, i, resultIndices[i], exp)
					}
				}
			}
		})
	}
}
