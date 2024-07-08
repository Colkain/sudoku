package sudoku

import (
	"fmt"
)

const ErrInvalidCoor = "invalid coordinates"
const ErrNumberExists = "there is already a number"
const ErrInvalidNumber = "invalid number"
const ErrNumberExistsInCell = "number exists in this cell"
const ErrNumberExistsInColumn = "number exists in this column"
const ErrNumberExistsInRow = "number exists in this row"

const BoardSize = 9

type Sudoku struct {
	Board [BoardSize][BoardSize]int
}

func Init() *Sudoku {
	return &Sudoku{
		Board: [BoardSize][BoardSize]int{},
	}
}

// solveSudoku solves the Sudoku puzzle using backtracking
func (s *Sudoku) Generate() bool {
	emptyCell := s.findEmptyCell()
	if emptyCell == nil {
		return true // Puzzle solved
	}

	row, col := emptyCell[0], emptyCell[1]

	for num := 1; num <= 9; num++ {
		isValid, _ := s.CheckValidity(row, col, num)
		if isValid {
			s.SetBoardValue(row, col, num)

			if s.Generate() {
				return true
			}

			s.SetBoardValue(row, col, 0) // Backtrack
		}
	}

	return false
}

func (s *Sudoku) SetBoardValue(x, y, number int) {
	s.Board[x][y] = number
}

func (s *Sudoku) CheckValidity(x, y, number int) (bool, error) {
	if number < 1 || number > 9 {
		return false, fmt.Errorf(ErrInvalidNumber)
	}

	if x < 0 || x > 8 || y < 0 || y > 8 {
		return false, fmt.Errorf(ErrInvalidCoor)
	}

	if s.Board[x][y] != 0 {
		return false, fmt.Errorf(ErrNumberExists)
	}

	for i := 0; i < BoardSize; i++ {
		if s.Board[x][i] == number {
			return false, fmt.Errorf(ErrNumberExistsInRow)
		}

		if s.Board[i][y] == number {
			return false, fmt.Errorf(ErrNumberExistsInColumn)
		}
	}

	// Check the 3x3 block
	blockRow := (x / 3) * 3
	blockCol := (y / 3) * 3
	for i := blockRow; i < blockRow+3; i++ {
		for j := blockCol; j < blockCol+3; j++ {
			if s.Board[i][j] == number {
				return false, fmt.Errorf(ErrNumberExistsInCell)
			}
		}
	}

	return true, nil
}

// findEmptyCell finds the first empty cell in the Sudoku board
func (s *Sudoku) findEmptyCell() *[2]int {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if s.Board[i][j] == 0 {
				return &[2]int{i, j}
			}
		}
	}
	return nil
}
