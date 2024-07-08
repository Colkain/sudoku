package main

import (
	"log"
	"net/http"

	"github.com/colkain/sudoku/v4/internal/server"
	"github.com/colkain/sudoku/v4/internal/sudoku"
)

func main() {
	game := sudoku.Init()
	game.Generate()

	server, err := server.NewPlayerServer(game)
	if err != nil {
		log.Fatalln(err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
