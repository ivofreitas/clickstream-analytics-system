package http

import (
	"encoding/json"
	"net/http"

	"github.com/ivofreitas/clickstream-analytics-system/internal/app"
	"github.com/ivofreitas/clickstream-analytics-system/internal/domain"
	"github.com/julienschmidt/httprouter"
)

type EventHandler struct {
	service *app.EventService
}

func NewEventHandler(service *app.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) TrackEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var event domain.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.TrackEvent(event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (h *EventHandler) GetPageViews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pageURL := ps.ByName("page_url")

	count, err := h.service.GetPageViews(pageURL)
	if err != nil {
		http.Error(w, "No data found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"page_url":   pageURL,
		"view_count": count,
	}

	json.NewEncoder(w).Encode(response)
}
