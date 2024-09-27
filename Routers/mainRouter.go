package routes

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	consumerhandlers "backend/Consumer"
)

func RegisterRoutes(r *mux.Router) {

	authorRouter := r.PathPrefix("/authors").Subrouter()
	RegisterAuthorRoutes(authorRouter)

	bookRouter := r.PathPrefix("/books").Subrouter()
	RegisterBookRoutes(bookRouter)

	consumerRouter := r.PathPrefix("/consume").Subrouter()
	consumerhandlers.RegisterConsumerRouter(consumerRouter)

	r.Handle("/metrics", promhttp.Handler())

	// Use the metrics middleware for your routes
}
