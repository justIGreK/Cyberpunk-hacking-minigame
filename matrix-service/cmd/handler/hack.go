package handler

import (
	"encoding/json"
	"matrix-service/internal/models"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Hack
// @Tags Main tools
// @Description Try to hack matrix by your own
// @Produce  json
// @Param matrix_id query int true "id of matrix"
// @Param path query string true "path with clear coordinates to hack matrix"
// @Router /Hack [post]
func (h *Handler) Hack(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("matrix_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	haskTry := models.HackAttempt{
		MatrixID: id,
		Path: r.URL.Query().Get("path"),
	}
	if strings.TrimSpace(haskTry.Path) == ""{
		http.Error(w, "Empty path field", http.StatusBadRequest)
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
		http.Error(w, "Failed", http.StatusBadRequest)
		return
	}
}
