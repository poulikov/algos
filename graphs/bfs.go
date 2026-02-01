package graphs

import (
	"fmt"
	"github.com/poulikov/algos/queues"
)

// BFSResult represents the result of a BFS traversal
type BFSResult[T comparable] struct {
	Order        []T       // Order of visited vertices
	Distances    map[T]int // Distance from start vertex
	Parents      map[T]T   // Parent of each vertex in BFS tree
	VisitedOrder []T       // Order in which vertices were first discovered
}

// BFS performs Breadth-First Search starting from a given vertex
// Returns BFSResult with traversal information
// Time complexity: O(V + E)
func BFS[T comparable](graph *Graph[T], start T) (*BFSResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	result := &BFSResult[T]{
		Order:     make([]T, 0, graph.VertexCount()),
		Distances: make(map[T]int),
		Parents:   make(map[T]T),
	}

	visited := make(map[T]bool)
	queue := queues.New[T]()
	queue.Enqueue(start)
	visited[start] = true
	result.Distances[start] = 0
	result.VisitedOrder = append(result.VisitedOrder, start)

	for !queue.IsEmpty() {
		vertex, _ := queue.Dequeue()
		result.Order = append(result.Order, vertex)

		for _, neighbor := range graph.GetNeighbors(vertex) {
			if !visited[neighbor] {
				visited[neighbor] = true
				result.Distances[neighbor] = result.Distances[vertex] + 1
				result.Parents[neighbor] = vertex
				result.VisitedOrder = append(result.VisitedOrder, neighbor)
				queue.Enqueue(neighbor)
			}
		}
	}

	return result, nil
}

// BFSWithPredicate performs BFS until a vertex satisfies the predicate
// Returns the path from start to the found vertex, or nil if not found
// Time complexity: O(V + E)
func BFSWithPredicate[T comparable](graph *Graph[T], start T, predicate func(T) bool) ([]T, error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if predicate(start) {
		return []T{start}, nil
	}

	visited := make(map[T]bool)
	parents := make(map[T]T)
	queue := queues.New[T]()
	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		vertex, _ := queue.Dequeue()

		for _, neighbor := range graph.GetNeighbors(vertex) {
			if !visited[neighbor] {
				visited[neighbor] = true
				parents[neighbor] = vertex

				if predicate(neighbor) {
					return reconstructPath(parents, start, neighbor), nil
				}

				queue.Enqueue(neighbor)
			}
		}
	}

	return nil, fmt.Errorf("no vertex satisfies the predicate")
}

// BFSAllComponents performs BFS on all disconnected components of the graph
// Returns a list of BFSResult for each component
// Time complexity: O(V + E)
func BFSAllComponents[T comparable](graph *Graph[T]) []*BFSResult[T] {
	results := []*BFSResult[T]{}
	visited := make(map[T]bool)

	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			result := BFSComponent(graph, vertex, visited)
			results = append(results, result)
		}
	}

	return results
}

// BFSComponent performs BFS on a single component, tracking visited vertices
func BFSComponent[T comparable](graph *Graph[T], start T, visited map[T]bool) *BFSResult[T] {
	result := &BFSResult[T]{
		Order:     make([]T, 0),
		Distances: make(map[T]int),
		Parents:   make(map[T]T),
	}

	queue := queues.New[T]()
	queue.Enqueue(start)
	visited[start] = true
	result.Distances[start] = 0
	result.VisitedOrder = append(result.VisitedOrder, start)

	for !queue.IsEmpty() {
		vertex, _ := queue.Dequeue()
		result.Order = append(result.Order, vertex)

		for _, neighbor := range graph.GetNeighbors(vertex) {
			if !visited[neighbor] {
				visited[neighbor] = true
				result.Distances[neighbor] = result.Distances[vertex] + 1
				result.Parents[neighbor] = vertex
				result.VisitedOrder = append(result.VisitedOrder, neighbor)
				queue.Enqueue(neighbor)
			}
		}
	}

	return result
}

// ShortestPath finds the shortest path between two vertices using BFS
// Returns the path as a slice of vertices, or nil if no path exists
// Time complexity: O(V + E)
func ShortestPath[T comparable](graph *Graph[T], start, end T) ([]T, error) {
	result, err := BFS(graph, start)
	if err != nil {
		return nil, err
	}

	if _, exists := result.Distances[end]; !exists {
		return nil, fmt.Errorf("no path exists between vertices")
	}

	return reconstructPath(result.Parents, start, end), nil
}

// ShortestPathUnweighted finds the shortest path in an unweighted graph
// Returns the path as a slice of vertices, or nil if no path exists
// Time complexity: O(V + E)
func ShortestPathUnweighted[T comparable](graph *Graph[T], start, end T) ([]T, error) {
	return ShortestPath(graph, start, end)
}

// reconstructPath reconstructs the path from start to end using parent map
func reconstructPath[T comparable](parents map[T]T, start, end T) []T {
	path := []T{end}
	current := end

	for current != start {
		parent, exists := parents[current]
		if !exists {
			return nil
		}
		path = append([]T{parent}, path...)
		current = parent
	}

	return path
}

// Distance returns the shortest distance between two vertices
// Returns -1 if no path exists
// Time complexity: O(V + E)
func Distance[T comparable](graph *Graph[T], start, end T) (int, error) {
	result, err := BFS(graph, start)
	if err != nil {
		return -1, err
	}

	distance, exists := result.Distances[end]
	if !exists {
		return -1, nil
	}

	return distance, nil
}

// ReachableVertices returns all vertices reachable from a start vertex
// Time complexity: O(V + E)
func ReachableVertices[T comparable](graph *Graph[T], start T) ([]T, error) {
	result, err := BFS(graph, start)
	if err != nil {
		return nil, err
	}
	return result.VisitedOrder, nil
}

// BFSLevelOrder returns vertices grouped by their distance from start
// Returns a slice where each element is a level (slice of vertices at same distance)
// Time complexity: O(V + E)
func BFSLevelOrder[T comparable](graph *Graph[T], start T) ([][]T, error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	visited := make(map[T]bool)
	levels := [][]T{{start}}
	visited[start] = true

	for currentLevel := 0; ; currentLevel++ {
		nextLevel := []T{}

		for _, vertex := range levels[currentLevel] {
			for _, neighbor := range graph.GetNeighbors(vertex) {
				if !visited[neighbor] {
					visited[neighbor] = true
					nextLevel = append(nextLevel, neighbor)
				}
			}
		}

		if len(nextLevel) == 0 {
			break
		}

		levels = append(levels, nextLevel)
	}

	return levels, nil
}

// IsReachable checks if there's a path from start to end
// Time complexity: O(V + E)
func IsReachable[T comparable](graph *Graph[T], start, end T) (bool, error) {
	result, err := BFS(graph, start)
	if err != nil {
		return false, err
	}

	_, exists := result.Distances[end]
	return exists, nil
}

// BFSCycleDetection detects if there's a cycle in an undirected graph using BFS
// Time complexity: O(V + E)
func BFSCycleDetection[T comparable](graph *Graph[T]) bool {
	if !graph.IsDirected() {
		return false
	}

	visited := make(map[T]bool)
	recursionStack := make(map[T]bool)

	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			if dfsHasCycle(graph, vertex, visited, recursionStack) {
				return true
			}
		}
	}

	return false
}
