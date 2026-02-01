package set

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

// HashMap represents a hash map implementation with chaining collision resolution
// It supports resizing for better performance
type HashMap[K comparable, V any] struct {
	buckets    []*entry[K, V]
	capacity   int
	size       int
	loadFactor float64
	threshold  int
}

type entry[K comparable, V any] struct {
	key   K
	value V
	next  *entry[K, V]
}

const (
	defaultCapacity   = 16
	defaultLoadFactor = 0.75
	minimumCapacity   = 8
	resizeMultiplier  = 2
)

// NewHashMap creates a new empty HashMap with default capacity
// Time complexity: O(1)
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return NewHashMapWithCapacity[K, V](defaultCapacity)
}

// NewHashMapWithCapacity creates a new empty HashMap with specified capacity
// Time complexity: O(1)
func NewHashMapWithCapacity[K comparable, V any](capacity int) *HashMap[K, V] {
	if capacity < minimumCapacity {
		capacity = minimumCapacity
	}

	return &HashMap[K, V]{
		buckets:    make([]*entry[K, V], capacity),
		capacity:   capacity,
		loadFactor: defaultLoadFactor,
		threshold:  int(float64(capacity) * defaultLoadFactor),
	}
}

// NewHashMapWithLoadFactor creates a new empty HashMap with specified capacity and load factor
// Time complexity: O(1)
func NewHashMapWithLoadFactor[K comparable, V any](capacity int, loadFactor float64) *HashMap[K, V] {
	if capacity < minimumCapacity {
		capacity = minimumCapacity
	}

	if loadFactor <= 0 {
		loadFactor = defaultLoadFactor
	}

	return &HashMap[K, V]{
		buckets:    make([]*entry[K, V], capacity),
		capacity:   capacity,
		loadFactor: loadFactor,
		threshold:  int(float64(capacity) * loadFactor),
	}
}

// Put stores a key-value pair in the map
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) Put(key K, value V) {
	if hm.size >= hm.threshold {
		hm.resize()
	}

	index := hm.hash(key)
	bucket := hm.buckets[index]

	if bucket == nil {
		hm.buckets[index] = &entry[K, V]{key: key, value: value}
		hm.size++
		return
	}

	current := bucket
	for current != nil {
		if current.key == key {
			current.value = value
			return
		}
		if current.next == nil {
			break
		}
		current = current.next
	}

	current.next = &entry[K, V]{key: key, value: value}
	hm.size++
}

// Get retrieves a value by key
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) Get(key K) (V, bool) {
	index := hm.hash(key)
	current := hm.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}

	var zero V
	return zero, false
}

// ContainsKey checks if a key exists in the map
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) ContainsKey(key K) bool {
	_, exists := hm.Get(key)
	return exists
}

// Remove deletes a key-value pair by key
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) Remove(key K) bool {
	index := hm.hash(key)
	current := hm.buckets[index]

	if current == nil {
		return false
	}

	if current.key == key {
		hm.buckets[index] = current.next
		hm.size--
		return true
	}

	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			hm.size--
			return true
		}
		current = current.next
	}

	return false
}

// Size returns the number of key-value pairs in the map
// Time complexity: O(1)
func (hm *HashMap[K, V]) Size() int {
	return hm.size
}

// IsEmpty returns true if the map is empty
// Time complexity: O(1)
func (hm *HashMap[K, V]) IsEmpty() bool {
	return hm.size == 0
}

// Capacity returns the current capacity of the map
// Time complexity: O(1)
func (hm *HashMap[K, V]) Capacity() int {
	return hm.capacity
}

// Clear removes all key-value pairs from the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Clear() {
	hm.buckets = make([]*entry[K, V], hm.capacity)
	hm.size = 0
}

// Keys returns a slice of all keys in the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Keys() []K {
	keys := make([]K, 0, hm.size)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			keys = append(keys, current.key)
			current = current.next
		}
	}

	return keys
}

// Values returns a slice of all values in the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Values() []V {
	values := make([]V, 0, hm.size)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			values = append(values, current.value)
			current = current.next
		}
	}

	return values
}

// Entries returns a slice of all key-value pairs in the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Entries() []mapEntry[K, V] {
	entries := make([]mapEntry[K, V], 0, hm.size)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			entries = append(entries, mapEntry[K, V]{Key: current.key, Value: current.value})
			current = current.next
		}
	}

	return entries
}

type mapEntry[K comparable, V any] struct {
	Key   K
	Value V
}

// PutAll puts all key-value pairs from another map into this map
// Time complexity: O(n)
func (hm *HashMap[K, V]) PutAll(other *HashMap[K, V]) {
	for key, value := range other.toMap() {
		hm.Put(key, value)
	}
}

// PutIfAbsent puts a key-value pair only if the key is not already present
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) PutIfAbsent(key K, value V) bool {
	if hm.ContainsKey(key) {
		return false
	}
	hm.Put(key, value)
	return true
}

// PutIfPresent updates a key-value pair only if the key is already present
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) PutIfPresent(key K, value V) bool {
	index := hm.hash(key)
	current := hm.buckets[index]

	for current != nil {
		if current.key == key {
			current.value = value
			return true
		}
		current = current.next
	}

	return false
}

// Replace updates the value for an existing key
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) Replace(key K, value V) bool {
	index := hm.hash(key)
	current := hm.buckets[index]

	for current != nil {
		if current.key == key {
			current.value = value
			return true
		}
		current = current.next
	}

	return false
}

// GetOrDefault retrieves a value by key or returns a default value if the key doesn't exist
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	if value, exists := hm.Get(key); exists {
		return value
	}
	return defaultValue
}

// ComputeIfAbsent computes a value for a key only if the key is not already present
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) ComputeIfAbsent(key K, computeFunc func(K) V) V {
	if value, exists := hm.Get(key); exists {
		return value
	}

	value := computeFunc(key)
	hm.Put(key, value)
	return value
}

// ComputeIfPresent computes a new value for a key only if the key is already present
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) ComputeIfPresent(key K, computeFunc func(K, V) V) {
	index := hm.hash(key)
	current := hm.buckets[index]

	for current != nil {
		if current.key == key {
			current.value = computeFunc(key, current.value)
			return
		}
		current = current.next
	}
}

// Merge merges the value with an existing value for a key
// Time complexity: O(1) average, O(n) worst case
func (hm *HashMap[K, V]) Merge(key K, value V, mergeFunc func(oldValue, newValue V) V) {
	index := hm.hash(key)
	current := hm.buckets[index]

	if current == nil {
		hm.buckets[index] = &entry[K, V]{key: key, value: value}
		hm.size++
		return
	}

	for current != nil {
		if current.key == key {
			current.value = mergeFunc(current.value, value)
			return
		}
		if current.next == nil {
			break
		}
		current = current.next
	}

	current.next = &entry[K, V]{key: key, value: value}
	hm.size++
}

// Copy creates a shallow copy of the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Copy() *HashMap[K, V] {
	newHM := NewHashMapWithCapacity[K, V](hm.capacity)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			newHM.Put(current.key, current.value)
			current = current.next
		}
	}

	return newHM
}

// ForEach applies a function to each key-value pair
// Time complexity: O(n)
func (hm *HashMap[K, V]) ForEach(action func(K, V)) {
	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			action(current.key, current.value)
			current = current.next
		}
	}
}

// Filter creates a new map with key-value pairs that satisfy the predicate
// Time complexity: O(n)
func (hm *HashMap[K, V]) Filter(predicate func(K, V) bool) *HashMap[K, V] {
	newHM := NewHashMapWithCapacity[K, V](hm.capacity)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			if predicate(current.key, current.value) {
				newHM.Put(current.key, current.value)
			}
			current = current.next
		}
	}

	return newHM
}

// Map transforms the values of the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) Map(mapper func(K, V) V) *HashMap[K, V] {
	newHM := NewHashMapWithCapacity[K, V](hm.capacity)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			newHM.Put(current.key, mapper(current.key, current.value))
			current = current.next
		}
	}

	return newHM
}

// toMap converts the HashMap to a Go map
// Time complexity: O(n)
func (hm *HashMap[K, V]) toMap() map[K]V {
	result := make(map[K]V)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			result[current.key] = current.value
			current = current.next
		}
	}

	return result
}

// hash computes the hash for a key
func (hm *HashMap[K, V]) hash(key K) int {
	h := fnv.New32a()
	h.Write([]byte(typeToString(key)))
	return int(h.Sum32()) % hm.capacity
}

// resize increases the capacity of the map
// Time complexity: O(n)
func (hm *HashMap[K, V]) resize() {
	newCapacity := hm.capacity * resizeMultiplier
	newHM := NewHashMapWithCapacity[K, V](newCapacity)

	for _, bucket := range hm.buckets {
		current := bucket
		for current != nil {
			newHM.Put(current.key, current.value)
			current = current.next
		}
	}

	hm.buckets = newHM.buckets
	hm.capacity = newHM.capacity
	hm.threshold = newHM.threshold
}

// typeToString is used for type conversion to string for hashing
func typeToString(a interface{}) string {
	switch v := a.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
