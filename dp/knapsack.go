package dp

// Item represents an item for the knapsack problem
type Item struct {
	Weight int
	Value  int
}

// KnapsackResult represents the result of the knapsack problem
type KnapsackResult struct {
	MaxValue      int    // Maximum value achievable
	SelectedItems []Item // Items selected to achieve the max value
	TotalWeight   int    // Total weight of selected items
}

// Knapsack01 solves the 0/1 knapsack problem using dynamic programming
// Each item can be used at most once (0 or 1 times)
// Time complexity: O(n * capacity)
// Space complexity: O(n * capacity)
func Knapsack01(items []Item, capacity int) *KnapsackResult {
	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if items[i-1].Weight <= w {
				exclude := dp[i-1][w]
				include := dp[i-1][w-items[i-1].Weight] + items[i-1].Value

				if include > exclude {
					dp[i][w] = include
				} else {
					dp[i][w] = exclude
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	result := &KnapsackResult{
		MaxValue:      dp[n][capacity],
		SelectedItems: []Item{},
		TotalWeight:   0,
	}

	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if w-items[i-1].Weight >= 0 && dp[i][w] != dp[i-1][w] {
			result.SelectedItems = append([]Item{items[i-1]}, result.SelectedItems...)
			result.TotalWeight += items[i-1].Weight
			w -= items[i-1].Weight
		}
	}

	return result
}

// Knapsack01Optimized solves the 0/1 knapsack with optimized space
// Time complexity: O(n * capacity)
// Space complexity: O(capacity)
func Knapsack01Optimized(items []Item, capacity int) *KnapsackResult {
	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	dp := make([]int, capacity+1)
	choice := make([]int, capacity+1)

	for i := 1; i <= n; i++ {
		for w := capacity; w >= items[i-1].Weight; w-- {
			exclude := dp[w]
			include := dp[w-items[i-1].Weight] + items[i-1].Value

			if include > exclude {
				dp[w] = include
				choice[w] = i
			}
		}
	}

	maxValue := dp[capacity]
	selectedItems := []Item{}
	w := capacity

	for w > 0 && choice[w] > 0 {
		itemIndex := choice[w] - 1
		selectedItems = append([]Item{items[itemIndex]}, selectedItems...)
		w -= items[itemIndex].Weight
	}

	totalWeight := 0
	for _, item := range selectedItems {
		totalWeight += item.Weight
	}

	return &KnapsackResult{
		MaxValue:      maxValue,
		SelectedItems: selectedItems,
		TotalWeight:   totalWeight,
	}
}

// KnapsackUnbounded solves the unbounded knapsack problem
// Each item can be used unlimited times
// Time complexity: O(n * capacity)
// Space complexity: O(capacity)
func KnapsackUnbounded(items []Item, capacity int) *KnapsackResult {
	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	dp := make([]int, capacity+1)
	choice := make([]int, capacity+1)

	for w := 0; w <= capacity; w++ {
		for i, item := range items {
			if item.Weight <= w {
				valueWithItem := dp[w-item.Weight] + item.Value
				if valueWithItem > dp[w] {
					dp[w] = valueWithItem
					choice[w] = i
				}
			}
		}
	}

	selectedItems := []Item{}
	w := capacity

	for w > 0 && choice[w] >= 0 {
		itemIndex := choice[w]
		selectedItems = append([]Item{items[itemIndex]}, selectedItems...)
		w -= items[itemIndex].Weight
	}

	totalWeight := 0
	for _, item := range selectedItems {
		totalWeight += item.Weight
	}

	return &KnapsackResult{
		MaxValue:      dp[capacity],
		SelectedItems: selectedItems,
		TotalWeight:   totalWeight,
	}
}

// KnapsackCount solves the knapsack problem with item counts
// Each item can be used at most count[i] times
// Time complexity: O(n * capacity * max_count)
// Space complexity: O(n * capacity)
func KnapsackCount(items []Item, counts []int, capacity int) *KnapsackResult {
	if len(items) != len(counts) {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	dp := make([][]int, n+1)
	choice := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
		choice[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			maxVal := 0
			bestK := 0

			for k := 0; k <= counts[i-1] && k*items[i-1].Weight <= w; k++ {
				val := dp[i-1][w-k*items[i-1].Weight] + k*items[i-1].Value
				if val > maxVal {
					maxVal = val
					bestK = k
				}
			}

			dp[i][w] = maxVal
			choice[i][w] = bestK
		}
	}

	result := &KnapsackResult{
		MaxValue:      dp[n][capacity],
		SelectedItems: []Item{},
		TotalWeight:   0,
	}

	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		k := choice[i][w]
		for j := 0; j < k; j++ {
			result.SelectedItems = append([]Item{items[i-1]}, result.SelectedItems...)
			result.TotalWeight += items[i-1].Weight
		}
		w -= k * items[i-1].Weight
	}

	return result
}

// KnapsackBinary solves the knapsack problem for binary items
// Each item can be used at most twice (2^1 times)
// Time complexity: O(n * capacity)
// Space complexity: O(capacity)
func KnapsackBinary(items []Item, capacity int) *KnapsackResult {
	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	dp := make([]int, capacity+1)
	choice := make([]int, capacity+1)

	for i := 1; i <= n; i++ {
		item := items[i-1]

		for w := capacity; w >= 0; w-- {
			maxVal := dp[w]
			bestCount := 0

			for k := 1; k <= 2 && k*item.Weight <= w; k++ {
				val := dp[w-k*item.Weight] + k*item.Value
				if val > maxVal {
					maxVal = val
					bestCount = k
				}
			}

			dp[w] = maxVal
			choice[w] = bestCount
		}
	}

	maxValue := dp[capacity]
	selectedItems := []Item{}
	w := capacity

	for w > 0 {
		currentMax := 0
		bestItemIndex := -1
		bestCount := 0

		for i, item := range items {
			for k := 1; k <= 2 && k*item.Weight <= w; k++ {
				val := dp[w-k*item.Weight] + k*item.Value
				if val > currentMax {
					currentMax = val
					bestItemIndex = i
					bestCount = k
				}
			}
		}

		if bestItemIndex >= 0 && currentMax == dp[w] {
			for k := 0; k < bestCount; k++ {
				selectedItems = append(selectedItems, items[bestItemIndex])
			}
			w -= bestCount * items[bestItemIndex].Weight
		} else {
			break
		}
	}

	totalWeight := 0
	for _, item := range selectedItems {
		totalWeight += item.Weight
	}

	return &KnapsackResult{
		MaxValue:      maxValue,
		SelectedItems: selectedItems,
		TotalWeight:   totalWeight,
	}
}

// KnapsackFractional solves the fractional knapsack problem
// Items can be broken into fractional parts
// Time complexity: O(n log n) for sorting
// Space complexity: O(1)
func KnapsackFractional(items []Item, capacity int) *KnapsackResult {
	if capacity <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	ratios := make([]float64, n)

	for i, item := range items {
		if item.Weight > 0 {
			ratios[i] = float64(item.Value) / float64(item.Weight)
		} else {
			ratios[i] = 0
		}
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if ratios[j] > ratios[i] {
				ratios[i], ratios[j] = ratios[j], ratios[i]
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	result := &KnapsackResult{
		MaxValue:      0,
		SelectedItems: []Item{},
		TotalWeight:   0,
	}

	remainingCapacity := capacity

	for i := 0; i < n && remainingCapacity > 0; i++ {
		item := items[i]

		if item.Weight <= remainingCapacity {
			result.SelectedItems = append(result.SelectedItems, item)
			result.TotalWeight += item.Weight
			result.MaxValue += item.Value
			remainingCapacity -= item.Weight
		} else {
			fraction := float64(remainingCapacity) / float64(item.Weight)
			fractionalItem := Item{
				Weight: remainingCapacity,
				Value:  int(float64(item.Value) * fraction),
			}
			result.SelectedItems = append(result.SelectedItems, fractionalItem)
			result.TotalWeight += remainingCapacity
			result.MaxValue += fractionalItem.Value
			break
		}
	}

	return result
}

// KnapsackSubset solves the knapsack to find a subset with exact weight
// Returns items with total weight equal to target weight (or as close as possible)
// Time complexity: O(n * capacity)
// Space complexity: O(n * capacity)
func KnapsackSubset(items []Item, targetWeight int) *KnapsackResult {
	if targetWeight <= 0 || len(items) == 0 {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	n := len(items)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, targetWeight+1)
	}
	dp[0][0] = true

	for i := 1; i <= n; i++ {
		for w := 0; w <= targetWeight; w++ {
			dp[i][w] = dp[i-1][w]

			if w >= items[i-1].Weight && dp[i-1][w-items[i-1].Weight] {
				dp[i][w] = true
			}
		}
	}

	if !dp[n][targetWeight] {
		return &KnapsackResult{
			MaxValue:      0,
			SelectedItems: []Item{},
			TotalWeight:   0,
		}
	}

	selectedItems := []Item{}
	w := targetWeight

	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			selectedItems = append([]Item{items[i-1]}, selectedItems...)
			w -= items[i-1].Weight
		}
	}

	totalValue := 0
	for _, item := range selectedItems {
		totalValue += item.Value
	}

	return &KnapsackResult{
		MaxValue:      totalValue,
		SelectedItems: selectedItems,
		TotalWeight:   targetWeight - w,
	}
}

// KnapsackMaxWeightUnderLimit finds maximum value with weight under a limit
// Returns maximum achievable value where total weight <= capacity
// Time complexity: O(n * capacity)
// Space complexity: O(capacity)
func KnapsackMaxWeightUnderLimit(items []Item, capacity int) *KnapsackResult {
	return Knapsack01(items, capacity)
}
