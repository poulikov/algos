package graphs

import (
	"testing"
)

func TestDFS(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	result, err := DFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 4 {
		t.Errorf("Expected 4 vertices in order, got %d", len(result.Order))
	}

	if result.DiscoveryTime[1] != 1 {
		t.Errorf("Discovery time for 1 should be 1, got %d", result.DiscoveryTime[1])
	}

	if parent, ok := result.Parents[2]; !ok || parent != 1 {
		t.Errorf("Expected parent 1 for vertex 2, got %d", parent)
	}

	if len(result.VisitedOrder) != 4 {
		t.Errorf("Expected 4 vertices in visited order, got %d", len(result.VisitedOrder))
	}
}

func TestDFSNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := DFS(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestDFSDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)

	result, err := DFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := result.DiscoveryTime[4]; !exists {
		t.Error("Vertex 4 should be reachable")
	}
}

func TestDFSDisconnectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	result, err := DFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := result.DiscoveryTime[3]; exists {
		t.Error("Vertex 3 should not be reachable from 1")
	}
}

func TestDFSParents(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	result, _ := DFS(g, 1)

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

func TestDFSFinishTime(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	result, _ := DFS(g, 1)

	if _, exists := result.FinishTime[3]; !exists {
		t.Error("Finish time should be set for all vertices")
	}

	if result.FinishTime[1] < result.FinishTime[3] {
		t.Error("Start vertex should finish last")
	}

	if result.FinishTime[3] < result.FinishTime[2] || result.FinishTime[2] < result.FinishTime[3] {
	}
}

func TestDFSSingleVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddVertex(1)

	result, err := DFS(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 1 {
		t.Errorf("Expected 1 vertex, got %d", len(result.Order))
	}

	if result.DiscoveryTime[1] != 1 {
		t.Errorf("Discovery time for 1 should be 1, got %d", result.DiscoveryTime[1])
	}
}

func TestDFSIterative(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	result, err := DFSIterative(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 4 {
		t.Errorf("Expected 4 vertices in order, got %d", len(result.Order))
	}

	if result.DiscoveryTime[1] != 1 {
		t.Errorf("Discovery time for 1 should be 1, got %d", result.DiscoveryTime[1])
	}
}

func TestDFSIterativeNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := DFSIterative(g, 999)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestDFSIterativeDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)

	result, err := DFSIterative(g, 1)
	if err != nil {
		t.Fatal(err)
	}

	if _, exists := result.DiscoveryTime[4]; !exists {
		t.Error("Vertex 4 should be reachable")
	}
}

func TestDFSAllComponents(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)
	g.AddEdge(4, 5, 0)

	result := DFSAllComponents(g)

	if len(result.Order) != 5 {
		t.Errorf("Expected 5 vertices total, got %d", len(result.Order))
	}

	if len(result.DiscoveryTime) != 5 {
		t.Errorf("Expected discovery times for 5 vertices, got %d", len(result.DiscoveryTime))
	}
}

func TestDFSAllComponentsSingleComponent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	result := DFSAllComponents(g)

	if len(result.Order) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(result.Order))
	}
}

func TestDFSWithPredicate(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 4, 0)

	path, err := DFSWithPredicate(g, 1, func(v int) bool {
		return v == 4
	})

	if err != nil {
		t.Fatal(err)
	}

	if path[0] != 1 {
		t.Errorf("Path should start at 1, got %d", path[0])
	}

	if path[len(path)-1] != 4 {
		t.Errorf("Path should end at 4, got %d", path[len(path)-1])
	}
}

func TestDFSWithPredicateNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := DFSWithPredicate(g, 999, func(v int) bool { return v == 2 })
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestDFSWithPredicateStartSatisfies(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	path, err := DFSWithPredicate(g, 1, func(v int) bool { return v == 1 })

	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 1 || path[0] != 1 {
		t.Errorf("Expected path [1], got %v", path)
	}
}

func TestDFSWithPredicateNotFound(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	_, err := DFSWithPredicate(g, 1, func(v int) bool { return v == 4 })

	if err == nil {
		t.Error("Expected error when no vertex satisfies predicate")
	}
}

func TestTopologicalSort(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)

	order, err := TopologicalSort(g)
	if err != nil {
		t.Fatal(err)
	}

	if len(order) != 4 {
		t.Errorf("Expected 4 vertices in order, got %d", len(order))
	}

	for i, v := range order {
		if v == 1 {
			if i != 0 {
				t.Errorf("Vertex 1 (no incoming edges) should be first, found at position %d", i)
			}
		}
	}
}

func TestTopologicalSortUndirectedGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := TopologicalSort(g)
	if err == nil {
		t.Error("Expected error for undirected graph")
	}
}

func TestTopologicalSortWithCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)

	_, err := TopologicalSort(g)
	if err == nil {
		t.Error("Expected error for graph with cycle")
	}
}

func TestDetectCycle(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)

	hasCycle := DetectCycle(g)
	if !hasCycle {
		t.Error("Graph with cycle should be detected")
	}
}

func TestDetectCycleAcyclic(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	hasCycle := DetectCycle(g)
	if hasCycle {
		t.Error("Acyclic graph should not be detected as having cycle")
	}
}

func TestDetectCycleUndirected(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)

	hasCycle := DetectCycle(g)
	if !hasCycle {
		t.Error("Undirected graph with cycle should be detected")
	}
}

func TestConnectedComponents(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)
	g.AddEdge(4, 5, 0)

	components := ConnectedComponents(g)

	if len(components) != 2 {
		t.Errorf("Expected 2 components, got %d", len(components))
	}

	totalVertices := 0
	for _, component := range components {
		totalVertices += len(component)
	}

	if totalVertices != 5 {
		t.Errorf("Expected total of 5 vertices, got %d", totalVertices)
	}
}

func TestConnectedComponentsSingleComponent(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	components := ConnectedComponents(g)

	if len(components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(components))
	}

	if len(components[0]) != 3 {
		t.Errorf("Expected 3 vertices in single component, got %d", len(components[0]))
	}
}

func TestConnectedComponentsDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	components := ConnectedComponents(g)

	if len(components) == 0 {
		t.Error("Should return strongly connected components for directed graph")
	}
}

func TestStronglyConnectedComponents(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(3, 1, 0)
	g.AddEdge(4, 5, 0)

	sccs := StronglyConnectedComponents(g)

	if len(sccs) < 2 {
		t.Errorf("Expected at least 2 strongly connected components, got %d", len(sccs))
	}

	foundCycleSCC := false
	for _, scc := range sccs {
		if len(scc) == 3 {
			hasAll := true
			for _, v := range []int{1, 2, 3} {
				found := false
				for _, sccV := range scc {
					if sccV == v {
						found = true
						break
					}
				}
				if !found {
					hasAll = false
					break
				}
			}
			if hasAll {
				foundCycleSCC = true
			}
		}
	}

	if !foundCycleSCC {
		t.Error("Expected to find SCC with vertices [1,2,3]")
	}
}

func TestStronglyConnectedComponentsAcyclic(t *testing.T) {
	g := NewDirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	sccs := StronglyConnectedComponents(g)

	if len(sccs) != 3 {
		t.Errorf("Expected 3 SCCs (each vertex is its own SCC), got %d", len(sccs))
	}

	for _, scc := range sccs {
		if len(scc) != 1 {
			t.Errorf("Each SCC should have 1 vertex, got %d", len(scc))
		}
	}
}

func TestFindPath(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(4, 3, 0)

	path, err := FindPath(g, 1, 3)
	if err != nil {
		t.Fatal(err)
	}

	if len(path) == 0 {
		t.Error("Path should not be empty")
	}

	if path[0] != 1 {
		t.Errorf("Path should start at 1, got %d", path[0])
	}

	if path[len(path)-1] != 3 {
		t.Errorf("Path should end at 3, got %d", path[len(path)-1])
	}
}

func TestFindPathNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	_, err := FindPath(g, 999, 2)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestFindPathNonExistentEnd(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	path, err := FindPath(g, 1, 999)
	if err != nil {
		t.Fatal(err)
	}

	if path != nil {
		t.Errorf("Expected nil path for non-existent end vertex, got %v", path)
	}
}

func TestFindPathNoPath(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(3, 4, 0)

	path, err := FindPath(g, 1, 4)
	if err != nil {
		t.Fatal(err)
	}

	if path != nil {
		t.Errorf("Expected nil path when no path exists, got %v", path)
	}
}

func TestFindPathSameVertex(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)

	path, err := FindPath(g, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(path) != 1 || path[0] != 1 {
		t.Errorf("Expected path [1], got %v", path)
	}
}

func TestDFSComplexGraph(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 3, 0)
	g.AddEdge(2, 4, 0)
	g.AddEdge(3, 5, 0)

	result, _ := DFS(g, 1)

	if len(result.Order) != 5 {
		t.Errorf("Expected 5 vertices, got %d", len(result.Order))
	}

	if len(result.Parents) < 4 {
		t.Errorf("Expected at least 4 vertices to have parents, got %d", len(result.Parents))
	}
}

func TestDFSStringVertices(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 0)
	g.AddEdge("B", "C", 0)

	result, err := DFS(g, "A")
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Order) != 3 {
		t.Errorf("Expected 3 vertices, got %d", len(result.Order))
	}

	if parent, ok := result.Parents["B"]; !ok || parent != "A" {
		t.Errorf("Expected parent 'A' for vertex 'B', got %s", parent)
	}
}

func BenchmarkDFS(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 1000; i++ {
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DFS(g, 0)
	}
}

func BenchmarkDFSIterative(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 1000; i++ {
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DFSIterative(g, 0)
	}
}

func BenchmarkDFSAllComponents(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i > 0 && i%10 == 0 {
			continue
		}
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DFSAllComponents(g)
	}
}

func BenchmarkDFSWithPredicate(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DFSWithPredicate(g, 0, func(v int) bool { return v == 99 })
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	g := NewDirectedGraph[int]()
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			g.AddEdge(i, j, 0)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TopologicalSort(g)
	}
}

func BenchmarkDetectCycle(b *testing.B) {
	g := NewDirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, (i+1)%100, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DetectCycle(g)
	}
}

func BenchmarkConnectedComponents(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		if i > 0 && i%10 == 0 {
			continue
		}
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConnectedComponents(g)
	}
}

func BenchmarkStronglyConnectedComponents(b *testing.B) {
	g := NewDirectedGraph[int]()
	for i := 0; i < 50; i++ {
		g.AddEdge(i, (i+1)%50, 0)
		g.AddEdge((i+1)%50, i, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StronglyConnectedComponents(g)
	}
}

func BenchmarkFindPath(b *testing.B) {
	g := NewUndirectedGraph[int]()
	for i := 0; i < 100; i++ {
		g.AddEdge(i, i+1, 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindPath(g, 0, 99)
	}
}
