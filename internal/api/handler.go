package api

import (
	"encoding/json"
	"net/http"

	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{storage: store}
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	results := h.storage.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}
