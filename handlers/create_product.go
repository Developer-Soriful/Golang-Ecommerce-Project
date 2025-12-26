package handlers

import (
	datalist "api-with-golang/data_list"
	"api-with-golang/databse"
	"encoding/json"
	"net/http"
	"time"
)

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
