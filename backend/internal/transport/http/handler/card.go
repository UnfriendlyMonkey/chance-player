package handler

import (
	"net/http"

	"github.com/unmo/chance-player/internal/usecase"
)

// CardHandler handles card-related HTTP requests.
type CardHandler struct {
	svc usecase.ChanceService
}

// NewCardHandler creates a CardHandler.
func NewCardHandler(svc usecase.ChanceService) *CardHandler {
	return &CardHandler{svc: svc}
}

// Draw handles GET /api/v1/card/draw
func (h *CardHandler) Draw(w http.ResponseWriter, r *http.Request) {
	result, err := h.svc.DrawCard(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, result)
}
