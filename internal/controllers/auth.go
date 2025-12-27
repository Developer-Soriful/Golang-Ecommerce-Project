package controllers

import (
	"api-with-golang/internal/database"
	"api-with-golang/internal/middlewares"
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

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

// LOGIN USER

func LoginUser(w http.ResponseWriter, r *http.Request) {
	middlewares.HandleCors(w, r)
	if r.Method != "POST" {
		http.Error(w, "You have not entered POST Method", http.StatusBadRequest)
	}
	// THIS IS FOR REQUEST DECODER
	var req internal.Login
	json.NewDecoder(r.Body).Decode(&req)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{
		"email": req.Email,
	}
	//THIS IS FOR FILTER USER
	var user internal.User
	err := database.UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		http.Error(w, "Invaild Email or Password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login Successful",
	})
}
