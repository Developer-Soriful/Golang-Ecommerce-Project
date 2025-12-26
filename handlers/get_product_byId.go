package handlers

import (
	datalist "api-with-golang/data_list"
	"encoding/json"
	"net/http"
	"strconv"
)

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
