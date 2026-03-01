package handlers

import (
	"net/http"

	"github.com/railgorail/web-backend/internal/services"
)

type Handlers struct {
	service *services.Service
}

func New(service *services.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) HandleHello(w http.ResponseWriter, r *http.Request) {
	message := h.service.GetHelloMessage()
	w.Write([]byte(message))
}
