package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"skillbox-diploma/internal/config"
	"skillbox-diploma/internal/models"
	"skillbox-diploma/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *service.Service
}

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", h.handleConn)

	return router
}

func (h *Handler) handleConn(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.Get()
	if err != nil {
		h.responseError(w, err, false)
		return
	}

	resT := models.ResultT{
		Data:   res,
		Status: true,
	}

	data, err := json.Marshal(resT)
	if err != nil {
		h.responseError(w, err, true)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Printf("[warning] failed to write response: %s", err.Error())
	}
}

func (h *Handler) responseError(w http.ResponseWriter, err error, isCritical bool) {
	if isCritical {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte{}); err != nil {
			log.Printf("[warning] failed to write response: %s", err.Error())
		}

		return
	}

	res := models.ResultT{
		Error: err.Error(),
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Printf("failed to marshal response data: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write(data); err != nil {
		log.Printf("[warning] failed to write response: %s", err.Error())
	}
}

func New(cfg config.Config) *Handler {
	return &Handler{
		service: service.New(cfg.TempDir, fmt.Sprintf("http://%s:%d", cfg.Simulator.Host, cfg.Simulator.Port)),
	}
}
