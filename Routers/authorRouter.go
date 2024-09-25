package routes

import (
	"github.com/gorilla/mux"

	handler "backend/Handlers"
)

func RegisterAuthorRoutes(router *mux.Router) {
	router.HandleFunc("", handler.GetAllAuthors).Methods("GET")
	router.HandleFunc("", handler.CreateAuthor).Methods("POST")
	router.HandleFunc("/{id}", handler.GetAuthorByID).Methods("GET")
	router.HandleFunc("/{id}", handler.UpdateAuthorByID).Methods("PUT")
	router.HandleFunc("/{id}", handler.DeleteAuthorByID).Methods("DELETE")
	router.HandleFunc("/{id}/books", handler.GetBooksByAuthor).Methods("GET")

}
