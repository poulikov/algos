# Consistent Hashing for Sharding

Consistent hashing is a technique used to distribute data across multiple servers in a way that minimizes data movement when servers are added or removed.

## Overview

This implementation provides:
- **Consistent hashing ring** with configurable virtual nodes
- **Minimal data remapping** when nodes are added/removed
- **Thread-safe** operations using RWMutex
- **Generic type support** for node identifiers (strings, integers, etc.)
- **Replication support** with multiple node selection

## How It Works

1. **Virtual Nodes**: Each physical node is represented by multiple virtual nodes on a ring
2. **Hash Ring**: All virtual nodes are placed on a ring (0 to 2^32-1)
3. **Key Assignment**: Keys are hashed and assigned to the next node clockwise on the ring
4. **Minimal Impact**: Adding/removing a node only affects keys that map to its virtual nodes

## Complexity

| Operation | Time Complexity | Space Complexity |
|-----------|---------------|------------------|
| AddNode | O(v * log v) | O(v) |
| RemoveNode | O(v * log v) | O(v) |
| GetNode | O(log v) | O(1) |
| GetNodes | O(log v + n) | O(n) |

Where:
- `v` = number of virtual nodes per physical node (default: 150)
- `n` = number of nodes requested for replication

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/poulikov/algos/sharding"
)

func main() {
    // Create a consistent hash ring with 150 virtual nodes per physical node
    ch := sharding.NewConsistentHash[string](150)

    // Add database nodes
    ch.AddNode("db-server-1")
    ch.AddNode("db-server-2")
    ch.AddNode("db-server-3")

    // Find which node should store a key
    node, ok := ch.GetNode("user:1234")
    if ok {
        fmt.Printf("Key 'user:1234' stored on: %s\n", node)
    }
}
```

### Replication Example

```go
// Get multiple nodes for replication (store on 3 different nodes)
replicas := ch.GetNodes("user:1234", 3)
for i, node := range replicas {
    fmt.Printf("Replica %d: %s\n", i+1, node)
}
```

### Dynamic Node Management

```go
// Add new node - only affects ~1/N keys where N is total nodes
ch.AddNode("db-server-4")

// Remove failed node - only affects keys on that node
ch.RemoveNode("db-server-2")
```

### Load Balancing Check

```go
// Generate sample keys
keys := []string{
    "user:1", "user:2", "user:3",
    "session:abc", "session:def", "session:ghi",
}

// Check distribution
distribution := ch.GetKeyDistribution(keys)
for node, count := range distribution {
    fmt.Printf("%s: %d keys\n", node, count)
}
```

## Choosing Virtual Node Count

The number of virtual nodes per physical node affects distribution quality:

| Virtual Nodes | Pros | Cons |
|---------------|-------|------|
| 50-100 | Lower memory usage, faster add/remove | Less balanced distribution |
| 150-200 | Good balance of performance and distribution | Recommended default |
| 300-500 | Better distribution, smoother scaling | Higher memory usage |

**Recommendation**: Use 150-200 virtual nodes for most production systems.

## Thread Safety

The `ConsistentHash` struct uses `sync.RWMutex` for concurrent access:
- **Read operations** (GetNode, GetNodes, ContainsNode): Multiple goroutines can read simultaneously
- **Write operations** (AddNode, RemoveNode, Clear): Exclusive lock held during operation

## Use Cases

1. **Database Sharding**: Distribute data across multiple database servers
2. **Caching**: Distribute cache keys across multiple cache servers
3. **Load Balancing**: Distribute requests across multiple servers
4. **File Storage**: Distribute files across storage nodes
5. **CDN**: Route requests to nearest/available edge servers

## Testing

Run tests:
```bash
go test ./sharding
```

Run tests with coverage:
```bash
go test ./sharding -cover
```

Run benchmarks:
```bash
go test ./sharding -bench=. -benchmem
```

## Performance Benchmarks

On Apple M3 Pro:
- `GetNode`: ~124 ns/op, 1 allocation
- `GetNodes` (3 nodes): ~271 ns/op, 2 allocations
- `AddNode`: ~24.8 µs/op (for 150 virtual nodes)
- `RemoveNode`: ~27.0 µs/op (for 150 virtual nodes)

## Limitations

1. **Hash Function**: Uses FNV-1a 32-bit hash, which may have collisions in very large deployments
2. **Memory**: Each virtual node adds memory overhead
3. **Cold Start**: Ring sorting happens on node addition

## Best Practices

1. **Monitor Distribution**: Regularly check key distribution to ensure even load
2. **Handle Failures**: Implement retry logic when GetNode fails (empty ring)
3. **Use Replication**: Use GetNodes(n) for critical data (n >= 3)
4. **Graceful Removal**: Remove nodes only after ensuring data has been replicated
5. **Test Distribution**: Generate test keys to verify distribution before production use

## API Reference

### Constructors

```go
func NewConsistentHash[T comparable](virtualNodesPerNode int) *ConsistentHash[T]
```

Creates a new consistent hashing ring.

- `virtualNodesPerNode`: Number of virtual nodes per physical node (default: 150 if 0)

### Node Management

```go
func (ch *ConsistentHash[T]) AddNode(node T)
```
Adds a physical node to the hash ring with virtual nodes.

```go
func (ch *ConsistentHash[T]) RemoveNode(node T)
```
Removes a physical node and all its virtual nodes from the ring.

```go
func (ch *ConsistentHash[T]) ContainsNode(node T) bool
```
Checks if a node exists in the ring.

### Key Operations

```go
func (ch *ConsistentHash[T]) GetNode(key string) (T, bool)
```
Returns the physical node responsible for the given key.

```go
func (ch *ConsistentHash[T]) GetNodes(key string, n int) []T
```
Returns n unique physical nodes responsible for the given key (useful for replication).

### Analysis

```go
func (ch *ConsistentHash[T]) GetKeyDistribution(keys []string) map[T]int
```
Returns the distribution of keys across nodes.

### Info

```go
func (ch *ConsistentHash[T]) GetNodeCount() int
```
Returns the number of physical nodes in the ring.

```go
func (ch *ConsistentHash[T]) GetVirtualNodeCount() int
```
Returns the total number of virtual nodes in the ring.

```go
func (ch *ConsistentHash[T]) GetNodesList() []T
```
Returns all physical nodes in the ring.

```go
func (ch *ConsistentHash[T]) Clear()
```
Removes all nodes from the ring.
