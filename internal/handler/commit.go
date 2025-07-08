package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (h *Handler) Commit(w http.ResponseWriter, r *http.Request) {
	event, code, err := validateCommitRequest(r)
	if err != nil {
		if code == http.StatusOK {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Ignored"))
			return
		}
		http.Error(w, err.Error(), code)
		return
	}

	if err := h.service.Commit(*event); err != nil {
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Commit event processed"))
}

func validateCommitRequest(r *http.Request) (*model.CommitEvent, int, error) {
	var event model.CommitEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if len(event.Commits) == 0 {
		return nil, http.StatusBadRequest, fmt.Errorf("no commits in payload")
	}

	return &event, http.StatusOK, nil
}
