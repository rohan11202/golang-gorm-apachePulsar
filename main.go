package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	schema "backend/Models"
	routes "backend/Routers"
	pulsarutils "backend/Utils"
)

func main() {
	schema.Initiate()
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	pulsarutils.SetupPulsar()

	log.Println("Starting server on :3000.....")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Failed to Start Server:%v", err)
	}

}
