package api

import (
	"log"
	"net/http"

	"github.com/DaVinciCodeCTF/status-checker/internal/crypto"
	"github.com/DaVinciCodeCTF/status-checker/internal/storage"
)

type Server struct {
	handler *Handler
	port    string
}

func NewServer(store storage.Storage, encryptor *crypto.Encryptor, port string) *Server {
	return &Server{
		handler: NewHandler(store, encryptor),
		port:    port,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/health", s.handler.Health)
	http.HandleFunc("/status", s.handler.Status)

	log.Printf("Starting API server on: %s...", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}

