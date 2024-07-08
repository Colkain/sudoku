package sudoku

import (
	"fmt"
	"log"
	"math/rand"
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
func (s *Sudoku) Generate() {
	cells := rand.Perm(BoardSize * BoardSize)
	log.Println("Solving")
	s.solve(cells, 0)
	// log.Println("Shuffling")
	// s.shuffle()
}

func (s *Sudoku) solve(cells []int, index int) bool {
	if index >= len(cells) {
		return true // Puzzle solved
	}

	cell := cells[index]
	row, col := cell/BoardSize, cell%BoardSize

	for num := 1; num <= 9; num++ {
		isValid, _ := s.CheckValidity(row, col, num)
		if isValid {
			s.SetBoardValue(row, col, num)

			if s.solve(cells, index+1) {
				log.Printf("Solved cell [%v, %v]=%v", row, col, num)
				return true
			}

			log.Printf("Backtracking for cell [%v,%v]=%v", row, col, num)
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

// shuffleSubgrids shuffles rows and columns within each 3x3 subgrid
func (s *Sudoku) shuffle() {
	// Shuffle rows within subgrids
	for startRow := 0; startRow < BoardSize; startRow += 3 {
		s.shuffleRows(startRow)
	}

	// Shuffle columns within subgrids
	for startCol := 0; startCol < BoardSize; startCol += 3 {
		s.shuffleColumns(startCol)
	}
}

func (s *Sudoku) shuffleRows(startRow int) {
	for i := startRow; i < startRow+3; i++ {
		rand.Shuffle(BoardSize, func(j, k int) {
			s.Board[i][j], s.Board[i][k] = s.Board[i][k], s.Board[i][j]
		})
	}
}

func (s *Sudoku) shuffleColumns(startCol int) {
	for i := startCol; i < startCol+3; i++ {
		rand.Shuffle(BoardSize, func(j, k int) {
			s.Board[j][i], s.Board[k][i] = s.Board[k][i], s.Board[j][i]
		})
	}
}

// // HideNumbers randomly hides numbers from the Sudoku board based on difficulty
// func HideNumbers(board *[BoardSize][BoardSize]int, difficulty float64) {
// 	numToHide := int(difficulty * float64(BoardSize*BoardSize))
// 	cells := rand.Perm(BoardSize * BoardSize)

// 	for _, cell := range cells[:numToHide] {
// 		row, col := cell/BoardSize, cell%BoardSize
// 		board[row][col] = 0
// 	}
// }
