package graphs

import (
	"testing"
)

func TestNewDAG(t *testing.T) {
	dag := NewDAG[int]()

	if dag == nil {
		t.Error("NewDAG should return a non-nil DAG")
	}

	if !dag.IsEmpty() {
		t.Error("New DAG should be empty")
	}

	if dag.VertexCount() != 0 {
		t.Errorf("New DAG should have 0 vertices, got %d", dag.VertexCount())
	}

	if dag.EdgeCount() != 0 {
		t.Errorf("New DAG should have 0 edges, got %d", dag.EdgeCount())
	}
}

func TestDAGAddVertex(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddVertex(1)
	dag.AddVertex(2)

	if !dag.HasVertex(1) {
		t.Error("Vertex 1 should exist")
	}

	if !dag.HasVertex(2) {
		t.Error("Vertex 2 should exist")
	}

	if dag.VertexCount() != 2 {
		t.Errorf("Expected 2 vertices, got %d", dag.VertexCount())
	}
}

func TestDAGAddVertexDuplicate(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddVertex(1)
	dag.AddVertex(1)

	if dag.VertexCount() != 1 {
		t.Errorf("Duplicate vertex should not be added, got %d", dag.VertexCount())
	}
}

func TestDAGAddEdge(t *testing.T) {
	dag := NewDAG[int]()
	err := dag.AddEdge(1, 2, 5.0)

	if err != nil {
		t.Fatalf("AddEdge should not return error: %v", err)
	}

	if !dag.HasEdge(1, 2) {
		t.Error("Edge (1, 2) should exist")
	}

	if dag.HasEdge(2, 1) {
		t.Error("Edge (2, 1) should not exist in directed graph")
	}

	weight, err := dag.GetEdgeWeight(1, 2)
	if err != nil {
		t.Fatalf("GetEdgeWeight should not return error: %v", err)
	}

	if weight != 5.0 {
		t.Errorf("Expected weight 5.0, got %f", weight)
	}
}

func TestDAGAddEdgeUnweighted(t *testing.T) {
	dag := NewDAG[int]()
	err := dag.AddEdgeUnweighted(1, 2)

	if err != nil {
		t.Fatalf("AddEdgeUnweighted should not return error: %v", err)
	}

	weight, err := dag.GetEdgeWeight(1, 2)
	if err != nil {
		t.Fatalf("GetEdgeWeight should not return error: %v", err)
	}

	if weight != 0 {
		t.Errorf("Expected weight 0 for unweighted edge, got %f", weight)
	}
}

func TestDAGAddEdgeDuplicate(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 1.0)
	err := dag.AddEdge(1, 2, 2.0)

	if err == nil {
		t.Error("Adding duplicate edge should return error")
	}
}

func TestDAGAddEdgeCreatesCycle(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)

	err := dag.AddEdge(3, 1, 0)
	if err == nil {
		t.Error("Adding edge that creates cycle should return error")
	}
}

func TestDAGAddEdgeCreatesComplexCycle(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)
	dag.AddEdge(1, 4, 0)
	dag.AddEdge(4, 3, 0)

	err := dag.AddEdge(3, 1, 0)
	if err == nil {
		t.Error("Adding edge (3, 1) should create a cycle through multiple paths")
	}
}

func TestDAGRemoveVertex(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddVertex(1)
	dag.AddVertex(2)
	dag.AddEdge(1, 2, 0)

	removed := dag.RemoveVertex(1)
	if !removed {
		t.Error("RemoveVertex should return true for existing vertex")
	}

	if dag.HasVertex(1) {
		t.Error("Vertex 1 should be removed")
	}

	if dag.HasEdge(1, 2) {
		t.Error("Edge (1, 2) should be removed when vertex 1 is removed")
	}

	if dag.VertexCount() != 1 {
		t.Errorf("Expected 1 vertex after removal, got %d", dag.VertexCount())
	}
}

func TestDAGRemoveVertexNonExistent(t *testing.T) {
	dag := NewDAG[int]()
	removed := dag.RemoveVertex(1)

	if removed {
		t.Error("RemoveVertex should return false for non-existent vertex")
	}
}

func TestDAGRemoveEdge(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)

	removed := dag.RemoveEdge(1, 2)
	if !removed {
		t.Error("RemoveEdge should return true for existing edge")
	}

	if dag.HasEdge(1, 2) {
		t.Error("Edge should be removed")
	}

	if dag.EdgeCount() != 0 {
		t.Errorf("Expected 0 edges after removal, got %d", dag.EdgeCount())
	}
}

func TestDAGRemoveEdgeNonExistent(t *testing.T) {
	dag := NewDAG[int]()
	removed := dag.RemoveEdge(1, 2)

	if removed {
		t.Error("RemoveEdge should return false for non-existent edge")
	}
}

func TestDAGGetNeighbors(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(1, 3, 0)
	dag.AddEdge(2, 4, 0)

	neighbors := dag.GetNeighbors(1)
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 neighbors for vertex 1, got %d", len(neighbors))
	}

	has2, has3 := false, false
	for _, v := range neighbors {
		if v == 2 {
			has2 = true
		}
		if v == 3 {
			has3 = true
		}
	}

	if !has2 || !has3 {
		t.Errorf("Neighbors should include 2 and 3, got %v", neighbors)
	}
}

func TestDAGGetIncomingNeighbors(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 3, 0)
	dag.AddEdge(2, 3, 0)
	dag.AddEdge(4, 1, 0)

	neighbors := dag.GetIncomingNeighbors(3)
	if len(neighbors) != 2 {
		t.Errorf("Expected 2 incoming neighbors for vertex 3, got %d", len(neighbors))
	}

	has1, has2 := false, false
	for _, v := range neighbors {
		if v == 1 {
			has1 = true
		}
		if v == 2 {
			has2 = true
		}
	}

	if !has1 || !has2 {
		t.Errorf("Incoming neighbors should include 1 and 2, got %v", neighbors)
	}
}

func TestDAGGetOutgoingEdges(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 5.0)
	dag.AddEdge(1, 3, 10.0)

	edges := dag.GetOutgoingEdges(1)
	if len(edges) != 2 {
		t.Errorf("Expected 2 outgoing edges from vertex 1, got %d", len(edges))
	}

	if edges[0].From != 1 {
		t.Errorf("Edge should start from 1, got %d", edges[0].From)
	}
}

func TestDAGGetIncomingEdges(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 3, 5.0)
	dag.AddEdge(2, 3, 10.0)

	edges := dag.GetIncomingEdges(3)
	if len(edges) != 2 {
		t.Errorf("Expected 2 incoming edges to vertex 3, got %d", len(edges))
	}

	for _, edge := range edges {
		if edge.To != 3 {
			t.Errorf("Edge should end at 3, got %d", edge.To)
		}
	}
}

func TestDAGGetVertices(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddVertex(1)
	dag.AddVertex(2)
	dag.AddVertex(3)

	vertices := dag.GetVertices()
	if len(vertices) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(vertices))
	}
}

func TestDAGGetEdges(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)

	edges := dag.GetEdges()
	if len(edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(edges))
	}
}

func TestDAGVertexCount(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddVertex(1)
	dag.AddVertex(2)
	dag.AddVertex(3)

	if dag.VertexCount() != 3 {
		t.Errorf("Expected 3 vertices, got %d", dag.VertexCount())
	}
}

func TestDAGEdgeCount(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)
	dag.AddEdge(3, 4, 0)

	if dag.EdgeCount() != 3 {
		t.Errorf("Expected 3 edges, got %d", dag.EdgeCount())
	}
}

func TestDAGInDegree(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 3, 0)
	dag.AddEdge(2, 3, 0)
	dag.AddEdge(4, 3, 0)

	if dag.InDegree(3) != 3 {
		t.Errorf("Expected in-degree 3 for vertex 3, got %d", dag.InDegree(3))
	}

	if dag.InDegree(1) != 0 {
		t.Errorf("Expected in-degree 0 for vertex 1, got %d", dag.InDegree(1))
	}
}

func TestDAGOutDegree(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(1, 3, 0)
	dag.AddEdge(1, 4, 0)

	if dag.OutDegree(1) != 3 {
		t.Errorf("Expected out-degree 3 for vertex 1, got %d", dag.OutDegree(1))
	}

	if dag.OutDegree(2) != 0 {
		t.Errorf("Expected out-degree 0 for vertex 2, got %d", dag.OutDegree(2))
	}
}

func TestDAGClear(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)

	dag.Clear()

	if !dag.IsEmpty() {
		t.Error("DAG should be empty after Clear")
	}

	if dag.VertexCount() != 0 {
		t.Errorf("Expected 0 vertices after Clear, got %d", dag.VertexCount())
	}

	if dag.EdgeCount() != 0 {
		t.Errorf("Expected 0 edges after Clear, got %d", dag.EdgeCount())
	}
}

func TestDAGCopy(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 5.0)
	dag.AddEdge(2, 3, 10.0)

	copyDAG := dag.Copy()

	if copyDAG.VertexCount() != dag.VertexCount() {
		t.Errorf("Copy should have same number of vertices, got %d vs %d",
			copyDAG.VertexCount(), dag.VertexCount())
	}

	if copyDAG.EdgeCount() != dag.EdgeCount() {
		t.Errorf("Copy should have same number of edges, got %d vs %d",
			copyDAG.EdgeCount(), dag.EdgeCount())
	}

	if !copyDAG.HasEdge(1, 2) {
		t.Error("Copy should have edge (1, 2)")
	}

	weight, _ := copyDAG.GetEdgeWeight(1, 2)
	if weight != 5.0 {
		t.Errorf("Copy should preserve edge weight, expected 5.0, got %f", weight)
	}

	copyDAG.RemoveVertex(1)
	if !dag.HasVertex(1) {
		t.Error("Modifying copy should not affect original")
	}
}

func TestDAGString(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)

	str := dag.String()
	if str == "" {
		t.Error("String should not be empty")
	}
}

func TestDAGIsEmpty(t *testing.T) {
	dag := NewDAG[int]()

	if !dag.IsEmpty() {
		t.Error("New DAG should be empty")
	}

	dag.AddVertex(1)
	if dag.IsEmpty() {
		t.Error("DAG with vertex should not be empty")
	}
}

func TestDAGIsDAG(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)
	dag.AddEdge(1, 4, 0)

	if !dag.IsDAG() {
		t.Error("Valid DAG should return true")
	}
}

func TestDAGTranspose(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)
	dag.AddEdge(2, 3, 0)

	transposed := dag.Transpose()

	if transposed.HasEdge(1, 2) {
		t.Error("Transposed DAG should not have original edge (1, 2)")
	}

	if !transposed.HasEdge(2, 1) {
		t.Error("Transposed DAG should have reversed edge (2, 1)")
	}

	if transposed.EdgeCount() != dag.EdgeCount() {
		t.Error("Transposed DAG should have same number of edges")
	}
}

func TestDAGToGraph(t *testing.T) {
	dag := NewDAG[int]()
	dag.AddEdge(1, 2, 0)

	graph := dag.ToGraph()

	if graph == nil {
		t.Error("ToGraph should return non-nil graph")
	}

	if !graph.HasEdge(1, 2) {
		t.Error("Returned graph should have edge (1, 2)")
	}
}

func TestDAGComplexStructure(t *testing.T) {
	dag := NewDAG[int]()

	err := dag.AddEdge(1, 2, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dag.AddEdge(1, 3, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dag.AddEdge(2, 4, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dag.AddEdge(2, 5, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = dag.AddEdge(3, 6, 0)
	if err != nil {
		t.Fatal(err)
	}

	if dag.VertexCount() != 6 {
		t.Errorf("Expected 6 vertices, got %d", dag.VertexCount())
	}

	if dag.EdgeCount() != 5 {
		t.Errorf("Expected 5 edges, got %d", dag.EdgeCount())
	}

	if !dag.IsDAG() {
		t.Error("Complex DAG should remain acyclic")
	}
}

func TestDAGLinearChain(t *testing.T) {
	dag := NewDAG[int]()

	for i := 1; i < 11; i++ {
		err := dag.AddEdge(i, i+1, 0)
		if err != nil {
			t.Fatalf("Failed to add edge %d -> %d: %v", i, i+1, err)
		}
	}

	if dag.VertexCount() != 11 {
		t.Errorf("Expected 11 vertices, got %d", dag.VertexCount())
	}

	if dag.EdgeCount() != 10 {
		t.Errorf("Expected 10 edges, got %d", dag.EdgeCount())
	}

	if dag.OutDegree(1) != 1 {
		t.Errorf("Expected out-degree 1 for vertex 1, got %d", dag.OutDegree(1))
	}

	if dag.InDegree(5) != 1 {
		t.Errorf("Expected in-degree 1 for vertex 5, got %d", dag.InDegree(5))
	}
}

func TestDAGMultipleSourcesSinks(t *testing.T) {
	dag := NewDAG[int]()

	dag.AddEdge(1, 4, 0)
	dag.AddEdge(2, 4, 0)
	dag.AddEdge(3, 4, 0)
	dag.AddEdge(4, 5, 0)
	dag.AddEdge(4, 6, 0)
	dag.AddEdge(4, 7, 0)

	sources := []int{}
	sinks := []int{}

	for _, v := range dag.GetVertices() {
		if dag.InDegree(v) == 0 {
			sources = append(sources, v)
		}
		if dag.OutDegree(v) == 0 {
			sinks = append(sinks, v)
		}
	}

	if len(sources) != 3 {
		t.Errorf("Expected 3 sources, got %d", len(sources))
	}

	if len(sinks) != 3 {
		t.Errorf("Expected 3 sinks, got %d", len(sinks))
	}
}

func TestDAGStringVertices(t *testing.T) {
	dag := NewDAG[string]()

	err := dag.AddEdge("A", "B", 1.0)
	if err != nil {
		t.Fatal(err)
	}

	err = dag.AddEdge("B", "C", 2.0)
	if err != nil {
		t.Fatal(err)
	}

	if !dag.HasVertex("A") {
		t.Error("Should have vertex 'A'")
	}

	if !dag.HasEdge("A", "B") {
		t.Error("Should have edge ('A', 'B')")
	}

	weight, _ := dag.GetEdgeWeight("B", "C")
	if weight != 2.0 {
		t.Errorf("Expected weight 2.0, got %f", weight)
	}
}

func BenchmarkDAGAddEdge(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 1000; i++ {
		dag.AddVertex(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.AddEdge(i%1000, (i+1)%1000, 0)
	}
}

func BenchmarkDAGIsDAG(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.IsDAG()
	}
}

func BenchmarkDAGCopy(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.Copy()
	}
}

func BenchmarkDAGTranspose(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.Transpose()
	}
}

func BenchmarkDAGGetNeighbors(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.GetNeighbors(i % 100)
	}
}

func BenchmarkDAGGetIncomingNeighbors(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.GetIncomingNeighbors(i % 100)
	}
}

func BenchmarkDAGRemoveVertex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		dag := NewDAG[int]()
		for j := 0; j < 100; j++ {
			for k := j + 1; k < 100; k++ {
				dag.AddEdge(j, k, 0)
			}
		}
		b.StartTimer()
		dag.RemoveVertex(i % 100)
	}
}

func BenchmarkDAGRemoveEdge(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.RemoveEdge(i%99, (i+1)%99+1)
	}
}

func BenchmarkDAGCopyLarge(b *testing.B) {
	dag := NewDAG[int]()
	for i := 0; i < 500; i++ {
		for j := i + 1; j < 500; j++ {
			dag.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dag.Copy()
	}
}
