package controllers

import (
	"api-with-golang/internal/database"
	"api-with-golang/internal/middlewares"
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	// "ShopSphere/internal/"
)

// AUTH CONTROLLERS
// FIND USER
func User(w http.ResponseWriter, r *http.Request) {
	middlewares.HandleCors(w, r)
	if r.Method != "GET" {
		http.Error(w, "You have not entered GET Method", http.StatusBadRequest)
	}
	var users []internal.User
	cursor, _ := database.UserCollection.Find(context.Background(), bson.M{})
	cursor.All(context.Background(), &users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
