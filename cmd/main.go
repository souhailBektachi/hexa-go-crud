package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/souhailBektachi/hexa-go-crud/internal/adapters/handler"
	service "github.com/souhailBektachi/hexa-go-crud/internal/core/services"

	. "github.com/souhailBektachi/hexa-go-crud/internal/adapters/repository"
)

var (
	HttpHandler *handler.HttpHandler
	svc         *service.SomthingService
)

func main() {
	repo := NewPostgressRepository()
	svc = service.NewSomthingService(repo)
	HttpHandler = handler.NewHttpHandler(*svc)
	router := mux.NewRouter()
	InitRoutes(router, HttpHandler)

	log.Println("Skkktarting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func InitRoutes(router *mux.Router, h *handler.HttpHandler) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}).Methods("GET")
	router.HandleFunc("/somthing", h.Save).Methods("POST")
	router.HandleFunc("/somthing/{id}", h.FindById).Methods("GET")
	router.HandleFunc("/somthing", h.FindAll).Methods("GET")
	router.HandleFunc("/somthing/{id}", h.DeleteById).Methods("DELETE")
	router.HandleFunc("/somthing/{id}", h.Update).Methods("PUT")
}
