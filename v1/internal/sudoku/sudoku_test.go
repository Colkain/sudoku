package sudoku_test

import (
	"testing"

	"github.com/colkain/sudoku/v1/internal/sudoku"
)

func TestSudoku_Init(t *testing.T) {
	t.Run("Initialize sudoku board", func(t *testing.T) {
		want := &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{},
		}
		got := sudoku.Init()

		assertBoard(t, got, want)
	})
}

func TestSudoku_CheckValidity(t *testing.T) {
	grid := sudoku.Init()

	t.Run("Place an invalid number", func(t *testing.T) {
		want := sudoku.ErrInvalidNumber
		_, got := grid.CheckValidity(-1, 0, 0)

		assertResponse(t, got.Error(), want)
	})

	t.Run("Place an invalid coordinate", func(t *testing.T) {
		want := sudoku.ErrInvalidCoor
		_, got := grid.CheckValidity(-1, 0, 1)

		assertResponse(t, got.Error(), want)
	})

	t.Run("Number already exists", func(t *testing.T) {
		grid = &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}

		want := sudoku.ErrNumberExists
		_, got := grid.CheckValidity(0, 0, 1)

		assertResponse(t, got.Error(), want)
	})

	t.Run("Number already exists in cell", func(t *testing.T) {
		grid = &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}

		want := sudoku.ErrNumberExistsInCell
		_, got := grid.CheckValidity(0, 0, 1)

		assertResponse(t, got.Error(), want)
	})

	t.Run("Number already exists in column", func(t *testing.T) {
		grid = &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}

		want := sudoku.ErrNumberExistsInColumn
		_, got := grid.CheckValidity(0, 0, 1)

		assertResponse(t, got.Error(), want)
	})

	t.Run("Number already exists in row", func(t *testing.T) {
		grid = &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{
				{0, 0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}

		want := sudoku.ErrNumberExistsInRow
		_, got := grid.CheckValidity(0, 0, 1)

		assertResponse(t, got.Error(), want)
	})
}

func TestSudoku_SetValue(t *testing.T) {
	t.Run("Initialize sudoku board", func(t *testing.T) {
		grid := sudoku.Init()
		want := &sudoku.Sudoku{
			Board: [sudoku.BoardSize][sudoku.BoardSize]int{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		}
		grid.SetBoardValue(0, 0, 1)

		assertBoard(t, grid, want)
	})
}

func TestSudoku_Generate(t *testing.T) {
	t.Run("Generate sudoku board", func(t *testing.T) {
		want := &sudoku.Sudoku{
			Board: [9][9]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{2, 1, 4, 3, 6, 5, 8, 9, 7},
				{3, 6, 5, 8, 9, 7, 2, 1, 4},
				{8, 9, 7, 2, 1, 4, 3, 6, 5},
				{5, 3, 1, 6, 4, 2, 9, 7, 8},
				{6, 4, 2, 9, 7, 8, 5, 3, 1},
				{9, 7, 8, 5, 3, 1, 6, 4, 2},
			},
		}
		grid := sudoku.Init()
		grid.Generate()

		assertBoard(t, want, grid)
	})
}

func assertResponse(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertBoard(t testing.TB, got, want *sudoku.Sudoku) {
	t.Helper()
	for x := range got.Board {
		for y := range got.Board {
			if got.Board[x][y] != want.Board[x][y] {
				t.Errorf("response body is wrong, got %q want %q", got, want)
			}
		}
	}
}
