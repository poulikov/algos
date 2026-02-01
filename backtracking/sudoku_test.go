package backtracking

import (
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	grid := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solved, ok := SolveSudoku(grid)

	if !ok {
		t.Errorf("Failed to solve sudoku")
	}

	if !IsValidSudoku(solved) {
		t.Errorf("Solution is not valid")
	}

	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			if grid[row][col] != SudokuEmpty && grid[row][col] != solved[row][col] {
				t.Errorf("Changed non-empty cell at (%d,%d): %d -> %d", row, col, grid[row][col], solved[row][col])
			}
		}
	}
}

func TestIsValidSudoku(t *testing.T) {
	validGrid := SudokuGrid{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	if !IsValidSudoku(validGrid) {
		t.Errorf("Valid grid is marked as invalid")
	}

	invalidGrid := SudokuGrid{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 7},
	}

	if IsValidSudoku(invalidGrid) {
		t.Errorf("Invalid grid is marked as valid")
	}
}

func TestGenerateSudoku(t *testing.T) {
	grid := GenerateSudoku()

	if !IsValidSudoku(grid) {
		t.Errorf("Generated sudoku is not valid")
	}

	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			if grid[row][col] == SudokuEmpty {
				t.Errorf("Generated sudoku has empty cells")
			}
		}
	}
}

func TestGenerateSudokuPuzzle(t *testing.T) {
	clues := 30
	puzzle := GenerateSudokuPuzzle(clues)

	clueCount := 0
	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			if puzzle[row][col] != SudokuEmpty {
				clueCount++
			}
		}
	}

	if clueCount != clues {
		t.Errorf("Expected %d clues, got %d", clues, clueCount)
	}

	solved, ok := SolveSudoku(puzzle)
	if !ok {
		t.Errorf("Generated puzzle is not solvable")
	}

	if !IsValidSudoku(solved) {
		t.Errorf("Solution to generated puzzle is not valid")
	}
}

func TestCopyGrid(t *testing.T) {
	original := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	copy := CopyGrid(original)

	original[0][0] = 9

	if copy[0][0] == 9 {
		t.Errorf("Copy is not deep")
	}

	if copy[0][0] != 5 {
		t.Errorf("Copy does not match original")
	}
}

func TestGridToString(t *testing.T) {
	grid := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	str := GridToString(grid)

	if len(str) == 0 {
		t.Errorf("GridToString returned empty string")
	}
}

func TestStringToGrid(t *testing.T) {
	str := "5 3 . . 7 . . . .\n6 . . 1 9 5 . . .\n. 9 8 . . . . 6 .\n8 . . . 6 . . . 3\n4 . . 8 . 3 . . 1\n7 . . . 2 . . . 6\n. 6 . . . . 2 8 .\n. . . 4 1 9 . . 5\n. . . . 8 . . 7 9\n"

	grid := StringToGrid(str)

	if grid[0][0] != 5 {
		t.Errorf("StringToGrid did not parse correctly")
	}

	if grid[0][2] != SudokuEmpty {
		t.Errorf("StringToGrid did not parse empty cells correctly")
	}
}

func TestSudokuSolver(t *testing.T) {
	grid := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	ss := NewSudokuSolver(grid)

	if !ss.Solve() {
		t.Errorf("SudokuSolver failed to solve")
	}

	solved := ss.GetGrid()

	if !IsValidSudoku(solved) {
		t.Errorf("Solution is not valid")
	}
}

func TestHasUniqueSolution(t *testing.T) {
	grid := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if !HasUniqueSolution(grid) {
		t.Errorf("Expected unique solution")
	}
}

func TestEmptyGrid(t *testing.T) {
	grid := SudokuGrid{}

	if !IsValidSudoku(grid) {
		t.Errorf("Empty grid should be valid (no conflicts)")
	}
}

func TestSudokuConsistency(t *testing.T) {
	grid := SudokuGrid{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	solved1, ok1 := SolveSudoku(grid)
	ss := NewSudokuSolver(grid)
	solved2 := ss.GetGrid()
	ss.Solve()
	solved2 = ss.GetGrid()

	if !ok1 {
		t.Errorf("First solve failed")
	}

	if !IsValidSudoku(solved1) {
		t.Errorf("First solution is not valid")
	}

	if !IsValidSudoku(solved2) {
		t.Errorf("Second solution is not valid")
	}
}
