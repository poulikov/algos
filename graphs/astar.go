package graphs

import (
	"fmt"
	"math"

	"github.com/poulikov/algos/heaps"
)

// AStarResult represents the result of A* search
type AStarResult[T comparable] struct {
	Path       []T     // The found path
	Cost       float64 // Total cost of the path
	Visited    []T     // Order of visited vertices
	Iterations int     // Number of iterations/vertices processed
}

// AStar performs A* search to find the shortest path from start to end
// Uses a heuristic function to guide the search
// Time complexity: O(b^d) where b is branching factor and d is depth of solution
func AStar[T comparable](graph *Graph[T], start, end T, heuristic func(T) float64) (*AStarResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if !graph.HasVertex(end) {
		return nil, fmt.Errorf("end vertex not found")
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < 0 {
			return nil, fmt.Errorf("A* does not support negative edge weights")
		}
	}

	openSet := heaps.NewMinHeap(func(a, b VertexScore[T]) bool {
		return a.FScore < b.FScore
	})

	startScore := VertexScore[T]{
		Vertex: start,
		GScore: 0,
		FScore: heuristic(start),
	}
	openSet.Insert(startScore)

	gScore := make(map[T]float64)
	fScore := make(map[T]float64)
	parents := make(map[T]T)

	for _, vertex := range graph.GetVertices() {
		gScore[vertex] = math.Inf(1)
		fScore[vertex] = math.Inf(1)
	}

	gScore[start] = 0
	fScore[start] = heuristic(start)

	visited := make(map[T]bool)
	visitedOrder := []T{}
	iterations := 0

	for openSet.Size() > 0 {
		current, _ := openSet.Extract()
		currentVertex := current.Vertex

		iterations++
		visitedOrder = append(visitedOrder, currentVertex)

		if currentVertex == end {
			path := []T{end}
			node := end
			for node != start {
				parent, exists := parents[node]
				if !exists {
					break
				}
				path = append([]T{parent}, path...)
				node = parent
			}

			return &AStarResult[T]{
				Path:       path,
				Cost:       gScore[end],
				Visited:    visitedOrder,
				Iterations: iterations,
			}, nil
		}

		visited[currentVertex] = true

		for _, edge := range graph.GetOutgoingEdges(currentVertex) {
			if visited[edge.To] {
				continue
			}

			tentativeG := gScore[currentVertex] + edge.Weight

			if tentativeG < gScore[edge.To] {
				parents[edge.To] = currentVertex
				gScore[edge.To] = tentativeG
				fScore[edge.To] = tentativeG + heuristic(edge.To)

				neighborScore := VertexScore[T]{
					Vertex: edge.To,
					GScore: tentativeG,
					FScore: fScore[edge.To],
				}

				openSet.Insert(neighborScore)
			}
		}
	}

	return nil, fmt.Errorf("no path found from %v to %v", start, end)
}

// VertexScore represents a vertex with its g-score and f-score (used in A* priority queue)
type VertexScore[T comparable] struct {
	Vertex T
	GScore float64
	FScore float64
}

// AStarWithLimit performs A* search with a maximum number of iterations
// Returns the best path found within the iteration limit
func AStarWithLimit[T comparable](graph *Graph[T], start, end T, heuristic func(T) float64, maxIterations int) (*AStarResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if !graph.HasVertex(end) {
		return nil, fmt.Errorf("end vertex not found")
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < 0 {
			return nil, fmt.Errorf("A* does not support negative edge weights")
		}
	}

	openSet := heaps.NewMinHeap(func(a, b VertexScore[T]) bool {
		return a.FScore < b.FScore
	})

	startScore := VertexScore[T]{
		Vertex: start,
		GScore: 0,
		FScore: heuristic(start),
	}
	openSet.Insert(startScore)

	gScore := make(map[T]float64)
	fScore := make(map[T]float64)
	parents := make(map[T]T)

	for _, vertex := range graph.GetVertices() {
		gScore[vertex] = math.Inf(1)
		fScore[vertex] = math.Inf(1)
	}

	gScore[start] = 0
	fScore[start] = heuristic(start)

	visited := make(map[T]bool)
	visitedOrder := []T{}
	iterations := 0

	for openSet.Size() > 0 && iterations < maxIterations {
		current, _ := openSet.Extract()
		currentVertex := current.Vertex

		iterations++
		visitedOrder = append(visitedOrder, currentVertex)

		if currentVertex == end {
			path := []T{end}
			node := end
			for node != start {
				parent, exists := parents[node]
				if !exists {
					break
				}
				path = append([]T{parent}, path...)
				node = parent
			}

			return &AStarResult[T]{
				Path:       path,
				Cost:       gScore[end],
				Visited:    visitedOrder,
				Iterations: iterations,
			}, nil
		}

		visited[currentVertex] = true

		for _, edge := range graph.GetOutgoingEdges(currentVertex) {
			if visited[edge.To] {
				continue
			}

			tentativeG := gScore[currentVertex] + edge.Weight

			if tentativeG < gScore[edge.To] {
				parents[edge.To] = currentVertex
				gScore[edge.To] = tentativeG
				fScore[edge.To] = tentativeG + heuristic(edge.To)

				neighborScore := VertexScore[T]{
					Vertex: edge.To,
					GScore: tentativeG,
					FScore: fScore[edge.To],
				}

				openSet.Insert(neighborScore)
			}
		}
	}

	if iterations >= maxIterations {
		return nil, fmt.Errorf("maximum iterations (%d) reached without finding a path", maxIterations)
	}

	return nil, fmt.Errorf("no path found from %v to %v", start, end)
}

// AStarMultipleTargets performs A* search to find the path to any of the target vertices
// Returns the path to the first target found
func AStarMultipleTargets[T comparable](graph *Graph[T], start T, targets []T, heuristic func(T) float64) (*AStarResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if len(targets) == 0 {
		return nil, fmt.Errorf("at least one target is required")
	}

	targetSet := make(map[T]bool)
	for _, target := range targets {
		if !graph.HasVertex(target) {
			return nil, fmt.Errorf("target vertex %v not found", target)
		}
		targetSet[target] = true
	}

	openSet := heaps.NewMinHeap(func(a, b VertexScore[T]) bool {
		return a.FScore < b.FScore
	})

	startScore := VertexScore[T]{
		Vertex: start,
		GScore: 0,
		FScore: heuristic(start),
	}
	openSet.Insert(startScore)

	gScore := make(map[T]float64)
	fScore := make(map[T]float64)
	parents := make(map[T]T)

	for _, vertex := range graph.GetVertices() {
		gScore[vertex] = math.Inf(1)
		fScore[vertex] = math.Inf(1)
	}

	gScore[start] = 0
	fScore[start] = heuristic(start)

	visited := make(map[T]bool)
	visitedOrder := []T{}
	iterations := 0

	for openSet.Size() > 0 {
		current, _ := openSet.Extract()
		currentVertex := current.Vertex

		iterations++
		visitedOrder = append(visitedOrder, currentVertex)

		if targetSet[currentVertex] {
			path := []T{currentVertex}
			node := currentVertex
			for node != start {
				parent, exists := parents[node]
				if !exists {
					break
				}
				path = append([]T{parent}, path...)
				node = parent
			}

			return &AStarResult[T]{
				Path:       path,
				Cost:       gScore[currentVertex],
				Visited:    visitedOrder,
				Iterations: iterations,
			}, nil
		}

		visited[currentVertex] = true

		for _, edge := range graph.GetOutgoingEdges(currentVertex) {
			if visited[edge.To] {
				continue
			}

			tentativeG := gScore[currentVertex] + edge.Weight

			if tentativeG < gScore[edge.To] {
				parents[edge.To] = currentVertex
				gScore[edge.To] = tentativeG
				fScore[edge.To] = tentativeG + heuristic(edge.To)

				neighborScore := VertexScore[T]{
					Vertex: edge.To,
					GScore: tentativeG,
					FScore: fScore[edge.To],
				}

				openSet.Insert(neighborScore)
			}
		}
	}

	return nil, fmt.Errorf("no path found from %v to any target", start)
}

// AStarWithReconstruction performs A* search and allows custom path reconstruction
// Useful when you need to track additional information during the search
func AStarWithReconstruction[T comparable](graph *Graph[T], start, end T, heuristic func(T) float64, reconstruct func(map[T]T, T, T) []T) (*AStarResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if !graph.HasVertex(end) {
		return nil, fmt.Errorf("end vertex not found")
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < 0 {
			return nil, fmt.Errorf("A* does not support negative edge weights")
		}
	}

	openSet := heaps.NewMinHeap(func(a, b VertexScore[T]) bool {
		return a.FScore < b.FScore
	})

	startScore := VertexScore[T]{
		Vertex: start,
		GScore: 0,
		FScore: heuristic(start),
	}
	openSet.Insert(startScore)

	gScore := make(map[T]float64)
	fScore := make(map[T]float64)
	parents := make(map[T]T)

	for _, vertex := range graph.GetVertices() {
		gScore[vertex] = math.Inf(1)
		fScore[vertex] = math.Inf(1)
	}

	gScore[start] = 0
	fScore[start] = heuristic(start)

	visited := make(map[T]bool)
	visitedOrder := []T{}
	iterations := 0

	for openSet.Size() > 0 {
		current, _ := openSet.Extract()
		currentVertex := current.Vertex

		iterations++
		visitedOrder = append(visitedOrder, currentVertex)

		if currentVertex == end {
			path := reconstruct(parents, start, end)

			return &AStarResult[T]{
				Path:       path,
				Cost:       gScore[end],
				Visited:    visitedOrder,
				Iterations: iterations,
			}, nil
		}

		visited[currentVertex] = true

		for _, edge := range graph.GetOutgoingEdges(currentVertex) {
			if visited[edge.To] {
				continue
			}

			tentativeG := gScore[currentVertex] + edge.Weight

			if tentativeG < gScore[edge.To] {
				parents[edge.To] = currentVertex
				gScore[edge.To] = tentativeG
				fScore[edge.To] = tentativeG + heuristic(edge.To)

				neighborScore := VertexScore[T]{
					Vertex: edge.To,
					GScore: tentativeG,
					FScore: fScore[edge.To],
				}

				openSet.Insert(neighborScore)
			}
		}
	}

	return nil, fmt.Errorf("no path found from %v to %v", start, end)
}
