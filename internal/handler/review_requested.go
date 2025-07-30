package handler

import (
	"encoding/json"
	"net/http"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (h *Handler) ReviewRequested(w http.ResponseWriter, r *http.Request) {
	event, code, err := validateReviewRequested(r)
	if err != nil {
		if code == http.StatusOK {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Ignored"))
			return
		}
		http.Error(w, err.Error(), code)
		return
	}

	if err := h.service.ReviewRequested(*event); err != nil {
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ReviewRequest event processed"))
}

func validateReviewRequested(r *http.Request) (*model.ReviewRequestedEvent, int, error) {
	var event model.ReviewRequestedEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if event.Action != "review_requested" {
		return nil, http.StatusOK, nil
	}

	return &event, http.StatusOK, nil
}
