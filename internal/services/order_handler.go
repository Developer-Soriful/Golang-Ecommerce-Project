package services

import (
	"api-with-golang/internal/database"
	internal "api-with-golang/internal/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ORDER CREATE

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	// METHOD
	if r.Method != http.MethodPost {
		http.Error(w, "You have not entered POST Method ", http.StatusBadRequest)
		return
	}
	var req internal.Order
	json.NewDecoder(r.Body).Decode(&req)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	orders := internal.Order{
		ID:        primitive.NewObjectID(),
		UserId:    req.UserId,
		TotalFees: req.TotalFees,
		Status:    req.Status,
		CreatedAt: time.Now().String(),
	}
	result, err := database.UserCollection.InsertOne(ctx, orders)
	if err != nil {
		http.Error(w, "Failed to create order. Please try again.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order Created Successful.",
		"order":   result.InsertedID,
	})
}
