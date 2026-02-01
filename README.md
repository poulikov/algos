# Algorithms and Data Structures in Go

Comprehensive collection of algorithms and data structures implemented in Go with full test coverage.

![Static Badge](https://img.shields.io/badge/vibecoded-blue?style=for-the-badge&logo=%F0%9F%A4%96&label=%F0%9F%A4%96)

## Requirements

- Go 1.24.0 or later

## Data Structures

### Linked Lists
- **Singly Linked List** - Basic linked list implementation with standard operations
  - Location: `linkedlist/linkedlist.go`
- **Doubly Linked List** - Efficient bidirectional traversal using generics
  - Location: `doublelist/doubly_linked_list.go`
  - Example: `examples/doublelist_example.go`

### Trees
- **Binary Search Tree (BST)** - Standard BST with insertion, search, and deletion
  - Location: `trees/bst.go`
- **AVL Tree** - Self-balancing BST for guaranteed O(log n) operations
  - Location: `trees/avl.go`
- **B-Tree** - Balanced tree optimized for disk storage and large datasets
  - Location: `trees/btree.go`
- **Red-Black Tree** - Self-balancing BST with efficient insertion/deletion
  - Location: `trees/redblack.go`
- **Trie** - Prefix tree for efficient string storage and retrieval
  - Location: `trees/trie.go`
- **Fenwick Tree (Binary Indexed Tree)** - Efficient range queries and updates
  - Location: `trees/fenwick.go`
- **Segment Tree** - Range query data structure with O(log n) operations
  - Location: `trees/segment.go`
- **Union Find (Disjoint Set)** - Efficiently tracks connected components
  - Location: `trees/unionfind.go`

### Heaps
- **Binary Heap** - Min-heap and max-heap implementations
  - Location: `heaps/heap.go`
- **Priority Queue** - Queue with priority-based ordering
  - Location: `heaps/priority.go`

### Queues
- **Queue** - FIFO queue with enqueue/dequeue operations
  - Location: `queues/queue.go`
- **Deque** - Double-ended queue with operations from both ends
  - Location: `queues/deque.go`

### Stacks
- **Stack** - LIFO stack with push/pop operations
  - Location: `stacks/stack.go`

### Sets
- **HashSet** - Hash-based set with O(1) average operations
  - Location: `set/hashset.go`
- **HashMap** - Hash-based map with key-value pairs
  - Location: `set/hashmap.go`

### Special Structures
- **Bloom Filter** - Probabilistic data structure for membership testing
  - Location: `structures/bloom.go`
- **Skip List** - Probabilistic alternative to balanced trees
  - Location: `structures/skiplist.go`
- **Cache** - Generic caching implementation
  - Location: `cache/cache.go`

## Algorithms

### Sorting Algorithms
- **Bubble Sort** - Simple comparison-based sorting
  - Location: `sorting/bubble.go`
- **Selection Sort** - In-place comparison sorting
  - Location: `sorting/selection.go`
- **Insertion Sort** - Efficient for small/nearly sorted arrays
  - Location: `sorting/insertion.go`
- **Merge Sort** - Divide-and-conquer, stable sorting
  - Location: `sorting/merge.go`
- **Quick Sort** - Efficient in-place divide-and-conquer
  - Location: `sorting/quick.go`
- **Heap Sort** - Uses heap data structure
  - Location: `sorting/heap.go`
- **Counting Sort** - Non-comparison sort for integer arrays
  - Location: `sorting/counting.go`
- **Radix Sort** - Non-comparison sort for integers
  - Location: `sorting/radix.go`

### Searching Algorithms
- **Binary Search** - Efficient search on sorted arrays
  - Location: `searching/binary.go`
- **Exponential Search** - For unbounded/infinite sorted arrays
  - Location: `searching/exponential.go`
- **Interpolation Search** - Improved binary search for uniformly distributed data
  - Location: `searching/interpolation.go`

### String Algorithms
- **Knuth-Morris-Pratt (KMP)** - Pattern matching with O(n + m) time
  - Location: `strings/kmp.go`
- **Rabin-Karp** - Rolling hash-based pattern matching
  - Location: `strings/rabin_karp.go`
- **Boyer-Moore** - Efficient pattern matching with bad character heuristics
  - Location: `strings/boyer_moore.go`

### Graph Algorithms
- **Breadth-First Search (BFS)** - Level-order graph traversal
  - Location: `graphs/bfs.go`
- **Depth-First Search (DFS)** - Graph traversal with recursive/iterative approaches
  - Location: `graphs/dfs.go`
- **Dijkstra's Algorithm** - Shortest paths from source to all vertices
  - Location: `graphs/dijkstra.go`
- **Bellman-Ford** - Shortest paths handling negative weights
  - Location: `graphs/bellman_ford.go`
- **Floyd-Warshall** - All-pairs shortest paths
  - Location: `graphs/floyd_warshall.go`
- **A* Search** - Heuristic-based pathfinding
  - Location: `graphs/astar.go`
- **Prim's Algorithm** - Minimum spanning tree
  - Location: `graphs/prim.go`
- **Kruskal's Algorithm** - Minimum spanning tree with union find
  - Location: `graphs/kruskal.go`
- **DAG (Directed Acyclic Graph)** - Topological sorting and cycle detection
  - Location: `graphs/dag.go`

### Dynamic Programming
- **Coin Change** - Minimum coins to make amount
  - Location: `dp/coin_change.go`
- **Knapsack** - 0/1 knapsack problem
  - Location: `dp/knapsack.go`
- **Longest Common Subsequence (LCS)** - Longest subsequence common to two sequences
  - Location: `dp/lcs.go`
- **Longest Increasing Subsequence (LIS)** - Longest strictly increasing subsequence
  - Location: `dp/lis.go`
- **Edit Distance** - Minimum operations to transform one string to another
  - Location: `dp/edit_distance.go`

### Greedy Algorithms
- **Activity Selection** - Maximize non-overlapping activities
  - Location: `greedy/activity.go`
- **Huffman Coding** - Optimal prefix coding for data compression
  - Location: `greedy/huffman.go`

### Backtracking
- **N-Queens** - Place N queens on N×N chessboard
  - Location: `backtracking/nqueens.go`
- **Sudoku Solver** - Solve Sudoku puzzles
  - Location: `backtracking/sudoku.go`

### Two Pointers
- **Two Pointers Technique** - Various problems solved with two-pointer approach
  - Location: `twopointers/twopointers.go`

### Sliding Window
- **Sliding Window** - Fixed/variable window techniques
  - Location: `slidingwindow/slidingwindow.go`

## Usage

All implementations use Go generics for type safety and flexibility. Example:

```go
import "github.com/poulikov/algos/stacks"

// Create a stack of integers
stack := stacks.New[int]()
stack.Push(10)
stack.Push(20)
val := stack.Pop()
```

For more examples, check the `examples/` directory.

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./sorting
go test ./graphs
go test ./trees
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Project Structure

```
algos/
├── backtracking/    # Backtracking algorithms
├── cache/          # Caching structures
├── dp/             # Dynamic programming
├── doublelist/     # Doubly linked list
├── graphs/         # Graph algorithms and structures
├── greedy/         # Greedy algorithms
├── heaps/          # Heap data structures
├── linkedlist/     # Singly linked list
├── queues/         # Queue implementations
├── searching/      # Searching algorithms
├── set/            # Hash-based sets and maps
├── slidingwindow/  # Sliding window techniques
├── sorting/        # Sorting algorithms
├── stacks/         # Stack implementation
├── strings/        # String algorithms
├── structures/     # Special data structures
├── trees/          # Tree data structures
├── twopointers/    # Two-pointer techniques
├── examples/       # Usage examples
└── README.md       # This file
```

All implementations include:
- Comprehensive test coverage
- Time and space complexity documentation
- Generic type support where applicable
- Clean, idiomatic Go code
