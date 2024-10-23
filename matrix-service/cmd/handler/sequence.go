package handler

import (
	"encoding/json"
	"fmt"
	"matrix-service/internal/models"
	"net/http"
	"strconv"
)

// @Summary GetSequence
// @Tags Main tools
// @Description Get new matrix and sequences for hacking
// @Produce  json
// @Param id query int true "id of Matrix"
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

// @Summary GetSequence For Human
// @Tags Additional methods
// @Description Get new matrix and sequences for hacking with human representation
// @Produce  json
// @Param id query int true "id of Matrix"
// @Router /GetSequenceSugar [get]
func (h *Handler) GetSequenceForHuman(w http.ResponseWriter, r *http.Request) {
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
	resp := h.printReadableResponse(*matrix)
	response := map[string]interface{}{
		"ID":        resp.ID,
		"Matrix":    resp.Matrix,
		"Sequences": resp.Sequence,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

type Response struct {
	ID       int
	Matrix   []string
	Sequence []string
}

func (h *Handler) printReadableResponse(matrix models.HackMatrix) Response {
	resp := Response{ID: matrix.ID}

	for _, rows := range matrix.Matrix {
		rowStr := ""
		for _, val := range rows {
			rowStr += fmt.Sprintf("%d ", val)
		}
		resp.Matrix = append(resp.Matrix, rowStr)
	}

	for _, sequence := range matrix.Sequence {
		seqStr := ""
		for _, val := range sequence {
			seqStr += fmt.Sprintf("%d ", val)
		}
		resp.Sequence = append(resp.Sequence, seqStr)
	}

	return resp
}
