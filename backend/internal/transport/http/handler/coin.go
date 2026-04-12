package handler

import (
	"net/http"
	"strconv"

	"github.com/unmo/chance-player/internal/usecase"
)

// CoinHandler handles coin-related HTTP requests.
type CoinHandler struct {
	svc usecase.ChanceService
}

// NewCoinHandler creates a CoinHandler.
func NewCoinHandler(svc usecase.ChanceService) *CoinHandler {
	return &CoinHandler{svc: svc}
}

// Flip handles GET /api/v1/coin/flip[?count=N]
// Without count: flips one coin. With count: flips N coins.
func (h *CoinHandler) Flip(w http.ResponseWriter, r *http.Request) {
	countStr := r.URL.Query().Get("count")
	if countStr == "" {
		result, err := h.svc.FlipCoin(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, result)
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil || count < 1 {
		respondError(w, http.StatusBadRequest, "count must be a positive integer")
		return
	}
	result, err := h.svc.FlipCoins(r.Context(), count)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, result)
}

// IChing handles GET /api/v1/coin/iching
func (h *CoinHandler) IChing(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.ConsultIChing(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, result)
}
