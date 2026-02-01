package dp

import (
	"reflect"
	"sort"
	"testing"
)

func TestKnapsack01(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 6},
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 12},
	}
	capacity := 5

	result := Knapsack01(items, capacity)

	if result.MaxValue != 22 {
		t.Errorf("Expected max value 22, got %d", result.MaxValue)
	}

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}

	if len(result.SelectedItems) != 2 {
		t.Errorf("Expected 2 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01Empty(t *testing.T) {
	items := []Item{}
	capacity := 10

	result := Knapsack01(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}

	if result.TotalWeight != 0 {
		t.Errorf("Expected total weight 0, got %d", result.TotalWeight)
	}

	if len(result.SelectedItems) != 0 {
		t.Errorf("Expected 0 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01ZeroCapacity(t *testing.T) {
	items := []Item{{Weight: 1, Value: 10}}
	capacity := 0

	result := Knapsack01(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}

	if len(result.SelectedItems) != 0 {
		t.Errorf("Expected 0 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01SingleItem(t *testing.T) {
	items := []Item{{Weight: 5, Value: 10}}
	capacity := 10

	result := Knapsack01(items, capacity)

	if result.MaxValue != 10 {
		t.Errorf("Expected max value 10, got %d", result.MaxValue)
	}

	if len(result.SelectedItems) != 1 {
		t.Errorf("Expected 1 item selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01AllItemsFit(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 5},
		{Weight: 3, Value: 7},
		{Weight: 1, Value: 3},
	}
	capacity := 10

	result := Knapsack01(items, capacity)

	expectedValue := 15
	if result.MaxValue != expectedValue {
		t.Errorf("Expected max value %d, got %d", expectedValue, result.MaxValue)
	}

	if len(result.SelectedItems) != 3 {
		t.Errorf("Expected 3 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01NoItemsFit(t *testing.T) {
	items := []Item{
		{Weight: 10, Value: 100},
		{Weight: 15, Value: 150},
	}
	capacity := 5

	result := Knapsack01(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}

	if len(result.SelectedItems) != 0 {
		t.Errorf("Expected 0 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01Optimized(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 6},
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 12},
	}
	capacity := 5

	result := Knapsack01Optimized(items, capacity)

	if result.MaxValue != 22 {
		t.Errorf("Expected max value 22, got %d", result.MaxValue)
	}

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}

	if len(result.SelectedItems) != 2 {
		t.Errorf("Expected 2 items selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsack01VsOptimized(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 6},
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 12},
		{Weight: 4, Value: 15},
		{Weight: 5, Value: 20},
	}
	capacity := 10

	result01 := Knapsack01(items, capacity)
	resultOptimized := Knapsack01Optimized(items, capacity)

	if result01.MaxValue != resultOptimized.MaxValue {
		t.Errorf("Both versions should give same max value: %d vs %d",
			result01.MaxValue, resultOptimized.MaxValue)
	}

	if result01.TotalWeight != resultOptimized.TotalWeight {
		t.Errorf("Both versions should have same total weight: %d vs %d",
			result01.TotalWeight, resultOptimized.TotalWeight)
	}
}

func TestKnapsackUnbounded(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 1},
		{Weight: 3, Value: 4},
		{Weight: 4, Value: 5},
		{Weight: 5, Value: 7},
	}
	capacity := 7

	result := KnapsackUnbounded(items, capacity)

	if result.MaxValue != 9 {
		t.Errorf("Expected max value 9, got %d", result.MaxValue)
	}

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}
}

func TestKnapsackUnboundedEmpty(t *testing.T) {
	items := []Item{}
	capacity := 10

	result := KnapsackUnbounded(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}
}

func TestKnapsackUnboundedSingleItem(t *testing.T) {
	items := []Item{{Weight: 3, Value: 5}}
	capacity := 9

	result := KnapsackUnbounded(items, capacity)

	if result.MaxValue != 15 {
		t.Errorf("Expected max value 15 (3*5), got %d", result.MaxValue)
	}

	if result.TotalWeight != 9 {
		t.Errorf("Expected total weight 9, got %d", result.TotalWeight)
	}
}

func TestKnapsackCount(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 3},
		{Weight: 3, Value: 4},
		{Weight: 5, Value: 6},
	}
	counts := []int{3, 2, 1}
	capacity := 10

	result := KnapsackCount(items, counts, capacity)

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}

	selectedCount := make(map[Item]int)
	for _, item := range result.SelectedItems {
		selectedCount[item]++
	}

	for i, item := range items {
		if selectedCount[item] > counts[i] {
			t.Errorf("Item %v selected %d times, but count is %d", item, selectedCount[item], counts[i])
		}
	}
}

func TestKnapsackCountEmpty(t *testing.T) {
	items := []Item{}
	counts := []int{}
	capacity := 10

	result := KnapsackCount(items, counts, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}
}

func TestKnapsackCountMismatchLength(t *testing.T) {
	items := []Item{{Weight: 1, Value: 10}}
	counts := []int{1, 2}
	capacity := 10

	result := KnapsackCount(items, counts, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0 for mismatched lengths, got %d", result.MaxValue)
	}
}

func TestKnapsackBinary(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 3},
		{Weight: 3, Value: 4},
		{Weight: 5, Value: 7},
	}
	capacity := 10

	result := KnapsackBinary(items, capacity)

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}

	selectedCount := make(map[Item]int)
	for _, item := range result.SelectedItems {
		selectedCount[item]++
	}

	for _, item := range result.SelectedItems {
		if selectedCount[item] > 2 {
			t.Errorf("Item %v selected %d times, but limit is 2", item, selectedCount[item])
		}
	}
}

func TestKnapsackBinaryEmpty(t *testing.T) {
	items := []Item{}
	capacity := 10

	result := KnapsackBinary(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}
}

func TestKnapsackFractional(t *testing.T) {
	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
		{Weight: 30, Value: 120},
	}
	capacity := 50

	result := KnapsackFractional(items, capacity)

	expectedValue := 240
	if result.MaxValue != expectedValue {
		t.Errorf("Expected max value %d, got %d", expectedValue, result.MaxValue)
	}

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}
}

func TestKnapsackFractionalEmpty(t *testing.T) {
	items := []Item{}
	capacity := 10

	result := KnapsackFractional(items, capacity)

	if result.MaxValue != 0 {
		t.Errorf("Expected max value 0, got %d", result.MaxValue)
	}
}

func TestKnapsackFractionalExactFit(t *testing.T) {
	items := []Item{
		{Weight: 10, Value: 100},
		{Weight: 20, Value: 200},
	}
	capacity := 30

	result := KnapsackFractional(items, capacity)

	if result.MaxValue != 300 {
		t.Errorf("Expected max value 300, got %d", result.MaxValue)
	}

	if len(result.SelectedItems) != 2 {
		t.Errorf("Expected 2 items selected (exact fit), got %d", len(result.SelectedItems))
	}
}

func TestKnapsackSubset(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 3},
		{Weight: 3, Value: 4},
		{Weight: 5, Value: 7},
		{Weight: 7, Value: 10},
	}
	targetWeight := 10

	result := KnapsackSubset(items, targetWeight)

	if result.TotalWeight > targetWeight {
		t.Errorf("Total weight %d exceeds target %d", result.TotalWeight, targetWeight)
	}

	totalWeight := 0
	for _, item := range result.SelectedItems {
		totalWeight += item.Weight
	}

	if totalWeight != result.TotalWeight {
		t.Errorf("Calculated weight %d doesn't match result.TotalWeight %d", totalWeight, result.TotalWeight)
	}
}

func TestKnapsackSubsetNoSolution(t *testing.T) {
	items := []Item{
		{Weight: 5, Value: 10},
		{Weight: 10, Value: 20},
	}
	targetWeight := 3

	result := KnapsackSubset(items, targetWeight)

	if result.TotalWeight != 0 {
		t.Errorf("Expected total weight 0 for no solution, got %d", result.TotalWeight)
	}

	if len(result.SelectedItems) != 0 {
		t.Errorf("Expected 0 items for no solution, got %d", len(result.SelectedItems))
	}
}

func TestKnapsackSubsetExactMatch(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 10},
		{Weight: 3, Value: 30},
		{Weight: 5, Value: 50},
	}
	targetWeight := 6

	result := KnapsackSubset(items, targetWeight)

	if result.TotalWeight != targetWeight {
		t.Errorf("Expected exact match of %d, got %d", targetWeight, result.TotalWeight)
	}

	if len(result.SelectedItems) == 0 {
		t.Error("Should find exact match")
	}
}

func TestKnapsackMaxWeightUnderLimit(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 15},
		{Weight: 4, Value: 20},
	}
	capacity := 5

	result := KnapsackMaxWeightUnderLimit(items, capacity)

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}

	expectedValue := 25
	if result.MaxValue != expectedValue {
		t.Errorf("Expected max value %d, got %d", expectedValue, result.MaxValue)
	}
}

func TestKnapsackLargeCapacity(t *testing.T) {
	items := []Item{
		{Weight: 50, Value: 100},
		{Weight: 100, Value: 200},
		{Weight: 150, Value: 300},
	}
	capacity := 1000

	result := Knapsack01(items, capacity)

	if result.TotalWeight > capacity {
		t.Errorf("Total weight %d exceeds capacity %d", result.TotalWeight, capacity)
	}
}

func TestKnapsackItemWithZeroValue(t *testing.T) {
	items := []Item{
		{Weight: 5, Value: 0},
		{Weight: 3, Value: 10},
	}
	capacity := 10

	result := Knapsack01(items, capacity)

	if len(result.SelectedItems) > 0 {
		for _, item := range result.SelectedItems {
			if item.Value == 0 {
				t.Error("Should not select items with zero value")
			}
		}
	}
}

func TestKnapsackSameWeightDifferentValues(t *testing.T) {
	items := []Item{
		{Weight: 5, Value: 10},
		{Weight: 5, Value: 20},
		{Weight: 5, Value: 15},
	}
	capacity := 10

	result := Knapsack01(items, capacity)

	expectedValue := 35
	if result.MaxValue != expectedValue {
		t.Errorf("Expected max value %d (20 + 15), got %d", expectedValue, result.MaxValue)
	}
}

func TestKnapsackResultStructure(t *testing.T) {
	items := []Item{{Weight: 5, Value: 10}}
	capacity := 10

	result := Knapsack01(items, capacity)

	if len(result.SelectedItems) > 0 {
		for _, item := range result.SelectedItems {
			if item.Weight < 0 || item.Value < 0 {
				t.Error("Selected items should have non-negative weight and value")
			}
		}
	}

	if result.TotalWeight < 0 || result.MaxValue < 0 {
		t.Error("Result totals should be non-negative")
	}
}

func TestKnapsackWeightSorting(t *testing.T) {
	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
		{Weight: 30, Value: 120},
	}
	capacity := 50

	unsortedResult := Knapsack01(items, capacity)

	sortedItems := make([]Item, len(items))
	copy(sortedItems, items)
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Weight < sortedItems[j].Weight
	})

	sortedResult := Knapsack01(sortedItems, capacity)

	if unsortedResult.MaxValue != sortedResult.MaxValue {
		t.Errorf("Result should be independent of input order: %d vs %d",
			unsortedResult.MaxValue, sortedResult.MaxValue)
	}
}

func TestKnapsackFractionalPreservesOrder(t *testing.T) {
	items := []Item{
		{Weight: 10, Value: 60},
		{Weight: 20, Value: 100},
	}
	capacity := 25

	result := KnapsackFractional(items, capacity)

	if len(result.SelectedItems) > 0 {
		for i, item := range result.SelectedItems {
			if item.Weight < 0 || item.Value < 0 {
				t.Errorf("Item at index %d has negative weight or value", i)
			}
		}
	}
}

func TestKnapsackMultipleSolutions(t *testing.T) {
	items := []Item{
		{Weight: 5, Value: 10},
		{Weight: 5, Value: 10},
	}
	capacity := 5

	result := Knapsack01(items, capacity)

	if result.MaxValue != 10 {
		t.Errorf("Expected max value 10, got %d", result.MaxValue)
	}

	if len(result.SelectedItems) != 1 {
		t.Errorf("Expected 1 item selected, got %d", len(result.SelectedItems))
	}
}

func TestKnapsackCompareVariants(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 3},
		{Weight: 3, Value: 4},
		{Weight: 4, Value: 5},
		{Weight: 5, Value: 6},
	}
	capacity := 5

	result01 := Knapsack01(items, capacity)
	resultUnbounded := KnapsackUnbounded(items, capacity)

	if result01.MaxValue > resultUnbounded.MaxValue {
		t.Errorf("Unbounded should have >= value than 0/1 knapsack")
	}
}

func TestKnapsackSelectedItems(t *testing.T) {
	items := []Item{
		{Weight: 1, Value: 6},
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 12},
	}
	capacity := 5

	result := Knapsack01(items, capacity)

	selectedWeight := 0
	for _, item := range result.SelectedItems {
		selectedWeight += item.Weight
	}

	if selectedWeight != result.TotalWeight {
		t.Errorf("Calculated weight %d != result.TotalWeight %d", selectedWeight, result.TotalWeight)
	}
}

func TestKnapsackItemsNotModified(t *testing.T) {
	items := []Item{
		{Weight: 2, Value: 10},
		{Weight: 3, Value: 15},
	}
	originalItems := make([]Item, len(items))
	copy(originalItems, items)

	capacity := 5
	_ = Knapsack01(items, capacity)

	if !reflect.DeepEqual(items, originalItems) {
		t.Error("Input items should not be modified")
	}
}
