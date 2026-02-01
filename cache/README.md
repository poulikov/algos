# Sharded LRU Cache

A sharded in-memory LRU cache with concurrent access in Go.

## Description

The cache implements the `Cache[K comparable, V any]` interface with two primary methods:

```go
type Cache[K comparable, V any] interface {
    Set(key K, val V)
    Get(key K) (V, error)
}
```

The cache is divided into multiple segments (shards), each with its own mutex for thread-safe concurrent access. Keys are distributed across shards using hashing, which allows for even load distribution and avoids lock contention between shards. Each shard maintains an LRU (Least Recently Used) eviction policy to manage memory usage.

## Core Components

1. **Cache Interface** - defines the main cache operations
2. **lruNode** - doubly-linked list node for LRU implementation (no boxing)
3. **shard** - cache segment containing data map, LRU list, and mutex for concurrent access
4. **shardCache** - main cache container with multiple shards
5. **getShardIndex** - calculates shard index for a given key using type-aware hashing
6. **Set** - sets a value with shard locking and LRU update
7. **Get** - retrieves a value with locking and LRU update

## Usage

```go
package main

import (
    "fmt"
    "algos/cache"
)

func main() {
    // Create a cache with 10 shards, each holding up to 1000 items
    myCache := cache.NewCache[string, int](10, 1000)
    
    // Set values
    myCache.Set("user:123", 456)
    myCache.Set("product:abc", 789)
    
    // Get values
    if val, err := myCache.Get("user:123"); err == nil {
        fmt.Printf("User ID: %d\n", val)
    }
    
    // Handle error when key doesn't exist
    if _, err := myCache.Get("nonexistent"); err != cache.ErrNotFound {
        fmt.Println("Error:", err)
    }
}
```

## Functions and Structures

### NewCache[K comparable, V any](shardsCount int, shardSize int) Cache[K, V]

Creates a new sharded LRU cache instance.

**Parameters:**
- `shardsCount` - number of shards for key distribution (default: 1 if ≤ 0)
- `shardSize` - maximum number of items per shard (default: 1000 if ≤ 0)

### Set(key K, val V)

Sets a value in the cache for the given key. Overwrites the value if the key already exists and moves it to the front of the LRU list. Evicts the least recently used item if the shard is at capacity.

**Parameters:**
- `key` - cache key (must be comparable type)
- `val` - value to store

### Get(key K) (V, error)

Retrieves a value from the cache for the given key. Moves the accessed item to the front of the LRU list.

**Returns:**
- Value of type V
- `ErrNotFound` error if the key is not found

### Error Types

```go
var ErrNotFound = &NotFoundError{}

type NotFoundError struct{}
func (e *NotFoundError) Error() string {
    return "key not found in cache"
}
```

## Testing

To run all tests:

```bash
go test ./cache
```

To run tests with coverage:

```bash
go test ./cache -cover
```

To run a specific test:

```bash
go test ./cache -run TestSetAndGet
```

To run benchmarks:

```bash
go test ./cache -bench .
```

## Performance Comparison

We compared our sharded LRU cache with other popular solutions at different data sizes:

### At 100 items

| Implementation | Read | Write | Allocations (write) |
|----------------|------|-------|-------------------|
| Sharded LRU Cache | ~248 ns/op | ~241 ns/op | 3 allocs/op |
| sync.Map | ~245 ns/op | ~271 ns/op | 5 allocs/op |
| Mutex Map | ~253 ns/op | ~770 ns/op | 2 allocs/op |

### At 100,000 items

| Implementation | Read | Write | Allocations (write) |
|----------------|------|-------|-------------------|
| Sharded LRU Cache | ~249 ns/op | ~254 ns/op | 3 allocs/op |
| sync.Map | ~247 ns/op | ~267 ns/op | 5 allocs/op |
| Mutex Map | ~255 ns/op | ~754 ns/op | 2 allocs/op |

### At 10,000,000 items

For very large data sizes (10M items), performance may vary significantly due to garbage collector pressure and memory usage. Benchmarks with such data volumes require significant time and resources to execute.

**Conclusions:**
- All three implementations show similar read performance
- After optimization, **sharded LRU cache is now faster** than others for writes (~241 ns/op vs ~271 ns/op for sync.Map)
- Optimizations reduced allocations from 7 to 3 (57% improvement)
- Our sharded cache offers an excellent balance: **best write performance** + LRU eviction + memory management
- Performance is relatively independent of cache size

**Key Optimizations:**
- Eliminated value boxing (storing `V` instead of `any`)
- Type-aware hashing for basic types without `fmt.Sprintf`
- Custom LRU list without `interface{}`
- Using `sync.Pool` for FNV hashers
- Fixed race condition in `Get`

## Features

- **Sharding**: keys distributed across shards using hashing
- **Concurrent Access**: each shard uses `sync.RWMutex` for thread-safe access
- **LRU Eviction**: automatic eviction of least recently used items when shard reaches capacity
- **Generics**: support for arbitrary key and value types (Go 1.18+)
- **Thread-Safe**: no race conditions under concurrent access
- **Memory Management**: configurable shard size limits memory usage

## Test Coverage

The cache includes comprehensive tests:

- `TestSetAndGet` - basic operations
- `TestGetNotFound` - handling non-existent keys
- `TestConcurrentAccess` - concurrent access patterns
- `TestDifferentKeyTypes` - different key types support
- `TestShardDistribution` - key distribution verification
- `TestOverwrite` - value overwriting
- `TestGetReturnsZeroValue` - zero value return
- `TestErrorType` - error type verification
- `BenchmarkConcurrentRead` - read performance
- `BenchmarkConcurrentWrite` - write performance
