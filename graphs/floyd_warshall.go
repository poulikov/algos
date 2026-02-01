package graphs

import (
	"fmt"
	"math"
)

// FloydWarshallResult represents the result of Floyd-Warshall algorithm
type FloydWarshallResult[T comparable] struct {
	Distances        map[T]map[T]float64 // Distance matrix: distance from i to j
	Next             map[T]map[T]T       // Next vertex in shortest path (for path reconstruction)
	Vertices         []T                 // List of all vertices
	VertexIndex      map[T]int           // Map from vertex to its index
	HasNegativeCycle bool                // Whether a negative cycle exists
}

// FloydWarshall computes all-pairs shortest paths
// Works with negative edge weights (but not negative cycles)
// Time complexity: O(V^3)
func FloydWarshall[T comparable](graph *Graph[T]) (*FloydWarshallResult[T], error) {
	vertices := graph.GetVertices()
	n := len(vertices)

	if n == 0 {
		return &FloydWarshallResult[T]{
			Distances:   make(map[T]map[T]float64),
			Next:        make(map[T]map[T]T),
			Vertices:    []T{},
			VertexIndex: make(map[T]int),
		}, nil
	}

	result := &FloydWarshallResult[T]{
		Distances:   make(map[T]map[T]float64),
		Next:        make(map[T]map[T]T),
		Vertices:    vertices,
		VertexIndex: make(map[T]int),
	}

	for i, v := range vertices {
		result.VertexIndex[v] = i
		result.Distances[v] = make(map[T]float64)
		result.Next[v] = make(map[T]T)

		for _, u := range vertices {
			result.Distances[v][u] = math.Inf(1)
		}

		result.Distances[v][v] = 0
	}

	for _, edge := range graph.GetEdges() {
		if edge.Weight < result.Distances[edge.From][edge.To] {
			result.Distances[edge.From][edge.To] = edge.Weight
			result.Next[edge.From][edge.To] = edge.To
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if result.Distances[vertices[i]][vertices[k]]+result.Distances[vertices[k]][vertices[j]] < result.Distances[vertices[i]][vertices[j]] {
					result.Distances[vertices[i]][vertices[j]] = result.Distances[vertices[i]][vertices[k]] + result.Distances[vertices[k]][vertices[j]]
					result.Next[vertices[i]][vertices[j]] = result.Next[vertices[i]][vertices[k]]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if result.Distances[vertices[i]][vertices[i]] < 0 {
			result.HasNegativeCycle = true
			break
		}
	}

	return result, nil
}

// GetPath reconstructs the shortest path between two vertices
// Returns the path as a slice of vertices
// Time complexity: O(V)
func (fw *FloydWarshallResult[T]) GetPath(from, to T) ([]T, error) {
	_, fromExists := fw.VertexIndex[from]
	_, toExists := fw.VertexIndex[to]

	if !fromExists {
		return nil, fmt.Errorf("from vertex not found")
	}

	if !toExists {
		return nil, fmt.Errorf("to vertex not found")
	}

	_, hasPath := fw.Next[from][to]
	if !hasPath {
		if from == to {
			return []T{from}, nil
		}
		return nil, fmt.Errorf("no path exists from %v to %v", from, to)
	}

	path := []T{from}
	current := from

	for current != to {
		current = fw.Next[current][to]
		path = append(path, current)
	}

	return path, nil
}

// Time complexity: O(1)
func (fw *FloydWarshallResult[T]) GetDistance(from, to T) (float64, error) {
	_, fromExists := fw.VertexIndex[from]
	_, toExists := fw.VertexIndex[to]

	if !fromExists {
		return 0, fmt.Errorf("from vertex not found")
	}

	if !toExists {
		return 0, fmt.Errorf("to vertex not found")
	}

	return fw.Distances[from][to], nil
}

// GetAllPairsDistances returns the complete distance matrix
// Returns a 2D slice where matrix[i][j] is distance from vertices[i] to vertices[j]
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetAllPairsDistances() [][]float64 {
	n := len(fw.Vertices)
	matrix := make([][]float64, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = fw.Distances[fw.Vertices[i]][fw.Vertices[j]]
		}
	}

	return matrix
}

// GetCenters finds all centers of the graph (vertices with minimum eccentricity)
// A center has the minimum maximum distance to any other vertex
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetCenters() []T {
	if len(fw.Vertices) == 0 {
		return []T{}
	}

	minEccentricity := math.Inf(1)
	centers := []T{}

	for _, v := range fw.Vertices {
		maxDist := 0.0
		for _, u := range fw.Vertices {
			if fw.Distances[v][u] > maxDist && !math.IsInf(fw.Distances[v][u], 1) {
				maxDist = fw.Distances[v][u]
			}
		}

		if maxDist < minEccentricity {
			minEccentricity = maxDist
			centers = []T{v}
		} else if maxDist == minEccentricity {
			centers = append(centers, v)
		}
	}

	return centers
}

// GetPeriphery finds all peripheral vertices (vertices with maximum eccentricity)
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetPeriphery() []T {
	if len(fw.Vertices) == 0 {
		return []T{}
	}

	maxEccentricity := 0.0
	periphery := []T{}

	for _, v := range fw.Vertices {
		maxDist := 0.0
		for _, u := range fw.Vertices {
			if fw.Distances[v][u] > maxDist && !math.IsInf(fw.Distances[v][u], 1) {
				maxDist = fw.Distances[v][u]
			}
		}

		if maxDist > maxEccentricity {
			maxEccentricity = maxDist
			periphery = []T{v}
		} else if maxDist == maxEccentricity {
			periphery = append(periphery, v)
		}
	}

	return periphery
}

// GetDiameter returns the diameter of the graph (maximum shortest path)
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetDiameter() float64 {
	if len(fw.Vertices) == 0 {
		return 0
	}

	maxDist := 0.0

	for _, v := range fw.Vertices {
		for _, u := range fw.Vertices {
			if fw.Distances[v][u] > maxDist && !math.IsInf(fw.Distances[v][u], 1) {
				maxDist = fw.Distances[v][u]
			}
		}
	}

	return maxDist
}

// GetRadius returns the radius of the graph (minimum eccentricity)
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetRadius() float64 {
	if len(fw.Vertices) == 0 {
		return 0
	}

	minEccentricity := math.Inf(1)

	for _, v := range fw.Vertices {
		maxDist := 0.0
		for _, u := range fw.Vertices {
			if fw.Distances[v][u] > maxDist && !math.IsInf(fw.Distances[v][u], 1) {
				maxDist = fw.Distances[v][u]
			}
		}

		if maxDist < minEccentricity {
			minEccentricity = maxDist
		}
	}

	return minEccentricity
}

// IsTransitiveClosure checks if the distance matrix satisfies the transitive property
// For all vertices i, j, k: dist[i][k] <= dist[i][j] + dist[j][k]
// Time complexity: O(V^3)
func (fw *FloydWarshallResult[T]) IsTransitiveClosure() bool {
	n := len(fw.Vertices)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				distIK := fw.Distances[fw.Vertices[i]][fw.Vertices[k]]
				distIJ := fw.Distances[fw.Vertices[i]][fw.Vertices[j]]
				distJK := fw.Distances[fw.Vertices[j]][fw.Vertices[k]]

				pathIJ := !math.IsInf(distIJ, 1)
				pathJK := !math.IsInf(distJK, 1)
				pathIK := !math.IsInf(distIK, 1)

				if pathIJ && pathJK && !pathIK {
					return false
				}

				if pathIJ && pathJK && pathIK {
					if distIK > distIJ+distJK {
						return false
					}
				}
			}
		}
	}

	return true
}

// GetTransitiveClosure returns the transitive closure matrix
// Returns a 2D slice where matrix[i][j] is true if path exists from vertices[i] to vertices[j]
// Time complexity: O(V^2)
func (fw *FloydWarshallResult[T]) GetTransitiveClosure() [][]bool {
	n := len(fw.Vertices)
	matrix := make([][]bool, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = !math.IsInf(fw.Distances[fw.Vertices[i]][fw.Vertices[j]], 1)
		}
	}

	return matrix
}
