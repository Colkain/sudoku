package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/colkain/sudoku/v4/internal/sudoku"
)

const JsonContentType = "application/json"
const HtmlTemplatePath = "web/templates/game.html"

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	game         *sudoku.Sudoku
	gameTemplate *template.Template

	http.Handler
}

func NewPlayerServer(game *sudoku.Sudoku) (*PlayerServer, error) {
	p := new(PlayerServer)

	tmpl, err := template.ParseFiles(HtmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", HtmlTemplatePath, err)
	}

	p.gameTemplate = tmpl
	p.game = game

	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(p.gameHandler))
	router.Handle("/generate", http.HandlerFunc(p.generateHandler))
	// router.Handle("/submit", http.HandlerFunc(p.gameHandler))
	// router.Handle("/solve", http.HandlerFunc(p.gameHandler))

	p.Handler = router

	return p, nil
}

func (p *PlayerServer) gameHandler(w http.ResponseWriter, r *http.Request) {
	p.gameTemplate.Execute(w, p.game.Game)
}

func (p *PlayerServer) generateHandler(w http.ResponseWriter, r *http.Request) {
	p.game.Generate()
	w.Header().Set("content-type", JsonContentType)
	json.NewEncoder(w).Encode(p.game.Game)
}
