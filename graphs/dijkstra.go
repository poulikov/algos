package graphs

import (
	"fmt"
	"math"

	"github.com/poulikov/algos/heaps"
)

// DijkstraResult represents the result of Dijkstra's algorithm
type DijkstraResult[T comparable] struct {
	Distances map[T]float64 // Shortest distance from start to each vertex
	Parents   map[T]T       // Parent of each vertex in the shortest path tree
	Start     T             // Starting vertex
}

// Dijkstra finds the shortest paths from a start vertex to all other vertices
// Uses a priority queue (min-heap) for efficient vertex selection
// Works only with non-negative edge weights
// Time complexity: O((V + E) log V) with binary heap
func Dijkstra[T comparable](graph *Graph[T], start T) (*DijkstraResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < 0 {
			return nil, fmt.Errorf("Dijkstra's algorithm does not support negative weights")
		}
	}

	result := &DijkstraResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range graph.GetVertices() {
		result.Distances[vertex] = math.Inf(1)
	}

	result.Distances[start] = 0

	heap := heaps.NewMinHeap(func(a, b VertexDistance[T]) bool {
		return a.Distance < b.Distance
	})
	heap.Insert(VertexDistance[T]{Vertex: start, Distance: 0})

	visited := make(map[T]bool)

	for heap.Size() > 0 {
		current, _ := heap.Extract()
		vertex := current.Vertex

		if visited[vertex] {
			continue
		}

		visited[vertex] = true

		for _, edge := range graph.GetOutgoingEdges(vertex) {
			if visited[edge.To] {
				continue
			}

			newDist := result.Distances[vertex] + edge.Weight
			if newDist < result.Distances[edge.To] {
				result.Distances[edge.To] = newDist
				result.Parents[edge.To] = vertex
				heap.Insert(VertexDistance[T]{Vertex: edge.To, Distance: newDist})
			}
		}
	}

	return result, nil
}

// VertexDistance represents a vertex with its distance (used in priority queue)
type VertexDistance[T comparable] struct {
	Vertex   T
	Distance float64
}

// ShortestPathDijkstra finds the shortest path from start to end using Dijkstra
// Returns the path as a slice of vertices, or nil if no path exists
// Time complexity: O((V + E) log V)
func ShortestPathDijkstra[T comparable](graph *Graph[T], start, end T) ([]T, error) {
	result, err := Dijkstra(graph, start)
	if err != nil {
		return nil, err
	}

	if math.IsInf(result.Distances[end], 1) {
		return nil, fmt.Errorf("no path exists from %v to %v", start, end)
	}

	path := []T{end}
	current := end

	for current != start {
		parent, exists := result.Parents[current]
		if !exists {
			return nil, fmt.Errorf("no path exists from %v to %v", start, end)
		}
		path = append([]T{parent}, path...)
		current = parent
	}

	return path, nil
}

// DijkstraToSpecific finds the shortest path to a specific target vertex
// Early termination when target is reached
// Time complexity: O((V + E) log V) in worst case, often less
func DijkstraToSpecific[T comparable](graph *Graph[T], start, target T) (*DijkstraResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if !graph.HasVertex(target) {
		return nil, fmt.Errorf("target vertex not found")
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < 0 {
			return nil, fmt.Errorf("Dijkstra's algorithm does not support negative weights")
		}
	}

	result := &DijkstraResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range graph.GetVertices() {
		result.Distances[vertex] = math.Inf(1)
	}

	result.Distances[start] = 0

	heap := heaps.NewMinHeap(func(a, b VertexDistance[T]) bool {
		return a.Distance < b.Distance
	})
	heap.Insert(VertexDistance[T]{Vertex: start, Distance: 0})

	visited := make(map[T]bool)

	for heap.Size() > 0 {
		current, _ := heap.Extract()
		vertex := current.Vertex

		if vertex == target {
			break
		}

		if visited[vertex] {
			continue
		}

		visited[vertex] = true

		for _, edge := range graph.GetOutgoingEdges(vertex) {
			if visited[edge.To] {
				continue
			}

			newDist := result.Distances[vertex] + edge.Weight
			if newDist < result.Distances[edge.To] {
				result.Distances[edge.To] = newDist
				result.Parents[edge.To] = vertex
				heap.Insert(VertexDistance[T]{Vertex: edge.To, Distance: newDist})
			}
		}
	}

	return result, nil
}

// DijkstraMultipleSources finds shortest paths from multiple source vertices
// Returns the minimum distance to each vertex from any of the sources
// Time complexity: O(k * (V + E) log V) where k is the number of sources
func DijkstraMultipleSources[T comparable](graph *Graph[T], sources []T) (*DijkstraResult[T], error) {
	if len(sources) == 0 {
		return nil, fmt.Errorf("at least one source vertex is required")
	}

	result := &DijkstraResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
	}

	for _, vertex := range graph.GetVertices() {
		result.Distances[vertex] = math.Inf(1)
	}

	for _, source := range sources {
		if !graph.HasVertex(source) {
			return nil, fmt.Errorf("source vertex %v not found", source)
		}

		sourceResult, err := Dijkstra(graph, source)
		if err != nil {
			return nil, err
		}

		for vertex, dist := range sourceResult.Distances {
			if dist < result.Distances[vertex] {
				result.Distances[vertex] = dist
				result.Parents[vertex] = sourceResult.Parents[vertex]
			}
		}
	}

	return result, nil
}

// DijkstraWithPathLimit finds shortest paths with a limit on path length
// Returns distances where path length is within the limit
// Time complexity: O(L * (V + E) log V) where L is the path limit
func DijkstraWithPathLimit[T comparable](graph *Graph[T], start T, maxSteps int) (*DijkstraResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	result := &DijkstraResult[T]{
		Distances: make(map[T]float64),
		Parents:   make(map[T]T),
		Start:     start,
	}

	for _, vertex := range graph.GetVertices() {
		result.Distances[vertex] = math.Inf(1)
	}

	result.Distances[start] = 0

	stepsRemaining := make(map[T]int)
	for _, vertex := range graph.GetVertices() {
		stepsRemaining[vertex] = maxSteps
	}

	heap := heaps.NewMinHeap(func(a, b VertexDistance[T]) bool {
		return a.Distance < b.Distance
	})
	heap.Insert(VertexDistance[T]{Vertex: start, Distance: 0})

	for heap.Size() > 0 {
		current, _ := heap.Extract()
		vertex := current.Vertex

		if stepsRemaining[vertex] <= 0 {
			continue
		}

		for _, edge := range graph.GetOutgoingEdges(vertex) {
			newDist := result.Distances[vertex] + edge.Weight
			newSteps := stepsRemaining[vertex] - 1

			if newSteps >= 0 && newDist < result.Distances[edge.To] {
				result.Distances[edge.To] = newDist
				result.Parents[edge.To] = vertex
				stepsRemaining[edge.To] = newSteps
				heap.Insert(VertexDistance[T]{Vertex: edge.To, Distance: newDist})
			}
		}
	}

	return result, nil
}
