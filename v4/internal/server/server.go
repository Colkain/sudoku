package server

import (
	"log"
	"net/http"
	"strconv"

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

	router.Handle("/", templ.Handler(web.Game(p.game.Game, false)))
	router.Handle("/new", http.HandlerFunc(p.generateHandler))
	router.Handle("/validate", http.HandlerFunc(p.validateHandler))
	router.Handle("/solve", http.HandlerFunc(p.solveHandler))

	p.Handler = router

	return p, nil
}

func (p *PlayerServer) generateHandler(w http.ResponseWriter, r *http.Request) {
	p.game = sudoku.Init()
	p.game.Generate()
	templ.Handler(web.Game(p.game.Game, false)).ServeHTTP(w, r)
}

func (p *PlayerServer) solveHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(web.Game(p.game.Board, true)).ServeHTTP(w, r)
}

func (p *PlayerServer) validateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for i, v := range r.PostForm {
		value, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatalln(err)
		}
		x, err := strconv.Atoi(i[:1])
		if err != nil {
			log.Fatalln(err)
		}
		y, err := strconv.Atoi(i[2:])
		if err != nil {
			log.Fatalln(err)
		}
		p.game.Game.SetBoardValue(x, y, int32(value))
	}

	for x := range p.game.Board {
		for y := range p.game.Board {
			if p.game.Board[x][y] != p.game.Game[x][y] {
				p.game.Game.SetBoardValue(x, y, int32(^uint32(int32(p.game.Game[x][y])-1)))
			}
		}
	}

	templ.Handler(web.Game(p.game.Game, true)).ServeHTTP(w, r)
}
