package handler

import (
	"context"
	_ "matrix-service/docs"
	"matrix-service/internal/models"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service Service
}

type Service interface {
	GenerateMatrix() [][]int
	GenerateSequences() [][]int
	GetMatrix(ctx context.Context, id int) (*models.HackMatrix, error)
	HackMatrix(ctx context.Context, attempt models.HackAttempt) (bool, error)
}

func NewHandler(matrixService Service) *Handler {
	return &Handler{service: matrixService}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Get("/GetSequence", h.GetSequence)
	r.Get("/GetSequenceSugar", h.GetSequenceForHuman)
	r.Post("/Hack", h.Hack)

	return r
}
