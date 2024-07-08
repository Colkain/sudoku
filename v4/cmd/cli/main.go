package main

import (
	"fmt"

	"github.com/colkain/sudoku/v4/internal/sudoku"
)

func main() {
	game := sudoku.Init()
	game.Generate()

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Printf("%d ", game.Game[x][y])
			if y%3 == 2 && y%9 != 8 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if x%3 == 2 && x%9 != 8 {
			fmt.Println("--------------------")
		}
	}
}
