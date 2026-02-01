package graphs

import (
	"testing"
)

func TestNewUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	if g == nil {
		t.Fatal("NewUndirectedGraph returned nil")
	}
	if g.IsDirected() {
		t.Error("Graph should be undirected")
	}
	if g.VertexCount() != 0 {
		t.Error("New graph should have 0 vertices")
	}
}

func TestNewDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	if g == nil {
		t.Fatal("NewDirectedGraph returned nil")
	}
	if !g.IsDirected() {
		t.Error("Graph should be directed")
	}
	if g.VertexCount() != 0 {
		t.Error("New graph should have 0 vertices")
	}
}

func TestAddVertex(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddVertex("A")
	g.AddVertex("B")

	if !g.HasVertex("A") {
		t.Error("Vertex A should exist")
	}
	if !g.HasVertex("B") {
		t.Error("Vertex B should exist")
	}
	if g.HasVertex("C") {
		t.Error("Vertex C should not exist")
	}
	if g.VertexCount() != 2 {
		t.Errorf("Expected 2 vertices, got %d", g.VertexCount())
	}
}

func TestAddVertexDuplicate(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(1)

	if g.VertexCount() != 1 {
		t.Errorf("Should not add duplicate vertex, got %d", g.VertexCount())
	}
}

func TestAddEdge(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 5.0)

	if !g.HasVertex(1) || !g.HasVertex(2) {
		t.Error("AddEdge should create vertices if they don't exist")
	}
	if !g.HasEdge(1, 2) {
		t.Error("Edge 1-2 should exist")
	}
	if !g.HasEdge(2, 1) {
		t.Error("Edge 2-1 should exist (undirected graph)")
	}
	if g.EdgeCount() != 1 {
		t.Errorf("Expected 1 edge (undirected), got %d", g.EdgeCount())
	}
}

func TestAddEdgeDirected(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 5.0)

	if !g.HasEdge(1, 2) {
		t.Error("Edge 1-2 should exist")
	}
	if g.HasEdge(2, 1) {
		t.Error("Edge 2-1 should not exist in directed graph")
	}
	if g.EdgeCount() != 1 {
		t.Errorf("Expected 1 edge (directed), got %d", g.EdgeCount())
	}
}

func TestAddEdgeWeighted(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 10.5)

	weight, err := g.GetEdgeWeight(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if weight != 10.5 {
		t.Errorf("Expected weight 10.5, got %f", weight)
	}
}

func TestAddEdgeUnweighted(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdgeUnweighted(1, 2)

	weight, err := g.GetEdgeWeight(1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if weight != 0 {
		t.Errorf("Expected weight 0, got %f", weight)
	}
}

func TestRemoveVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 5.0)
	g.AddEdge(2, 3, 3.0)
	g.AddEdge(1, 4, 2.0)

	removed := g.RemoveVertex(2)
	if !removed {
		t.Error("RemoveVertex should return true")
	}

	if g.HasVertex(2) {
		t.Error("Vertex 2 should be removed")
	}

	if g.HasEdge(1, 2) || g.HasEdge(2, 1) {
		t.Error("Edge 1-2 should be removed")
	}

	if g.HasEdge(2, 3) || g.HasEdge(3, 2) {
		t.Error("Edge 2-3 should be removed")
	}

	if !g.HasEdge(1, 4) && !g.HasEdge(4, 1) {
		t.Error("Edge 1-4 should still exist")
	}

	if g.EdgeCount() != 1 {
		t.Errorf("Expected 1 edge, got %d", g.EdgeCount())
	}
}

func TestRemoveVertexNonExistent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	removed := g.RemoveVertex(999)
	if removed {
		t.Error("RemoveVertex should return false for non-existent vertex")
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 5.0)

	removed := g.RemoveEdge(1, 2)
	if !removed {
		t.Error("RemoveEdge should return true")
	}
	if g.HasEdge(1, 2) || g.HasEdge(2, 1) {
		t.Error("Edge should be removed")
	}
	if g.EdgeCount() != 0 {
		t.Errorf("Expected 0 edges, got %d", g.EdgeCount())
	}
}

func TestRemoveEdgeDirected(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 5.0)

	removed := g.RemoveEdge(1, 2)
	if !removed {
		t.Error("RemoveEdge should return true")
	}
	if g.HasEdge(1, 2) {
		t.Error("Edge should be removed")
	}
}

func TestRemoveEdgeNonExistent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	removed := g.RemoveEdge(1, 2)
	if removed {
		t.Error("RemoveEdge should return false for non-existent edge")
	}
}

func TestGetNeighbors(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 1.0)
	g.AddEdge(1, 4, 1.0)

	neighbors := g.GetNeighbors(1)
	if len(neighbors) != 3 {
		t.Errorf("Expected 3 neighbors, got %d", len(neighbors))
	}
}

func TestGetOutgoingEdges(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 2.0)

	edges := g.GetOutgoingEdges(1)
	if len(edges) != 2 {
		t.Errorf("Expected 2 outgoing edges, got %d", len(edges))
	}
	if edges[0].To != 2 && edges[0].To != 3 {
		t.Error("Outgoing edges should point to 2 or 3")
	}
}

func TestGetIncomingEdges(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(3, 2, 2.0)

	edges := g.GetIncomingEdges(2)
	if len(edges) != 2 {
		t.Errorf("Expected 2 incoming edges, got %d", len(edges))
	}
}

func TestDegree(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 1.0)

	if g.Degree(1) != 2 {
		t.Errorf("Expected degree 2, got %d", g.Degree(1))
	}
	if g.Degree(2) != 1 {
		t.Errorf("Expected degree 1, got %d", g.Degree(2))
	}
}

func TestInDegreeOutDegree(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(3, 2, 1.0)
	g.AddEdge(2, 4, 1.0)

	if g.InDegree(2) != 2 {
		t.Errorf("Expected in-degree 2, got %d", g.InDegree(2))
	}
	if g.OutDegree(2) != 1 {
		t.Errorf("Expected out-degree 1, got %d", g.OutDegree(2))
	}
}

func TestClear(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)

	g.Clear()
	if !g.IsEmpty() {
		t.Error("Graph should be empty after Clear")
	}
	if g.VertexCount() != 0 || g.EdgeCount() != 0 {
		t.Error("Vertex and edge count should be 0")
	}
}

func TestCopy(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)

	copied := g.Copy()
	if copied == g {
		t.Error("Copy should return a new instance")
	}
	if copied.VertexCount() != g.VertexCount() {
		t.Error("Copied graph should have same number of vertices")
	}
	if copied.EdgeCount() != g.EdgeCount() {
		t.Error("Copied graph should have same number of edges")
	}
}

func TestReverse(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)

	reversed := g.Reverse()
	if !reversed.HasEdge(2, 1) || !reversed.HasEdge(3, 2) {
		t.Error("Reverse should flip edge directions")
	}
	if reversed.HasEdge(1, 2) || reversed.HasEdge(2, 3) {
		t.Error("Original edges should not exist in reversed graph")
	}
}

func TestReverseUndirected(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)

	reversed := g.Reverse()
	if !reversed.HasEdge(1, 2) || !reversed.HasEdge(2, 1) {
		t.Error("Reverse of undirected graph should be identical")
	}
}

func TestFromAdjacencyMatrix(t *testing.T) {
	matrix := [][]float64{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}
	mapping := []int{1, 2, 3}

	g, err := FromAdjacencyMatrix(matrix, mapping, Undirected)
	if err != nil {
		t.Fatal(err)
	}
	if g.VertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", g.VertexCount())
	}
	if !g.HasEdge(1, 2) || !g.HasEdge(2, 3) {
		t.Error("Matrix edges should be correctly converted")
	}
}

func TestFromAdjacencyMatrixErrors(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]float64
		mapping []int
	}{
		{"non-square", [][]float64{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2}},
		{"mismatch length", [][]float64{{1, 2}, {3, 4}}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FromAdjacencyMatrix(tt.matrix, tt.mapping, Undirected)
			if err == nil {
				t.Error("Expected error for invalid matrix")
			}
		})
	}
}

func TestToAdjacencyMatrix(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)

	matrix, vertices, err := g.ToAdjacencyMatrix()
	if err != nil {
		t.Fatal(err)
	}
	if len(matrix) != 3 || len(matrix[0]) != 3 {
		t.Error("Matrix should be 3x3")
	}
	if len(vertices) != 3 {
		t.Error("Should have 3 vertices")
	}
}

func TestIsEmpty(t *testing.T) {
	g := NewUndirectedGraph[int]()
	if !g.IsEmpty() {
		t.Error("New graph should be empty")
	}

	g.AddVertex(1)
	if g.IsEmpty() {
		t.Error("Graph with vertex should not be empty")
	}
}

func TestGetVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	vertices := g.GetVertices()
	if len(vertices) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(vertices))
	}
}

func TestGetEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 2.0)

	edges := g.GetEdges()
	if len(edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(edges))
	}
}

func TestGetEdgesDirected(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 1, 1.0)

	edges := g.GetEdges()
	if len(edges) != 2 {
		t.Errorf("Expected 2 edges in directed graph, got %d", len(edges))
	}
}

func TestGraphString(t *testing.T) {
	g := NewUndirectedGraph[int]()
	str := g.String()
	if str != "{}" {
		t.Errorf("Empty graph string should be '{}', got %s", str)
	}

	g.AddEdge(1, 2, 1.0)
	str = g.String()
	if str == "{}" {
		t.Error("Non-empty graph string should not be '{}'")
	}
}

func TestEdgeSelfLoop(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 1, 1.0)

	if !g.HasEdge(1, 1) {
		t.Error("Self-loop should be allowed")
	}
}

func TestMultipleEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 2, 2.0)

	edges := g.GetOutgoingEdges(1)
	if len(edges) != 2 {
		t.Errorf("Multiple edges should be allowed, got %d", len(edges))
	}
}
