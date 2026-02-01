package set

import (
	"testing"
)

func TestNewHashMap(t *testing.T) {
	hm := NewHashMap[int, string]()
	if hm == nil {
		t.Fatal("NewHashMap() returned nil")
	}
	if !hm.IsEmpty() {
		t.Fatal("NewHashMap should be empty")
	}
	if hm.Size() != 0 {
		t.Fatalf("Expected size 0, got %d", hm.Size())
	}
}

func TestNewHashMapWithCapacity(t *testing.T) {
	hm := NewHashMapWithCapacity[int, string](32)
	if hm.Capacity() != 32 {
		t.Fatalf("Expected capacity 32, got %d", hm.Capacity())
	}
}

func TestNewHashMapWithLoadFactor(t *testing.T) {
	hm := NewHashMapWithLoadFactor[int, string](16, 0.5)
	if hm.Capacity() != 16 {
		t.Fatalf("Expected capacity 16, got %d", hm.Capacity())
	}
}

func TestPutAndGet(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	value, exists := hm.Get(1)
	if !exists {
		t.Fatal("Get(1) should return true")
	}
	if value != "one" {
		t.Fatalf("Expected 'one', got '%s'", value)
	}

	value, exists = hm.Get(2)
	if !exists {
		t.Fatal("Get(2) should return true")
	}
	if value != "two" {
		t.Fatalf("Expected 'two', got '%s'", value)
	}
}

func TestPutUpdate(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(1, "ONE")

	value, _ := hm.Get(1)
	if value != "ONE" {
		t.Fatalf("Expected 'ONE', got '%s'", value)
	}

	if hm.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", hm.Size())
	}
}

func TestContainsKey(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	if !hm.ContainsKey(1) {
		t.Fatal("ContainsKey(1) should return true")
	}

	if hm.ContainsKey(2) {
		t.Fatal("ContainsKey(2) should return false")
	}
}

func TestHMRemove(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	if !hm.Remove(1) {
		t.Fatal("Remove(1) should return true")
	}

	if hm.ContainsKey(1) {
		t.Fatal("Key 1 should not exist after removal")
	}

	if hm.Size() != 1 {
		t.Fatalf("Expected size 1, got %d", hm.Size())
	}
}

func TestHMRemoveNonExistent(t *testing.T) {
	hm := NewHashMap[int, string]()
	if hm.Remove(1) {
		t.Fatal("Remove(1) should return false for non-existent key")
	}
}

func TestSize(t *testing.T) {
	hm := NewHashMap[int, string]()

	for i := 0; i < 100; i++ {
		hm.Put(i, "value")
		if hm.Size() != i+1 {
			t.Fatalf("Expected size %d, got %d", i+1, hm.Size())
		}
	}
}

func TestHMClear(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	hm.Clear()

	if !hm.IsEmpty() {
		t.Fatal("Clear() should make map empty")
	}
	if hm.Size() != 0 {
		t.Fatalf("Clear() should result in size 0, got %d", hm.Size())
	}
}

func TestKeys(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")
	hm.Put(3, "three")

	keys := hm.Keys()
	if len(keys) != 3 {
		t.Fatalf("Expected 3 keys, got %d", len(keys))
	}

	keySet := map[int]bool{1: true, 2: true, 3: true}
	for _, key := range keys {
		if !keySet[key] {
			t.Fatalf("Unexpected key: %d", key)
		}
	}
}

func TestValues(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")
	hm.Put(3, "three")

	values := hm.Values()
	if len(values) != 3 {
		t.Fatalf("Expected 3 values, got %d", len(values))
	}

	valueSet := map[string]bool{"one": true, "two": true, "three": true}
	for _, value := range values {
		if !valueSet[value] {
			t.Fatalf("Unexpected value: %s", value)
		}
	}
}

func TestEntries(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	entries := hm.Entries()
	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}
}

func TestPutAll(t *testing.T) {
	hm1 := NewHashMap[int, string]()
	hm1.Put(1, "one")
	hm1.Put(2, "two")

	hm2 := NewHashMap[int, string]()
	hm2.Put(3, "three")
	hm2.Put(4, "four")

	hm1.PutAll(hm2)

	if hm1.Size() != 4 {
		t.Fatalf("Expected size 4, got %d", hm1.Size())
	}

	if !hm1.ContainsKey(3) || !hm1.ContainsKey(4) {
		t.Fatal("PutAll should add all entries from other map")
	}
}

func TestPutIfAbsent(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	if hm.PutIfAbsent(1, "ONE") {
		t.Fatal("PutIfAbsent should return false for existing key")
	}

	if !hm.PutIfAbsent(2, "two") {
		t.Fatal("PutIfAbsent should return true for new key")
	}

	if !hm.ContainsKey(2) {
		t.Fatal("PutIfAbsent should add key if absent")
	}
}

func TestPutIfPresent(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	if !hm.PutIfPresent(1, "ONE") {
		t.Fatal("PutIfPresent should return true for existing key")
	}

	value, _ := hm.Get(1)
	if value != "ONE" {
		t.Fatalf("Expected 'ONE', got '%s'", value)
	}

	if hm.PutIfPresent(2, "two") {
		t.Fatal("PutIfPresent should return false for non-existent key")
	}
}

func TestReplace(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	if !hm.Replace(1, "ONE") {
		t.Fatal("Replace should return true for existing key")
	}

	value, _ := hm.Get(1)
	if value != "ONE" {
		t.Fatalf("Expected 'ONE', got '%s'", value)
	}

	if hm.Replace(2, "two") {
		t.Fatal("Replace should return false for non-existent key")
	}
}

func TestGetOrDefault(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	value := hm.GetOrDefault(1, "default")
	if value != "one" {
		t.Fatalf("Expected 'one', got '%s'", value)
	}

	value = hm.GetOrDefault(2, "default")
	if value != "default" {
		t.Fatalf("Expected 'default', got '%s'", value)
	}
}

func TestHMComputeIfAbsent(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	value := hm.ComputeIfAbsent(1, func(k int) string {
		return "ONE"
	})

	if value != "one" {
		t.Fatalf("Expected 'one', got '%s'", value)
	}

	value = hm.ComputeIfAbsent(2, func(k int) string {
		return "two"
	})

	if value != "two" {
		t.Fatalf("Expected 'two', got '%s'", value)
	}

	if !hm.ContainsKey(2) {
		t.Fatal("ComputeIfAbsent should add key if absent")
	}
}

func TestComputeIfPresent(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	hm.ComputeIfPresent(1, func(k int, v string) string {
		return v + "!"
	})

	value, _ := hm.Get(1)
	if value != "one!" {
		t.Fatalf("Expected 'one!', got '%s'", value)
	}

	hm.ComputeIfPresent(2, func(k int, v string) string {
		return "two"
	})

	if hm.ContainsKey(2) {
		t.Fatal("ComputeIfPresent should not add key if absent")
	}
}

func TestMerge(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")

	hm.Merge(1, "ONE", func(old, new string) string {
		return old + new
	})

	value, _ := hm.Get(1)
	if value != "oneONE" {
		t.Fatalf("Expected 'oneONE', got '%s'", value)
	}

	hm.Merge(2, "two", func(old, new string) string {
		return old + new
	})

	value, _ = hm.Get(2)
	if value != "two" {
		t.Fatalf("Expected 'two', got '%s'", value)
	}
}

func TestHMCopy(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	hmCopy := hm.Copy()

	if hmCopy.Size() != hm.Size() {
		t.Fatalf("Copy should have same size, got %d vs %d", hmCopy.Size(), hm.Size())
	}

	hm.Put(3, "three")
	hm.Remove(1)

	if hmCopy.ContainsKey(3) {
		t.Fatal("Copy should be independent")
	}

	if !hmCopy.ContainsKey(1) {
		t.Fatal("Copy should have original keys")
	}
}

func TestHMForEach(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	count := 0
	hm.ForEach(func(k int, v string) {
		count++
	})

	if count != 2 {
		t.Fatalf("ForEach should visit all entries, got %d", count)
	}
}

func TestHMFilter(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")
	hm.Put(3, "three")

	filtered := hm.Filter(func(k int, v string) bool {
		return k%2 == 1
	})

	if filtered.Size() != 2 {
		t.Fatalf("Expected size 2, got %d", filtered.Size())
	}

	if !filtered.ContainsKey(1) || !filtered.ContainsKey(3) {
		t.Fatal("Filter should keep only odd keys")
	}
}

func TestHMMap(t *testing.T) {
	hm := NewHashMap[int, string]()
	hm.Put(1, "one")
	hm.Put(2, "two")

	mapped := hm.Map(func(k int, v string) string {
		return v + "!"
	})

	value, _ := mapped.Get(1)
	if value != "one!" {
		t.Fatalf("Expected 'one!', got '%s'", value)
	}
}

func TestResize(t *testing.T) {
	hm := NewHashMapWithCapacity[int, string](4)
	initialCapacity := hm.Capacity()

	for i := 0; i < 100; i++ {
		hm.Put(i, "value")
	}

	if hm.Capacity() <= initialCapacity {
		t.Fatal("Capacity should increase after adding many elements")
	}

	for i := 0; i < 100; i++ {
		if !hm.ContainsKey(i) {
			t.Fatalf("Key %d not found after resize", i)
		}
	}
}

func TestCollisionHandling(t *testing.T) {
	hm := NewHashMap[int, string]()

	for i := 0; i < 1000; i++ {
		hm.Put(i, "value")
	}

	for i := 0; i < 1000; i++ {
		value, exists := hm.Get(i)
		if !exists || value != "value" {
			t.Fatalf("Failed to get key %d", i)
		}
	}
}

func TestStringTypeKey(t *testing.T) {
	hm := NewHashMap[string, int]()
	hm.Put("one", 1)
	hm.Put("two", 2)

	value, _ := hm.Get("one")
	if value != 1 {
		t.Fatalf("Expected 1, got %d", value)
	}
}

func BenchmarkPut(b *testing.B) {
	hm := NewHashMap[int, string]()
	for i := 0; i < b.N; i++ {
		hm.Put(i, "value")
	}
}

func BenchmarkGet(b *testing.B) {
	hm := NewHashMap[int, string]()
	for i := 0; i < 1000; i++ {
		hm.Put(i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Get(i % 1000)
	}
}

func BenchmarkHMRemove(b *testing.B) {
	hm := NewHashMap[int, string]()
	for i := 0; i < 1000; i++ {
		hm.Put(i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.Remove(i % 1000)
		hm.Put(i%1000, "value")
	}
}

func BenchmarkHMForEach(b *testing.B) {
	hm := NewHashMap[int, string]()
	for i := 0; i < 1000; i++ {
		hm.Put(i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hm.ForEach(func(k int, v string) {})
	}
}

func TestTypeToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{int(65), "65"},
		{int(0), "0"},
		{int(-1), "-1"},
		{int(65535), "65535"},
		{int8(65), "65"},
		{int8(127), "127"},
		{int8(-128), "-128"},
		{int16(65), "65"},
		{int16(32767), "32767"},
		{int16(-32768), "-32768"},
		{int32(65), "65"},
		{int32(2147483647), "2147483647"},
		{int32(-2147483648), "-2147483648"},
		{int64(65), "65"},
		{int64(9223372036854775807), "9223372036854775807"},
		{int64(-9223372036854775808), "-9223372036854775808"},
		{uint(65), "65"},
		{uint(0), "0"},
		{uint(65535), "65535"},
		{uint8(65), "65"},
		{uint8(255), "255"},
		{uint8(0), "0"},
		{uint16(65), "65"},
		{uint16(65535), "65535"},
		{uint16(0), "0"},
		{uint32(65), "65"},
		{uint32(4294967295), "4294967295"},
		{uint32(0), "0"},
		{uint64(65), "65"},
		{uint64(18446744073709551615), "18446744073709551615"},
		{uint64(0), "0"},
		{float32(3.14), "3.14"},
		{float32(0.0), "0"},
		{float64(3.14), "3.14"},
		{float64(0.0), "0"},
		{"hello", "hello"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := typeToString(test.input)
			if result != test.expected {
				t.Errorf("typeToString(%v) = %q, want %q", test.input, result, test.expected)
			}
		})
	}
}

func TestHashMapNumericKeys(t *testing.T) {
	hm := NewHashMap[int, string]()

	testCases := []struct {
		key      int
		value    string
		expected string
	}{
		{65, "sixty-five", "sixty-five"},
		{0, "zero", "zero"},
		{1, "one", "one"},
		{65535, "large", "large"},
	}

	for _, tc := range testCases {
		hm.Put(tc.key, tc.value)
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if value, exists := hm.Get(tc.key); !exists {
				t.Errorf("key %d not found", tc.key)
			} else if value != tc.expected {
				t.Errorf("key %d: got %q, want %q", tc.key, value, tc.expected)
			}
		})
	}

	if hm.Size() != len(testCases) {
		t.Errorf("Size() = %d, want %d", hm.Size(), len(testCases))
	}
}

func TestHashMapFloatKeys(t *testing.T) {
	hm := NewHashMap[float64, string]()

	testCases := []struct {
		key      float64
		value    string
		expected string
	}{
		{3.14, "pi", "pi"},
		{0.0, "zero", "zero"},
		{-1.5, "negative", "negative"},
	}

	for _, tc := range testCases {
		hm.Put(tc.key, tc.value)
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if value, exists := hm.Get(tc.key); !exists {
				t.Errorf("key %f not found", tc.key)
			} else if value != tc.expected {
				t.Errorf("key %f: got %q, want %q", tc.key, value, tc.expected)
			}
		})
	}

	if hm.Size() != len(testCases) {
		t.Errorf("Size() = %d, want %d", hm.Size(), len(testCases))
	}
}
