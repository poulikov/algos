package dp

// LCSResult represents the result of Longest Common Subsequence algorithm
type LCSResult struct {
	Length      int    // Length of the LCS
	Subsequence string // The actual LCS string
}

// LCS finds the length of the Longest Common Subsequence between two strings
// Time complexity: O(m * n) where m and n are the lengths of the two strings
// Space complexity: O(m * n)
func LCS(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}

// LCSString finds the actual Longest Common Subsequence between two strings
// Returns the LCS string
// Time complexity: O(m * n)
// Space complexity: O(m * n)
func LCSString(s1, s2 string) string {
	m := len(s1)
	n := len(s2)

	if m == 0 || n == 0 {
		return ""
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	lcs := ""
	i := m
	j := n

	for i > 0 && j > 0 {
		if s1[i-1] == s2[j-1] {
			lcs = string(s1[i-1]) + lcs
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs
}

// LCSOptimized finds the length of LCS with optimized space complexity
// Time complexity: O(m * n)
// Space complexity: O(min(m, n))
func LCSOptimized(s1, s2 string) int {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	m := len(s1)
	n := len(s2)

	if n == 0 {
		return 0
	}

	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				if prev[j] > curr[j-1] {
					curr[j] = prev[j]
				} else {
					curr[j] = curr[j-1]
				}
			}
		}

		for j := 0; j <= n; j++ {
			prev[j] = curr[j]
		}
	}

	return curr[n]
}

// LCSAll finds all possible LCS strings between two strings
// Returns a slice of all possible LCS strings
// Time complexity: O(m * n * 2^(m+n)) in worst case
func LCSAll(s1, s2 string) []string {
	m := len(s1)
	n := len(s2)

	if m == 0 || n == 0 {
		return []string{""}
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	uniqueResults := make(map[string]bool)
	findAllLCS(dp, s1, s2, m, n, "", uniqueResults)

	results := make([]string, 0, len(uniqueResults))
	for result := range uniqueResults {
		results = append(results, result)
	}

	return results
}

// findAllLCS recursively finds all possible LCS strings
func findAllLCS(dp [][]int, s1, s2 string, i, j int, current string, results map[string]bool) {
	if i == 0 || j == 0 {
		results[current] = true
		return
	}

	if s1[i-1] == s2[j-1] {
		findAllLCS(dp, s1, s2, i-1, j-1, string(s1[i-1])+current, results)
	} else {
		if dp[i-1][j] >= dp[i][j-1] {
			findAllLCS(dp, s1, s2, i-1, j, current, results)
		}
		if dp[i][j-1] >= dp[i-1][j] {
			findAllLCS(dp, s1, s2, i, j-1, current, results)
		}
	}
}

// LCSDistance computes the edit distance between two strings
// Edit distance is the minimum number of insertions, deletions, or substitutions
// required to transform s1 into s2
// Time complexity: O(m * n)
// Space complexity: O(m * n)
func LCSDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,
				min(
					dp[i][j-1]+1,
					dp[i-1][j-1]+cost,
				),
			)
		}
	}

	return dp[m][n]
}

// LCSDistanceWithOperations computes edit distance and the sequence of operations
// Returns the distance and a description of operations
// Time complexity: O(m * n)
// Space complexity: O(m * n)
type EditOperation struct {
	Operation string // "insert", "delete", "replace", "match"
	Char      byte   // Character involved
}

func LCSDistanceWithOperations(s1, s2 string) (int, []EditOperation) {
	m := len(s1)
	n := len(s2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,
				min(
					dp[i][j-1]+1,
					dp[i-1][j-1]+cost,
				),
			)
		}
	}

	operations := []EditOperation{}
	i := m
	j := n

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && s1[i-1] == s2[j-1] {
			operations = append([]EditOperation{{Operation: "match", Char: s1[i-1]}}, operations...)
			i--
			j--
		} else if j > 0 && (i == 0 || dp[i-1][j] >= dp[i][j-1]) {
			operations = append([]EditOperation{{Operation: "insert", Char: s2[j-1]}}, operations...)
			j--
		} else {
			operations = append([]EditOperation{{Operation: "delete", Char: s1[i-1]}}, operations...)
			i--
		}
	}

	return dp[m][n], operations
}

// LCSSimilarity computes similarity ratio between two strings
// Returns a value between 0 and 1 (1 = identical, 0 = completely different)
// Time complexity: O(m * n)
func LCSSimilarity(s1, s2 string) float64 {
	lcsLength := LCS(s1, s2)
	maxLength := len(s1)
	if len(s2) > maxLength {
		maxLength = len(s2)
	}

	if maxLength == 0 {
		return 0.0
	}

	return float64(lcsLength) / float64(maxLength)
}

// min returns the minimum of multiple integers
func min(values ...int) int {
	minValue := values[0]

	for _, v := range values[1:] {
		if v < minValue {
			minValue = v
		}
	}

	return minValue
}

// LCSMultiple finds the LCS of multiple strings by sequential application
// IMPORTANT: This function finds the LCS by sequentially applying LCSString
// to pairs of strings, which may NOT give the optimal result for finding
// the LCS of all strings simultaneously. For the exact LCS of all strings,
// use DP with n-dimensional tables where n is the number of strings.
//
// Example where sequential application fails:
// LCSMultiple(["ABCD", "ABDC", "ACBD"]) may return different results
// compared to optimal DP solution for all three strings simultaneously.
//
// Time complexity: O(k * m * n) where k is number of strings,
// and m, n are average string lengths
func LCSMultiple(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	if len(strings) == 1 {
		return strings[0]
	}

	lcs := strings[0]

	for i := 1; i < len(strings); i++ {
		lcs = LCSString(lcs, strings[i])
	}

	return lcs
}
