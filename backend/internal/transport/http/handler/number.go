package handler

import (
	"net/http"

	"github.com/unmo/chance-player/internal/usecase"
)

// NumberHandler handles random number HTTP requests.
type NumberHandler struct {
	svc usecase.ChanceService
}

// NewNumberHandler creates a NumberHandler.
func NewNumberHandler(svc usecase.ChanceService) *NumberHandler {
	return &NumberHandler{svc: svc}
}

// Random handles GET /api/v1/number/random[?min=N&max=N]
// Defaults: min=1, max=100.
func (h *NumberHandler) Random(w http.ResponseWriter, r *http.Request) {
	min := parseIntQuery(r, "min", 1)
	max := parseIntQuery(r, "max", 100)

	if min >= max {
		respondError(w, http.StatusBadRequest, "min must be less than max")
		return
	}
	result, err := h.svc.RandomNumber(r.Context(), min, max)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, result)
}
