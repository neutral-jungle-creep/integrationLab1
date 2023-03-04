package delivery

import (
	"IntegrationLab1/internal/service"
	"net/http"
)

type handler struct {
	service *service.Service
}

type Handler interface {
	SubmitCompletedDoc(w http.ResponseWriter, r *http.Request)
}

func NewHandler(service *service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) SubmitCompletedDoc(w http.ResponseWriter, r *http.Request) {
}
