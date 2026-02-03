# AGENTS.md - Agent Guide for algos Repository

This document helps AI agents work effectively in this algorithms and data structures codebase.

## Project Overview

Comprehensive collection of algorithms and data structures implemented in Go with full test coverage. Each implementation uses Go generics for type safety and flexibility.

**Module**: `github.com/poulikov/algos`
**Go Version**: 1.24.0 or later
**Total Files**: 110+ Go files (implementations + tests)

## Commands

### Testing
```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./graphs
go test ./sorting
go test ./trees

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test ./... -cover

# Run benchmarks
go test ./... -bench .

# Run specific benchmark
go test ./graphs -bench BenchmarkBFS
```

### Building
```bash
# Build module (useful to catch compilation errors)
go build ./...
```

### Dependencies
```bash
# Download dependencies
go mod download

# Tidy up go.mod
go mod tidy
```

## Project Structure

```
algos/
├── backtracking/     # Backtracking algorithms (nqueens, sudoku)
├── cache/           # Caching structures
├── dp/              # Dynamic programming (coin_change, knapsack, LCS, LIS, edit_distance)
├── doublelist/      # Doubly linked list
├── graphs/          # Graph algorithms (BFS, DFS, Dijkstra, etc.)
├── greedy/          # Greedy algorithms (activity, huffman)
├── heaps/           # Heap data structures (min-heap, max-heap, priority queue)
├── linkedlist/      # Singly linked list
├── queues/          # Queue implementations (FIFO queue, deque)
├── searching/       # Searching algorithms (binary, exponential, interpolation)
├── set/             # Hash-based sets and maps
├── slidingwindow/   # Sliding window techniques
├── sorting/         # Sorting algorithms (merge, quick, heap, etc.)
├── stacks/          # Stack implementation (LIFO)
├── strings/         # String algorithms (KMP, Rabin-Karp, Boyer-Moore)
├── structures/      # Special data structures (bloom filter, skip list)
├── trees/           # Tree data structures (BST, AVL, Red-Black, Trie, etc.)
├── twopointers/     # Two-pointer techniques
├── go.mod           # Module definition
├── go.sum           # Dependency checksums
└── README.md        # Project documentation
```

## Code Conventions

### Generics
- Use Go generics extensively for type safety
- Type constraints:
  - `T comparable` - for types that support equality comparison (used in maps, sets, graphs)
  - `T constraints.Ordered` - for types that support ordering (used in sorting, search trees, heaps)
  - `T any` - when no specific constraint is needed (requires custom comparator function)

Example:
```go
// For comparable types (can be used as map keys)
type Graph[T comparable] struct { ... }

// For ordered types (can be compared with <, >, <=, >=)
type BST[T constraints.Ordered] struct { ... }

// For any type with custom comparison
type Heap[T any] struct {
    less func(T, T) bool
}
```

### Package Imports
- Internal imports use full module path: `github.com/poulikov/algos/queues`
- Example from `graphs/bfs.go`:
```go
import (
    "github.com/poulikov/algos/queues"
)
```

### Naming Conventions
- **Constructors**: `NewTypeName()` or `New()` - e.g., `NewBST()`, `NewHeap()`
- **Methods**: PascalCase for exported - e.g., `Insert()`, `Search()`, `Inorder()`
- **Constants**: PascalCase for exported - e.g., `MinHeap`, `Directed`, `Undirected`
- **Errors**: PascalCase with `Err` prefix - e.g., `ErrHeapEmpty`, `ErrStackEmpty`
- **Private types**: lowerCase for unexported - e.g., `bstNode[T]`, `heapNode[T]`

### Documentation Comments
Every exported function includes:
- Description of what it does
- Time complexity comment: `// Time complexity: O(...)`
- Space complexity comment (if notable): `// Space complexity: O(...)`
- Parameter descriptions if non-obvious
- Return value descriptions if non-obvious

Example:
```go
// Insert adds a value to the tree
// Time complexity: O(h) where h is the height of the tree (O(log n) average, O(n) worst case)
func (bst *BST[T]) Insert(value T) { ... }
```

### Error Handling
- Define package-level error variables for common errors:
```go
var (
    ErrHeapEmpty = errors.New("heap is empty")
    ErrStackEmpty = errors.New("stack is empty")
)
```
- Return errors from functions that can fail:
```go
func (h *Heap[T]) Extract() (T, error) {
    if h.IsEmpty() {
        var zero T
        return zero, ErrHeapEmpty
    }
    // ...
}
```

### File Organization
Each package typically has:
1. **Main implementation file**: `{topic}.go` - e.g., `bfs.go`, `bst.go`, `merge.go`
2. **Test file**: `{topic}_test.go` - e.g., `bfs_test.go`, `bst_test.go`, `merge_test.go`
3. **README.md** (optional): Some packages have detailed README with usage examples (e.g., `stacks/README.md`, `cache/README.md`)

## Testing Patterns

### Test Structure
- Test files follow Go conventions with `_test.go` suffix
- Tests use `testing` package
- Test functions named `Test{FunctionName}` - e.g., `TestBFS`, `TestInsert`
- Benchmark functions named `Benchmark{FunctionName}` - e.g., `BenchmarkBFS`

### Test Coverage
- Comprehensive test coverage is expected
- Test edge cases: empty structures, single elements, errors
- Test both happy path and error paths
- Example from `graphs/bfs_test.go`:
```go
func TestBFS(t *testing.T) {
    g := NewUndirectedGraph[int]()
    g.AddEdge(1, 2, 0)
    g.AddEdge(2, 3, 0)

    result, err := BFS(g, 1)
    if err != nil {
        t.Fatal(err)
    }

    if len(result.Order) != 3 {
        t.Errorf("Expected 3 vertices in order, got %d", len(result.Order))
    }
}
```

### Benchmark Pattern
```go
func BenchmarkBFS(b *testing.B) {
    g := NewUndirectedGraph[int]()
    for i := 0; i < 1000; i++ {
        g.AddEdge(i, i+1, 1)
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        BFS(g, 0)
    }
}
```

## Common Patterns

### Type Constraints Selection
| Constraint | Use When |
|-----------|----------|
| `T comparable` | Types stored in maps, used as keys, checked for equality (graphs, sets) |
| `T constraints.Ordered` | Types that can be compared/ordered (sorting, trees, heaps) |
| `T any` | Any type, require custom comparator function |

### Constructor Pattern
```go
// Generic constructor with type parameter
func New[T constraints.Ordered]() *BST[T] {
    return &BST[T]{}
}

// Constructor with parameters
func NewGraph[T comparable](graphType GraphType) *Graph[T] {
    return &Graph[T]{
        vertices:  make(map[T]struct{}),
        graphType: graphType,
    }
}
```

### Helper Functions
- Private helper functions are lowerCase - e.g., `insertRecursive()`, `heapifyUp()`
- Recursive helpers often used in tree/graph algorithms
- Prefix/suffix methods for variants - e.g., `MergeSortCopy()`, `MergeSortDescending()`

### String Representation
Most types implement `String() string` for debugging:
```go
func (g *Graph[T]) String() string {
    // Return human-readable representation
}
```

## Important Gotchas

### Zero Values
When returning errors for empty structures, use zero value of type `T`:
```go
func (h *Heap[T]) Extract() (T, error) {
    if h.IsEmpty() {
        var zero T  // Get zero value of type T
        return zero, ErrHeapEmpty
    }
    // ...
}
```

### Map Iteration
When iterating over maps (e.g., for graph vertices), the order is non-deterministic. Use slices when order matters:
```go
// Good - deterministic order
vertices := g.GetVertices()  // Returns []T
for _, v := range vertices {
    // ...
}

// Be careful - non-deterministic order
for v := range g.vertices {
    // ...
}
```

### Slice Capacity
Pre-allocate slice capacity when size is known for better performance:
```go
result := make([]T, 0, bst.size)  // Pre-allocate capacity
```

### Recursion Depth
Tree algorithms use recursion. For very large trees, consider iterative approaches to avoid stack overflow.

### Undirected vs Directed Graphs
Graph package supports both types. Many algorithms work differently:
- Undirected: edges work both ways
- Directed: edges have direction (from → to)
- Check `graph.IsDirected()` when behavior differs

## Dependencies

### Required
- `go 1.24.0` or later

### External
- `golang.org/x/exp/constraints` - For generic type constraints (Ordered)

Import example:
```go
import "golang.org/x/exp/constraints"
```

## Adding New Implementations

When adding a new algorithm or data structure:

1. **Create directory**: `/category/{name}/`
2. **Create implementation**: `{name}.go`
   - Use appropriate type constraint (`comparable` or `Ordered`)
   - Include time/space complexity in comments
   - Define constructor functions
   - Implement standard methods (Insert, Search, Delete, etc.)
3. **Create tests**: `{name}_test.go`
   - Test happy path
   - Test error cases
   - Test edge cases (empty, single element)
   - Add benchmarks for performance-critical code
4. **Run tests**: `go test ./category/{name}`
5. **Update README.md** (if applicable): Add entry in main README

## Module Details

**Module Name**: `github.com/poulikov/algos`
**Version**: No explicit versioning in go.mod (uses semantic versions in Git)
**Dependency Management**: Standard Go modules

When importing this module in external projects:
```go
import "github.com/poulikov/algos/stacks"
import "github.com/poulikov/algos/graphs"
// etc.
```

## Language Notes

- Most documentation is in English
- Some README files are in Russian (e.g., `stacks/README.md`)
- Code comments and identifiers are in English

## Testing Checklist

Before considering an implementation complete:
- [ ] All public functions tested
- [ ] Edge cases covered (empty, single element, max size)
- [ ] Error paths tested
- [ ] Time/space complexity documented
- [ ] Benchmarks added for performance-critical algorithms
- [ ] `go test ./package` passes
- [ ] `go test ./...` still passes (no regressions)

## Debugging Tips

### Test Failures
- Use `-v` flag for verbose output: `go test -v ./package`
- Use `-run` flag to test specific function: `go test -v ./package -run TestFunctionName`

### Benchmark Comparison
- Run benchmarks multiple times to account for variance
- Use `benchstat` tool to compare benchmarks if available
- Use `b.ResetTimer()` in benchmarks to exclude setup time

### Type Issues
- Check that type constraint (`comparable` vs `Ordered`) matches use case
- Verify custom comparator functions are correct when using `T any`
- Ensure zero value handling is correct for generic types
