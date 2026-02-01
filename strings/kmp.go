package strings

// KMP searches for a pattern in a text using the Knuth-Morris-Pratt algorithm
// Builds a prefix function (also called failure function) to avoid rechecking characters
// Time complexity: O(n + m) where n is text length and m is pattern length
func KMP(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	if len(text) == 0 || len(pattern) > len(text) {
		return []int{}
	}

	matches := []int{}
	prefixFunc := buildPrefixFunction(pattern)

	n := len(text)
	m := len(pattern)
	i := 0
	j := 0

	for i < n {
		if j < m && pattern[j] == text[i] {
			i++
			j++
		} else {
			if j != 0 {
				j = prefixFunc[j-1]
			} else {
				i++
			}
		}

		if j == m {
			matches = append(matches, i-j)
			j = prefixFunc[j-1]
		}
	}

	return matches
}

// buildPrefixFunction constructs the prefix function for KMP algorithm
// The prefix function at index i tells us the length of the longest proper prefix
// of pattern[0:i+1] that is also a suffix of pattern[0:i+1]
// Time complexity: O(m) where m is pattern length
func buildPrefixFunction(pattern string) []int {
	m := len(pattern)
	prefixFunc := make([]int, m)
	prefixFunc[0] = 0

	i := 1
	j := 0

	for i < m {
		if pattern[i] == pattern[j] {
			j++
			prefixFunc[i] = j
			i++
		} else {
			if j > 0 {
				j = prefixFunc[j-1]
			} else {
				prefixFunc[i] = 0
				i++
			}
		}
	}

	return prefixFunc
}

// KMPFirst returns the first occurrence of pattern in text
// Returns -1 if pattern is not found
// Time complexity: O(n + m)
func KMPFirst(text, pattern string) int {
	matches := KMP(text, pattern)

	if len(matches) == 0 {
		return -1
	}

	return matches[0]
}

// KMPCount returns the count of pattern occurrences in text
// Time complexity: O(n + m)
func KMPCount(text, pattern string) int {
	return len(KMP(text, pattern))
}

// KMPSearchAll returns all occurrences of pattern in text with overlap allowed
// For example, pattern "aa" in text "aaaa" returns 3 matches (indices 0, 1, 2)
// Time complexity: O(n + m)
func KMPSearchAll(text, pattern string) []int {
	return KMP(text, pattern)
}

// KMPSearchNonOverlapping returns all occurrences of pattern in text without overlap
// For example, pattern "aa" in text "aaaa" returns 2 matches (indices 0, 2)
// Time complexity: O(n + m)
func KMPSearchNonOverlapping(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	matches := []int{}
	prefixFunc := buildPrefixFunction(pattern)

	i := 0
	j := 0

	for i < len(text) {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == len(pattern) {
			matches = append(matches, i-j)
			i += len(pattern)
			j = 0
		} else if i < len(text) && pattern[j] != text[i] {
			if j > 0 {
				j = prefixFunc[j-1]
			} else {
				i++
			}
		}
	}

	return matches
}

// KMPReplace replaces all occurrences of pattern in text with replacement
// Returns the modified text
// Time complexity: O(n + m)
func KMPReplace(text, pattern, replacement string) string {
	if len(pattern) == 0 {
		return text
	}

	matches := KMPSearchNonOverlapping(text, pattern)

	if len(matches) == 0 {
		return text
	}

	result := ""
	lastIndex := 0

	for _, match := range matches {
		result += text[lastIndex:match] + replacement
		lastIndex = match + len(pattern)
	}

	result += text[lastIndex:]

	return result
}

// KMPCanFindPattern checks if pattern exists in text
// Returns true if pattern is found, false otherwise
// Time complexity: O(n + m) in worst case, often less
func KMPCanFindPattern(text, pattern string) bool {
	return KMPFirst(text, pattern) >= 0
}

// KMPMultiplePatterns searches for multiple patterns in text sequentially
// Returns a map from pattern to list of indices where it occurs
// Note: This function runs KMP separately for each pattern.
// For simultaneous multi-pattern search with optimal O(n + m) complexity,
// use the Aho-Corasick algorithm instead.
// Time complexity: O(k * (n + m)) where k is number of patterns,
// n is text length, and m is average pattern length
func KMPMultiplePatterns(text string, patterns []string) map[string][]int {
	if len(patterns) == 0 || len(text) == 0 {
		return map[string][]int{}
	}

	result := make(map[string][]int)

	for _, pattern := range patterns {
		matches := KMP(text, pattern)
		if len(matches) > 0 {
			result[pattern] = matches
		}
	}

	return result
}

// KMPFindAllPatternsWithMatches finds all patterns that exist in text
// Returns a map from pattern to the count of its occurrences
// Time complexity: O(k * (n + m)) where k is number of patterns
func KMPFindAllPatternsWithMatches(text string, patterns []string) map[string]int {
	if len(patterns) == 0 || len(text) == 0 {
		return map[string]int{}
	}

	result := make(map[string]int)

	for _, pattern := range patterns {
		count := KMPCount(text, pattern)
		if count > 0 {
			result[pattern] = count
		}
	}

	return result
}
