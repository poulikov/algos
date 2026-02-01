package graphs

import (
	"testing"
)

func TestIsTransitiveClosure(t *testing.T) {
	tests := []struct {
		name  string
		edges []Edge[int]
		want  bool
	}{
		{
			name: "Complete graph (transitive)",
			edges: []Edge[int]{
				{1, 2, 1},
				{2, 3, 1},
				{1, 3, 1},
			},
			want: true,
		},
		{
			name: "Transitive - triangle inequality satisfied",
			edges: []Edge[int]{
				{1, 2, 1},
				{2, 3, 1},
				{1, 3, 2},
			},
			want: true,
		},
		{
			name: "Not transitive - violates triangle inequality with negative edge",
			edges: []Edge[int]{
				{1, 2, 1},
				{2, 3, 1},
				{1, 3, 4},
				{3, 1, -5},
			},
			want: false,
		},
		{
			name: "Single vertex (transitive)",
			edges: []Edge[int]{
				{1, 1, 0},
			},
			want: true,
		},
		{
			name:  "Isolated vertices",
			edges: []Edge[int]{},
			want:  true,
		},
		{
			name: "Triangle with equal weights",
			edges: []Edge[int]{
				{1, 2, 1},
				{2, 3, 1},
				{3, 1, 1},
			},
			want: true,
		},
		{
			name: "Complex transitive graph",
			edges: []Edge[int]{
				{1, 2, 1},
				{2, 3, 1},
				{3, 4, 1},
				{1, 3, 2},
				{2, 4, 2},
				{1, 4, 3},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := NewDirectedGraph[int]()
			for _, edge := range tt.edges {
				graph.AddEdge(edge.From, edge.To, edge.Weight)
			}

			fw, _ := FloydWarshall(graph)
			got := fw.IsTransitiveClosure()

			if got != tt.want {
				t.Errorf("IsTransitiveClosure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTransitiveClosure(t *testing.T) {
	graph := NewDirectedGraph[int]()
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(2, 3, 1)
	graph.AddEdge(1, 3, 1)

	fw, _ := FloydWarshall(graph)
	matrix := fw.GetTransitiveClosure()

	vertexToIndex := make(map[int]int)
	for i, v := range fw.Vertices {
		vertexToIndex[v] = i
	}

	testCases := []struct {
		from     int
		to       int
		expected bool
	}{
		{1, 1, true},
		{1, 2, true},
		{1, 3, true},
		{2, 1, false},
		{2, 2, true},
		{2, 3, true},
		{3, 1, false},
		{3, 2, false},
		{3, 3, true},
	}

	for _, tc := range testCases {
		fromIdx := vertexToIndex[tc.from]
		toIdx := vertexToIndex[tc.to]
		got := matrix[fromIdx][toIdx]

		if got != tc.expected {
			t.Errorf("GetTransitiveClosure()[%d][%d] = %v, want %v (from %d to %d)", fromIdx, toIdx, got, tc.expected, tc.from, tc.to)
		}
	}
}

func TestGetTransitiveClosureSingleEdge(t *testing.T) {
	graph := NewDirectedGraph[int]()
	graph.AddEdge(1, 2, 1)

	fw, _ := FloydWarshall(graph)
	matrix := fw.GetTransitiveClosure()

	if len(matrix) != 2 {
		t.Fatalf("Expected 2 vertices, got %d", len(matrix))
	}

	vertexToIndex := make(map[int]int)
	for i, v := range fw.Vertices {
		vertexToIndex[v] = i
	}

	testCases := []struct {
		from     int
		to       int
		expected bool
	}{
		{1, 1, true},
		{1, 2, true},
		{2, 1, false},
		{2, 2, true},
	}

	for _, tc := range testCases {
		fromIdx := vertexToIndex[tc.from]
		toIdx := vertexToIndex[tc.to]
		got := matrix[fromIdx][toIdx]

		if got != tc.expected {
			t.Errorf("GetTransitiveClosure()[%d][%d] = %v, want %v (from %d to %d)", fromIdx, toIdx, got, tc.expected, tc.from, tc.to)
		}
	}
}
