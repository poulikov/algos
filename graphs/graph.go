package graphs

import (
	"fmt"
)

// Edge represents a directed or undirected edge in a graph
type Edge[T comparable] struct {
	From   T
	To     T
	Weight float64
}

// GraphType defines whether the graph is directed or undirected
type GraphType int

const (
	// Undirected represents an undirected graph (edges work both ways)
	Undirected GraphType = iota
	// Directed represents a directed graph (edges have direction)
	Directed
)

// Graph represents a graph using adjacency list representation
// Supports both directed and undirected graphs
type Graph[T comparable] struct {
	vertices  map[T]struct{}
	edges     map[T][]Edge[T]
	graphType GraphType
	edgeCount int
}

// NewGraph creates a new graph with specified type
// Time complexity: O(1)
func NewGraph[T comparable](graphType GraphType) *Graph[T] {
	return &Graph[T]{
		vertices:  make(map[T]struct{}),
		edges:     make(map[T][]Edge[T]),
		graphType: graphType,
	}
}

// NewUndirectedGraph creates a new undirected graph
// Time complexity: O(1)
func NewUndirectedGraph[T comparable]() *Graph[T] {
	return NewGraph[T](Undirected)
}

// NewDirectedGraph creates a new directed graph
// Time complexity: O(1)
func NewDirectedGraph[T comparable]() *Graph[T] {
	return NewGraph[T](Directed)
}

// AddVertex adds a vertex to the graph
// Time complexity: O(1)
func (g *Graph[T]) AddVertex(vertex T) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = struct{}{}
		g.edges[vertex] = []Edge[T]{}
	}
}

// AddEdge adds an edge between two vertices
// If graph is undirected, adds edges in both directions
// Time complexity: O(1)
func (g *Graph[T]) AddEdge(from, to T, weight float64) {
	g.AddVertex(from)
	g.AddVertex(to)

	g.edges[from] = append(g.edges[from], Edge[T]{From: from, To: to, Weight: weight})

	if g.graphType == Undirected {
		g.edges[to] = append(g.edges[to], Edge[T]{From: to, To: from, Weight: weight})
	}

	g.edgeCount++
}

// AddEdgeUnweighted adds an unweighted edge (weight = 0)
// Time complexity: O(1)
func (g *Graph[T]) AddEdgeUnweighted(from, to T) {
	g.AddEdge(from, to, 0)
}

// RemoveVertex removes a vertex and all its edges
// Time complexity: O(V + E) where V is vertices, E is edges
func (g *Graph[T]) RemoveVertex(vertex T) bool {
	if _, exists := g.vertices[vertex]; !exists {
		return false
	}

	delete(g.vertices, vertex)

	for v := range g.edges {
		originalCount := len(g.edges[v])
		for i := len(g.edges[v]) - 1; i >= 0; i-- {
			if g.edges[v][i].To == vertex {
				g.edges[v] = append(g.edges[v][:i], g.edges[v][i+1:]...)
			}
		}
		removedCount := originalCount - len(g.edges[v])
		if g.graphType == Undirected && removedCount > 0 {
			g.edgeCount--
		} else if g.graphType == Directed {
			g.edgeCount -= removedCount
		}
	}

	delete(g.edges, vertex)
	return true
}

// RemoveEdge removes an edge between two vertices
// Time complexity: O(deg(v)) where deg(v) is the degree of the vertex
func (g *Graph[T]) RemoveEdge(from, to T) bool {
	edges := g.edges[from]
	for i, edge := range edges {
		if edge.To == to {
			g.edges[from] = append(edges[:i], edges[i+1:]...)
			g.edgeCount--

			if g.graphType == Undirected {
				edgesTo := g.edges[to]
				for j, edgeTo := range edgesTo {
					if edgeTo.To == from {
						g.edges[to] = append(edgesTo[:j], edgesTo[j+1:]...)
						break
					}
				}
			}
			return true
		}
	}
	return false
}

// HasVertex checks if a vertex exists in the graph
// Time complexity: O(1)
func (g *Graph[T]) HasVertex(vertex T) bool {
	_, exists := g.vertices[vertex]
	return exists
}

// HasEdge checks if an edge exists between two vertices
// Time complexity: O(deg(from)) where deg(from) is the degree of the vertex
func (g *Graph[T]) HasEdge(from, to T) bool {
	for _, edge := range g.edges[from] {
		if edge.To == to {
			return true
		}
	}
	return false
}

// GetEdgeWeight returns the weight of the edge between two vertices
// Time complexity: O(deg(from))
func (g *Graph[T]) GetEdgeWeight(from, to T) (float64, error) {
	for _, edge := range g.edges[from] {
		if edge.To == to {
			return edge.Weight, nil
		}
	}
	return 0, fmt.Errorf("edge not found")
}

// GetNeighbors returns all neighbors of a vertex
// Time complexity: O(1)
func (g *Graph[T]) GetNeighbors(vertex T) []T {
	neighbors := make([]T, 0, len(g.edges[vertex]))
	for _, edge := range g.edges[vertex] {
		neighbors = append(neighbors, edge.To)
	}
	return neighbors
}

// GetOutgoingEdges returns all outgoing edges from a vertex
// Time complexity: O(1)
func (g *Graph[T]) GetOutgoingEdges(vertex T) []Edge[T] {
	edges := make([]Edge[T], len(g.edges[vertex]))
	copy(edges, g.edges[vertex])
	return edges
}

// GetIncomingEdges returns all incoming edges to a vertex
// Time complexity: O(V + E)
func (g *Graph[T]) GetIncomingEdges(vertex T) []Edge[T] {
	incoming := []Edge[T]{}
	for _, edges := range g.edges {
		for _, edge := range edges {
			if edge.To == vertex {
				incoming = append(incoming, edge)
			}
		}
	}
	return incoming
}

// GetVertices returns all vertices in the graph
// Time complexity: O(V)
func (g *Graph[T]) GetVertices() []T {
	vertices := make([]T, 0, len(g.vertices))
	for v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}

// GetEdges returns all edges in the graph
// For undirected graphs, returns each edge only once
// Time complexity: O(V + E)
func (g *Graph[T]) GetEdges() []Edge[T] {
	if g.graphType == Directed {
		edges := make([]Edge[T], 0, g.edgeCount)
		for _, vertexEdges := range g.edges {
			edges = append(edges, vertexEdges...)
		}
		return edges
	}

	seen := make(map[string]struct{})
	edges := make([]Edge[T], 0, g.edgeCount/2)

	for _, vertexEdges := range g.edges {
		for _, edge := range vertexEdges {
			key := fmt.Sprintf("%v-%v", edge.From, edge.To)
			key2 := fmt.Sprintf("%v-%v", edge.To, edge.From)
			if _, exists := seen[key]; !exists && edge.From != edge.To {
				seen[key] = struct{}{}
				seen[key2] = struct{}{}
				edges = append(edges, edge)
			}
		}
	}

	return edges
}

// VertexCount returns the number of vertices in the graph
// Time complexity: O(1)
func (g *Graph[T]) VertexCount() int {
	return len(g.vertices)
}

// EdgeCount returns the number of edges in the graph
// Time complexity: O(1)
func (g *Graph[T]) EdgeCount() int {
	return g.edgeCount
}

// Degree returns the degree of a vertex (number of connected edges)
// For directed graphs, returns out-degree
// Time complexity: O(1)
func (g *Graph[T]) Degree(vertex T) int {
	return len(g.edges[vertex])
}

// InDegree returns the in-degree of a vertex (for directed graphs)
// Time complexity: O(V + E)
func (g *Graph[T]) InDegree(vertex T) int {
	count := 0
	for _, edges := range g.edges {
		for _, edge := range edges {
			if edge.To == vertex {
				count++
			}
		}
	}
	return count
}

// OutDegree returns the out-degree of a vertex (for directed graphs)
// Time complexity: O(1)
func (g *Graph[T]) OutDegree(vertex T) int {
	return len(g.edges[vertex])
}

// IsDirected checks if the graph is directed
// Time complexity: O(1)
func (g *Graph[T]) IsDirected() bool {
	return g.graphType == Directed
}

// Clear removes all vertices and edges from the graph
// Time complexity: O(1)
func (g *Graph[T]) Clear() {
	g.vertices = make(map[T]struct{})
	g.edges = make(map[T][]Edge[T])
	g.edgeCount = 0
}

// Copy creates a deep copy of the graph
// Time complexity: O(V + E)
func (g *Graph[T]) Copy() *Graph[T] {
	newGraph := NewGraph[T](g.graphType)

	for v := range g.vertices {
		newGraph.AddVertex(v)
	}

	for from, edges := range g.edges {
		for _, edge := range edges {
			newGraph.edges[from] = append(newGraph.edges[from], edge)
		}
	}

	newGraph.edgeCount = g.edgeCount
	return newGraph
}

// String returns a string representation of the graph
func (g *Graph[T]) String() string {
	if g.IsEmpty() {
		return "{}"
	}

	result := "{\n"
	for v, edges := range g.edges {
		result += fmt.Sprintf("  %v: [", v)
		for i, edge := range edges {
			if i > 0 {
				result += ", "
			}
			result += fmt.Sprintf("%v", edge.To)
			if edge.Weight != 0 {
				result += fmt.Sprintf("(%.2f)", edge.Weight)
			}
		}
		result += "]\n"
	}
	result += "}"
	return result
}

// IsEmpty checks if the graph has no vertices
// Time complexity: O(1)
func (g *Graph[T]) IsEmpty() bool {
	return len(g.vertices) == 0
}

// FromAdjacencyMatrix creates a graph from an adjacency matrix
// matrix[i][j] represents the weight of edge from vertex i to vertex j
// For unweighted graphs, use 0 for no edge and 1 for edge
// Time complexity: O(V^2)
func FromAdjacencyMatrix[T comparable](matrix [][]float64, mapping []T, graphType GraphType) (*Graph[T], error) {
	n := len(matrix)
	if n == 0 {
		return NewGraph[T](graphType), nil
	}

	if len(matrix[0]) != n {
		return nil, fmt.Errorf("matrix must be square")
	}

	if len(mapping) != n {
		return nil, fmt.Errorf("mapping length must match matrix dimensions")
	}

	graph := NewGraph[T](graphType)

	for i := 0; i < n; i++ {
		graph.AddVertex(mapping[i])
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				graph.AddEdge(mapping[i], mapping[j], matrix[i][j])
			}
		}
	}

	return graph, nil
}

// ToAdjacencyMatrix converts the graph to an adjacency matrix
// Returns error if vertices are not convertible to indices
// Time complexity: O(V^2)
func (g *Graph[T]) ToAdjacencyMatrix() ([][]float64, []T, error) {
	vertices := g.GetVertices()
	n := len(vertices)

	if n == 0 {
		return [][]float64{}, []T{}, nil
	}

	vertexToIndex := make(map[T]int)
	for i, v := range vertices {
		vertexToIndex[v] = i
	}

	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	for from, edges := range g.edges {
		for _, edge := range edges {
			matrix[vertexToIndex[from]][vertexToIndex[edge.To]] = edge.Weight
		}
	}

	return matrix, vertices, nil
}

// Reverse creates a new graph with all edges reversed
// For undirected graphs, returns a copy
// Time complexity: O(V + E)
func (g *Graph[T]) Reverse() *Graph[T] {
	if g.graphType == Undirected {
		return g.Copy()
	}

	reversed := NewDirectedGraph[T]()

	for v := range g.vertices {
		reversed.AddVertex(v)
	}

	for _, edges := range g.edges {
		for _, edge := range edges {
			reversed.AddEdge(edge.To, edge.From, edge.Weight)
		}
	}

	return reversed
}
