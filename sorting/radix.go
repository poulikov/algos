package sorting

// RadixSort performs radix sort on a slice of integers (non-negative only)
// Time complexity: O(d * (n + k)) where d is the number of digits and k=10 for base-10
// Space complexity: O(n + k)
func RadixSort(slice []int) {
	if len(slice) <= 1 {
		return
	}

	max := slice[0]
	for _, v := range slice {
		if v < 0 {
			panic("RadixSort does not support negative numbers")
		}
		if v > max {
			max = v
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(slice, exp)
	}
}

// countingSortByDigit sorts slice by a specific digit using counting sort
func countingSortByDigit(slice []int, exp int) {
	const base = 10
	output := make([]int, len(slice))
	count := make([]int, base)

	for i := 0; i < len(slice); i++ {
		digit := (slice[i] / exp) % base
		count[digit]++
	}

	for i := 1; i < base; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		digit := (slice[i] / exp) % base
		output[count[digit]-1] = slice[i]
		count[digit]--
	}

	copy(slice, output)
}

// RadixSortDescending performs radix sort on a slice of integers in descending order (non-negative only)
// Time complexity: O(d * (n + k)) where d is the number of digits and k=10 for base-10
// Space complexity: O(n + k)
func RadixSortDescending(slice []int) {
	RadixSort(slice)
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}

// RadixSortCopy returns a new sorted slice without modifying original (non-negative only)
// Time complexity: O(d * (n + k))
// Space complexity: O(n + k)
func RadixSortCopy(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	RadixSort(result)
	return result
}

// RadixSortDescendingCopy returns a new sorted slice in descending order (non-negative only)
// Time complexity: O(d * (n + k))
// Space complexity: O(n + k)
func RadixSortDescendingCopy(slice []int) []int {
	result := RadixSortCopy(slice)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

// RadixSortWithBase performs radix sort with a custom base (2-36) on non-negative integers
// Higher base means fewer passes but more memory per pass
// Time complexity: O(d * (n + k)) where k is the base
// Space complexity: O(n + k)
func RadixSortWithBase(slice []int, base int) {
	if len(slice) <= 1 {
		return
	}

	if base < 2 || base > 36 {
		panic("base must be between 2 and 36")
	}

	max := slice[0]
	for _, v := range slice {
		if v < 0 {
			panic("RadixSortWithBase does not support negative numbers")
		}
		if v > max {
			max = v
		}
	}

	for exp := 1; max/exp > 0; exp *= base {
		countingSortByDigitWithBase(slice, exp, base)
	}
}

// countingSortByDigitWithBase sorts slice by a specific digit using counting sort with custom base
func countingSortByDigitWithBase(slice []int, exp int, base int) {
	output := make([]int, len(slice))
	count := make([]int, base)

	for i := 0; i < len(slice); i++ {
		digit := (slice[i] / exp) % base
		count[digit]++
	}

	for i := 1; i < base; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		digit := (slice[i] / exp) % base
		output[count[digit]-1] = slice[i]
		count[digit]--
	}

	copy(slice, output)
}

// RadixSortUint32 performs radix sort on a slice of uint32 values
// Optimized for 32-bit unsigned integers
// Time complexity: O(d * (n + k)) with base=256
// Space complexity: O(n + k)
func RadixSortUint32(slice []uint32) {
	if len(slice) <= 1 {
		return
	}

	const base = 256
	const maxBytes = 4

	for byteIndex := 0; byteIndex < maxBytes; byteIndex++ {
		countingSortByByte(slice, byteIndex, base)
	}
}

// countingSortByByte sorts slice by a specific byte using counting sort
func countingSortByByte(slice []uint32, byteIndex int, base int) {
	output := make([]uint32, len(slice))
	count := make([]int, base)

	shift := byteIndex * 8
	for i := 0; i < len(slice); i++ {
		digit := byte((slice[i] >> shift) & 0xFF)
		count[digit]++
	}

	for i := 1; i < base; i++ {
		count[i] += count[i-1]
	}

	for i := len(slice) - 1; i >= 0; i-- {
		digit := byte((slice[i] >> shift) & 0xFF)
		output[count[digit]-1] = slice[i]
		count[digit]--
	}

	copy(slice, output)
}
