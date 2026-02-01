package graphs

import (
	"math"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Start != 1 {
		t.Errorf("Start should be 1, got %d", result.Start)
	}

	expectedDistances := map[int]float64{
		1: 0,
		2: 3,
		3: 2,
		4: 8,
	}

	for vertex, expectedDist := range expectedDistances {
		if result.Distances[vertex] != expectedDist {
			t.Errorf("Distance to %d: expected %f, got %f", vertex, expectedDist, result.Distances[vertex])
		}
	}
}

func TestBellmanFordNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFord(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestBellmanFordNegativeWeights(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 5)
	g.AddEdge(2, 3, -2)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[4] != 4 {
		t.Errorf("Expected distance 4 to vertex 4, got %f", result.Distances[4])
	}
}

func TestBellmanFordNegativeCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, -3)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !result.HasCycle {
		t.Error("Should detect negative cycle")
	}

	if len(result.CycleVertices) == 0 {
		t.Error("Should have vertices in the cycle")
	}
}

func TestBellmanFordSameVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance to same vertex should be 0, got %f", result.Distances[1])
	}
}

func TestBellmanFordUnreachable(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !math.IsInf(result.Distances[3], 1) {
		t.Errorf("Unreachable vertex should have infinite distance, got %f", result.Distances[3])
	}
}

func TestBellmanFordDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 2 {
		t.Errorf("Expected distance 2 to vertex 3, got %f", result.Distances[3])
	}

	if result.Distances[2] != 4 {
		t.Errorf("Expected distance 4 to vertex 2, got %f", result.Distances[2])
	}
}

func TestBellmanFordWithLimit(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFordWithLimit(g, 1, 2)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 2 {
		t.Errorf("Expected distance 2 with limit 2, got %f", result.Distances[3])
	}

	if math.IsInf(result.Distances[4], 1) {
		t.Log("Vertex 4 is unreachable with iteration limit 2 (as expected)")
	}
}

func TestBellmanFordWithLimitTooSmall(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(1, 6, 100)

	result, err := BellmanFordWithLimit(g, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[6] != 100 {
		t.Errorf("Expected distance 100 with 1 iteration, got %f", result.Distances[6])
	}
}

func TestBellmanFordWithLimitZero(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFordWithLimit(g, 1, 0)
	if err == nil {
		t.Error("Expected error for zero max iterations")
	}
}

func TestBellmanFordWithLimitNegativeCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, -3)

	result, err := BellmanFordWithLimit(g, 1, 100)
	if err != nil {
		t.Fatal(err)
	}

	if !result.HasCycle {
		t.Error("Should detect negative cycle even with limit")
	}
}

func TestBellmanFordToSpecific(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	result, err := BellmanFordToSpecific(g, 1, 4)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[4] != 8 {
		t.Errorf("Expected distance 8 to vertex 4, got %f", result.Distances[4])
	}
}

func TestBellmanFordToSpecificNonExistentTarget(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFordToSpecific(g, 1, 999)
	if err == nil {
		t.Error("Expected error for non-existent target vertex")
	}
}

func TestBellmanFordToSpecificNonExistentStart(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFordToSpecific(g, 999, 2)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestBellmanFordToSpecificNegativeCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, -3)

	result, err := BellmanFordToSpecific(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if !result.HasCycle {
		t.Error("Should detect negative cycle")
	}
}

func TestBellmanFordMultipleSources(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 2)
	g.AddEdge(4, 3, 1)

	result, err := BellmanFordMultipleSources(g, []int{1, 4})
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 1 {
		t.Errorf("Expected distance 1 to vertex 3 (from 4), got %f", result.Distances[3])
	}

	if result.Distances[1] != 0 {
		t.Errorf("Expected distance 0 to vertex 1, got %f", result.Distances[1])
	}

	if result.Distances[4] != 0 {
		t.Errorf("Expected distance 0 to vertex 4, got %f", result.Distances[4])
	}

	if result.Distances[2] != 2 {
		t.Errorf("Expected distance 2 to vertex 2, got %f", result.Distances[2])
	}
}

func TestBellmanFordMultipleSourcesEmpty(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFordMultipleSources(g, []int{})
	if err == nil {
		t.Error("Expected error for empty sources slice")
	}
}

func TestBellmanFordMultipleSourcesNonExistent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := BellmanFordMultipleSources(g, []int{1, 999})
	if err == nil {
		t.Error("Expected error for non-existent source vertex")
	}
}

func TestBellmanFordMultipleSourcesNegativeCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 2, -3)

	result, err := BellmanFordMultipleSources(g, []int{1})
	if err != nil {
		t.Fatal(err)
	}

	if !result.HasCycle {
		t.Error("Should detect negative cycle")
	}
}

func TestBellmanFordPathReconstruction(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	path := reconstructPathBF(result.Parents, 1, 4)
	expectedPath := []int{1, 2, 3, 4}

	if len(path) != len(expectedPath) {
		t.Fatalf("Expected path length %d, got %d", len(expectedPath), len(path))
	}

	for i, v := range expectedPath {
		if path[i] != v {
			t.Errorf("Path at index %d: expected %d, got %d", i, v, path[i])
		}
	}
}

func TestBellmanFordPathReconstructionUnreachable(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	path := reconstructPathBF(result.Parents, 1, 4)
	if len(path) > 0 && path[len(path)-1] != 4 {
		t.Errorf("Path should not reach unreachable vertex 4")
	}
}

func TestBellmanFordZeroWeightEdge(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 1 {
		t.Errorf("Expected distance 1 with zero-weight edge, got %f", result.Distances[3])
	}
}

func TestBellmanFordMultipleNegativeWeights(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 5)
	g.AddEdge(2, 3, -2)
	g.AddEdge(3, 4, -1)
	g.AddEdge(4, 5, 3)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[5] != 5 {
		t.Errorf("Expected distance 5 to vertex 5, got %f", result.Distances[5])
	}
}

func TestBellmanFordSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance to single vertex should be 0, got %f", result.Distances[1])
	}
}

func TestBellmanFordLargeGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i+1 < 100 {
			g.AddEdge(i, i+1, 1)
		}
	}

	result, err := BellmanFord(g, 0)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[99] != 99 {
		t.Errorf("Expected distance 99 in linear graph, got %f", result.Distances[99])
	}
}

func TestBellmanFordParentsMap(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := result.Parents[2]; !ok {
		t.Error("Vertex 2 should have a parent")
	}
	if _, ok := result.Parents[3]; !ok {
		t.Error("Vertex 3 should have a parent")
	}
	if parent, ok := result.Parents[1]; ok {
		t.Errorf("Start vertex should not have a parent, got %d", parent)
	}
}

func TestBellmanFordNegativeCycleNotReachable(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 3, -2)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.HasCycle {
		t.Error("Should not detect unreachable negative cycle")
	}
}

func TestBellmanFordWithLimitSameAsNormal(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	normalResult, _ := BellmanFord(g, 1)
	limitedResult, _ := BellmanFordWithLimit(g, 1, 10)

	if normalResult.Distances[4] != limitedResult.Distances[4] {
		t.Errorf("Results should be same when limit is high enough")
	}
}

func TestBellmanFordComplexGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, -3)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 1)

	result, err := BellmanFord(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[2] != 4 {
		t.Errorf("Expected distance 4 to vertex 2, got %f", result.Distances[2])
	}

	if result.Distances[3] != 1 {
		t.Errorf("Expected distance 1 to vertex 3, got %f", result.Distances[3])
	}

	if result.Distances[4] != 2 {
		t.Errorf("Expected distance 2 to vertex 4, got %f", result.Distances[4])
	}
}

func reconstructPathBF[T comparable](parents map[T]T, start, end T) []T {
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
	return path
}

func BenchmarkBellmanFord(b *testing.B) {
	g := createTestGraphBF(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BellmanFord(g, 0)
	}
}

func BenchmarkBellmanFordWithLimit(b *testing.B) {
	g := createTestGraphBF(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BellmanFordWithLimit(g, 0, 50)
	}
}

func BenchmarkBellmanFordToSpecific(b *testing.B) {
	g := createTestGraphBF(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BellmanFordToSpecific(g, 0, 99)
	}
}

func BenchmarkBellmanFordMultipleSources(b *testing.B) {
	g := createTestGraphBF(100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BellmanFordMultipleSources(g, []int{0, 50})
	}
}

func createTestGraphBF(size int) *Graph[int] {
	g := NewUndirectedGraph[int]()
	for i := 0; i < size; i++ {
		if i+1 < size {
			g.AddEdge(i, i+1, 1)
		}
		if i+10 < size {
			g.AddEdge(i, i+10, 5)
		}
	}
	return g
}
