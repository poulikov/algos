package structures

import (
	"testing"
)

func TestBloomFilterNew(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	if bf.Size() == 0 {
		t.Errorf("Expected non-zero size")
	}
	if bf.HashFunctions() == 0 {
		t.Errorf("Expected non-zero hash functions")
	}
}

func TestBloomFilterAddContains(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	data := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("test"),
	}

	for _, d := range data {
		bf.Add(d)
	}

	for _, d := range data {
		if !bf.Contains(d) {
			t.Errorf("Contains(%s) should return true", d)
		}
	}

	if bf.Contains([]byte("notpresent")) {
		t.Errorf("Contains(\"notpresent\") should return false")
	}
}

func TestBloomFilterAddString(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	strings := []string{"hello", "world", "test"}

	for _, s := range strings {
		bf.AddString(s)
	}

	for _, s := range strings {
		if !bf.ContainsString(s) {
			t.Errorf("ContainsString(%s) should return true", s)
		}
	}

	if bf.ContainsString("notpresent") {
		t.Errorf("ContainsString(\"notpresent\") should return false")
	}
}

func TestBloomFilterFalsePositive(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	n := 500
	for i := 0; i < n; i++ {
		bf.AddString(string(rune('a' + i)))
	}

	falsePositives := 0
	tests := 1000
	for i := n; i < n+tests; i++ {
		if bf.ContainsString(string(rune('a' + i))) {
			falsePositives++
		}
	}

	rate := float64(falsePositives) / float64(tests)
	expectedRate := bf.FalsePositiveRate(n)

	if rate > 0.1 {
		t.Errorf("False positive rate %f is too high (expected ~%f)", rate, expectedRate)
	}
}

func TestBloomFilterBitsSet(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	initialBits := bf.BitsSet()
	if initialBits != 0 {
		t.Errorf("Initial bits set should be 0, got %d", initialBits)
	}

	bf.AddString("test")

	afterBits := bf.BitsSet()
	if afterBits == 0 {
		t.Errorf("After adding, bits set should be > 0")
	}
	if afterBits != bf.HashFunctions() {
		t.Errorf("After one add, bits set should equal hash functions: %d != %d", afterBits, bf.HashFunctions())
	}
}

func TestBloomFilterReset(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	bf.AddString("test")

	bf.Reset()

	if bf.BitsSet() != 0 {
		t.Errorf("After reset, bits set should be 0, got %d", bf.BitsSet())
	}
	if bf.ContainsString("test") {
		t.Errorf("After reset, filter should not contain \"test\"")
	}
}

func TestBloomFilterUnion(t *testing.T) {
	bf1 := NewBloomFilterWithParams(100, 5)
	bf2 := NewBloomFilterWithParams(100, 5)

	bf1.AddString("hello")
	bf1.AddString("world")

	bf2.AddString("world")
	bf2.AddString("test")

	bf1.Union(bf2)

	if !bf1.ContainsString("hello") {
		t.Errorf("Union should contain \"hello\"")
	}
	if !bf1.ContainsString("world") {
		t.Errorf("Union should contain \"world\"")
	}
	if !bf1.ContainsString("test") {
		t.Errorf("Union should contain \"test\"")
	}
}

func TestBloomFilterUnionInvalid(t *testing.T) {
	bf1 := NewBloomFilterWithParams(100, 5)
	bf2 := NewBloomFilterWithParams(200, 5)

	if bf1.Union(bf2) {
		t.Errorf("Union with different sizes should return false")
	}
}

func TestBloomFilterIntersection(t *testing.T) {
	bf1 := NewBloomFilterWithParams(100, 5)
	bf2 := NewBloomFilterWithParams(100, 5)

	bf1.AddString("hello")
	bf1.AddString("world")
	bf1.AddString("test")

	bf2.AddString("world")
	bf2.AddString("test")
	bf2.AddString("data")

	bf1.Intersection(bf2)

	if bf1.ContainsString("hello") {
		t.Errorf("Intersection should not contain \"hello\"")
	}
	if !bf1.ContainsString("world") {
		t.Errorf("Intersection should contain \"world\"")
	}
	if !bf1.ContainsString("test") {
		t.Errorf("Intersection should contain \"test\"")
	}
	if bf1.ContainsString("data") {
		t.Errorf("Intersection should not contain \"data\"")
	}
}

func TestBloomFilterIntersectionInvalid(t *testing.T) {
	bf1 := NewBloomFilterWithParams(100, 5)
	bf2 := NewBloomFilterWithParams(200, 5)

	if bf1.Intersection(bf2) {
		t.Errorf("Intersection with different sizes should return false")
	}
}

func TestBloomFilterLarge(t *testing.T) {
	bf := NewBloomFilter(10000, 0.001)

	n := 5000
	for i := 0; i < n; i++ {
		bf.AddString(string(rune('a' + i%26)))
	}

	falsePositives := 0
	tests := 1000
	for i := n; i < n+tests; i++ {
		if bf.ContainsString(string(rune('z' + i%26))) {
			falsePositives++
		}
	}

	rate := float64(falsePositives) / float64(tests)
	expectedRate := bf.FalsePositiveRate(n)

	if rate > 0.05 {
		t.Errorf("False positive rate %f is too high (expected ~%f)", rate, expectedRate)
	}
}

func TestBloomFilterConsistency(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	strings := []string{"hello", "world", "test", "data", "code"}

	for _, s := range strings {
		bf.AddString(s)
		if !bf.ContainsString(s) {
			t.Errorf("After adding, should contain %s", s)
		}
	}

	for _, s := range strings {
		if !bf.ContainsString(s) {
			t.Errorf("Should still contain %s", s)
		}
	}
}

func TestBloomFilterEstimatedBitsSet(t *testing.T) {
	bf := NewBloomFilter(1000, 0.01)

	n := 100
	for i := 0; i < n; i++ {
		bf.AddString(string(rune(i)))
	}

	actual := bf.BitsSet()
	estimated := bf.EstimatedBitsSet(n)

	ratio := float64(actual) / estimated
	if ratio < 0.5 || ratio > 1.5 {
		t.Errorf("Estimated bits %f is not close to actual %d", estimated, actual)
	}
}
