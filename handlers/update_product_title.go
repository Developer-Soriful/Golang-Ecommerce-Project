package handlers

import (
	datalist "api-with-golang/data_list"
	"encoding/json"
	"net/http"
	"strconv"
)

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
