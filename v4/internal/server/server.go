package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/colkain/sudoku/v4/internal/sudoku"
	"github.com/colkain/sudoku/v4/web"
)

const JsonContentType = "application/json"

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	game *sudoku.Sudoku

	http.Handler
}

func NewPlayerServer(game *sudoku.Sudoku) (*PlayerServer, error) {
	p := new(PlayerServer)
	p.game = game

	router := http.NewServeMux()

	router.Handle("/", templ.Handler(web.Game(p.game.Game)))
	router.Handle("/generate", http.HandlerFunc(p.generateHandler))
	router.Handle("/validate", http.HandlerFunc(p.validateHandler))
	router.Handle("/solve", http.HandlerFunc(p.solveHandler))

	p.Handler = router

	return p, nil
}

func (p *PlayerServer) generateHandler(w http.ResponseWriter, r *http.Request) {
	p.game = sudoku.Init()
	p.game.Generate()
	templ.Handler(web.Game(p.game.Game))
}

func (p *PlayerServer) solveHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(web.Game(p.game.Board))
}

func (p *PlayerServer) validateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for _, v := range r.Form {
		if v == nil {
			// p.gameTemplate.Execute(w, "You lose!")
		}
		// convert key cell-x-y to x
	}

	// for x := range p.game.Board {
	// 	for y := range p.game.Board {
	// 		if p.game.Board[x][y] != p.game.Game[x][y] {
	// 			p.gameTemplate.Execute(w, "You lose!")
	// 		}
	// 	}
	// }

	// p.gameTemplate.Execute(w, "<p>You win!</p>")
}
