package graphs

import (
	"fmt"
)

// DAG represents a Directed Acyclic Graph
// It ensures that no cycles can be created during edge additions
type DAG[T comparable] struct {
	graph *Graph[T]
}

// NewDAG creates a new empty Directed Acyclic Graph
// Time complexity: O(1)
func NewDAG[T comparable]() *DAG[T] {
	return &DAG[T]{
		graph: NewDirectedGraph[T](),
	}
}

// AddVertex adds a vertex to the DAG
// Time complexity: O(1)
func (dag *DAG[T]) AddVertex(vertex T) {
	dag.graph.AddVertex(vertex)
}

// AddEdge adds an edge between two vertices
// Returns an error if adding the edge would create a cycle
// Time complexity: O(V + E)
func (dag *DAG[T]) AddEdge(from, to T, weight float64) error {
	if dag.graph.HasEdge(from, to) {
		return fmt.Errorf("edge already exists")
	}

	if dag.hasCycleFrom(to, from) {
		return fmt.Errorf("adding edge would create a cycle")
	}

	dag.graph.AddEdge(from, to, weight)
	return nil
}

// AddEdgeUnweighted adds an unweighted edge (weight = 0)
// Returns an error if adding the edge would create a cycle
// Time complexity: O(V + E)
func (dag *DAG[T]) AddEdgeUnweighted(from, to T) error {
	return dag.AddEdge(from, to, 0)
}

// hasCycleFrom checks if adding an edge from 'to' to 'from' would create a cycle
// Uses DFS to check if there's a path from 'to' to 'from'
// Time complexity: O(V + E)
func (dag *DAG[T]) hasCycleFrom(from, to T) bool {
	visited := make(map[T]bool)
	return dag.hasPath(from, to, visited)
}

// hasPath checks if there's a path from 'from' to 'to' using DFS
func (dag *DAG[T]) hasPath(from, to T, visited map[T]bool) bool {
	if from == to {
		return true
	}

	if visited[from] {
		return false
	}

	visited[from] = true

	for _, neighbor := range dag.graph.GetNeighbors(from) {
		if dag.hasPath(neighbor, to, visited) {
			return true
		}
	}

	return false
}

// RemoveVertex removes a vertex and all its edges
// Time complexity: O(V + E)
func (dag *DAG[T]) RemoveVertex(vertex T) bool {
	return dag.graph.RemoveVertex(vertex)
}

// RemoveEdge removes an edge between two vertices
// Time complexity: O(deg(v))
func (dag *DAG[T]) RemoveEdge(from, to T) bool {
	return dag.graph.RemoveEdge(from, to)
}

// HasVertex checks if a vertex exists in the DAG
// Time complexity: O(1)
func (dag *DAG[T]) HasVertex(vertex T) bool {
	return dag.graph.HasVertex(vertex)
}

// HasEdge checks if an edge exists between two vertices
// Time complexity: O(deg(from))
func (dag *DAG[T]) HasEdge(from, to T) bool {
	return dag.graph.HasEdge(from, to)
}

// GetEdgeWeight returns the weight of the edge between two vertices
// Time complexity: O(deg(from))
func (dag *DAG[T]) GetEdgeWeight(from, to T) (float64, error) {
	return dag.graph.GetEdgeWeight(from, to)
}

// GetNeighbors returns all outgoing neighbors of a vertex
// Time complexity: O(1)
func (dag *DAG[T]) GetNeighbors(vertex T) []T {
	return dag.graph.GetNeighbors(vertex)
}

// GetIncomingNeighbors returns all incoming neighbors of a vertex
// Time complexity: O(V + E)
func (dag *DAG[T]) GetIncomingNeighbors(vertex T) []T {
	neighbors := []T{}
	for v := range dag.graph.vertices {
		if dag.graph.HasEdge(v, vertex) {
			neighbors = append(neighbors, v)
		}
	}
	return neighbors
}

// GetOutgoingEdges returns all outgoing edges from a vertex
// Time complexity: O(1)
func (dag *DAG[T]) GetOutgoingEdges(vertex T) []Edge[T] {
	return dag.graph.GetOutgoingEdges(vertex)
}

// GetIncomingEdges returns all incoming edges to a vertex
// Time complexity: O(V + E)
func (dag *DAG[T]) GetIncomingEdges(vertex T) []Edge[T] {
	return dag.graph.GetIncomingEdges(vertex)
}

// GetVertices returns all vertices in the DAG
// Time complexity: O(V)
func (dag *DAG[T]) GetVertices() []T {
	return dag.graph.GetVertices()
}

// GetEdges returns all edges in the DAG
// Time complexity: O(V + E)
func (dag *DAG[T]) GetEdges() []Edge[T] {
	return dag.graph.GetEdges()
}

// VertexCount returns the number of vertices in the DAG
// Time complexity: O(1)
func (dag *DAG[T]) VertexCount() int {
	return dag.graph.VertexCount()
}

// EdgeCount returns the number of edges in the DAG
// Time complexity: O(1)
func (dag *DAG[T]) EdgeCount() int {
	return dag.graph.EdgeCount()
}

// InDegree returns the in-degree of a vertex
// Time complexity: O(V + E)
func (dag *DAG[T]) InDegree(vertex T) int {
	return dag.graph.InDegree(vertex)
}

// OutDegree returns the out-degree of a vertex
// Time complexity: O(1)
func (dag *DAG[T]) OutDegree(vertex T) int {
	return dag.graph.OutDegree(vertex)
}

// Clear removes all vertices and edges from the DAG
// Time complexity: O(1)
func (dag *DAG[T]) Clear() {
	dag.graph.Clear()
}

// Copy creates a deep copy of the DAG
// Time complexity: O(V + E)
func (dag *DAG[T]) Copy() *DAG[T] {
	return &DAG[T]{
		graph: dag.graph.Copy(),
	}
}

// String returns a string representation of the DAG
func (dag *DAG[T]) String() string {
	return fmt.Sprintf("DAG{\n%s}", dag.graph.String())
}

// IsEmpty checks if the DAG has no vertices
// Time complexity: O(1)
func (dag *DAG[T]) IsEmpty() bool {
	return dag.graph.IsEmpty()
}

// IsDAG validates that the graph is indeed acyclic
// Time complexity: O(V + E)
func (dag *DAG[T]) IsDAG() bool {
	visited := make(map[T]bool)
	recursionStack := make(map[T]bool)

	for _, vertex := range dag.GetVertices() {
		if !visited[vertex] {
			if dag.isCyclic(vertex, visited, recursionStack) {
				return false
			}
		}
	}

	return true
}

// isCyclic checks if there's a cycle reachable from a vertex
func (dag *DAG[T]) isCyclic(vertex T, visited, recursionStack map[T]bool) bool {
	visited[vertex] = true
	recursionStack[vertex] = true

	for _, neighbor := range dag.GetNeighbors(vertex) {
		if !visited[neighbor] {
			if dag.isCyclic(neighbor, visited, recursionStack) {
				return true
			}
		} else if recursionStack[neighbor] {
			return true
		}
	}

	recursionStack[vertex] = false
	return false
}

// Transpose returns the transpose of the DAG (all edges reversed)
// Time complexity: O(V + E)
func (dag *DAG[T]) Transpose() *DAG[T] {
	return &DAG[T]{
		graph: dag.graph.Reverse(),
	}
}

// ToGraph returns the underlying Graph structure
// Use with caution as modifications could break the DAG property
// Time complexity: O(1)
func (dag *DAG[T]) ToGraph() *Graph[T] {
	return dag.graph
}
