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

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []schema.Author
	db := database.GetDB()

	result := db.Preload("Books").Find(&authors)
	if result.Error != nil {
		http.Error(w, "Failed to fetch authors", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(authors)

	logMessage := map[string]string{
		"message": "Fetched all authors",
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsAuthorProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending fetch all authors log message to Pulsar: %v", err)
	}
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author schema.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	result := db.Create(&author)
	if result.Error != nil {
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)

	logMessage := map[string]string{
		"message": "Author created",
		"author":  author.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsAuthorProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending create author log message to Pulsar: %v", err)
	}
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var author schema.Author
	db := database.GetDB()

	result := db.Preload("Books").First(&author, id)
	if result.Error != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(author)

	logMessage := map[string]string{
		"message": "Fetched author by ID",
		"author":  author.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsAuthorProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending fetch author by ID log message to Pulsar: %v", err)
	}
}

func UpdateAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var author schema.Author
	db := database.GetDB()

	if err := db.First(&author, id).Error; err != nil {
		http.Error(w, "Author not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.Save(&author)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(author)

	logMessage := map[string]string{
		"message": "Updated author",
		"author":  author.Name,
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsAuthorProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending update author log message to Pulsar: %v", err)
	}
}

func DeleteAuthorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db := database.GetDB()

	if err := db.Delete(&schema.Author{}, id).Error; err != nil {
		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	logMessage := map[string]string{
		"message": "Deleted author",
		"id":      strconv.Itoa(id),
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsAuthorProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending delete author log message to Pulsar: %v", err)
	}
}

func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID, _ := strconv.Atoi(vars["id"])

	var books []schema.Book
	db := database.GetDB()

	result := db.Where("author_id = ?", authorID).Find(&books)
	if result.Error != nil {
		http.Error(w, "No books found for this author", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books)

	logMessage := map[string]string{
		"message":  "Fetched books by author ID",
		"authorID": strconv.Itoa(authorID),
	}
	msgData, _ := json.Marshal(logMessage)
	_, err := pulsarutils.LogsBookProducer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msgData),
	})
	if err != nil {
		log.Printf("Error sending fetch books by author log message to Pulsar: %v", err)
	}
}
