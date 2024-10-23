package handler

import (
	"encoding/json"
	"net/http"
)

// @Summary Reports
// @Tags Additional methods
// @Description Check reports of hacked matrixes
// @Produce  json
// @Router /GetReports [get]
func (h *Handler) GetReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.service.GetReports(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"Reports": reports,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
