package handlers

import (
	datalist "api-with-golang/data_list"
	"api-with-golang/databse"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	// Handle CORS preflight requests
	HandleCors(w, r)
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	// this is for client data from request body
	var updatedData databse.ItemList
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// find and update the item
	for i, item := range datalist.Items {
		if item.ID == int64(id) {
			datalist.Items[i].Title = updatedData.Title
			datalist.Items[i].Content = updatedData.Content
			datalist.Items[i].Updated = time.Now().Format(time.RFC3339)
			// Update logic goes here
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(datalist.Items[i])
			return
		}
	}
	http.Error(w, "Product Not update yet..!", http.StatusNotFound)
}
