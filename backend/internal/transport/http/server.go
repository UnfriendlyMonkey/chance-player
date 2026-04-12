// Package http provides the HTTP transport layer for the Chance Player API.
package http

import (
	"net/http"

	"github.com/unmo/chance-player/internal/transport/http/handler"
	"github.com/unmo/chance-player/internal/transport/http/middleware"
	"github.com/unmo/chance-player/internal/usecase"
	"github.com/unmo/chance-player/pkg/config"
)

// Server wraps the HTTP mux and its dependencies.
type Server struct {
	Mux *http.ServeMux
	cfg *config.Config
}

// New wires all handlers and middleware, returning a ready-to-serve Server.
func New(svc usecase.ChanceService, cfg *config.Config) *Server {
	mux := http.NewServeMux()

	coinH := handler.NewCoinHandler(svc)
	diceH := handler.NewDiceHandler(svc)
	cardH := handler.NewCardHandler(svc)
	numH := handler.NewNumberHandler(svc)

	// Route table
	mux.HandleFunc("GET /api/v1/coin/flip", coinH.Flip)
	mux.HandleFunc("GET /api/v1/coin/iching", coinH.IChing)
	mux.HandleFunc("GET /api/v1/dice/roll", diceH.Roll)
	mux.HandleFunc("GET /api/v1/card/draw", cardH.Draw)
	mux.HandleFunc("GET /api/v1/number/random", numH.Random)

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// Stack middleware: CORS → Logger → Auth(stub) → mux
	var h http.Handler = mux
	h = middleware.Auth(h)
	h = middleware.Logger(h)
	h = middleware.CORS(h)

	srv := &Server{cfg: cfg}
	srv.Mux = http.NewServeMux()
	srv.Mux.Handle("/", h)

	return srv
}
