package slidingwindow

import (
	"math"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	data := []int{1, 3, -1, -3, 5, 3, 6, 7}
	windowSize := 3

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != 7 {
		t.Errorf("Expected max 7, got %d", result.Max)
	}
}

func TestMaxSlidingWindowEmpty(t *testing.T) {
	data := []int{}
	windowSize := 3

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != 0 || result.MaxIdx != 0 {
		t.Error("Empty data should return zero result")
	}
}

func TestMaxSlidingWindowInvalidSize(t *testing.T) {
	data := []int{1, 2, 3}

	result1 := MaxSlidingWindow(data, 0)
	if result1.Max != 0 {
		t.Error("Window size 0 should return zero result")
	}

	result2 := MaxSlidingWindow(data, 5)
	if result2.Max != 0 {
		t.Error("Window size > data length should return zero result")
	}

	result3 := MaxSlidingWindow(data, -1)
	if result3.Max != 0 {
		t.Error("Negative window size should return zero result")
	}
}

func TestMinSlidingWindow(t *testing.T) {
	data := []int{1, 3, -1, -3, 5, 3, 6, 7}
	windowSize := 3

	result := MinSlidingWindow(data, windowSize)

	if result.Min != -3 {
		t.Errorf("Expected min -3, got %d", result.Min)
	}
}

func TestMinSlidingWindowEmpty(t *testing.T) {
	data := []int{}
	windowSize := 3

	result := MinSlidingWindow(data, windowSize)

	if result.Min != 0 || result.MinIdx != 0 {
		t.Error("Empty data should return zero result")
	}
}

func TestSumSlidingWindow(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	start := 1
	end := 3

	sum := SumSlidingWindow(data, start, end)

	if sum != 9 {
		t.Errorf("Expected sum 9 (2+3+4), got %d", sum)
	}
}

func TestSumSlidingWindowInvalidBounds(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	sum1 := SumSlidingWindow(data, -1, 3)
	if sum1 != 0 {
		t.Error("Negative start should return 0")
	}

	sum2 := SumSlidingWindow(data, 0, 10)
	if sum2 != 0 {
		t.Error("End >= length should return 0")
	}

	sum3 := SumSlidingWindow(data, 3, 1)
	if sum3 != 0 {
		t.Error("Start > end should return 0")
	}
}

func TestAverageSlidingWindow(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	start := 0
	end := 4

	avg := AverageSlidingWindow(data, start, end)

	expected := 3.0
	if avg != expected {
		t.Errorf("Expected average %f, got %f", expected, avg)
	}
}

func TestAverageSlidingWindowEmpty(t *testing.T) {
	data := []int{}

	avg := AverageSlidingWindow(data, 0, 2)

	if avg != 0 {
		t.Error("Empty data should return 0")
	}
}

func TestContainsInWindow(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	value := 3
	start := 1
	end := 3

	contains := ContainsInWindow(data, value, start, end)

	if !contains {
		t.Error("Window should contain value 3")
	}
}

func TestContainsInWindowNotFound(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	value := 6
	start := 0
	end := 4

	contains := ContainsInWindow(data, value, start, end)

	if contains {
		t.Error("Window should not contain value 6")
	}
}

func TestContainsInWindowInvalidBounds(t *testing.T) {
	data := []int{1, 2, 3}

	contains1 := ContainsInWindow(data, 2, -1, 2)
	if contains1 {
		t.Error("Negative start should return false")
	}

	contains2 := ContainsInWindow(data, 2, 0, 5)
	if contains2 {
		t.Error("End >= length should return false")
	}
}

func TestCountInWindow(t *testing.T) {
	data := []int{1, 2, 2, 3, 2, 4}
	value := 2
	start := 0
	end := 4

	count := CountInWindow(data, value, start, end)

	if count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}
}

func TestCountInWindowZero(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	value := 6
	start := 0
	end := 4

	count := CountInWindow(data, value, start, end)

	if count != 0 {
		t.Error("Should count 0 for non-existent value")
	}
}

func TestFindInWindow(t *testing.T) {
	data := []int{1, 2, 2, 3, 2, 4}
	value := 2
	start := 0
	end := 4

	indices := FindInWindow(data, value, start, end)

	if len(indices) != 3 {
		t.Errorf("Expected 3 occurrences, got %d", len(indices))
	}

	expected := []int{1, 2, 4}
	for i, idx := range expected {
		if indices[i] != idx {
			t.Errorf("Expected index %d at position %d, got %d", idx, i, indices[i])
		}
	}
}

func TestFindInWindowEmpty(t *testing.T) {
	data := []int{1, 2, 3}
	value := 4
	start := 0
	end := 2

	indices := FindInWindow(data, value, start, end)

	if len(indices) != 0 {
		t.Error("Should return empty slice for non-existent value")
	}
}

func TestFirstOccurrence(t *testing.T) {
	data := []int{1, 2, 3, 2, 4}
	value := 2
	start := 0
	end := 4

	idx, found := FirstOccurrence(data, value, start, end)

	if !found {
		t.Error("Should find value 2")
	}

	if idx != 1 {
		t.Errorf("Expected first occurrence at index 1, got %d", idx)
	}
}

func TestFirstOccurrenceNotFound(t *testing.T) {
	data := []int{1, 2, 3}
	value := 4
	start := 0
	end := 2

	_, found := FirstOccurrence(data, value, start, end)

	if found {
		t.Error("Should not find non-existent value")
	}
}

func TestAllOccurrences(t *testing.T) {
	data := []int{1, 2, 2, 3, 2}
	value := 2
	start := 0
	end := 4

	indices := AllOccurrences(data, value, start, end)

	if len(indices) != 3 {
		t.Errorf("Expected 3 occurrences, got %d", len(indices))
	}
}

func TestMaxSlidingWindowSum(t *testing.T) {
	data := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	windowSize := 4

	result := MaxSlidingWindowSum(data, windowSize)

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	if result.Sum != 39 {
		t.Errorf("Expected max sum 39 (4+2+10+23), got %d", result.Sum)
	}

	if result.StartIdx != 1 {
		t.Errorf("Expected start index 1, got %d", result.StartIdx)
	}

	if result.EndIdx != 4 {
		t.Errorf("Expected end index 4, got %d", result.EndIdx)
	}
}

func TestMaxSlidingWindowSumEmpty(t *testing.T) {
	data := []int{}
	windowSize := 3

	result := MaxSlidingWindowSum(data, windowSize)

	if result != nil {
		t.Error("Empty data should return nil")
	}
}

func TestMinSlidingWindowSum(t *testing.T) {
	data := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	windowSize := 4

	result := MinSlidingWindowSum(data, windowSize)

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	if result.Sum != 17 {
		t.Errorf("Expected min sum 17 (1+4+2+10), got %d", result.Sum)
	}
}

func TestFixedSizeSlidingWindow(t *testing.T) {
	data := []int{1, 4, 2, 10, 23, 3}
	windowSize := 2

	result := FixedSizeSlidingWindow(data, windowSize)

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	if result.Sum != 33 {
		t.Errorf("Expected sum 33 (10+23), got %d", result.Sum)
	}
}

func TestAllMaxSlidingWindow(t *testing.T) {
	data := []int{1, 2, 3, 4}
	windowSize := 2

	results := AllMaxSlidingWindow(data, windowSize)

	if len(results) != 3 {
		t.Errorf("Expected 3 windows, got %d", len(results))
	}

	if results[0].Sum != 3 {
		t.Errorf("First window sum should be 3, got %d", results[0].Sum)
	}

	if results[1].Sum != 5 {
		t.Errorf("Second window sum should be 5, got %d", results[1].Sum)
	}

	if results[2].Sum != 7 {
		t.Errorf("Third window sum should be 7, got %d", results[2].Sum)
	}
}

func TestAllMaxSlidingWindowEmpty(t *testing.T) {
	data := []int{}
	windowSize := 2

	results := AllMaxSlidingWindow(data, windowSize)

	if len(results) != 0 {
		t.Error("Empty data should return empty slice")
	}
}

func TestAllMinSlidingWindow(t *testing.T) {
	data := []int{1, 2, 3, 4}
	windowSize := 2

	results := AllMinSlidingWindow(data, windowSize)

	if len(results) != 3 {
		t.Errorf("Expected 3 windows, got %d", len(results))
	}
}

func TestSlidingWindowResultString(t *testing.T) {
	result := &SlidingWindowResult[int]{
		Max:    10,
		MaxIdx: 5,
		Min:    1,
		MinIdx: 2,
	}

	str := result.String()
	expected := "Max: 10 (index: 5), Min: 1 (index: 2)"
	if str != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str)
	}
}

func TestSlidingWindowSumString(t *testing.T) {
	result := &SlidingWindowSum[int]{
		Sum:      15,
		StartIdx: 2,
		EndIdx:   4,
	}

	str := result.String()
	expected := "Sum: 15 (idx: 2-4)"
	if str != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str)
	}
}

func TestSlidingWindowSumNilString(t *testing.T) {
	var result *SlidingWindowSum[int]

	str := result.String()
	expected := "nil"
	if str != expected {
		t.Errorf("Expected '%s', got '%s'", expected, str)
	}
}

func TestSlidingWindowFloats(t *testing.T) {
	data := []float64{1.5, 2.5, 3.5}
	windowSize := 2

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != 3.5 {
		t.Errorf("Expected max 3.5, got %f", result.Max)
	}
}

func TestSlidingWindowNegativeNumbers(t *testing.T) {
	data := []int{-1, -2, -3, -4}
	windowSize := 2

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != -1 {
		t.Errorf("Expected max -1, got %d", result.Max)
	}

	resultMin := MinSlidingWindow(data, windowSize)
	if resultMin.Min != -4 {
		t.Errorf("Expected min -4, got %d", resultMin.Min)
	}
}

func TestSlidingWindowSingleElement(t *testing.T) {
	data := []int{5}
	windowSize := 1

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != 5 {
		t.Errorf("Expected max 5, got %d", result.Max)
	}
}

func TestSlidingWindowFullArray(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	windowSize := 5

	result := MaxSlidingWindow(data, windowSize)

	if result.Max != 5 {
		t.Errorf("Expected max 5, got %d", result.Max)
	}
}

func TestSlidingWindowLargeData(t *testing.T) {
	data := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}
	windowSize := 100

	result := MaxSlidingWindowSum(data, windowSize)

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	expectedSum := int(0)
	for i := 900; i < 1000; i++ {
		expectedSum += i
	}

	if result.Sum != expectedSum {
		t.Errorf("Expected sum %d, got %d", expectedSum, result.Sum)
	}
}

func TestSlidingWindowAllSame(t *testing.T) {
	data := []int{5, 5, 5, 5, 5}
	windowSize := 2

	resultMax := MaxSlidingWindow(data, windowSize)
	if resultMax.Max != 5 {
		t.Errorf("Expected max 5, got %d", resultMax.Max)
	}

	resultMin := MinSlidingWindow(data, windowSize)
	if resultMin.Min != 5 {
		t.Errorf("Expected min 5, got %d", resultMin.Min)
	}
}

func TestAverageSlidingWindowFloats(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	start := 0
	end := 4

	avg := AverageSlidingWindow(data, start, end)

	if math.Abs(avg-3.0) > 0.0001 {
		t.Errorf("Expected average 3.0, got %f", avg)
	}
}
