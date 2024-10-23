package handler

import (
	"encoding/json"
	"matrix-service/internal/models"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Hack
// @Tags HackTools
// @Description Try to hack matrix by your own
// @Accept  json
// @Produce  json
// @Param matrix_id query int true "id of matrix"
// @Param attempts query string true "clear coordinates of attempts to hack matrix"
// @Router /Hack [post]
func (h *Handler) Hack(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("matrix_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	haskTry := models.HackAttempts{
		MatrixID: id,
		Attempts: r.URL.Query().Get("attempts"),
	}
	if strings.TrimSpace(haskTry.Attempts) == ""{
		http.Error(w, "Empty attemtps field", http.StatusUnprocessableEntity)
		return
	}

	isHacked, err := h.service.HackMatrix(r.Context(), haskTry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if isHacked {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Matrix is successfully hacked")
		return
	} else {
		http.Error(w, "Failed", http.StatusUnprocessableEntity)
		return
	}
}
