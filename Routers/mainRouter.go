package routes

import (
	"github.com/gorilla/mux"

	consumerhandlers "backend/Consumer"
)

func RegisterRoutes(r *mux.Router) {

	authorRouter := r.PathPrefix("/authors").Subrouter()
	RegisterAuthorRoutes(authorRouter)

	bookRouter := r.PathPrefix("/books").Subrouter()
	RegisterBookRoutes(bookRouter)

	consumerRouter := r.PathPrefix("/consume").Subrouter()
	consumerhandlers.RegisterConsumerRouter(consumerRouter)

}
