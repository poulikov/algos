package dp

type Operation int

const (
	Insert Operation = iota
	Delete
	Replace
	Match
)

type EditPath struct {
	Char rune
	Op   Operation
	Pos  int
}

func EditDistance(s1, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				replace := dp[i-1][j-1] + 1
				del := dp[i-1][j] + 1
				insert := dp[i][j-1] + 1
				dp[i][j] = replace
				if del < dp[i][j] {
					dp[i][j] = del
				}
				if insert < dp[i][j] {
					dp[i][j] = insert
				}
			}
		}
	}

	return dp[m][n]
}

func EditDistanceOptimized(s1, s2 string) int {
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	m, n := len(s1), len(s2)
	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for j := range prev {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr[0] = i
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1]
			} else {
				replace := prev[j-1] + 1
				del := prev[j] + 1
				insert := curr[j-1] + 1
				curr[j] = replace
				if del < curr[j] {
					curr[j] = del
				}
				if insert < curr[j] {
					curr[j] = insert
				}
			}
		}
		prev, curr = curr, prev
	}

	return prev[n]
}

func EditDistanceWithPath(s1, s2 string) (int, []EditPath) {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				replace := dp[i-1][j-1] + 1
				del := dp[i-1][j] + 1
				insert := dp[i][j-1] + 1
				dp[i][j] = replace
				if del < dp[i][j] {
					dp[i][j] = del
				}
				if insert < dp[i][j] {
					dp[i][j] = insert
				}
			}
		}
	}

	i, j := m, n
	var path []EditPath

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && s1[i-1] == s2[j-1] {
			path = append([]EditPath{{Char: rune(s1[i-1]), Op: Match}}, path...)
			i--
			j--
		} else if i > 0 && j > 0 && dp[i-1][j-1]+1 == dp[i][j] {
			path = append([]EditPath{{Char: rune(s2[j-1]), Op: Replace}}, path...)
			i--
			j--
		} else if i > 0 && dp[i-1][j]+1 == dp[i][j] {
			path = append([]EditPath{{Char: rune(s1[i-1]), Op: Delete}}, path...)
			i--
		} else if j > 0 && dp[i][j-1]+1 == dp[i][j] {
			path = append([]EditPath{{Char: rune(s2[j-1]), Op: Insert}}, path...)
			j--
		} else if i > 0 {
			path = append([]EditPath{{Char: rune(s1[i-1]), Op: Delete}}, path...)
			i--
		} else if j > 0 {
			path = append([]EditPath{{Char: rune(s2[j-1]), Op: Insert}}, path...)
			j--
		}
	}

	return dp[m][n], path
}
