package handler

import (
	"context"
	_ "hacker-service/docs"
	"hacker-service/internal/models"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service Service
}

type Service interface {
	HackMatrix(ctx context.Context, id int) ([]string, error)
	GetReports(ctx context.Context)([]models.HackReport, error)
}

func NewHandler(matrixService Service) *Handler {
	return &Handler{service: matrixService}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Post("/HelpHack", h.Hack)
	r.Get("/GetReports", h.GetReports)

	return r
}
