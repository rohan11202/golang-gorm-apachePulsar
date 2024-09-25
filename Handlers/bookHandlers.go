package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gorilla/mux"

	database "backend/Database"
	schema "backend/Models"
	pulsarutils "backend/Utils"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []schema.Book
	db := database.GetDB()

	result := db.Find(&books)
	if result.Error != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

	logMessage := map[string]string{
		"message": "Fetched all books",
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending fetch all books log message to Pulsar: %v", err)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book schema.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	result := db.Create(&book)
	if result.Error != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

	logMessage := map[string]string{
		"message": "Book created",
		"book":    book.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending create book log message to Pulsar: %v", err)
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book schema.Book
	db := database.GetDB()

	result := db.First(&book, id)
	if result.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)

	logMessage := map[string]string{
		"message": "Fetched book by ID",
		"book":    book.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending fetch book by ID log message to Pulsar: %v", err)
	}
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var book schema.Book
	db := database.GetDB()

	if err := db.First(&book, id).Error; err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.Save(&book)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)

	logMessage := map[string]string{
		"message": "Updated book",
		"book":    book.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending update book log message to Pulsar: %v", err)
	}
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db := database.GetDB()

	if err := db.Delete(&schema.Book{}, id).Error; err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	logMessage := map[string]string{
		"message": "Deleted book",
		"id":      strconv.Itoa(id),
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending delete book log message to Pulsar: %v", err)
	}
}
