package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary GetSequence
// @Tags Matrix_service
// @Description Get new matrix and sequences for hacking
// @Accept  json
// @Produce  json
// @Param id query int true "id of Hack"
// @Router /GetSequence [get]
func (h *Handler) GetSequence(w http.ResponseWriter, r *http.Request) {
	hackStr := r.URL.Query().Get("id")
	hackId, err := strconv.Atoi(hackStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	matrix, err := h.service.GetMatrix(r.Context(), hackId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"ID":        matrix.ID,
		"Matrix":    matrix.Matrix,
		"Sequences": matrix.Sequence,
	}
	json.NewEncoder(w).Encode(response)

}
