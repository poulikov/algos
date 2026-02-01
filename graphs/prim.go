package graphs

import (
	"fmt"
	"math"

	"github.com/poulikov/algos/heaps"
)

// MSTResult represents the result of a Minimum Spanning Tree algorithm
type MSTResult[T comparable] struct {
	Edges     []Edge[T] // Edges in the MST
	TotalCost float64   // Total weight of the MST
}

// Prim finds the Minimum Spanning Tree using Prim's algorithm
// Uses a priority queue (min-heap) for efficient edge selection
// Works on undirected, connected graphs
// Time complexity: O((V + E) log V)
func Prim[T comparable](graph *Graph[T], start T) (*MSTResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if graph.IsDirected() {
		return nil, fmt.Errorf("Prim's algorithm only works on undirected graphs")
	}

	vertices := graph.GetVertices()
	vertexCount := len(vertices)

	if vertexCount == 0 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	mst := &MSTResult[T]{
		Edges:     make([]Edge[T], 0, vertexCount-1),
		TotalCost: 0,
	}

	inMST := make(map[T]bool)
	minEdge := make(map[T]Edge[T])

	for _, vertex := range vertices {
		minEdge[vertex] = Edge[T]{Weight: math.Inf(1)}
	}

	minEdge[start] = Edge[T]{Weight: 0}

	for i := 0; i < vertexCount; i++ {
		u, minWeight := getVertexWithMinEdgeNotInMST(vertices, minEdge, inMST)

		if math.IsInf(minWeight, 1) {
			return nil, fmt.Errorf("graph is not connected")
		}

		inMST[u] = true

		if i > 0 {
			mst.Edges = append(mst.Edges, minEdge[u])
			mst.TotalCost += minEdge[u].Weight
		}

		for _, edge := range graph.GetOutgoingEdges(u) {
			if !inMST[edge.To] && edge.Weight < minEdge[edge.To].Weight {
				minEdge[edge.To] = edge
			}
		}
	}

	if len(mst.Edges) != vertexCount-1 {
		return nil, fmt.Errorf("graph is not connected")
	}

	return mst, nil
}

// getVertexWithMinEdgeNotInMST finds the vertex with minimum edge weight not yet in MST
func getVertexWithMinEdgeNotInMST[T comparable](vertices []T, minEdge map[T]Edge[T], inMST map[T]bool) (T, float64) {
	minWeight := math.Inf(1)
	var minVertex T
	found := false

	for _, vertex := range vertices {
		if !inMST[vertex] && minEdge[vertex].Weight < minWeight {
			minWeight = minEdge[vertex].Weight
			minVertex = vertex
			found = true
		}
	}

	if !found {
		return minVertex, minWeight
	}

	return minVertex, minWeight
}

// PrimWithHeap finds the MST using Prim's algorithm with a heap
// More efficient implementation using priority queue
// Time complexity: O((V + E) log V)
func PrimWithHeap[T comparable](graph *Graph[T], start T) (*MSTResult[T], error) {
	if !graph.HasVertex(start) {
		return nil, fmt.Errorf("start vertex not found")
	}

	if graph.IsDirected() {
		return nil, fmt.Errorf("Prim's algorithm only works on undirected graphs")
	}

	vertices := graph.GetVertices()
	vertexCount := len(vertices)

	if vertexCount == 0 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	mst := &MSTResult[T]{
		Edges:     make([]Edge[T], 0, vertexCount-1),
		TotalCost: 0,
	}

	inMST := make(map[T]bool)
	key := make(map[T]float64)

	for _, vertex := range vertices {
		key[vertex] = math.Inf(1)
	}

	key[start] = 0

	heap := heaps.NewMinHeap(func(a, b VertexKey[T]) bool {
		return a.Key < b.Key
	})

	heap.Insert(VertexKey[T]{Vertex: start, Key: 0})

	for heap.Size() > 0 {
		current, _ := heap.Extract()
		u := current.Vertex

		if inMST[u] {
			continue
		}

		inMST[u] = true

		for _, edge := range graph.GetOutgoingEdges(u) {
			if !inMST[edge.To] && edge.Weight < key[edge.To] {
				key[edge.To] = edge.Weight
				heap.Insert(VertexKey[T]{Vertex: edge.To, Key: edge.Weight})
			}
		}
	}

	for u := range key {
		if u != start {
			for _, edge := range graph.GetOutgoingEdges(u) {
				if edge.Weight == key[u] && inMST[edge.To] {
					mst.Edges = append(mst.Edges, edge)
					mst.TotalCost += edge.Weight
					break
				}
			}
		}
	}

	if len(mst.Edges) != vertexCount-1 {
		return nil, fmt.Errorf("graph is not connected")
	}

	return mst, nil
}

// VertexKey represents a vertex with its key (used in Prim's priority queue)
type VertexKey[T comparable] struct {
	Vertex T
	Key    float64
}

// PrimAll finds MST for each connected component
// Returns a list of MSTResult, one for each component
// Time complexity: O(V * ((V + E) log V))
func PrimAll[T comparable](graph *Graph[T]) ([]*MSTResult[T], error) {
	if graph.IsDirected() {
		return nil, fmt.Errorf("Prim's algorithm only works on undirected graphs")
	}

	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return []*MSTResult[T]{}, nil
	}

	visited := make(map[T]bool)
	results := []*MSTResult[T]{}

	for _, vertex := range vertices {
		if !visited[vertex] {
			componentVertices := getConnectedComponent(graph, vertex, visited)
			if len(componentVertices) > 1 {
				mst, err := Prim(graph, vertex)
				if err != nil {
					return nil, err
				}
				results = append(results, mst)
			}
		}
	}

	return results, nil
}

// getConnectedComponent finds all vertices in a connected component using BFS
func getConnectedComponent[T comparable](graph *Graph[T], start T, visited map[T]bool) []T {
	component := []T{}
	queue := []T{start}
	visited[start] = true

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		component = append(component, vertex)

		for _, neighbor := range graph.GetNeighbors(vertex) {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return component
}
