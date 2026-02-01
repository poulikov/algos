package dp

import (
	"sort"
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 0, 0},
		{[]int{1, 2, 5}, 100, 20},
		{[]int{2, 5, 10, 1}, 27, 4},
		{[]int{2}, 1, -1},
		{[]int{1}, 2, 2},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}

	for _, test := range tests {
		result := CoinChange(test.coins, test.amount)
		if result != test.expected {
			t.Errorf("CoinChange(%v, %d) = %d, expected %d", test.coins, test.amount, result, test.expected)
		}
	}
}

func TestCoinChangeWithSolution(t *testing.T) {
	tests := []struct {
		coins         []int
		amount        int
		expectedCount int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 0, 0},
		{[]int{1, 2, 5}, 100, 20},
		{[]int{2, 5, 10, 1}, 27, 4},
	}

	for _, test := range tests {
		count, solution := CoinChangeWithSolution(test.coins, test.amount)
		if count != test.expectedCount {
			t.Errorf("CoinChangeWithSolution(%v, %d) = %d, expected %d", test.coins, test.amount, count, test.expectedCount)
		}
		if count >= 0 {
			sum := 0
			for _, c := range solution {
				sum += c
			}
			if sum != test.amount {
				t.Errorf("CoinChangeWithSolution(%v, %d): solution %v sums to %d, expected %d",
					test.coins, test.amount, solution, sum, test.amount)
			}
			if len(solution) != count {
				t.Errorf("CoinChangeWithSolution(%v, %d): solution length %d != count %d",
					test.coins, test.amount, len(solution), count)
			}
		}
	}
}

func TestCoinChangeCount(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 5, 4},
		{[]int{2}, 3, 0},
		{[]int{1}, 0, 1},
		{[]int{1, 2, 5}, 0, 1},
		{[]int{1, 2, 5}, 10, 10},
		{[]int{2, 3, 5}, 8, 3},
		{[]int{1, 2, 3}, 4, 4},
	}

	for _, test := range tests {
		result := CoinChangeCount(test.coins, test.amount)
		if result != test.expected {
			t.Errorf("CoinChangeCount(%v, %d) = %d, expected %d", test.coins, test.amount, result, test.expected)
		}
	}
}

func TestCoinChangeMaxCoins(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 11},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 0, 0},
		{[]int{1, 2, 5}, 100, 100},
		{[]int{2, 5, 10, 1}, 27, 27},
		{[]int{2}, 4, 2},
	}

	for _, test := range tests {
		result := CoinChangeMaxCoins(test.coins, test.amount)
		if result != test.expected {
			t.Errorf("CoinChangeMaxCoins(%v, %d) = %d, expected %d", test.coins, test.amount, result, test.expected)
		}
	}
}

func TestCoinChangeCombinations(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected [][]int
	}{
		{[]int{1, 2, 5}, 5, [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 2}, {1, 2, 2}, {5}}},
		{[]int{2}, 3, nil},
		{[]int{1}, 0, [][]int{{}}},
	}

	for _, test := range tests {
		result := CoinChangeCombinations(test.coins, test.amount)
		if len(result) != len(test.expected) {
			t.Errorf("CoinChangeCombinations(%v, %d) = %v combinations, expected %v",
				test.coins, test.amount, len(result), len(test.expected))
			continue
		}

		for i, comb := range result {
			sum := 0
			for _, c := range comb {
				sum += c
			}
			if sum != test.amount {
				t.Errorf("CoinChangeCombinations(%v, %d)[%d] = %v sums to %d, expected %d",
					test.coins, test.amount, i, comb, sum, test.amount)
			}
		}
	}
}

func TestCoinChangeConsistency(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2, 5, 10, 1}, 27},
		{[]int{1, 2, 5}, 100},
	}

	for _, test := range tests {
		minCoins := CoinChange(test.coins, test.amount)
		count, solution := CoinChangeWithSolution(test.coins, test.amount)

		if minCoins != count {
			t.Errorf("Inconsistent results: CoinChange = %d, CoinChangeWithSolution = %d", minCoins, count)
		}

		if minCoins >= 0 && len(solution) != minCoins {
			t.Errorf("Solution length %d != minCoins %d", len(solution), minCoins)
		}
	}
}

func TestCoinChangeValidSolution(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2, 5, 10, 1}, 27},
		{[]int{1, 2, 5}, 100},
	}

	for _, test := range tests {
		_, solution := CoinChangeWithSolution(test.coins, test.amount)

		if solution != nil {
			sum := 0
			for _, c := range solution {
				sum += c
			}
			if sum != test.amount {
				t.Errorf("Invalid solution: coins %v sum to %d, expected %d", solution, sum, test.amount)
			}

			for _, c := range solution {
				valid := false
				for _, coin := range test.coins {
					if c == coin {
						valid = true
						break
					}
				}
				if !valid {
					t.Errorf("Invalid coin %d in solution %v for coins %v", c, solution, test.coins)
				}
			}
		}
	}
}

func normalizeSolution(solution []int) []int {
	result := make([]int, len(solution))
	copy(result, solution)
	sort.Ints(result)
	return result
}
