package graphs

import (
	"fmt"
	"sort"

	"github.com/poulikov/algos/trees"
)

// Kruskal finds the Minimum Spanning Tree using Kruskal's algorithm
// Uses Union-Find (Disjoint Set Union) for efficient cycle detection
// Works on undirected graphs
// Time complexity: O(E log E) with Union-Find
func Kruskal[T comparable](graph *Graph[T]) (*MSTResult[T], error) {
	if graph.IsDirected() {
		return nil, fmt.Errorf("Kruskal's algorithm only works on undirected graphs")
	}

	edges := graph.GetEdges()
	vertices := graph.GetVertices()
	vertexCount := len(vertices)

	if vertexCount == 0 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	mst := &MSTResult[T]{
		Edges:     make([]Edge[T], 0, vertexCount-1),
		TotalCost: 0,
	}

	uf := trees.NewUnionFind[T]()

	for _, vertex := range vertices {
		uf.MakeSet(vertex)
	}

	for _, edge := range edges {
		fromRoot, _ := uf.Find(edge.From)
		toRoot, _ := uf.Find(edge.To)

		if fromRoot != toRoot {
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += edge.Weight
			uf.Union(edge.From, edge.To)

			if len(mst.Edges) == vertexCount-1 {
				break
			}
		}
	}

	if len(mst.Edges) != vertexCount-1 && vertexCount > 1 {
		return nil, fmt.Errorf("graph is not connected")
	}

	return mst, nil
}

// KruskalWithLimit finds MST with a maximum number of edges
// Useful when you want to limit the size of the MST
// Time complexity: O(E log E)
func KruskalWithLimit[T comparable](graph *Graph[T], maxEdges int) (*MSTResult[T], error) {
	if graph.IsDirected() {
		return nil, fmt.Errorf("Kruskal's algorithm only works on undirected graphs")
	}

	edges := graph.GetEdges()
	vertices := graph.GetVertices()

	if len(vertices) == 0 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	if maxEdges < 1 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	mst := &MSTResult[T]{
		Edges:     make([]Edge[T], 0, maxEdges),
		TotalCost: 0,
	}

	uf := trees.NewUnionFind[T]()

	for _, vertex := range vertices {
		uf.MakeSet(vertex)
	}

	edgesAdded := 0

	for _, edge := range edges {
		if edgesAdded >= maxEdges {
			break
		}

		fromRoot, _ := uf.Find(edge.From)
		toRoot, _ := uf.Find(edge.To)

		if fromRoot != toRoot {
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += edge.Weight
			uf.Union(edge.From, edge.To)
			edgesAdded++
		}
	}

	return mst, nil
}

// KruskalByWeightLimit finds MST with a maximum total weight
// Stops adding edges when total cost reaches the limit
// Time complexity: O(E log E)
func KruskalByWeightLimit[T comparable](graph *Graph[T], maxWeight float64) (*MSTResult[T], error) {
	if graph.IsDirected() {
		return nil, fmt.Errorf("Kruskal's algorithm only works on undirected graphs")
	}

	edges := graph.GetEdges()
	vertices := graph.GetVertices()

	if len(vertices) == 0 {
		return &MSTResult[T]{
			Edges:     []Edge[T]{},
			TotalCost: 0,
		}, nil
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	mst := &MSTResult[T]{
		Edges:     []Edge[T]{},
		TotalCost: 0,
	}

	uf := trees.NewUnionFind[T]()

	for _, vertex := range vertices {
		uf.MakeSet(vertex)
	}

	for _, edge := range edges {
		if mst.TotalCost+edge.Weight > maxWeight {
			break
		}

		fromRoot, _ := uf.Find(edge.From)
		toRoot, _ := uf.Find(edge.To)

		if fromRoot != toRoot {
			mst.Edges = append(mst.Edges, edge)
			mst.TotalCost += edge.Weight
			uf.Union(edge.From, edge.To)
		}
	}

	return mst, nil
}

// KruskalAllComponents finds MST for each connected component
// Returns a list of MSTResult, one for each component
// Time complexity: O(E log E)
func KruskalAllComponents[T comparable](graph *Graph[T]) ([]*MSTResult[T], error) {
	if graph.IsDirected() {
		return nil, fmt.Errorf("Kruskal's algorithm only works on undirected graphs")
	}

	vertices := graph.GetVertices()
	if len(vertices) == 0 {
		return []*MSTResult[T]{}, nil
	}

	visited := make(map[T]bool)
	results := []*MSTResult[T]{}

	for _, vertex := range vertices {
		if !visited[vertex] {
			componentVertices := getConnectedComponentForKruskal(graph, vertex, visited)

			if len(componentVertices) > 1 {
				componentGraph := createComponentGraph(graph, componentVertices)
				mst, err := Kruskal(componentGraph)
				if err != nil {
					return nil, err
				}
				results = append(results, mst)
			}
		}
	}

	return results, nil
}

// getConnectedComponentForKruskal finds all vertices in a connected component using BFS
func getConnectedComponentForKruskal[T comparable](graph *Graph[T], start T, visited map[T]bool) []T {
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

// createComponentGraph creates a subgraph with only the specified vertices
func createComponentGraph[T comparable](graph *Graph[T], vertices []T) *Graph[T] {
	componentGraph := NewUndirectedGraph[T]()

	for _, vertex := range vertices {
		componentGraph.AddVertex(vertex)
	}

	for _, edge := range graph.GetEdges() {
		if contains(vertices, edge.From) && contains(vertices, edge.To) {
			componentGraph.AddEdge(edge.From, edge.To, edge.Weight)
		}
	}

	return componentGraph
}

// contains checks if a vertex is in a slice of vertices
func contains[T comparable](vertices []T, vertex T) bool {
	for _, v := range vertices {
		if v == vertex {
			return true
		}
	}
	return false
}
