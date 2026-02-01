package greedy

import (
	"sort"
)

type Activity struct {
	Start int
	End   int
}

type Activities []Activity

func (a Activities) Len() int           { return len(a) }
func (a Activities) Less(i, j int) bool { return a[i].End < a[j].End }
func (a Activities) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ActivitySelection(activities []Activity) []Activity {
	if len(activities) == 0 {
		return []Activity{}
	}

	sorted := make(Activities, len(activities))
	copy(sorted, activities)
	sort.Sort(sorted)

	selected := []Activity{sorted[0]}
	lastEnd := sorted[0].End

	for i := 1; i < len(sorted); i++ {
		if sorted[i].Start >= lastEnd {
			selected = append(selected, sorted[i])
			lastEnd = sorted[i].End
		}
	}

	return selected
}

func ActivitySelectionCount(activities []Activity) int {
	return len(ActivitySelection(activities))
}

func MaxNonOverlappingActivities(activities []Activity) []Activity {
	return ActivitySelection(activities)
}

func ActivitySelectionByWeighted(weightedActivities []WeightedActivity) []WeightedActivity {
	if len(weightedActivities) == 0 {
		return []WeightedActivity{}
	}

	n := len(weightedActivities)

	sorted := make([]WeightedActivity, len(weightedActivities))
	copy(sorted, weightedActivities)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].End < sorted[j].End
	})

	latestCompatible := make([]int, n)
	for i := 0; i < n; i++ {
		latestCompatible[i] = -1
		for j := i - 1; j >= 0; j-- {
			if sorted[j].End <= sorted[i].Start {
				latestCompatible[i] = j
				break
			}
		}
	}

	dp := make([]int, n+1)
	choice := make([]bool, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		exclude := dp[i-1]
		include := sorted[i-1].Weight
		if latestCompatible[i-1] != -1 {
			include += dp[latestCompatible[i-1]+1]
		}

		if include > exclude {
			dp[i] = include
			choice[i] = true
		} else {
			dp[i] = exclude
			choice[i] = false
		}
	}

	var result []WeightedActivity
	i := n
	for i > 0 {
		if choice[i] {
			result = append([]WeightedActivity{sorted[i-1]}, result...)
			i = latestCompatible[i-1] + 1
		} else {
			i--
		}
	}

	return result
}

type WeightedActivity struct {
	Start  int
	End    int
	Weight int
}

type WeightedActivities []WeightedActivity

func (wa WeightedActivities) Len() int           { return len(wa) }
func (wa WeightedActivities) Less(i, j int) bool { return wa[i].End < wa[j].End }
func (wa WeightedActivities) Swap(i, j int)      { wa[i], wa[j] = wa[j], wa[i] }

func WeightedActivitySelection(activities []WeightedActivity) int {
	if len(activities) == 0 {
		return 0
	}

	n := len(activities)

	sorted := make(WeightedActivities, len(activities))
	copy(sorted, activities)
	sort.Sort(sorted)

	latestCompatible := make([]int, n)
	for i := 0; i < n; i++ {
		latestCompatible[i] = -1
		for j := i - 1; j >= 0; j-- {
			if sorted[j].End <= sorted[i].Start {
				latestCompatible[i] = j
				break
			}
		}
	}

	dp := make([]int, n+1)
	dp[0] = 0

	for i := 1; i <= n; i++ {
		exclude := dp[i-1]
		include := sorted[i-1].Weight
		if latestCompatible[i-1] != -1 {
			include += dp[latestCompatible[i-1]+1]
		}
		dp[i] = max(exclude, include)
	}

	return dp[n]
}

func Compatible(activities []Activity) bool {
	for i := 0; i < len(activities)-1; i++ {
		if activities[i].End > activities[i+1].Start {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
