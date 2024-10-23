package handler

import (
	"context"
	_ "hacker-service/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service Service
}

type Service interface {
	HackMatrix(ctx context.Context, id int) ([]string, error)
}

func NewHandler(matrixService Service) *Handler {
	return &Handler{service: matrixService}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	r.Post("/Hack", h.Hack)

	return r
}
