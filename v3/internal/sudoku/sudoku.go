package sudoku

import (
	"fmt"
	"math"
	"math/rand"
)

const ErrInvalidCoor = "invalid coordinates"
const ErrNumberExists = "there is already a number"
const ErrInvalidNumber = "invalid number"
const ErrNumberExistsInCell = "number exists in this cell"
const ErrNumberExistsInColumn = "number exists in this column"
const ErrNumberExistsInRow = "number exists in this row"

const BoardSize = 9
const NumbersToHide = 12

type Grid [BoardSize][BoardSize]int

type Sudoku struct {
	Board Grid
	Game  Grid
	SRN   int
}

func Init() *Sudoku {
	return &Sudoku{
		Game:  [BoardSize][BoardSize]int{},
		Board: [BoardSize][BoardSize]int{},
		SRN:   int(float64(BoardSize) / math.Sqrt(float64(BoardSize))),
	}
}

// solveSudoku solves the Sudoku puzzle using backtracking
func (s *Sudoku) Generate() {
	s.fillDiagonal()
	s.fillRemaining(0, s.SRN)
	s.HideNumbers()
}

func (g *Grid) SetBoardValue(x, y, number int) {
	g[x][y] = number
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

func (s *Sudoku) fillDiagonal() {
	for i := 0; i < BoardSize; i += s.SRN {
		s.fillBox(i, i)
	}
}

func (s *Sudoku) fillRemaining(i, j int) bool {
	if i == BoardSize-1 && j == BoardSize {
		return true
	}
	if j == BoardSize {
		i++
		j = 0
	}
	if s.Board[i][j] != 0 {
		return s.fillRemaining(i, j+1)
	}

	for num := 1; num <= BoardSize; num++ {
		if isValid, _ := s.CheckValidity(i, j, num); isValid {
			s.Board.SetBoardValue(i, j, num)
			if s.fillRemaining(i, j+1) {
				return true
			}

			s.Board.SetBoardValue(i, j, 0) // Backtrack
		}
	}

	return false
}

func (s *Sudoku) fillBox(row, col int) {
	var num int
	for i := 0; i < s.SRN; i++ {
		for j := 0; j < s.SRN; j++ {
			for {
				num = rand.Intn(BoardSize) + 1
				if s.unUsedInBox(row, col, num) {
					break
				}
			}
			s.Board.SetBoardValue(row+i, col+j, num)
		}
	}
}

func (s *Sudoku) unUsedInBox(rowStart, colStart, num int) bool {
	for i := 0; i < s.SRN; i++ {
		for j := 0; j < s.SRN; j++ {
			if s.Board[rowStart+i][colStart+j] == num {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) HideNumbers() {
	count := NumbersToHide
	s.Game = s.Board
	for count != 0 {
		i := rand.Intn(BoardSize)
		j := rand.Intn(BoardSize)
		if s.Game[i][j] != 0 {
			s.Game.SetBoardValue(i, j, 0)
			count--
		}
	}
}
