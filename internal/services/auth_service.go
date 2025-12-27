package services

import (
	"api-with-golang/internal/database"
	"api-with-golang/internal/middlewares"
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	result, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, "User Createation faild", http.DefaultMaxHeaderBytes)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.InsertedID)
}

// DELETE USER FROM DATA BASE
func UserDelete(w http.ResponseWriter, r *http.Request) {
	// 	THIS IS FOR HANDLE CORS
	middlewares.HandleCors(w, r)
	if r.Method != http.MethodDelete {
		http.Error(w, "You have not entered DELETE Request", http.StatusBadRequest)
		return
	}
	// THIS IS FOR CONTEXT
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// UserID := r.PathValue("id")
	vars := mux.Vars(r)
	UserID := vars["id"]
	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}
	result, err := database.UserCollection.DeleteOne(
		ctx,
		bson.M{"_id": objID},
	)
	fmt.Println(UserID)
	if result.DeletedCount == 0 {
		http.Error(w, "no user found with this id", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted Successful",
	})

}
