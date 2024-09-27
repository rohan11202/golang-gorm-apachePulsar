package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	middleware "backend/Middleware"
	schema "backend/Models"
	routes "backend/Routers"
	pulsarutils "backend/Utils"
)

func main() {
	middleware.Init()

	schema.Initiate()
	router := mux.NewRouter()
	router.Use(middleware.MetricsMiddleware)

	routes.RegisterRoutes(router)

	pulsarutils.SetupPulsar()

	log.Println("Starting server on :3000.....")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Failed to Start Server:%v", err)
	}

}
