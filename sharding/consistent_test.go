package sharding

import (
	"fmt"
	"math"
	"testing"
)

func TestNewConsistentHash(t *testing.T) {
	ch := NewConsistentHash[string](150)

	if ch.virtualNodes != 150 {
		t.Errorf("Expected virtual nodes 150, got %d", ch.virtualNodes)
	}

	if ch.GetNodeCount() != 0 {
		t.Errorf("Expected 0 nodes, got %d", ch.GetNodeCount())
	}
}

func TestNewConsistentHashDefault(t *testing.T) {
	ch := NewConsistentHash[string](0)

	if ch.virtualNodes != 150 {
		t.Errorf("Expected default virtual nodes 150, got %d", ch.virtualNodes)
	}
}

func TestAddNode(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")

	if ch.GetNodeCount() != 1 {
		t.Errorf("Expected 1 node, got %d", ch.GetNodeCount())
	}

	if ch.GetVirtualNodeCount() != 10 {
		t.Errorf("Expected 10 virtual nodes, got %d", ch.GetVirtualNodeCount())
	}

	if !ch.ContainsNode("node1") {
		t.Error("Expected node1 to be in ring")
	}
}

func TestAddMultipleNodes(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	if ch.GetNodeCount() != 3 {
		t.Errorf("Expected 3 nodes, got %d", ch.GetNodeCount())
	}

	if ch.GetVirtualNodeCount() != 30 {
		t.Errorf("Expected 30 virtual nodes, got %d", ch.GetVirtualNodeCount())
	}
}

func TestAddDuplicateNode(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node1") // Try to add same node again

	if ch.GetNodeCount() != 1 {
		t.Errorf("Expected 1 node, got %d", ch.GetNodeCount())
	}

	if ch.GetVirtualNodeCount() != 10 {
		t.Errorf("Expected 10 virtual nodes, got %d", ch.GetVirtualNodeCount())
	}
}

func TestRemoveNode(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")

	ch.RemoveNode("node1")

	if ch.GetNodeCount() != 1 {
		t.Errorf("Expected 1 node after removal, got %d", ch.GetNodeCount())
	}

	if ch.ContainsNode("node1") {
		t.Error("Expected node1 to be removed")
	}

	if !ch.ContainsNode("node2") {
		t.Error("Expected node2 to still be in ring")
	}
}

func TestRemoveNonExistentNode(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")

	ch.RemoveNode("node2") // Try to remove non-existent node

	if ch.GetNodeCount() != 1 {
		t.Errorf("Expected 1 node, got %d", ch.GetNodeCount())
	}
}

func TestGetNode(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	node, ok := ch.GetNode("key1")
	if !ok {
		t.Error("Expected to find a node")
	}

	if node != "node1" && node != "node2" && node != "node3" {
		t.Errorf("Unexpected node: %s", node)
	}
}

func TestGetNodeEmptyRing(t *testing.T) {
	ch := NewConsistentHash[string](10)

	_, ok := ch.GetNode("key1")
	if ok {
		t.Error("Expected false for empty ring")
	}
}

func TestGetNodeConsistency(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")

	// Same key should always map to same node
	node1, _ := ch.GetNode("key1")
	node2, _ := ch.GetNode("key1")

	if node1 != node2 {
		t.Error("Same key should map to same node")
	}
}

func TestGetNodes(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	nodes := ch.GetNodes("key1", 2)

	if len(nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(nodes))
	}

	// Check that nodes are unique
	seen := make(map[string]bool)
	for _, node := range nodes {
		if seen[node] {
			t.Error("Expected unique nodes")
		}
		seen[node] = true
	}
}

func TestGetNodesMoreThanAvailable(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")

	nodes := ch.GetNodes("key1", 5)

	if len(nodes) != 2 {
		t.Errorf("Expected 2 nodes (max available), got %d", len(nodes))
	}
}

func TestGetNodesEmptyRing(t *testing.T) {
	ch := NewConsistentHash[string](10)

	nodes := ch.GetNodes("key1", 2)

	if len(nodes) != 0 {
		t.Errorf("Expected 0 nodes, got %d", len(nodes))
	}
}

func TestGetNodesZeroRequest(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")

	nodes := ch.GetNodes("key1", 0)

	if len(nodes) != 0 {
		t.Errorf("Expected 0 nodes, got %d", len(nodes))
	}
}

func TestKeyDistribution(t *testing.T) {
	ch := NewConsistentHash[string](100)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	// Generate 1000 test keys
	keys := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
	}

	distribution := ch.GetKeyDistribution(keys)

	// All keys should be distributed
	total := 0
	for _, count := range distribution {
		total += count
	}

	if total != 1000 {
		t.Errorf("Expected 1000 keys distributed, got %d", total)
	}

	// Check distribution is reasonably balanced (no node has more than 60% or less than 10%)
	maxCount := 0
	minCount := 1000
	for _, count := range distribution {
		if count > maxCount {
			maxCount = count
		}
		if count < minCount {
			minCount = count
		}
	}

	maxRatio := float64(maxCount) / 1000
	minRatio := float64(minCount) / 1000

	if maxRatio > 0.6 {
		t.Errorf("Distribution too skewed: max node has %.2f%% of keys", maxRatio*100)
	}

	if minRatio < 0.1 {
		t.Errorf("Distribution too skewed: min node has %.2f%% of keys", minRatio*100)
	}
}

func TestNodeAdditionMinimalRemapping(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")

	// Record initial mappings
	keys := make([]string, 100)
	initialMappings := make(map[string]string)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		keys[i] = key
		node, _ := ch.GetNode(key)
		initialMappings[key] = node
	}

	// Add a new node
	ch.AddNode("node3")

	// Check that most keys still map to the same node
	remapped := 0
	for _, key := range keys {
		newNode, _ := ch.GetNode(key)
		if initialMappings[key] != newNode {
			remapped++
		}
	}

	// With 3 nodes and good distribution, remapping should be minimal
	// In a perfect distribution, ~1/3 of keys would remap
	remappingRatio := float64(remapped) / 100

	if remappingRatio > 0.5 {
		t.Errorf("Too many keys remapped: %.2f%%", remappingRatio*100)
	}
}

func TestNodeRemovalMinimalRemapping(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	// Record initial mappings
	keys := make([]string, 100)
	initialMappings := make(map[string]string)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		keys[i] = key
		node, _ := ch.GetNode(key)
		initialMappings[key] = node
	}

	// Remove a node
	ch.RemoveNode("node3")

	// Check that only keys from removed node are remapped
	remapped := 0
	for _, key := range keys {
		newNode, _ := ch.GetNode(key)
		if initialMappings[key] != newNode {
			remapped++
		}
	}

	// Only keys from removed node should be remapped
	// With 3 nodes and good distribution, ~1/3 of keys should remap
	remappingRatio := float64(remapped) / 100

	if remappingRatio > 0.5 {
		t.Errorf("Too many keys remapped: %.2f%%", remappingRatio*100)
	}
}

func TestGetNodesList(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")
	ch.AddNode("node3")

	nodes := ch.GetNodesList()

	if len(nodes) != 3 {
		t.Errorf("Expected 3 nodes, got %d", len(nodes))
	}

	// Check all nodes are present
	nodeSet := make(map[string]bool)
	for _, node := range nodes {
		nodeSet[node] = true
	}

	if !nodeSet["node1"] || !nodeSet["node2"] || !nodeSet["node3"] {
		t.Error("Not all nodes returned")
	}
}

func TestClear(t *testing.T) {
	ch := NewConsistentHash[string](10)
	ch.AddNode("node1")
	ch.AddNode("node2")

	ch.Clear()

	if ch.GetNodeCount() != 0 {
		t.Errorf("Expected 0 nodes after clear, got %d", ch.GetNodeCount())
	}

	if ch.GetVirtualNodeCount() != 0 {
		t.Errorf("Expected 0 virtual nodes after clear, got %d", ch.GetVirtualNodeCount())
	}
}

func TestIntegerNodes(t *testing.T) {
	ch := NewConsistentHash[int](10)
	ch.AddNode(1)
	ch.AddNode(2)
	ch.AddNode(3)

	node, ok := ch.GetNode("key1")
	if !ok {
		t.Error("Expected to find a node")
	}

	if node != 1 && node != 2 && node != 3 {
		t.Errorf("Unexpected node: %d", node)
	}
}

func TestConcurrentAccess(t *testing.T) {
	ch := NewConsistentHash[string](100)

	// Add nodes concurrently
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch.AddNode(fmt.Sprintf("node%d", i))
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if ch.GetNodeCount() != 10 {
		t.Errorf("Expected 10 nodes, got %d", ch.GetNodeCount())
	}

	// Get nodes concurrently
	for i := 0; i < 100; i++ {
		go func(i int) {
			ch.GetNode(fmt.Sprintf("key%d", i))
		}(i)
	}
}

func TestLoadBalancing(t *testing.T) {
	ch := NewConsistentHash[string](150)

	// Add multiple nodes
	for i := 0; i < 10; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i))
	}

	// Generate many keys
	keys := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		keys[i] = fmt.Sprintf("user:%d", i%1000)
	}

	distribution := ch.GetKeyDistribution(keys)

	// Calculate coefficient of variation
	counts := make([]int, 0, len(distribution))
	for _, count := range distribution {
		counts = append(counts, count)
	}

	sum := 0
	for _, count := range counts {
		sum += count
	}
	mean := float64(sum) / float64(len(counts))

	variance := 0.0
	for _, count := range counts {
		diff := float64(count) - mean
		variance += diff * diff
	}
	variance /= float64(len(counts))
	coefVar := math.Sqrt(variance) / mean // Coefficient of variation

	// Coefficient of variation should be less than 0.3 (30%)
	if coefVar > 0.3 {
		t.Errorf("Load balancing not good: coefficient of variation %.2f", coefVar)
	}
}

func BenchmarkGetNode(b *testing.B) {
	ch := NewConsistentHash[string](150)
	for i := 0; i < 100; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.GetNode(fmt.Sprintf("key%d", i%1000))
	}
}

func BenchmarkAddNode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := NewConsistentHash[string](150)
		ch.AddNode(fmt.Sprintf("node%d", i))
	}
}

func BenchmarkRemoveNode(b *testing.B) {
	ch := NewConsistentHash[string](150)
	for i := 0; i < 100; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i%100))
		ch.RemoveNode(fmt.Sprintf("node%d", i%100))
	}
}

func BenchmarkGetNodes(b *testing.B) {
	ch := NewConsistentHash[string](150)
	for i := 0; i < 100; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.GetNodes(fmt.Sprintf("key%d", i%1000), 3)
	}
}

func BenchmarkKeyDistribution(b *testing.B) {
	ch := NewConsistentHash[string](150)
	for i := 0; i < 100; i++ {
		ch.AddNode(fmt.Sprintf("node%d", i))
	}

	keys := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.GetKeyDistribution(keys)
	}
}
