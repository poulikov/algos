package graphs

import (
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	result, err := Dijkstra(g, 1)
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

func TestDijkstraDirected(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(3, 2, 1)
	g.AddEdge(2, 4, 5)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[2] != 3 {
		t.Errorf("Expected distance 3 to vertex 2, got %f", result.Distances[2])
	}
}

func TestDijkstraUnreachable(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !math.IsInf(result.Distances[3], 1) {
		t.Errorf("Unreachable vertex should have infinite distance, got %f", result.Distances[3])
	}
}

func TestDijkstraNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := Dijkstra(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestDijkstraNegativeWeight(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, -1)

	_, err := Dijkstra(g, 1)
	if err == nil {
		t.Error("Expected error for negative weight")
	}
}

func TestShortestPathDijkstra(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	path, err := ShortestPathDijkstra(g, 1, 4)
	if err != nil {
		t.Fatal(err)
	}

	expectedPath := []int{1, 3, 2, 4}
	if len(path) != len(expectedPath) {
		t.Fatalf("Expected path length %d, got %d", len(expectedPath), len(path))
	}

	for i, v := range expectedPath {
		if path[i] != v {
			t.Errorf("Path at index %d: expected %d, got %d", i, v, path[i])
		}
	}
}

func TestShortestPathDijkstraNoPath(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(3, 4, 1)

	_, err := ShortestPathDijkstra(g, 1, 4)
	if err == nil {
		t.Error("Expected error when no path exists")
	}
}

func TestShortestPathDijkstraSameVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	path, err := ShortestPathDijkstra(g, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 1 || path[0] != 1 {
		t.Errorf("Expected path [1], got %v", path)
	}
}

func TestDijkstraToSpecific(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	result, err := DijkstraToSpecific(g, 1, 2)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[2] != 3 {
		t.Errorf("Expected distance 3 to vertex 2, got %f", result.Distances[2])
	}
}

func TestDijkstraToSpecificNonExistentTarget(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := DijkstraToSpecific(g, 1, 999)
	if err == nil {
		t.Error("Expected error for non-existent target vertex")
	}
}

func TestDijkstraMultipleSources(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 2)
	g.AddEdge(4, 3, 1)

	result, err := DijkstraMultipleSources(g, []int{1, 4})
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 1 {
		t.Errorf("Expected distance 1 to vertex 3 (from 4), got %f", result.Distances[3])
	}
}

func TestDijkstraMultipleSourcesEmpty(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := DijkstraMultipleSources(g, []int{})
	if err == nil {
		t.Error("Expected error for empty sources slice")
	}
}

func TestDijkstraMultipleSourcesNonExistent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)

	_, err := DijkstraMultipleSources(g, []int{1, 999})
	if err == nil {
		t.Error("Expected error for non-existent source vertex")
	}
}

func TestDijkstraWithPathLimit(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	result, err := DijkstraWithPathLimit(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 2 {
		t.Errorf("Expected distance 2 with path limit 3, got %f", result.Distances[3])
	}

	if math.IsInf(result.Distances[4], 1) {
		t.Log("Vertex 4 is unreachable with path limit 3 (as expected)")
	}
}

func TestDijkstraWithPathLimitExceeded(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	result, err := DijkstraWithPathLimit(g, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[4] != math.Inf(1) {
		t.Errorf("Expected infinite distance when path limit exceeded, got %f", result.Distances[4])
	}
}

func TestDijkstraSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance to single vertex should be 0, got %f", result.Distances[1])
	}
}

func TestDijkstraPathReconstruction(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 5)

	result, _ := Dijkstra(g, 1)

	parent, ok := result.Parents[4]
	if !ok {
		t.Error("Should have parent for vertex 4")
	}
	if parent != 2 {
		t.Errorf("Expected parent 2 for vertex 4, got %d", parent)
	}
}

func TestDijkstraEqualWeights(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 3, 2)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 2 {
		t.Errorf("Expected distance 2 (either path), got %f", result.Distances[3])
	}
}

func TestDijkstraLargeGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	for i := 1; i <= 100; i++ {
		g.AddEdge(i, i+1, 1)
	}

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[100] != 99 {
		t.Errorf("Expected distance 99 in linear graph, got %f", result.Distances[100])
	}
}

func TestShortestPathDijkstraMultiple(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 3, 3)

	path, err := ShortestPathDijkstra(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 3 {
		t.Errorf("Expected path length 3, got %d", len(path))
	}

	if path[0] != 1 || path[1] != 2 || path[2] != 3 {
		t.Errorf("Expected path [1, 2, 3], got %v", path)
	}
}

func TestDijkstraZeroWeightEdge(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 1)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if result.Distances[3] != 1 {
		t.Errorf("Expected distance 1 with zero-weight edge, got %f", result.Distances[3])
	}
}

func TestDijkstraParentsMap(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)

	result, err := Dijkstra(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := result.Parents[2]; !ok {
		t.Error("Vertex 2 should have a parent")
	}
	if _, ok := result.Parents[3]; !ok {
		t.Error("Vertex 3 should have a parent")
	}
	if _, ok := result.Parents[1]; ok {
		t.Error("Start vertex should not have a parent")
	}
}
