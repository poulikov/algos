package backtracking

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	SudokuSize    = 9
	SudokuBoxSize = 3
	SudokuEmpty   = 0
)

type SudokuGrid [SudokuSize][SudokuSize]int

type SudokuSolver struct {
	grid SudokuGrid
}

func NewSudokuSolver(grid SudokuGrid) *SudokuSolver {
	return &SudokuSolver{grid: grid}
}

func (ss *SudokuSolver) Solve() bool {
	return ss.solve(0, 0)
}

func (ss *SudokuSolver) solve(row, col int) bool {
	if row == SudokuSize {
		return true
	}

	nextRow, nextCol := ss.nextCell(row, col)

	if ss.grid[row][col] != SudokuEmpty {
		return ss.solve(nextRow, nextCol)
	}

	for num := 1; num <= SudokuSize; num++ {
		if ss.isValid(row, col, num) {
			ss.grid[row][col] = num
			if ss.solve(nextRow, nextCol) {
				return true
			}
			ss.grid[row][col] = SudokuEmpty
		}
	}

	return false
}

func (ss *SudokuSolver) nextCell(row, col int) (int, int) {
	col++
	if col == SudokuSize {
		col = 0
		row++
	}
	return row, col
}

func (ss *SudokuSolver) isValid(row, col, num int) bool {
	for i := 0; i < SudokuSize; i++ {
		if ss.grid[row][i] == num {
			return false
		}
	}

	for i := 0; i < SudokuSize; i++ {
		if ss.grid[i][col] == num {
			return false
		}
	}

	boxRow := (row / SudokuBoxSize) * SudokuBoxSize
	boxCol := (col / SudokuBoxSize) * SudokuBoxSize

	for i := 0; i < SudokuBoxSize; i++ {
		for j := 0; j < SudokuBoxSize; j++ {
			if ss.grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

func (ss *SudokuSolver) GetGrid() SudokuGrid {
	return ss.grid
}

func (ss *SudokuSolver) SetGrid(grid SudokuGrid) {
	ss.grid = grid
}

func SolveSudoku(grid SudokuGrid) (SudokuGrid, bool) {
	ss := NewSudokuSolver(grid)
	solved := ss.Solve()
	if solved {
		return ss.GetGrid(), true
	}
	return grid, false
}

func IsValidSudoku(grid SudokuGrid) bool {
	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			num := grid[row][col]
			if num != SudokuEmpty {
				grid[row][col] = SudokuEmpty
				if !NewSudokuSolver(grid).isValid(row, col, num) {
					grid[row][col] = num
					return false
				}
				grid[row][col] = num
			}
		}
	}
	return true
}

func GenerateSudoku() SudokuGrid {
	rand.Seed(time.Now().UnixNano())

	grid := SudokuGrid{}
	solved := false

	for !solved {
		grid = SudokuGrid{}
		fillDiagonalBoxes(&grid)
		ss := NewSudokuSolver(grid)
		solved = ss.solve(0, 0)
		grid = ss.GetGrid()
	}

	return grid
}

func GenerateSudokuPuzzle(clues int) SudokuGrid {
	solution := GenerateSudoku()
	puzzle := solution

	cells := make([]struct {
		row int
		col int
	}, SudokuSize*SudokuSize)

	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			cells[row*SudokuSize+col] = struct {
				row int
				col int
			}{row, col}
		}
	}

	rand.Shuffle(len(cells), func(i, j int) {
		cells[i], cells[j] = cells[j], cells[i]
	})

	for i := 0; i < len(cells)-clues; i++ {
		puzzle[cells[i].row][cells[i].col] = SudokuEmpty
	}

	return puzzle
}

func fillDiagonalBoxes(grid *SudokuGrid) {
	for i := 0; i < SudokuSize; i += SudokuBoxSize {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rand.Shuffle(len(nums), func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})

		for r := 0; r < SudokuBoxSize; r++ {
			for c := 0; c < SudokuBoxSize; c++ {
				grid[i+r][i+c] = nums[r*SudokuBoxSize+c]
			}
		}
	}
}

func CountSudokuSolutions(grid SudokuGrid) int {
	ss := NewSudokuSolver(grid)
	return ss.countSolutions(0, 0)
}

func (ss *SudokuSolver) countSolutions(row, col int) int {
	if row == SudokuSize {
		return 1
	}

	nextRow, nextCol := ss.nextCell(row, col)

	if ss.grid[row][col] != SudokuEmpty {
		return ss.countSolutions(nextRow, nextCol)
	}

	count := 0
	for num := 1; num <= SudokuSize; num++ {
		if ss.isValid(row, col, num) {
			ss.grid[row][col] = num
			count += ss.countSolutions(nextRow, nextCol)
			ss.grid[row][col] = SudokuEmpty
		}
	}

	return count
}

func HasUniqueSolution(grid SudokuGrid) bool {
	return CountSudokuSolutions(grid) == 1
}

func CopyGrid(grid SudokuGrid) SudokuGrid {
	var copy SudokuGrid
	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			copy[row][col] = grid[row][col]
		}
	}
	return copy
}

func GridToString(grid SudokuGrid) string {
	str := ""
	for row := 0; row < SudokuSize; row++ {
		for col := 0; col < SudokuSize; col++ {
			if grid[row][col] == SudokuEmpty {
				str += ". "
			} else {
				str += strconv.Itoa(grid[row][col]) + " "
			}
			if (col+1)%SudokuBoxSize == 0 && col < SudokuSize-1 {
				str += "| "
			}
		}
		str += "\n"
		if (row+1)%SudokuBoxSize == 0 && row < SudokuSize-1 {
			str += "------+-------+------\n"
		}
	}
	return str
}

func StringToGrid(str string) SudokuGrid {
	var grid SudokuGrid
	row := 0
	col := 0

	for _, ch := range str {
		if ch >= '1' && ch <= '9' {
			grid[row][col] = int(ch - '0')
			col++
		} else if ch == '.' || ch == '0' || ch == ' ' {
			if ch == '.' || ch == '0' {
				col++
			}
		}

		if col == SudokuSize {
			col = 0
			row++
		}

		if row == SudokuSize {
			break
		}
	}

	return grid
}
