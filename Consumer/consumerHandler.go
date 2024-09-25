package consumerhandlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	pulsarutils "backend/Utils"
)

func RegisterConsumerRouter(router *mux.Router) {
	router.HandleFunc("/BookLogs", ConsumeBookLogs).Methods("GET")
	router.HandleFunc("/AuthorLogs", ConsumeAuthorLogs).Methods("GET")
}

func ConsumeAuthorLogs(w http.ResponseWriter, r *http.Request) {
	msg, err := pulsarutils.LogsAuthorConsumer.Receive(context.Background())
	if err != nil {
		log.Printf("Error receiving logs for Author: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Received log Author message: %s", string(msg.Payload()))
	pulsarutils.LogsAuthorConsumer.Ack(msg)
}

func ConsumeBookLogs(w http.ResponseWriter, r *http.Request) {
	msg, err := pulsarutils.LogsBookConsumer.Receive(context.Background())
	if err != nil {
		log.Printf("Error receiving logs for Book: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Received log Book message: %s", string(msg.Payload()))
	pulsarutils.LogsBookConsumer.Ack(msg)
}
