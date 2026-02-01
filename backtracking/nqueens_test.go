package backtracking

import (
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, test := range tests {
		solutions := SolveNQueens(test.n)
		if len(solutions) != test.expected {
			t.Errorf("SolveNQueens(%d) returned %d solutions, expected %d", test.n, len(solutions), test.expected)
		}
		for _, solution := range solutions {
			if !IsValidSolution(solution, test.n) {
				t.Errorf("Invalid solution for N=%d", test.n)
			}
		}
	}
}

func TestCountSolutions(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{4, 2},
		{5, 10},
		{8, 92},
	}

	for _, test := range tests {
		count := CountSolutions(test.n)
		if count != test.expected {
			t.Errorf("CountSolutions(%d) = %d, expected %d", test.n, count, test.expected)
		}
	}
}

func TestHasSolution(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{0, false},
		{1, true},
		{2, false},
		{3, false},
		{4, true},
		{5, true},
		{8, true},
	}

	for _, test := range tests {
		has := HasSolution(test.n)
		if has != test.expected {
			t.Errorf("HasSolution(%d) = %v, expected %v", test.n, has, test.expected)
		}
	}
}

func TestFindOneSolution(t *testing.T) {
	tests := []struct {
		n int
	}{
		{1},
		{4},
		{8},
	}

	for _, test := range tests {
		solution := FindOneSolution(test.n)
		if solution == nil && test.n > 0 {
			t.Errorf("FindOneSolution(%d) returned nil", test.n)
		}
		if solution != nil && !IsValidSolution(solution, test.n) {
			t.Errorf("FindOneSolution(%d) returned invalid solution", test.n)
		}
	}

	solution := FindOneSolution(2)
	if solution != nil {
		t.Errorf("FindOneSolution(2) should return nil, got %v", solution)
	}
}

func TestIsValidSolution(t *testing.T) {
	solution := Solution{
		{0, 1},
		{1, 3},
		{2, 0},
		{3, 2},
	}

	if !IsValidSolution(solution, 4) {
		t.Errorf("Valid solution is marked as invalid")
	}

	invalidSolution := Solution{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 3},
	}

	if IsValidSolution(invalidSolution, 4) {
		t.Errorf("Invalid solution is marked as valid")
	}
}

func TestPrintSolution(t *testing.T) {
	solution := Solution{
		{0, 1},
		{1, 3},
		{2, 0},
		{3, 2},
	}

	grid := PrintSolution(solution, 4)

	if len(grid) != 4 {
		t.Errorf("Expected grid with 4 rows, got %d", len(grid))
	}

	queenCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "Q" {
				queenCount++
			}
		}
	}

	if queenCount != 4 {
		t.Errorf("Expected 4 queens, got %d", queenCount)
	}
}

func TestSolutionToGrid(t *testing.T) {
	solution := Solution{
		{0, 1},
		{1, 3},
		{2, 0},
		{3, 2},
	}

	grid := solution.ToGrid(4)

	if len(grid) != 4 {
		t.Errorf("Expected grid with 4 rows, got %d", len(grid))
	}

	trueCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				trueCount++
			}
		}
	}

	if trueCount != 4 {
		t.Errorf("Expected 4 true values, got %d", trueCount)
	}
}

func TestNQueensWithConstraints(t *testing.T) {
	n := 4
	forbidden := []Position{{0, 0}, {0, 1}}

	solutions := NQueensWithConstraints(n, forbidden)

	for _, solution := range solutions {
		if !IsValidSolution(solution, n) {
			t.Errorf("Invalid solution with constraints")
		}
		for _, pos := range forbidden {
			for _, queen := range solution {
				if queen.Row == pos.Row && queen.Col == pos.Col {
					t.Errorf("Solution includes forbidden position")
				}
			}
		}
	}
}

func TestIsSafe(t *testing.T) {
	queens := []Position{
		{0, 1},
		{1, 3},
	}

	if !isSafe(queens, 2, 0) {
		t.Errorf("Position (2,0) should be safe")
	}

	if isSafe(queens, 2, 1) {
		t.Errorf("Position (2,1) should be unsafe (same column)")
	}

	if isSafe(queens, 2, 2) {
		t.Errorf("Position (2,2) should be unsafe (same diagonal)")
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x        int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-100, 100},
	}

	for _, test := range tests {
		result := abs(test.x)
		if result != test.expected {
			t.Errorf("abs(%d) = %d, expected %d", test.x, result, test.expected)
		}
	}
}

func TestNQueensConsistency(t *testing.T) {
	n := 8

	solutions := SolveNQueens(n)
	count := CountSolutions(n)

	if len(solutions) != count {
		t.Errorf("Inconsistent: SolveNQueens returned %d, CountSolutions returned %d", len(solutions), count)
	}

	for i, solution := range solutions {
		if !IsValidSolution(solution, n) {
			t.Errorf("Solution %d is invalid", i)
		}

		grid := PrintSolution(solution, n)
		if len(grid) != n {
			t.Errorf("Solution %d has invalid grid size", i)
		}
	}
}

func TestNQueensLarge(t *testing.T) {
	n := 10
	solutions := SolveNQueens(n)

	if len(solutions) == 0 {
		t.Errorf("Expected at least one solution for N=%d", n)
	}

	for i, solution := range solutions {
		if !IsValidSolution(solution, n) {
			t.Errorf("Solution %d is invalid", i)
		}
	}
}
