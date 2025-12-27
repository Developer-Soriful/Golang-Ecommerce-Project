package authservice

import (
	"api-with-golang/internal/database"
	"api-with-golang/internal/middlewares"
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// REGISTER CONTROLLER
func Register(w http.ResponseWriter, r *http.Request) {

	// IMPLEMENTATION FOR REGISTATION GOES HERE
	middlewares.HandleCors(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "You have entered different Method", http.StatusBadRequest)
	}

	// REGISTRATION LOGIC
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	// VALIDATION AND INSERTION LOGIC GOES HERE
	ctx := context.Background()
	count, _ := database.UserCollection.CountDocuments(ctx, bson.M{"email": req.Email})
	if count > 0 {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}
	// CREATE USER OBJECT
	user := internal.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	database.UserCollection.InsertOne(ctx, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User Created")
}
