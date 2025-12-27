package controllers

import (
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitDB(db *mongo.Database) {
	userCollection = db.Collection("users")
}

// AUTH CONTROLLERS
// REGISTER CONTROLLER
func Register(w http.ResponseWriter, r *http.Request) {

	// IMPLEMENTATION FOR REGISTATION GOES HERE
	HandleCors(w, r)
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
	count, _ := userCollection.CountDocuments(ctx, bson.M{"email": req.Email})
	if count > 0 {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	// HASH PASSWORD
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	// CREATE USER OBJECT
	user := internal.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hash),
		Role:      req.Role,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	userCollection.InsertOne(ctx, user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User Created")
}

// FIND USER
func User(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method != "GET" {
		http.Error(w, "You have not entered GET Method", http.StatusBadRequest)
	}
	var users []internal.User
	cursor, _ := userCollection.Find(context.Background(), bson.M{})
	cursor.All(context.Background(), &users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
