package api

import (
	"encoding/json"
	"net/http"

	"github.com/DaVinciCodeCTF/status-checker/internal/crypto"
	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
)

type Handler struct {
	storage   storage.Storage
	encryptor *crypto.Encryptor
}

func NewHandler(store storage.Storage, encryptor *crypto.Encryptor) *Handler {
	return &Handler{
		storage:   store,
		encryptor: encryptor,
	}
}

func (h *Handler) Status(w http.ResponseWriter, r *http.Request) {
	results := h.storage.GetAll()

	plainJSON, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "failed to serialize status", http.StatusInternalServerError)
		return
	}

	encrypted, err := h.encryptor.Encrypt(plainJSON)
	if err != nil {
		http.Error(w, "failed to encrypt status", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(encrypted)
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`))
}

