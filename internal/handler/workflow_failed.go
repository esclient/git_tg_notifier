package handler

import (
	"encoding/json"
	"net/http"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (h *Handler) WorkflowFailed(w http.ResponseWriter, r *http.Request) {
	event, code, err := validateWorkflowFailed(r)
	if err != nil {
		if code == http.StatusOK {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Ignored"))
			return
		}
		http.Error(w, err.Error(), code)
		return
	}

	if event == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ignored"))
		return
	}

	if err := h.service.WorkflowFailed(*event); err != nil {
		http.Error(w, "service error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("WorkflowFailed event processed"))
}

func validateWorkflowFailed(r *http.Request) (*model.WorkflowFailedEvent, int, error) {
	var event model.WorkflowFailedEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, http.StatusBadRequest, err
	}

	if event.Action != "completed" {
		return nil, http.StatusOK, nil
	}

	if event.WorkflowJob.Conclusion != "failure" {
		return nil, http.StatusOK, nil
	}

	return &event, http.StatusOK, nil
}
