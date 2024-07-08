package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/colkain/sudoku/v4/internal/server"
	"github.com/colkain/sudoku/v4/internal/sudoku"
)

func TestServer(t *testing.T) {
	t.Run("GET / returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, sudoku.Sudoku{})

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})
}

func newGameRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/", nil)
	return req
}

func mustMakePlayerServer(t *testing.T, game sudoku.Sudoku) *server.PlayerServer {
	server, err := server.NewPlayerServer(&game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}

	return server
}

func assertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("did not get correct status, got %d, want %d", got.Code, want)
	}
}
