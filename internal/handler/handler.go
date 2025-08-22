package handler

import (
	"net/http"

	"github.com/esclient/git_tg_notifier/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/commit", h.Commit)
	mux.HandleFunc("/review_requested", h.ReviewRequested)
	mux.HandleFunc("/workflow_failed", h.WorkflowFailed)
}
