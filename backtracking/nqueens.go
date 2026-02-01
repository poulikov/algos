package backtracking

type Position struct {
	Row int
	Col int
}

type Solution []Position

func (s Solution) ToGrid(n int) [][]bool {
	grid := make([][]bool, n)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	for _, pos := range s {
		grid[pos.Row][pos.Col] = true
	}
	return grid
}

func SolveNQueens(n int) []Solution {
	if n <= 0 {
		return []Solution{}
	}

	var solutions []Solution
	queens := make([]Position, 0, n)

	nQueensHelper(n, 0, queens, &solutions)

	return solutions
}

func nQueensHelper(n, row int, queens []Position, solutions *[]Solution) {
	if row == n {
		solution := make(Solution, len(queens))
		copy(solution, queens)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(queens, row, col) {
			queens = append(queens, Position{row, col})
			nQueensHelper(n, row+1, queens, solutions)
			queens = queens[:len(queens)-1]
		}
	}
}

func isSafe(queens []Position, row, col int) bool {
	for _, q := range queens {
		if q.Col == col {
			return false
		}
		if q.Row-q.Col == row-col {
			return false
		}
		if q.Row+q.Col == row+col {
			return false
		}
	}
	return true
}

func CountSolutions(n int) int {
	return len(SolveNQueens(n))
}

func HasSolution(n int) bool {
	return CountSolutions(n) > 0
}

func FindOneSolution(n int) Solution {
	solutions := SolveNQueens(n)
	if len(solutions) == 0 {
		return nil
	}
	return solutions[0]
}

func PrintSolution(solution Solution, n int) [][]string {
	grid := make([][]string, n)
	for i := range grid {
		grid[i] = make([]string, n)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, pos := range solution {
		grid[pos.Row][pos.Col] = "Q"
	}

	return grid
}

func IsValidSolution(solution Solution, n int) bool {
	if len(solution) != n {
		return false
	}

	for i, pos1 := range solution {
		if pos1.Row != i {
			return false
		}
		for j, pos2 := range solution {
			if i != j {
				if pos1.Col == pos2.Col {
					return false
				}
				if abs(pos1.Row-pos2.Row) == abs(pos1.Col-pos2.Col) {
					return false
				}
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func NQueensWithConstraints(n int, forbidden []Position) []Solution {
	var solutions []Solution
	queens := make([]Position, 0, n)

	nQueensWithConstraintsHelper(n, 0, queens, forbidden, &solutions)

	return solutions
}

func nQueensWithConstraintsHelper(n, row int, queens, forbidden []Position, solutions *[]Solution) {
	if row == n {
		solution := make(Solution, len(queens))
		copy(solution, queens)
		*solutions = append(*solutions, solution)
		return
	}

	for col := 0; col < n; col++ {
		if !isForbidden(row, col, forbidden) && isSafe(queens, row, col) {
			queens = append(queens, Position{row, col})
			nQueensWithConstraintsHelper(n, row+1, queens, forbidden, solutions)
			queens = queens[:len(queens)-1]
		}
	}
}

func isForbidden(row, col int, forbidden []Position) bool {
	for _, pos := range forbidden {
		if pos.Row == row && pos.Col == col {
			return true
		}
	}
	return false
}
