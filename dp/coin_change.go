package dp

import (
	"math"
)

const MAX_COINS = math.MaxInt

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	return dp[amount]
}

func CoinChangeWithSolution(coins []int, amount int) (int, []int) {
	if amount == 0 {
		return 0, []int{}
	}
	if len(coins) == 0 {
		return -1, nil
	}

	dp := make([]int, amount+1)
	prev := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		prev[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
					prev[i] = coin
				}
			}
		}
	}

	if dp[amount] == -1 {
		return -1, nil
	}

	solution := []int{}
	remaining := amount
	for remaining > 0 {
		coin := prev[remaining]
		solution = append(solution, coin)
		remaining -= coin
	}

	return dp[amount], solution
}

func CoinChangeCount(coins []int, amount int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func CoinChangeMaxCoins(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i-coin]+1 > dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	return dp[amount]
}

func CoinChangeCombinations(coins []int, amount int) [][]int {
	if amount == 0 {
		return [][]int{{}}
	}
	if len(coins) == 0 {
		return nil
	}

	dp := make([][][]int, amount+1)
	dp[0] = [][]int{{}}

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if len(dp[i-coin]) > 0 {
				for _, comb := range dp[i-coin] {
					newComb := make([]int, len(comb))
					copy(newComb, comb)
					newComb = append(newComb, coin)
					dp[i] = append(dp[i], newComb)
				}
			}
		}
	}

	return dp[amount]
}
