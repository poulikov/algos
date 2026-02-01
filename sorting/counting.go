package sorting

// CountingSort performs counting sort on a slice of integers
// Time complexity: O(n + k) where k is the range of input
// Space complexity: O(n + k)
func CountingSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	min := slice[0]
	max := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	span := max - min + 1
	count := make([]int, span)
	output := make([]int, len(slice))

	for _, v := range slice {
		count[v-min]++
	}

	for i := 1; i < span; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		output[count[slice[i]-min]-1] = slice[i]
		count[slice[i]-min]--
	}

	copy(slice, output)
}

// CountingSortDescending performs counting sort on a slice of integers in descending order
// Time complexity: O(n + k) where k is the range of input
// Space complexity: O(n + k)
func CountingSortDescending(slice []int) {
	if len(slice) <= 1 {
		return
	}

	min := slice[0]
	max := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	span := max - min + 1
	count := make([]int, span)
	output := make([]int, len(slice))

	for _, v := range slice {
		count[v-min]++
	}

	for i := 1; i < span; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		output[count[slice[i]-min]-1] = slice[i]
		count[slice[i]-min]--
	}

	for i := 0; i < len(output)/2; i++ {
		output[i], output[len(output)-1-i] = output[len(output)-1-i], output[i]
	}

	copy(slice, output)
}

// CountingSortCopy returns a new sorted slice without modifying original
// Time complexity: O(n + k)
// Space complexity: O(n + k)
func CountingSortCopy(slice []int) []int {
	if len(slice) <= 1 {
		result := make([]int, len(slice))
		copy(result, slice)
		return result
	}

	min := slice[0]
	max := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	span := max - min + 1
	count := make([]int, span)
	output := make([]int, len(slice))

	for _, v := range slice {
		count[v-min]++
	}

	for i := 1; i < span; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		output[count[slice[i]-min]-1] = slice[i]
		count[slice[i]-min]--
	}

	return output
}

// CountingSortDescendingCopy returns a new sorted slice in descending order
// Time complexity: O(n + k)
// Space complexity: O(n + k)
func CountingSortDescendingCopy(slice []int) []int {
	result := CountingSortCopy(slice)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

// CountingSortUint8 performs counting sort on a slice of uint8 values
// Optimized for the specific range of uint8 (0-255)
// Time complexity: O(n)
// Space complexity: O(n)
func CountingSortUint8(slice []uint8) {
	if len(slice) <= 1 {
		return
	}

	const rangeSize = 256
	count := [rangeSize]int{}
	output := make([]uint8, len(slice))

	for _, v := range slice {
		count[v]++
	}

	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		output[count[slice[i]]-1] = slice[i]
		count[slice[i]]--
	}

	copy(slice, output)
}
