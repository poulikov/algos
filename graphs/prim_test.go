package graphs

import (
	"math"
	"testing"
)

func TestPrim(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 5)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges in MST (V-1), got %d", len(mst.Edges))
	}

	expectedCost := 4.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	_, err := Prim(g, 1)
	if err == nil {
		t.Error("Prim should return error for directed graph")
	}
}

func TestPrimEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error for graph with single vertex: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for single vertex, got %d", len(mst.Edges))
	}

	if mst.TotalCost != 0 {
		t.Errorf("Expected total cost 0 for single vertex, got %f", mst.TotalCost)
	}
}

func TestPrimSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error for single vertex: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for single vertex, got %d", len(mst.Edges))
	}
}

func TestPrimDisconnectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	_, err := Prim(g, 1)
	if err == nil {
		t.Error("Prim should return error for disconnected graph")
	}
}

func TestPrimNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := Prim(g, 999)
	if err == nil {
		t.Error("Prim should return error for non-existent start vertex")
	}
}

func TestPrimLinearGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 4 {
		t.Errorf("Expected 4 edges, got %d", len(mst.Edges))
	}

	expectedCost := 10.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimCompleteGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	weights := [][]float64{
		{0, 1, 2},
		{1, 0, 3},
		{2, 3, 0},
	}

	for i := 1; i <= 3; i++ {
		for j := i + 1; j <= 3; j++ {
			g.AddEdge(i, j, weights[i-1][j-1])
		}
	}

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimTriangle(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 3, 3)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimNegativeWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, -2)
	g.AddEdge(2, 3, 3)
	g.AddEdge(1, 3, 4)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 1.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimZeroWeightEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 5)
	g.AddEdge(1, 3, 3)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimAllSameWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(1, 4, 1)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimDifferentStartVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 1, 4)

	mst1, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	mst2, err := Prim(g, 2)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	mst3, err := Prim(g, 3)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	mst4, err := Prim(g, 4)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if mst1.TotalCost != mst2.TotalCost || mst2.TotalCost != mst3.TotalCost || mst3.TotalCost != mst4.TotalCost {
		t.Error("MST cost should be the same regardless of start vertex")
	}

	expectedCost := 6.0
	if mst1.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst1.TotalCost)
	}
}

func TestPrimComplexGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 6)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 3, 15)
	g.AddEdge(3, 4, 4)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}

	expectedCost := 19.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeap(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 5)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges in MST (V-1), got %d", len(mst.Edges))
	}

	expectedCost := 4.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeapDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	_, err := PrimWithHeap(g, 1)
	if err == nil {
		t.Error("PrimWithHeap should return error for directed graph")
	}
}

func TestPrimWithHeapEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error for graph with single vertex: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for single vertex, got %d", len(mst.Edges))
	}
}

func TestPrimWithHeapSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error for single vertex: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for single vertex, got %d", len(mst.Edges))
	}
}

func TestPrimWithHeapDisconnectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	_, err := PrimWithHeap(g, 1)
	if err == nil {
		t.Error("PrimWithHeap should return error for disconnected graph")
	}
}

func TestPrimWithHeapNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := PrimWithHeap(g, 999)
	if err == nil {
		t.Error("PrimWithHeap should return error for non-existent start vertex")
	}
}

func TestPrimWithHeapSameAsPrim(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 5)

	mst1, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	mst2, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	if mst1.TotalCost != mst2.TotalCost {
		t.Errorf("Prim and PrimWithHeap should produce same cost: %.1f vs %.1f",
			mst1.TotalCost, mst2.TotalCost)
	}

	if len(mst1.Edges) != len(mst2.Edges) {
		t.Errorf("Prim and PrimWithHeap should produce same number of edges: %d vs %d",
			len(mst1.Edges), len(mst2.Edges))
	}
}

func TestPrimWithHeapLinearGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	if len(mst.Edges) != 4 {
		t.Errorf("Expected 4 edges, got %d", len(mst.Edges))
	}

	expectedCost := 10.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeapNegativeWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, -2)
	g.AddEdge(2, 3, 3)
	g.AddEdge(1, 3, 4)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	expectedCost := 1.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimAll(t *testing.T) {
	g := NewUndirectedGraph[int]()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)

	msts, err := PrimAll(g)
	if err != nil {
		t.Fatalf("PrimAll should not return error for connected graph: %v", err)
	}

	if len(msts) != 1 {
		t.Errorf("Expected 1 MST for single connected component, got %d", len(msts))
	}

	if len(msts[0].Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(msts[0].Edges))
	}
}

func TestPrimAllEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	msts, err := PrimAll(g)
	if err != nil {
		t.Fatalf("PrimAll should not return error for empty graph: %v", err)
	}

	if len(msts) != 0 {
		t.Errorf("Expected 0 MSTs for empty graph, got %d", len(msts))
	}
}

func TestPrimAllSingleVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	msts, err := PrimAll(g)
	if err != nil {
		t.Fatalf("PrimAll should not return error: %v", err)
	}

	if len(msts) != 0 {
		t.Errorf("Expected 0 MSTs for graph with only single vertices, got %d", len(msts))
	}
}

func TestPrimAllDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	_, err := PrimAll(g)
	if err == nil {
		t.Error("PrimAll should return error for directed graph")
	}
}

func TestPrimAllWithIsolatedVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)

	msts, err := PrimAll(g)
	if err != nil {
		t.Fatalf("PrimAll should not return error: %v", err)
	}

	if len(msts) != 1 {
		t.Errorf("Expected 1 MST, got %d", len(msts))
	}
}

func TestPrimStringVertices(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)
	g.AddEdge("A", "C", 3)

	mst, err := Prim(g, "A")
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeapStringVertices(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)
	g.AddEdge("A", "C", 3)

	mst, err := PrimWithHeap(g, "A")
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimLargeGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	for i := 1; i <= 10; i++ {
		for j := i + 1; j <= 10; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	if len(mst.Edges) != 9 {
		t.Errorf("Expected 9 edges, got %d", len(mst.Edges))
	}

	expectedCost := 9.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeapLargeGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	for i := 1; i <= 10; i++ {
		for j := i + 1; j <= 10; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	if len(mst.Edges) != 9 {
		t.Errorf("Expected 9 edges, got %d", len(mst.Edges))
	}

	expectedCost := 9.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimMSTProperty(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error: %v", err)
	}

	visited := make(map[int]bool)
	visited[mst.Edges[0].From] = true
	visited[mst.Edges[0].To] = true

	for _, edge := range mst.Edges[1:] {
		connected := visited[edge.From] || visited[edge.To]
		if !connected {
			t.Error("MST should be connected")
		}
		visited[edge.From] = true
		visited[edge.To] = true
	}

	if len(visited) != 5 {
		t.Error("MST should visit all vertices")
	}
}

func TestPrimWithHeapMSTProperty(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error: %v", err)
	}

	visited := make(map[int]bool)
	for _, edge := range mst.Edges {
		visited[edge.From] = true
		visited[edge.To] = true
	}

	if len(visited) != 5 {
		t.Error("MST should visit all vertices")
	}
}

func TestPrimNoEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	_, err := Prim(g, 1)
	if err == nil {
		t.Error("Prim should return error for graph with vertices but no edges")
	}
}

func TestPrimWithHeapNoEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	_, err := PrimWithHeap(g, 1)
	if err == nil {
		t.Error("PrimWithHeap should return error for graph with vertices but no edges")
	}
}

func TestPrimInfWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, math.Inf(1))
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 3, 2)

	mst, err := Prim(g, 1)
	if err != nil {
		t.Fatalf("Prim should not return error for graph with Inf weight: %v", err)
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestPrimWithHeapInfWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, math.Inf(1))
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 3, 2)

	mst, err := PrimWithHeap(g, 1)
	if err != nil {
		t.Fatalf("PrimWithHeap should not return error for graph with Inf weight: %v", err)
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}

	if math.IsInf(mst.TotalCost, 1) {
		t.Error("MST should not have infinite total cost if there's a finite path")
	}
}

func BenchmarkPrim(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Prim(g, 0)
	}
}

func BenchmarkPrimWithHeap(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimWithHeap(g, 0)
	}
}

func BenchmarkPrimAll(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i > 0 && i%10 == 0 {
			continue
		}
		g.AddEdge(i, i+1, 1.0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimAll(g)
	}
}

func BenchmarkPrimLarge(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Prim(g, 0)
	}
}

func BenchmarkPrimWithHeapLarge(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimWithHeap(g, 0)
	}
}
