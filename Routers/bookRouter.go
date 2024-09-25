package routes

import (
	"github.com/gorilla/mux"

	handler "backend/Handlers"
)

func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("", handler.GetAllBooks).Methods("GET")
	router.HandleFunc("", handler.CreateBook).Methods("POST")
	router.HandleFunc("/{id}", handler.GetBookByID).Methods("GET")
	router.HandleFunc("/{id}", handler.UpdateBookByID).Methods("PUT")
	router.HandleFunc("/{id}", handler.DeleteBookByID).Methods("DELETE")

}
