package graphs

import (
	"testing"
)

func TestKruskal(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 5)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges in MST (V-1), got %d", len(mst.Edges))
	}

	expectedCost := 4.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	_, err := Kruskal(g)
	if err == nil {
		t.Error("Kruskal should return error for directed graph")
	}
}

func TestKruskalEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error for empty graph: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for empty graph, got %d", len(mst.Edges))
	}

	if mst.TotalCost != 0 {
		t.Errorf("Expected total cost 0 for empty graph, got %f", mst.TotalCost)
	}
}

func TestKruskalSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error for single vertex: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for single vertex, got %d", len(mst.Edges))
	}
}

func TestKruskalDisconnectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	_, err := Kruskal(g)
	if err == nil {
		t.Error("Kruskal should return error for disconnected graph")
	}
}

func TestKruskalLinearGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 4 {
		t.Errorf("Expected 4 edges, got %d", len(mst.Edges))
	}

	expectedCost := 10.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalCompleteGraph(t *testing.T) {
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

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalAllSameWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(1, 4, 1)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalNegativeWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, -2)
	g.AddEdge(2, 3, 3)
	g.AddEdge(1, 3, 4)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 1.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalZeroWeightEdges(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 5)
	g.AddEdge(1, 3, 3)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalComplexGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 6)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 3, 15)
	g.AddEdge(3, 4, 4)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}

	expectedCost := 19.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalWithLimit(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)
	g.AddEdge(5, 1, 5)

	mst, err := KruskalWithLimit(g, 2)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error: %v", err)
	}

	if len(mst.Edges) > 2 {
		t.Errorf("Expected at most 2 edges, got %d", len(mst.Edges))
	}
}

func TestKruskalWithLimitZero(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	mst, err := KruskalWithLimit(g, 0)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for limit 0, got %d", len(mst.Edges))
	}

	if mst.TotalCost != 0 {
		t.Errorf("Expected total cost 0 for limit 0, got %f", mst.TotalCost)
	}
}

func TestKruskalWithLimitEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	mst, err := KruskalWithLimit(g, 5)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error for empty graph: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for empty graph, got %d", len(mst.Edges))
	}
}

func TestKruskalWithLimitNegative(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	mst, err := KruskalWithLimit(g, -1)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error for negative limit: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for negative limit, got %d", len(mst.Edges))
	}
}

func TestKruskalWithLimitGreaterThanMST(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)

	mst, err := KruskalWithLimit(g, 100)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}
}

func TestKruskalByWeightLimit(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)
	g.AddEdge(4, 5, 4)

	mst, err := KruskalByWeightLimit(g, 5.5)
	if err != nil {
		t.Fatalf("KruskalByWeightLimit should not return error: %v", err)
	}

	if mst.TotalCost > 5.5 {
		t.Errorf("Expected total cost <= 5.5, got %f", mst.TotalCost)
	}
}

func TestKruskalByWeightLimitZero(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	mst, err := KruskalByWeightLimit(g, 0)
	if err != nil {
		t.Fatalf("KruskalByWeightLimit should not return error: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for limit 0, got %d", len(mst.Edges))
	}

	if mst.TotalCost != 0 {
		t.Errorf("Expected total cost 0 for limit 0, got %f", mst.TotalCost)
	}
}

func TestKruskalByWeightLimitEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	mst, err := KruskalByWeightLimit(g, 5)
	if err != nil {
		t.Fatalf("KruskalByWeightLimit should not return error for empty graph: %v", err)
	}

	if len(mst.Edges) != 0 {
		t.Errorf("Expected 0 edges for empty graph, got %d", len(mst.Edges))
	}
}

func TestKruskalByWeightLimitGreaterThanMST(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 3)

	mst, err := KruskalByWeightLimit(g, 100)
	if err != nil {
		t.Fatalf("KruskalByWeightLimit should not return error: %v", err)
	}

	if len(mst.Edges) != 3 {
		t.Errorf("Expected 3 edges, got %d", len(mst.Edges))
	}

	expectedCost := 6.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalAllComponents(t *testing.T) {
	g := NewUndirectedGraph[int]()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	g.AddEdge(4, 5, 3)
	g.AddEdge(5, 6, 4)

	msts, err := KruskalAllComponents(g)
	if err != nil {
		t.Fatalf("KruskalAllComponents should not return error: %v", err)
	}

	if len(msts) != 2 {
		t.Errorf("Expected 2 MSTs (one per component), got %d", len(msts))
	}

	totalEdges := 0
	for _, mst := range msts {
		totalEdges += len(mst.Edges)
	}

	if totalEdges != 4 {
		t.Errorf("Expected total of 4 edges (2 per component), got %d", totalEdges)
	}
}

func TestKruskalAllComponentsEmptyGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	msts, err := KruskalAllComponents(g)
	if err != nil {
		t.Fatalf("KruskalAllComponents should not return error for empty graph: %v", err)
	}

	if len(msts) != 0 {
		t.Errorf("Expected 0 MSTs for empty graph, got %d", len(msts))
	}
}

func TestKruskalAllComponentsSingleVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	msts, err := KruskalAllComponents(g)
	if err != nil {
		t.Fatalf("KruskalAllComponents should not return error: %v", err)
	}

	if len(msts) != 0 {
		t.Errorf("Expected 0 MSTs for graph with only single vertices, got %d", len(msts))
	}
}

func TestKruskalAllComponentsDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	_, err := KruskalAllComponents(g)
	if err == nil {
		t.Error("KruskalAllComponents should return error for directed graph")
	}
}

func TestKruskalAllComponentsWithIsolatedVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()

	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)

	g.AddVertex(4)
	g.AddVertex(5)

	msts, err := KruskalAllComponents(g)
	if err != nil {
		t.Fatalf("KruskalAllComponents should not return error: %v", err)
	}

	if len(msts) != 1 {
		t.Errorf("Expected 1 MST (isolated vertices should not create MSTs), got %d", len(msts))
	}
}

func TestKruskalTriangle(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 3, 3)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}

	for _, edge := range mst.Edges {
		if edge.Weight == 3 {
			t.Error("MST should not include the heaviest edge (3)")
		}
	}
}

func TestKruskalStringVertices(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)
	g.AddEdge("A", "C", 3)

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 2 {
		t.Errorf("Expected 2 edges, got %d", len(mst.Edges))
	}

	expectedCost := 3.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func TestKruskalWithLimitRespectsCycle(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 1, 1)
	g.AddEdge(3, 4, 5)

	mst, err := KruskalWithLimit(g, 3)
	if err != nil {
		t.Fatalf("KruskalWithLimit should not return error: %v", err)
	}

	if len(mst.Edges) > 3 {
		t.Errorf("Expected at most 3 edges, got %d", len(mst.Edges))
	}

	hasCycle := false
	visited := make(map[int]bool)
	for _, edge := range mst.Edges {
		visited[edge.From] = true
		visited[edge.To] = true
	}

	if len(visited) > len(mst.Edges)+1 {
		hasCycle = true
	}

	if hasCycle {
		t.Error("MST should not contain cycles")
	}
}

func TestKruskalLargeGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()

	for i := 1; i <= 10; i++ {
		for j := i + 1; j <= 10; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	mst, err := Kruskal(g)
	if err != nil {
		t.Fatalf("Kruskal should not return error: %v", err)
	}

	if len(mst.Edges) != 9 {
		t.Errorf("Expected 9 edges, got %d", len(mst.Edges))
	}

	expectedCost := 9.0
	if mst.TotalCost != expectedCost {
		t.Errorf("Expected total cost %.1f, got %.1f", expectedCost, mst.TotalCost)
	}
}

func BenchmarkKruskal(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Kruskal(g)
	}
}

func BenchmarkKruskalWithLimit(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KruskalWithLimit(g, 50)
	}
}

func BenchmarkKruskalByWeightLimit(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KruskalByWeightLimit(g, 50.0)
	}
}

func BenchmarkKruskalAllComponents(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i > 0 && i%10 == 0 {
			continue
		}
		g.AddEdge(i, i+1, 1.0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KruskalAllComponents(g)
	}
}

func BenchmarkKruskalLarge(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 200; i++ {
		for j := i + 1; j < 200; j++ {
			g.AddEdge(i, j, float64(j-i))
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Kruskal(g)
	}
}
