package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bookstore-api/database"

	"bookstore-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	cur, err := database.BooksCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	var books []models.Book
	for cur.Next(context.TODO()) {
		var book models.Book
		err := cur.Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	response := struct {
		Data []models.Book `json:"data"`
	}{
		Data: books,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var book models.Book
	err = database.BooksCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := struct {
		Data models.Book `json:"data"`
	}{
		Data: book,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	result, err := database.BooksCollection.InsertOne(context.TODO(), book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		ID string `json:"id"`
	}{
		ID: result.InsertedID.(primitive.ObjectID).Hex(),
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)

	_, err = database.BooksCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updatedBook})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Book updated successfully",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.BooksCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Book with ID %s has been deleted", id.Hex()),
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
