package dp

import (
	"sort"
)

func LISLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

func LISLengthOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := []int{nums[0]}

	for _, num := range nums[1:] {
		if num > tails[len(tails)-1] {
			tails = append(tails, num)
		} else {
			idx := sort.SearchInts(tails, num)
			tails[idx] = num
		}
	}

	return len(tails)
}

func LIS(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	dp := make([]int, n)
	prev := make([]int, n)
	maxLen := 1
	maxIdx := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIdx = i
		}
	}

	result := make([]int, 0, maxLen)
	for i := maxIdx; i >= 0; i = prev[i] {
		result = append([]int{nums[i]}, result...)
	}

	return result
}

func LISOptimized(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	tails := make([]int, 0, n)
	tailIndices := make([]int, 0, n)
	prevIndices := make([]int, n)
	for i := range prevIndices {
		prevIndices[i] = -1
	}

	tails = append(tails, nums[0])
	tailIndices = append(tailIndices, 0)

	for i := 1; i < n; i++ {
		if nums[i] > tails[len(tails)-1] {
			prevIndices[i] = tailIndices[len(tailIndices)-1]
			tails = append(tails, nums[i])
			tailIndices = append(tailIndices, i)
		} else {
			idx := sort.SearchInts(tails, nums[i])
			if idx > 0 {
				prevIndices[i] = tailIndices[idx-1]
			}
			tails[idx] = nums[i]
			tailIndices[idx] = i
		}
	}

	result := make([]int, 0, len(tails))
	for idx := tailIndices[len(tailIndices)-1]; idx != -1; idx = prevIndices[idx] {
		result = append([]int{nums[idx]}, result...)
	}

	return result
}

func LNDSLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] <= nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}
