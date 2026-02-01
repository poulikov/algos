package graphs

import (
	"fmt"
)

// DFSResult represents the result of a DFS traversal
type DFSResult[T comparable] struct {
	Order         []T       // Order of visited vertices
	DiscoveryTime map[T]int // Time when vertex was first discovered
	FinishTime    map[T]int // Time when DFS finished processing vertex
	Parents       map[T]T   // Parent of each vertex in DFS tree
	VisitedOrder  []T       // Order in which vertices were first discovered
	IsBackEdge    bool      // Whether a back edge was found (indicates cycle)
}

// DFS performs Depth-First Search starting from a given vertex
// Returns DFSResult with traversal information
// Time complexity: O(V + E)
func DFS[T comparable](graph *Graph[T], start T) (*DFSResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	result := &DFSResult[T]{
		Order:         make([]T, 0, graph.VertexCount()),
		DiscoveryTime: make(map[T]int),
		FinishTime:    make(map[T]int),
		Parents:       make(map[T]T),
	}

	visited := make(map[T]bool)
	time := 0

	dfsVisit(graph, start, visited, result, &time)

	return result, nil
}

// dfsVisit is the recursive DFS helper function
func dfsVisit[T comparable](graph *Graph[T], vertex T, visited map[T]bool, result *DFSResult[T], time *int) {
	visited[vertex] = true
	*time++
	result.DiscoveryTime[vertex] = *time
	result.VisitedOrder = append(result.VisitedOrder, vertex)

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			result.Parents[neighbor] = vertex
			dfsVisit(graph, neighbor, visited, result, time)
		}
	}

	*time++
	result.FinishTime[vertex] = *time
	result.Order = append(result.Order, vertex)
}

// DFSAllComponents performs DFS on all disconnected components of the graph
// Returns a DFSResult for the entire graph
// Time complexity: O(V + E)
func DFSAllComponents[T comparable](graph *Graph[T]) *DFSResult[T] {
	result := &DFSResult[T]{
		Order:         make([]T, 0, graph.VertexCount()),
		DiscoveryTime: make(map[T]int),
		FinishTime:    make(map[T]int),
		Parents:       make(map[T]T),
	}

	visited := make(map[T]bool)
	time := 0

	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			dfsVisit(graph, vertex, visited, result, &time)
		}
	}

	return result
}

// DFSIterative performs Depth-First Search iteratively (using a stack)
// Time complexity: O(V + E)
func DFSIterative[T comparable](graph *Graph[T], start T) (*DFSResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	result := &DFSResult[T]{
		Order:         make([]T, 0, graph.VertexCount()),
		DiscoveryTime: make(map[T]int),
		FinishTime:    make(map[T]int),
		Parents:       make(map[T]T),
		VisitedOrder:  make([]T, 0),
	}

	visited := make(map[T]bool)
	stack := []T{start}
	time := 0

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			visited[vertex] = true
			time++
			result.DiscoveryTime[vertex] = time
			result.VisitedOrder = append(result.VisitedOrder, vertex)
			result.Order = append(result.Order, vertex)

			neighbors := graph.GetNeighbors(vertex)
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					result.Parents[neighbors[i]] = vertex
					stack = append(stack, neighbors[i])
				}
			}
		}
	}

	return result, nil
}

// DFSWithPredicate performs DFS until a vertex satisfies the predicate
// Returns the path from start to the found vertex, or nil if not found
// Time complexity: O(V + E)
func DFSWithPredicate[T comparable](graph *Graph[T], start T, predicate func(T) bool) ([]T, error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if predicate(start) {
		return []T{start}, nil
	}

	visited := make(map[T]bool)
	parents := make(map[T]T)

	if dfsFindPath(graph, start, predicate, visited, parents) {
		end := findVertexSatisfying(parents, predicate)
		if end != nil {
			return reconstructDFSPath(parents, start, *end), nil
		}
	}

	return nil, fmt.Errorf("no vertex satisfies the predicate")
}

// dfsFindPath performs DFS to find a path to a vertex that satisfies the predicate
func dfsFindPath[T comparable](graph *Graph[T], vertex T, predicate func(T) bool, visited map[T]bool, parents map[T]T) bool {
	visited[vertex] = true

	if predicate(vertex) {
		return true
	}

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			parents[neighbor] = vertex
			if dfsFindPath(graph, neighbor, predicate, visited, parents) {
				return true
			}
		}
	}

	return false
}

// findVertexSatisfying finds a vertex in the parents map that satisfies the predicate
func findVertexSatisfying[T comparable](parents map[T]T, predicate func(T) bool) *T {
	for vertex := range parents {
		if predicate(vertex) {
			return &vertex
		}
	}
	return nil
}

// reconstructDFSPath reconstructs the path from start to end using parent map
func reconstructDFSPath[T comparable](parents map[T]T, start, end T) []T {
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

// TopologicalSort performs topological sorting of a DAG using DFS
// Returns vertices in topological order, or error if graph has a cycle
// Time complexity: O(V + E)
func TopologicalSort[T comparable](graph *Graph[T]) ([]T, error) {
	if !graph.IsDirected() {
		return nil, fmt.Errorf("topological sort only works on directed graphs")
	}

	visited := make(map[T]bool)
	stack := make([]T, 0)
	recursionStack := make(map[T]bool)

	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			if dfsHasCycle(graph, vertex, visited, recursionStack) {
				return nil, fmt.Errorf("graph has a cycle, cannot perform topological sort")
			}
		}
	}

	visited = make(map[T]bool)
	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			dfsTopologicalSort(graph, vertex, visited, &stack)
		}
	}

	reverseSlice(stack)
	return stack, nil
}

// dfsTopologicalSort performs DFS and builds topological order
func dfsTopologicalSort[T comparable](graph *Graph[T], vertex T, visited map[T]bool, stack *[]T) {
	visited[vertex] = true

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			dfsTopologicalSort(graph, neighbor, visited, stack)
		}
	}

	*stack = append(*stack, vertex)
}

// dfsHasCycle checks if there's a cycle reachable from a vertex (for topological sort)
func dfsHasCycle[T comparable](graph *Graph[T], vertex T, visited, recursionStack map[T]bool) bool {
	visited[vertex] = true
	recursionStack[vertex] = true

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			if dfsHasCycle(graph, neighbor, visited, recursionStack) {
				return true
			}
		} else if recursionStack[neighbor] {
			return true
		}
	}

	recursionStack[vertex] = false
	return false
}

// DetectCycle detects if there's a cycle in the graph using DFS
// Returns true if a cycle is found
// Time complexity: O(V + E)
func DetectCycle[T comparable](graph *Graph[T]) bool {
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

// ConnectedComponents finds all connected components in an undirected graph
// Returns a slice where each element is a component (slice of vertices)
// Time complexity: O(V + E)
func ConnectedComponents[T comparable](graph *Graph[T]) [][]T {
	if graph.IsDirected() {
		return StronglyConnectedComponents(graph)
	}

	visited := make(map[T]bool)
	components := [][]T{}

	for _, vertex := range graph.GetVertices() {
		if !visited[vertex] {
			component := []T{}
			dfsFindComponent(graph, vertex, visited, &component)
			components = append(components, component)
		}
	}

	return components
}

// dfsFindComponent finds all vertices in a connected component
func dfsFindComponent[T comparable](graph *Graph[T], vertex T, visited map[T]bool, component *[]T) {
	visited[vertex] = true
	*component = append(*component, vertex)

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			dfsFindComponent(graph, neighbor, visited, component)
		}
	}
}

// StronglyConnectedComponents finds all strongly connected components in a directed graph
// Uses Kosaraju's algorithm
// Returns a slice where each element is an SCC (slice of vertices)
// Time complexity: O(V + E)
func StronglyConnectedComponents[T comparable](graph *Graph[T]) [][]T {
	result := DFSAllComponents(graph)
	order := make([]T, len(result.Order))
	copy(order, result.Order)
	reverseSlice(order)

	transposed := graph.Reverse()
	visited := make(map[T]bool)
	sccs := [][]T{}

	for _, vertex := range order {
		if !visited[vertex] {
			scc := []T{}
			dfsFindSCC(transposed, vertex, visited, &scc)
			sccs = append(sccs, scc)
		}
	}

	return sccs
}

// dfsFindSCC finds all vertices in a strongly connected component
func dfsFindSCC[T comparable](graph *Graph[T], vertex T, visited map[T]bool, scc *[]T) {
	visited[vertex] = true
	*scc = append(*scc, vertex)

	for _, neighbor := range graph.GetNeighbors(vertex) {
		if !visited[neighbor] {
			dfsFindSCC(graph, neighbor, visited, scc)
		}
	}
}

// FindPath finds any path between two vertices using DFS
// Returns the path as a slice of vertices, or nil if no path exists
// Time complexity: O(V + E)
func FindPath[T comparable](graph *Graph[T], start, end T) ([]T, error) {
	result, err := DFS(graph, start)
	if err != nil {
		return nil, err
	}

	return dfsFindPathTo(result.Parents, start, end), nil
}

// dfsFindPathTo finds a path from start to end using parent map
func dfsFindPathTo[T comparable](parents map[T]T, start, end T) []T {
	if start == end {
		return []T{start}
	}

	if _, exists := parents[end]; !exists {
		return nil
	}

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

// reverseSlice reverses a slice in place
func reverseSlice[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
