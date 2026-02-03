package sharding

import (
	"fmt"
	"hash/fnv"
	"slices"
	"sort"
	"sync"
)

// ConsistentHash implements consistent hashing with virtual nodes
// for distributing keys across nodes in a distributed system.
// Time complexity for GetNode: O(log n) where n is number of virtual nodes
// Space complexity: O(v) where v is number of virtual nodes
type ConsistentHash[T comparable] struct {
	sync.RWMutex
	virtualNodes int            // Number of virtual nodes per physical node
	ring         []uint32       // Sorted ring of hash values
	nodeMap      map[uint32]T   // Hash -> Physical node mapping
	virtualMap   map[T][]uint32 // Physical node -> Virtual node hashes
}

// NewConsistentHash creates a new consistent hashing ring
// virtualNodesPerNode specifies how many virtual nodes to create per physical node
// More virtual nodes = better distribution but more memory
// Recommended values: 100-500 for production systems
// Time complexity: O(1)
func NewConsistentHash[T comparable](virtualNodesPerNode int) *ConsistentHash[T] {
	if virtualNodesPerNode <= 0 {
		virtualNodesPerNode = 150 // Default value
	}
	return &ConsistentHash[T]{
		virtualNodes: virtualNodesPerNode,
		ring:         []uint32{},
		nodeMap:      make(map[uint32]T),
		virtualMap:   make(map[T][]uint32),
	}
}

// AddNode adds a physical node to the hash ring
// Creates virtualNodes virtual nodes for better distribution
// Time complexity: O(v * log(v)) where v is number of virtual nodes
func (ch *ConsistentHash[T]) AddNode(node T) {
	ch.Lock()
	defer ch.Unlock()

	// Check if node already exists
	if _, exists := ch.virtualMap[node]; exists {
		return
	}

	// Create virtual nodes for this physical node
	virtualHashes := make([]uint32, ch.virtualNodes)
	for i := 0; i < ch.virtualNodes; i++ {
		hash := ch.hashNode(node, i)
		virtualHashes[i] = hash
		ch.nodeMap[hash] = node
		ch.ring = append(ch.ring, hash)
	}

	// Sort the ring
	slices.Sort(ch.ring)

	// Store mapping
	ch.virtualMap[node] = virtualHashes
}

// RemoveNode removes a physical node and all its virtual nodes from the ring
// Time complexity: O(v * log(v)) where v is number of virtual nodes
func (ch *ConsistentHash[T]) RemoveNode(node T) {
	ch.Lock()
	defer ch.Unlock()

	virtualHashes, exists := ch.virtualMap[node]
	if !exists {
		return
	}

	// Remove all virtual nodes
	for _, hash := range virtualHashes {
		delete(ch.nodeMap, hash)

		// Remove from ring using binary search
		idx := sort.Search(len(ch.ring), func(i int) bool {
			return ch.ring[i] >= hash
		})

		if idx < len(ch.ring) && ch.ring[idx] == hash {
			ch.ring = append(ch.ring[:idx], ch.ring[idx+1:]...)
		}
	}

	// Clean up mappings
	delete(ch.virtualMap, node)
}

// GetNode returns the physical node responsible for the given key
// Uses consistent hashing to find the node with the next highest hash
// Time complexity: O(log v) where v is number of virtual nodes
func (ch *ConsistentHash[T]) GetNode(key string) (T, bool) {
	ch.RLock()
	defer ch.RUnlock()

	if len(ch.ring) == 0 {
		var zero T
		return zero, false
	}

	hash := ch.hashKey(key)

	// Find the first virtual node with hash >= key hash
	idx := sort.Search(len(ch.ring), func(i int) bool {
		return ch.ring[i] >= hash
	})

	// Wrap around to the first node if we've reached the end
	if idx == len(ch.ring) {
		idx = 0
	}

	return ch.nodeMap[ch.ring[idx]], true
}

// GetNodes returns n nodes responsible for the given key
// Useful for replication (e.g., store data on multiple nodes)
// Time complexity: O(log v + n) where v is virtual nodes, n is requested nodes
func (ch *ConsistentHash[T]) GetNodes(key string, n int) []T {
	ch.RLock()
	defer ch.RUnlock()

	if len(ch.ring) == 0 || n <= 0 {
		return []T{}
	}

	hash := ch.hashKey(key)

	// Find starting position
	idx := sort.Search(len(ch.ring), func(i int) bool {
		return ch.ring[i] >= hash
	})

	result := make([]T, 0, n)
	seen := make(map[T]bool)

	// Collect n unique nodes, wrapping around if necessary
	for len(result) < n && len(seen) < len(ch.virtualMap) {
		if idx >= len(ch.ring) {
			idx = 0
		}

		node := ch.nodeMap[ch.ring[idx]]
		if !seen[node] {
			seen[node] = true
			result = append(result, node)
		}

		idx++
	}

	return result
}

// GetNodeCount returns the number of physical nodes in the ring
// Time complexity: O(1)
func (ch *ConsistentHash[T]) GetNodeCount() int {
	ch.RLock()
	defer ch.RUnlock()
	return len(ch.virtualMap)
}

// GetVirtualNodeCount returns the total number of virtual nodes in the ring
// Time complexity: O(1)
func (ch *ConsistentHash[T]) GetVirtualNodeCount() int {
	ch.RLock()
	defer ch.RUnlock()
	return len(ch.ring)
}

// ContainsNode checks if a physical node exists in the ring
// Time complexity: O(1)
func (ch *ConsistentHash[T]) ContainsNode(node T) bool {
	ch.RLock()
	defer ch.RUnlock()
	_, exists := ch.virtualMap[node]
	return exists
}

// GetNodes returns all physical nodes in the ring
// Time complexity: O(n) where n is number of physical nodes
func (ch *ConsistentHash[T]) GetNodesList() []T {
	ch.RLock()
	defer ch.RUnlock()

	nodes := make([]T, 0, len(ch.virtualMap))
	for node := range ch.virtualMap {
		nodes = append(nodes, node)
	}
	return nodes
}

// GetKeyDistribution returns the distribution of keys across nodes
// Useful for monitoring and rebalancing
// Time complexity: O(k * log v) where k is number of keys, v is virtual nodes
func (ch *ConsistentHash[T]) GetKeyDistribution(keys []string) map[T]int {
	ch.RLock()
	defer ch.RUnlock()

	distribution := make(map[T]int)

	for _, key := range keys {
		if node, ok := ch.GetNode(key); ok {
			distribution[node]++
		}
	}

	return distribution
}

// Clear removes all nodes from the ring
// Time complexity: O(1)
func (ch *ConsistentHash[T]) Clear() {
	ch.Lock()
	defer ch.Unlock()

	ch.ring = []uint32{}
	ch.nodeMap = make(map[uint32]T)
	ch.virtualMap = make(map[T][]uint32)
}

// hashKey hashes a string key to a uint32 value
func (ch *ConsistentHash[T]) hashKey(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32()
}

// hashNode hashes a node with its virtual node index
// Using different hash for each virtual node ensures better distribution
func (ch *ConsistentHash[T]) hashNode(node T, index int) uint32 {
	hash := fnv.New32a()
	// Write virtual node index first to differentiate virtual nodes
	hash.Write([]byte{byte(index), byte(index >> 8), byte(index >> 16), byte(index >> 24)})
	// Use fmt.Sprintf to convert any type to string consistently
	hash.Write([]byte(fmt.Sprintf("%v", node)))
	return hash.Sum32()
}
