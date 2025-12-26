package handlers

import (
	datalist "api-with-golang/internal/data"
	databse "api-with-golang/internal/database"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// CREATE NEW PRODUCT
func UserPost(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method != "POST" {
		http.Error(w, "Please send post request ..!", http.StatusBadRequest)
		return
	}
	var NewItem databse.ItemList
	err := json.NewDecoder(r.Body).Decode(&NewItem)
	if err != nil {
		http.Error(w, "Please enter valid json", http.StatusBadRequest)
		return
	}
	NewItem.ID = int64(len(datalist.Items) + 1)
	NewItem.Created = time.Now().Format(time.RFC3339)
	NewItem.Updated = time.Now().Format(time.RFC3339)
	datalist.Items = append(datalist.Items, NewItem)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datalist.Items)
}

// GET ALL OF THE PRODUCTS
func GetProduct(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	json.NewEncoder(w).Encode(datalist.Items)
}

// GET PRODUCT BY ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product Id", http.StatusBadRequest)
		return
	}
	for _, item := range datalist.Items {
		if item.ID == int64(id) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Product Not Found yet...!", http.StatusNotFound)
}

// UPDATE PRODUCT BY ID
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

// DELETE PRODUCT BY ID
func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	// this is for CORS handleing preflight requests
	HandleCors(w, r)
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Product Id", http.StatusBadRequest)
		return
	}
	//find and delete the item
	for i, item := range datalist.Items {
		if item.ID == int64(id) {
			datalist.Items = append(datalist.Items[:i], datalist.Items[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Delete Request Product Successfull"))
			return
		}
	}
	http.Error(w, "Delete Request Product NotFounded", http.StatusNotFound)
}

// UPDATE PRODUCT TITLE BY ID
func UpdateTitleByID(w http.ResponseWriter, r *http.Request) {
	// this is for CORS handleing preflight requests
	HandleCors(w, r)
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid Request Title ID.", http.StatusBadRequest)
		return
	}
	var updatedTitle struct {
		Title string `json:"title"`
	}
	err = json.NewDecoder(r.Body).Decode(&updatedTitle)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	}
	//find and update the title of the item
	for i, item := range datalist.Items {
		if item.ID == int64(id) {
			datalist.Items[i].Title = updatedTitle.Title
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(datalist.Items[i])
		}
	}
	http.Error(w, "Requested Product Not Found.", http.StatusNotFound)
}
