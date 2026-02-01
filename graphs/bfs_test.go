package graphs

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	result, err := BFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 4 {
		t.Errorf("Expected 4 vertices in order, got %d", len(result.Order))
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance from 1 to 1 should be 0, got %d", result.Distances[1])
	}

	if result.Distances[2] != 1 {
		t.Errorf("Distance from 1 to 2 should be 1, got %d", result.Distances[2])
	}

	if result.Distances[4] != 3 {
		t.Errorf("Distance from 1 to 4 should be 3, got %d", result.Distances[4])
	}
}

func TestBFSNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := BFS(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestBFSDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)

	result, err := BFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := result.Distances[4]; !exists {
		t.Error("Vertex 4 should be reachable")
	}

	if result.Distances[4] != 1 {
		t.Errorf("Distance from 1 to 4 should be 1, got %d", result.Distances[4])
	}
}

func TestBFSDisconnectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	result, err := BFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := result.Distances[3]; exists {
		t.Error("Vertex 3 should not be reachable from 1")
	}
}

func TestBFSParents(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	result, _ := BFS(g, 1)

	if parent, ok := result.Parents[2]; !ok || parent != 1 {
		t.Errorf("Expected parent 1 for vertex 2, got %d", parent)
	}

	if parent, ok := result.Parents[3]; !ok || parent != 2 {
		t.Errorf("Expected parent 2 for vertex 3, got %d", parent)
	}

	if _, ok := result.Parents[1]; ok {
		t.Error("Start vertex should not have a parent")
	}
}

func TestBFSVisitedOrder(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 4, 0)

	result, _ := BFS(g, 1)

	if len(result.VisitedOrder) != 4 {
		t.Errorf("Expected 4 vertices in visited order, got %d", len(result.VisitedOrder))
	}

	if result.VisitedOrder[0] != 1 {
		t.Errorf("First visited should be 1, got %d", result.VisitedOrder[0])
	}
}

func TestBFSWithPredicate(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	path, err := BFSWithPredicate(g, 1, func(v int) bool {
		return v == 4
	})

	if err != nil {
		t.Fatal(err)
	}

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

func TestBFSWithPredicateNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := BFSWithPredicate(g, 999, func(v int) bool { return v == 2 })
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestBFSWithPredicateStartSatisfies(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	path, err := BFSWithPredicate(g, 1, func(v int) bool { return v == 1 })

	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 1 || path[0] != 1 {
		t.Errorf("Expected path [1], got %v", path)
	}
}

func TestBFSWithPredicateNotFound(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	_, err := BFSWithPredicate(g, 1, func(v int) bool { return v == 4 })

	if err == nil {
		t.Error("Expected error when no vertex satisfies predicate")
	}
}

func TestBFSAllComponents(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)
	g.AddEdge(4, 5, 0)

	results := BFSAllComponents(g)

	if len(results) != 2 {
		t.Errorf("Expected 2 components, got %d", len(results))
	}

	totalVertices := 0
	for _, result := range results {
		totalVertices += len(result.Order)
	}

	if totalVertices != 5 {
		t.Errorf("Expected total of 5 vertices, got %d", totalVertices)
	}
}

func TestBFSAllComponentsSingleComponent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	results := BFSAllComponents(g)

	if len(results) != 1 {
		t.Errorf("Expected 1 component, got %d", len(results))
	}
}

func TestBFSComponent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	result := BFSComponent(g, 1, make(map[int]bool))

	if len(result.Order) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(result.Order))
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance to start should be 0, got %d", result.Distances[1])
	}
}

func TestShortestPath(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(4, 3, 0)

	path, err := ShortestPath(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	expectedPath := []int{1, 2, 3}
	if len(path) != len(expectedPath) {
		t.Fatalf("Expected path length %d, got %d", len(expectedPath), len(path))
	}

	for i, v := range expectedPath {
		if path[i] != v {
			t.Errorf("Path at index %d: expected %d, got %d", i, v, path[i])
		}
	}
}

func TestShortestPathNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := ShortestPath(g, 999, 2)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestShortestPathNonExistentEnd(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := ShortestPath(g, 1, 999)
	if err == nil {
		t.Error("Expected error for non-existent end vertex")
	}
}

func TestShortestPathNoPath(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	_, err := ShortestPath(g, 1, 4)
	if err == nil {
		t.Error("Expected error when no path exists")
	}
}

func TestShortestPathSameVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	path, err := ShortestPath(g, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 1 || path[0] != 1 {
		t.Errorf("Expected path [1], got %v", path)
	}
}

func TestShortestPathUnweighted(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	path, err := ShortestPathUnweighted(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	expectedPath := []int{1, 2, 3}
	if len(path) != len(expectedPath) {
		t.Fatalf("Expected path length %d, got %d", len(expectedPath), len(path))
	}

	for i, v := range expectedPath {
		if path[i] != v {
			t.Errorf("Path at index %d: expected %d, got %d", i, v, path[i])
		}
	}
}

func TestDistance(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	distance, err := Distance(g, 1, 4)
	if err != nil {
		t.Fatal(err)
	}

	if distance != 3 {
		t.Errorf("Expected distance 3, got %d", distance)
	}
}

func TestDistanceNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := Distance(g, 999, 2)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestDistanceNonExistentEnd(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	distance, err := Distance(g, 1, 999)
	if err != nil {
		t.Fatal(err)
	}

	if distance != -1 {
		t.Errorf("Expected distance -1, got %d", distance)
	}
}

func TestReachableVertices(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(4, 5, 0)

	vertices, err := ReachableVertices(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(vertices) != 3 {
		t.Errorf("Expected 3 reachable vertices, got %d", len(vertices))
	}
}

func TestReachableVerticesNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := ReachableVertices(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestBFSLevelOrder(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 4, 0)
	g.AddEdge(3, 5, 0)

	levels, err := BFSLevelOrder(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(levels) < 3 {
		t.Errorf("Expected at least 3 levels, got %d", len(levels))
	}

	if len(levels[0]) != 1 || levels[0][0] != 1 {
		t.Errorf("Level 0 should be [1], got %v", levels[0])
	}

	if len(levels[1]) != 2 {
		t.Errorf("Level 1 should have 2 vertices, got %d", len(levels[1]))
	}
}

func TestBFSLevelOrderNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := BFSLevelOrder(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestIsReachable(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	reachable, err := IsReachable(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if !reachable {
		t.Error("Vertex 3 should be reachable from 1")
	}
}

func TestIsReachableNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := IsReachable(g, 999, 2)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestIsReachableNonExistentEnd(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	reachable, err := IsReachable(g, 1, 999)
	if err != nil {
		t.Fatal(err)
	}

	if reachable {
		t.Error("Vertex 999 should not be reachable from 1")
	}
}

func TestIsReachableDisconnected(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	reachable, err := IsReachable(g, 1, 4)
	if err != nil {
		t.Fatal(err)
	}

	if reachable {
		t.Error("Vertex 4 should not be reachable from 1")
	}
}

func TestBFSCycleDetection(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)

	hasCycle := BFSCycleDetection(g)
	if !hasCycle {
		t.Error("Directed graph with cycle should be detected as having a cycle")
	}
}

func TestBFSCycleDetectionAcyclic(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	hasCycle := BFSCycleDetection(g)
	if hasCycle {
		t.Error("Acyclic directed graph should not be detected as having a cycle")
	}
}

func TestBFSCycleDetectionUndirected(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)

	hasCycle := BFSCycleDetection(g)
	if hasCycle {
		t.Error("BFSCycleDetection should not detect cycles in undirected graphs (as per implementation)")
	}
}

func TestBFSSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	result, err := BFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 1 {
		t.Errorf("Expected 1 vertex, got %d", len(result.Order))
	}

	if result.Distances[1] != 0 {
		t.Errorf("Distance to self should be 0, got %d", result.Distances[1])
	}
}

func TestBFSComplexGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 4, 0)
	g.AddEdge(3, 5, 0)

	result, _ := BFS(g, 1)

	if len(result.Order) != 5 {
		t.Errorf("Expected 5 vertices, got %d", len(result.Order))
	}

	if result.Distances[5] != 2 {
		t.Errorf("Distance to vertex 5 should be 2, got %d", result.Distances[5])
	}

	if result.Distances[4] != 2 {
		t.Errorf("Distance to vertex 4 should be 2, got %d", result.Distances[4])
	}
}

func BenchmarkBFS(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 1000; i++ {
		g.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFS(g, 0)
	}
}

func BenchmarkBFSWithPredicate(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFSWithPredicate(g, 0, func(v int) bool { return v == 99 })
	}
}

func BenchmarkShortestPath(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortestPath(g, 0, 99)
	}
}

func BenchmarkBFSAllComponents(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i > 0 && i%10 == 0 {
			continue
		}
		g.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFSAllComponents(g)
	}
}

func BenchmarkBFSLevelOrder(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1, 1)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BFSLevelOrder(g, 0)
	}
}
