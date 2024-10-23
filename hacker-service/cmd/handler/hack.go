package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// @Summary HelpHack
// @Tags Main tools
// @Description Get ready answers how to hack matrixes of matrix_service
// @Produce  json
// @Param matrix_id query int true "id of the matrix desired to be hacked"
// @Router /HelpHack [post]
func (h *Handler) Hack(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("matrix_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	answers, err := h.service.HackMatrix(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"Answers": answers,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
