package strings

// RabinKarp searches for a pattern in a text using Rabin-Karp algorithm
// Uses rolling hash for efficient pattern matching
// Time complexity: O(n + m) average, O(nm) worst case where n is text length and m is pattern length
func RabinKarp(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{}
	}

	if len(text) == 0 || len(pattern) > len(text) {
		return []int{}
	}

	matches := []int{}
	const base int64 = 256
	const prime int64 = 101

	m := len(pattern)
	n := len(text)

	patternHash := computeHash(pattern, base, prime)

	for i := 0; i <= n-m; i++ {
		textHash := computeHash(text[i:i+m], base, prime)
		if patternHash == textHash {
			if checkEqual(text, pattern, i, i+m) {
				matches = append(matches, i)
			}
		}
	}

	return matches
}

func computeHash(s string, base, prime int64) int64 {
	hash := int64(0)
	for i := 0; i < len(s); i++ {
		hash = (hash*base + int64(s[i])) % prime
	}
	return hash
}

// checkEqual checks if text[index1:index2] equals pattern
// Used to handle hash collisions
func checkEqual(text, pattern string, index1, index2 int) bool {
	for i := 0; i < len(pattern); i++ {
		if text[index1+i] != pattern[i] {
			return false
		}
	}
	return true
}

// RabinKarpMultiplePatterns searches for multiple patterns in text sequentially
// Returns a map from pattern to list of indices where it occurs
// Time complexity: O(k * (n + m)) where k is number of patterns
func RabinKarpMultiplePatterns(text string, patterns []string) map[string][]int {
	if len(patterns) == 0 || len(text) == 0 {
		return map[string][]int{}
	}

	result := make(map[string][]int)

	for _, pattern := range patterns {
		matches := RabinKarp(text, pattern)
		if len(matches) > 0 {
			result[pattern] = matches
		}
	}

	return result
}
