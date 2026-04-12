package handler

import (
	"net/http"
	"strconv"

	"github.com/unmo/chance-player/internal/usecase"
)

// DiceHandler handles dice-related HTTP requests.
type DiceHandler struct {
	svc usecase.ChanceService
}

// NewDiceHandler creates a DiceHandler.
func NewDiceHandler(svc usecase.ChanceService) *DiceHandler {
	return &DiceHandler{svc: svc}
}

// Roll handles GET /api/v1/dice/roll[?sides=N][&count=N]
// Defaults: sides=6, count=1.
func (h *DiceHandler) Roll(w http.ResponseWriter, r *http.Request) {
	sides := parseIntQuery(r, "sides", 6)
	count := parseIntQuery(r, "count", 1)

	if sides < 2 {
		respondError(w, http.StatusBadRequest, "sides must be at least 2")
		return
	}
	if count < 1 {
		respondError(w, http.StatusBadRequest, "count must be at least 1")
		return
	}

	if count == 1 {
		result, err := h.svc.RollDice(r.Context(), sides)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, result)
	} else {
		result, err := h.svc.RollMultipleDice(r.Context(), count, sides)
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, result)
	}
}

func parseIntQuery(r *http.Request, key string, fallback int) int {
	v := r.URL.Query().Get(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
