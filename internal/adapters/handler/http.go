package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/souhailBektachi/hexa-go-crud/internal/core/domain"
	service "github.com/souhailBektachi/hexa-go-crud/internal/core/services"
)

type HttpHandler struct {
	svc service.SomthingService
}

func NewHttpHandler(svc service.SomthingService) *HttpHandler {
	return &HttpHandler{svc: svc}
}

func (h *HttpHandler) Save(w http.ResponseWriter, r *http.Request) {
	var somthing domain.Somthing
	if err := json.NewDecoder(r.Body).Decode(&somthing); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.svc.Save(somthing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *HttpHandler) FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	somthing, err := h.svc.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(somthing)
}

func (h *HttpHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	somthings, err := h.svc.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(somthings)
}

func (h *HttpHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.svc.DeleteById(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HttpHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var somthing domain.Somthing
	if err := json.NewDecoder(r.Body).Decode(&somthing); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	somthing.ID = id
	if err := h.svc.Update(somthing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
