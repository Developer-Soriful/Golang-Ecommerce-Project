package authservice

import (
	"api-with-golang/internal/database"
	"api-with-golang/internal/middlewares"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
