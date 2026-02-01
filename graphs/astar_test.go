package graphs

import (
	"math"
	"testing"
)

func TestAStar(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 4)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "C", 1)
	g.AddEdge("B", "D", 5)
	g.AddEdge("C", "D", 8)

	heuristic := func(v string) float64 {
		distances := map[string]float64{
			"A": 5,
			"B": 5,
			"C": 8,
			"D": 0,
		}
		return distances[v]
	}

	result, err := AStar(g, "A", "D", heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Path[0] != "A" || result.Path[len(result.Path)-1] != "D" {
		t.Errorf("Path should start at A and end at D, got %v", result.Path)
	}

	if result.Cost != 9 {
		t.Errorf("Expected cost 9, got %f", result.Cost)
	}
}

func TestAStarNonExistentStart(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStar(g, "Z", "A", heuristic)
	if err == nil {
		t.Error("Expected error for non-existent start vertex")
	}
}

func TestAStarNonExistentEnd(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStar(g, "A", "Z", heuristic)
	if err == nil {
		t.Error("Expected error for non-existent end vertex")
	}
}

func TestAStarNegativeWeight(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", -1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStar(g, "A", "B", heuristic)
	if err == nil {
		t.Error("Expected error for negative edge weight")
	}
}

func TestAStarNoPath(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("C", "D", 1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStar(g, "A", "D", heuristic)
	if err == nil {
		t.Error("Expected error when no path exists")
	}
}

func TestAStarSameVertex(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddVertex("A")

	heuristic := func(v string) float64 { return 0 }

	result, err := AStar(g, "A", "A", heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Path) != 1 || result.Path[0] != "A" {
		t.Errorf("Expected path [A], got %v", result.Path)
	}

	if result.Cost != 0 {
		t.Errorf("Expected cost 0, got %f", result.Cost)
	}
}

func TestAStarGrid(t *testing.T) {
	g := NewUndirectedGraph[string]()

	edges := [][3]string{
		{"A", "B"}, {"A", "C"}, {"B", "D"}, {"B", "E"},
		{"C", "F"}, {"C", "G"}, {"D", "H"}, {"E", "H"},
		{"F", "I"}, {"G", "I"}, {"H", "I"},
	}

	for _, e := range edges {
		g.AddEdge(e[0], e[1], 1)
	}

	heuristic := func(v string) float64 {
		distances := map[string]float64{
			"A": 2, "B": 1, "C": 1, "D": 1, "E": 1,
			"F": 1, "G": 1, "H": 0.5, "I": 0,
		}
		return distances[v]
	}

	result, err := AStar(g, "A", "I", heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Path[0] != "A" || result.Path[len(result.Path)-1] != "I" {
		t.Errorf("Path should start at A and end at I, got %v", result.Path)
	}
}

func TestAStarWithLimit(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 5, 1)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 4, 2: 3, 3: 2, 4: 1, 5: 0,
		}
		return distances[v]
	}

	result, err := AStarWithLimit(g, 1, 5, heuristic, 2)
	if err == nil {
		t.Error("Expected error when iteration limit is reached")
	}

	if result != nil {
		t.Error("Should return nil result when limit is reached")
	}
}

func TestAStarWithLimitSuccess(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 2, 2: 1, 3: 0,
		}
		return distances[v]
	}

	result, err := AStarWithLimit(g, 1, 3, heuristic, 10)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Path) != 3 {
		t.Errorf("Expected path length 3, got %d", len(result.Path))
	}
}

func TestAStarMultipleTargets(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 2)
	g.AddEdge("B", "C", 2)
	g.AddEdge("C", "D", 2)
	g.AddEdge("A", "D", 10)

	heuristic := func(v string) float64 {
		distances := map[string]float64{
			"A": 2, "B": 1, "C": 1, "D": 0,
		}
		return distances[v]
	}

	result, err := AStarMultipleTargets(g, "A", []string{"C", "D"}, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Path[len(result.Path)-1] != "C" && result.Path[len(result.Path)-1] != "D" {
		t.Errorf("Path should end at C or D, got %v", result.Path)
	}

	if result.Cost != 4 {
		t.Errorf("Expected cost 4, got %f", result.Cost)
	}
}

func TestAStarMultipleTargetsNoTargets(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStarMultipleTargets(g, "A", []string{}, heuristic)
	if err == nil {
		t.Error("Expected error for empty targets slice")
	}
}

func TestAStarMultipleTargetsNonExistent(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 1)

	heuristic := func(v string) float64 { return 0 }

	_, err := AStarMultipleTargets(g, "A", []string{"Z"}, heuristic)
	if err == nil {
		t.Error("Expected error for non-existent target vertex")
	}
}

func TestAStarWithReconstruction(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(1, 3, 5)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 2, 2: 1, 3: 0,
		}
		return distances[v]
	}

	reconstruct := func(parents map[int]int, start, end int) []int {
		path := []int{end}
		node := end
		for node != start {
			parent, exists := parents[node]
			if !exists {
				break
			}
			path = append([]int{parent}, path...)
			node = parent
		}
		return path
	}

	result, err := AStarWithReconstruction(g, 1, 3, heuristic, reconstruct)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Path) != 3 {
		t.Errorf("Expected path length 3, got %d", len(result.Path))
	}

	if result.Path[0] != 1 || result.Path[2] != 3 {
		t.Errorf("Path should be [1, 2, 3], got %v", result.Path)
	}

	if result.Cost != 2 {
		t.Errorf("Expected cost 2, got %f", result.Cost)
	}
}

func TestAStarZeroHeuristic(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 3, 1)
	g.AddEdge(2, 4, 5)
	g.AddEdge(3, 4, 8)

	heuristic := func(v int) float64 { return 0 }

	result, err := AStar(g, 1, 4, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Cost != 8 {
		t.Errorf("Expected cost 8 with zero heuristic, got %f", result.Cost)
	}
}

func TestAStarPathValidity(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 4)
	g.AddEdge(3, 4, 5)
	g.AddEdge(1, 4, 20)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 0, 2: 7, 3: 4, 4: 0,
		}
		return distances[v]
	}

	result, err := AStar(g, 1, 4, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if !isValidPath(g, result.Path) {
		t.Errorf("Path %v is not valid", result.Path)
	}
}

func TestAStarVisited(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 2, 2: 1, 3: 0,
		}
		return distances[v]
	}

	result, err := AStar(g, 1, 3, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Visited) == 0 {
		t.Error("Visited vertices should not be empty")
	}

	if result.Iterations == 0 {
		t.Error("Iterations should be greater than 0")
	}
}

func TestAStarDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[string]()
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 1)
	g.AddEdge("C", "A", 10)

	heuristic := func(v string) float64 {
		distances := map[string]float64{
			"A": 2, "B": 1, "C": 0,
		}
		return distances[v]
	}

	result, err := AStar(g, "A", "C", heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Cost != 2 {
		t.Errorf("Expected cost 2 in directed graph, got %f", result.Cost)
	}
}

func TestAStarMultipleEqualPaths(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 4, 1)
	g.AddEdge(3, 4, 1)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 2, 2: 1, 3: 1, 4: 0,
		}
		return distances[v]
	}

	result, err := AStar(g, 1, 4, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Cost != 2 {
		t.Errorf("Expected cost 2, got %f", result.Cost)
	}

	if !isValidPath(g, result.Path) {
		t.Errorf("Path %v is not valid", result.Path)
	}
}

func TestAStarSingleEdge(t *testing.T) {
	g := NewUndirectedGraph[string]()
	g.AddEdge("A", "B", 5)

	heuristic := func(v string) float64 {
		if v == "A" {
			return 5
		}
		return 0
	}

	result, err := AStar(g, "A", "B", heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Path) != 2 {
		t.Errorf("Expected path length 2, got %d", len(result.Path))
	}

	if result.Path[0] != "A" || result.Path[1] != "B" {
		t.Errorf("Expected path [A, B], got %v", result.Path)
	}

	if result.Cost != 5 {
		t.Errorf("Expected cost 5, got %f", result.Cost)
	}
}

func TestAStarOptimalHeuristic(t *testing.T) {
	g := NewUndirectedGraph[int]()
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 4)
	g.AddEdge(1, 4, 5)
	g.AddEdge(4, 3, 6)

	heuristic := func(v int) float64 {
		distances := map[int]float64{
			1: 7, 2: 4, 3: 0, 4: 6,
		}
		return distances[v]
	}

	result, err := AStar(g, 1, 3, heuristic)
	if err != nil {
		t.Fatal(err)
	}

	if result.Cost != 7 {
		t.Errorf("Expected cost 7, got %f", result.Cost)
	}
}

func isValidPath[T comparable](g *Graph[T], path []T) bool {
	if len(path) < 1 {
		return false
	}

	for i := 0; i < len(path)-1; i++ {
		edges := g.GetOutgoingEdges(path[i])
		found := false
		for _, edge := range edges {
			if edge.To == path[i+1] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func BenchmarkAStar(b *testing.B) {
	g := createTestGraph(100)
	heuristic := func(v int) float64 { return math.Abs(float64(v - 99)) }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AStar(g, 0, 99, heuristic)
	}
}

func BenchmarkAStarWithLimit(b *testing.B) {
	g := createTestGraph(100)
	heuristic := func(v int) float64 { return math.Abs(float64(v - 99)) }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AStarWithLimit(g, 0, 99, heuristic, 1000)
	}
}

func BenchmarkAStarMultipleTargets(b *testing.B) {
	g := createTestGraph(100)
	heuristic := func(v int) float64 { return math.Abs(float64(v - 99)) }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AStarMultipleTargets(g, 0, []int{50, 99}, heuristic)
	}
}

func createTestGraph(size int) *Graph[int] {
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
