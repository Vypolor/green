package handler

import (
	"net/http"

	greenapi "green/internal/clients/green-api"
)

type Handler struct {
	greenAPI *greenapi.Client
}

func New(greenAPI *greenapi.Client) *Handler {
	return &Handler{greenAPI: greenAPI}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("static")))

	// Green API proxy routes
	mux.HandleFunc("/api/getSettings", h.GetSettings)
	mux.HandleFunc("/api/getStateInstance", h.GetStateInstance)
	mux.HandleFunc("/api/sendMessage", h.SendMessage)
	mux.HandleFunc("/api/sendFileByUrl", h.SendFileByUrl)
}
