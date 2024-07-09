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

		request := newGetRequest("/")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("GET /new returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, sudoku.Sudoku{})

		request := newGetRequest("/new")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("GET /solve returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, sudoku.Sudoku{})

		request := newGetRequest("/solve")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("convert Post form error value", func(t *testing.T) {
		srv := mustMakePlayerServer(t, sudoku.Sudoku{})
		data := map[string][]string{
			"puzzle": {"..9..38.0.101.76....2432......1...69.83.9.....6.62.71...9......1945....89.207.132.170....6.."},
		}

		want := server.ErrConvertVal
		_, got := srv.ConvertFormToGame(data)

		assertResponse(t, got.Error(), want)
	})

	t.Run("convert Post form error row", func(t *testing.T) {
		srv := mustMakePlayerServer(t, sudoku.Sudoku{})

		data := map[string][]string{
			"a,0": {"1"},
		}

		want := server.ErrConvertRow
		_, got := srv.ConvertFormToGame(data)

		assertResponse(t, got.Error(), want)
	})

	t.Run("convert Post form error col", func(t *testing.T) {
		srv := mustMakePlayerServer(t, sudoku.Sudoku{})
		data := map[string][]string{
			"0,a": {"1"},
		}

		want := server.ErrConvertCol
		_, got := srv.ConvertFormToGame(data)

		assertResponse(t, got.Error(), want)
	})

	t.Run("convert Post form error col", func(t *testing.T) {
		srv := mustMakePlayerServer(t, sudoku.Sudoku{})
		data := map[string][]string{
			"0,0": {"1"},
		}
		want := &sudoku.Sudoku{
			Game: [sudoku.BoardSize][sudoku.BoardSize]int32{
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

		got, _ := srv.ConvertFormToGame(data)

		assertGame(t, got, want)
	})

	t.Run("POST /validate returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, sudoku.Sudoku{})
		request := newPostRequest("/validate")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})
}

func newGetRequest(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	return req
}

func newPostRequest(url string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, url, nil)
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

func assertGame(t testing.TB, got, want *sudoku.Sudoku) {
	t.Helper()
	for x := range got.Board {
		for y := range got.Board {
			if got.Game[x][y] != want.Game[x][y] {
				t.Errorf("response body is wrong, got %q want %q", got, want)
			}
		}
	}
}

func assertResponse(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
