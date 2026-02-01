package greedy

import (
	"testing"
)

func TestActivitySelection(t *testing.T) {
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 7},
		{6, 9},
		{8, 10},
	}

	selected := ActivitySelection(activities)

	if len(selected) != 3 {
		t.Errorf("Expected 3 activities, got %d", len(selected))
	}

	if !Compatible(selected) {
		t.Errorf("Selected activities are not compatible")
	}
}

func TestActivitySelectionCount(t *testing.T) {
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 7},
		{6, 9},
		{8, 10},
	}

	count := ActivitySelectionCount(activities)
	expected := 3

	if count != expected {
		t.Errorf("Expected count %d, got %d", expected, count)
	}
}

func TestActivitySelectionEmpty(t *testing.T) {
	activities := []Activity{}
	selected := ActivitySelection(activities)

	if len(selected) != 0 {
		t.Errorf("Expected 0 activities, got %d", len(selected))
	}
}

func TestActivitySelectionSingle(t *testing.T) {
	activities := []Activity{{1, 3}}
	selected := ActivitySelection(activities)

	if len(selected) != 1 {
		t.Errorf("Expected 1 activity, got %d", len(selected))
	}
}

func TestActivitySelectionAllCompatible(t *testing.T) {
	activities := []Activity{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}

	selected := ActivitySelection(activities)

	if len(selected) != 4 {
		t.Errorf("Expected 4 activities, got %d", len(selected))
	}
}

func TestActivitySelectionNoneCompatible(t *testing.T) {
	activities := []Activity{
		{1, 5},
		{2, 6},
		{3, 7},
	}

	selected := ActivitySelection(activities)

	if len(selected) != 1 {
		t.Errorf("Expected 1 activity, got %d", len(selected))
	}
}

func TestActivitySelectionByWeighted(t *testing.T) {
	activities := []WeightedActivity{
		{1, 3, 5},
		{2, 5, 6},
		{4, 7, 5},
		{6, 9, 8},
		{5, 8, 7},
	}

	selected := ActivitySelectionByWeighted(activities)

	if len(selected) == 0 {
		t.Errorf("Expected some activities, got none")
	}

	totalWeight := 0
	for _, a := range selected {
		totalWeight += a.Weight
	}

	if totalWeight == 0 {
		t.Errorf("Expected non-zero total weight")
	}
}

func TestWeightedActivitySelection(t *testing.T) {
	activities := []WeightedActivity{
		{1, 3, 5},
		{2, 5, 6},
		{4, 7, 5},
		{6, 9, 8},
		{5, 8, 7},
	}

	maxWeight := WeightedActivitySelection(activities)

	if maxWeight <= 0 {
		t.Errorf("Expected positive max weight, got %d", maxWeight)
	}
}

func TestCompatible(t *testing.T) {
	tests := []struct {
		activities []Activity
		expected   bool
	}{
		{
			[]Activity{
				{1, 2},
				{2, 3},
				{3, 4},
			},
			true,
		},
		{
			[]Activity{
				{1, 3},
				{2, 4},
				{3, 5},
			},
			false,
		},
		{
			[]Activity{
				{1, 2},
			},
			true,
		},
		{
			[]Activity{},
			true,
		},
	}

	for _, test := range tests {
		result := Compatible(test.activities)
		if result != test.expected {
			t.Errorf("Compatible(%v) = %v, expected %v", test.activities, result, test.expected)
		}
	}
}

func TestActivitySelectionConsistency(t *testing.T) {
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 7},
		{6, 9},
		{8, 10},
	}

	selected := ActivitySelection(activities)
	count := ActivitySelectionCount(activities)

	if len(selected) != count {
		t.Errorf("Inconsistent: ActivitySelection returned %d, ActivitySelectionCount returned %d",
			len(selected), count)
	}

	if !Compatible(selected) {
		t.Errorf("Selected activities are not compatible")
	}
}

func TestMaxNonOverlappingActivities(t *testing.T) {
	activities := []Activity{
		{1, 3},
		{2, 5},
		{4, 7},
		{6, 9},
		{8, 10},
	}

	selected := MaxNonOverlappingActivities(activities)

	if len(selected) != 3 {
		t.Errorf("Expected 3 activities, got %d", len(selected))
	}

	if !Compatible(selected) {
		t.Errorf("Selected activities are not compatible")
	}
}

func TestActivitySelectionLarge(t *testing.T) {
	n := 100
	activities := make([]Activity, n)
	for i := 0; i < n; i++ {
		activities[i] = Activity{i * 2, i*2 + 1}
	}

	selected := ActivitySelection(activities)

	if len(selected) != n {
		t.Errorf("Expected %d activities, got %d", n, len(selected))
	}
}
