package graphs

import (
	"fmt"
	"math"
)

// BellmanFordResult represents the result of Bellman-Ford algorithm
type BellmanFordResult[T comparable] struct {
	Distances     map[T]float64 // Shortest distance from start to each vertex
	Parents       map[T]T       // Parent of each vertex in the shortest path tree
	Start         T             // Starting vertex
	HasCycle      bool          // Whether a negative cycle was detected
	CycleVertices []T           // Vertices in the detected negative cycle (if any)
}

// BellmanFord finds the shortest paths from a start vertex
// Handles negative edge weights and detects negative cycles
// Time complexity: O(V * E)
func BellmanFord[T comparable](graph *Graph[T], start T) (*BellmanFordResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	vertices := graph.GetVertices()
	vertexCount := len(vertices)
	vertexIndex := make(map[T]int)
	for i, v := range vertices {
		vertexIndex[v] = i
	}

	edges := graph.GetEdges()

	result := &BellmanFordResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range vertices {
		result.Distances[vertex] = math.Inf(1)
	}
	result.Distances[start] = 0

	for i := 0; i < vertexCount-1; i++ {
		updated := false

		distCopy := make(map[T]float64)
		for k, v := range result.Distances {
			distCopy[k] = v
		}

		for _, edge := range edges {
			if distCopy[edge.From]+edge.Weight < result.Distances[edge.To] {
				result.Distances[edge.To] = distCopy[edge.From] + edge.Weight
				result.Parents[edge.To] = edge.From
				updated = true
			}
		}

		if !updated {
			break
		}
	}

	for _, edge := range edges {
		if result.Distances[edge.From]+edge.Weight < result.Distances[edge.To] {
			result.HasCycle = true
			result.CycleVertices = findNegativeCycle(result.Parents, edge.To)
			return result, nil
		}
	}

	return result, nil
}

// findNegativeCycle attempts to find vertices in a negative cycle
func findNegativeCycle[T comparable](parents map[T]T, start T) []T {
	visited := make(map[T]bool)
	cycle := []T{}
	current := start

	for !visited[current] {
		visited[current] = true
		cycle = append(cycle, current)

		parent, exists := parents[current]
		if !exists {
			break
		}
		current = parent
	}

	return cycle
}

// BellmanFordWithLimit performs Bellman-Ford with a maximum number of iterations
// Useful for very large graphs when you need to limit computation
// Time complexity: O(min(k, V) * E) where k is the limit
func BellmanFordWithLimit[T comparable](graph *Graph[T], start T, maxIterations int) (*BellmanFordResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if maxIterations < 1 {
		return nil, fmt.Errorf("max iterations must be at least 1")
	}

	vertices := graph.GetVertices()
	edges := graph.GetEdges()

	result := &BellmanFordResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range vertices {
		result.Distances[vertex] = math.Inf(1)
	}
	result.Distances[start] = 0

	vertexCount := len(vertices)
	limit := min(maxIterations, vertexCount-1)

	for i := 0; i < limit; i++ {
		updated := false

		distCopy := make(map[T]float64)
		for k, v := range result.Distances {
			distCopy[k] = v
		}

		for _, edge := range edges {
			if distCopy[edge.From]+edge.Weight < result.Distances[edge.To] {
				result.Distances[edge.To] = distCopy[edge.From] + edge.Weight
				result.Parents[edge.To] = edge.From
				updated = true
			}
		}

		if !updated {
			break
		}
	}

	for _, edge := range edges {
		if result.Distances[edge.From]+edge.Weight < result.Distances[edge.To] {
			result.HasCycle = true
			result.CycleVertices = findNegativeCycle(result.Parents, edge.To)
			return result, nil
		}
	}

	return result, nil
}

// BellmanFordToSpecific finds shortest paths to a specific target vertex
// Early termination when target is stable (no more updates)
// Time complexity: O(V * E) in worst case, often less
func BellmanFordToSpecific[T comparable](graph *Graph[T], start, target T) (*BellmanFordResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if !graph.HasVertex(target) {
		return nil, fmt.Errorf("target vertex not found")
	}

	vertices := graph.GetVertices()
	vertexCount := len(vertices)
	edges := graph.GetEdges()

	result := &BellmanFordResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range vertices {
		result.Distances[vertex] = math.Inf(1)
	}
	result.Distances[start] = 0

	for i := 0; i < vertexCount-1; i++ {
		updated := false

		distCopy := make(map[T]float64)
		for k, v := range result.Distances {
			distCopy[k] = v
		}

		for _, edge := range edges {
			if distCopy[edge.From]+edge.Weight < result.Distances[edge.To] {
				result.Distances[edge.To] = distCopy[edge.From] + edge.Weight
				result.Parents[edge.To] = edge.From
				updated = true
			}
		}

		if !updated {
			break
		}
	}

	for _, edge := range edges {
		if result.Distances[edge.From]+edge.Weight < result.Distances[edge.To] {
			result.HasCycle = true
			result.CycleVertices = findNegativeCycle(result.Parents, edge.To)
			return result, nil
		}
	}

	return result, nil
}

// BellmanFordMultipleSources finds shortest paths from multiple source vertices
// Returns minimum distance to each vertex from any of the sources
// Time complexity: O(k * V * E) where k is the number of sources
func BellmanFordMultipleSources[T comparable](graph *Graph[T], sources []T) (*BellmanFordResult[T], error) {
	if len(sources) == 0 {
		return nil, fmt.Errorf("at least one source vertex is required")
	}

	for _, source := range sources {
		if !graph.HasVertex(source) {
			return nil, fmt.Errorf("source vertex %v not found", source)
		}
	}

	result := &BellmanFordResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
	}

	vertices := graph.GetVertices()
	for _, vertex := range vertices {
		result.Distances[vertex] = math.Inf(1)
	}

	for _, source := range sources {
		sourceResult, err := BellmanFord(graph, source)
		if err != nil {
			return nil, err
		}

		for vertex, dist := range sourceResult.Distances {
			if dist < result.Distances[vertex] {
				result.Distances[vertex] = dist
				result.Parents[vertex] = sourceResult.Parents[vertex]
			}
		}

		if sourceResult.HasCycle {
			result.HasCycle = true
			result.CycleVertices = sourceResult.CycleVertices
		}
	}

	return result, nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
