package handlers

import (
	datalist "api-with-golang/data_list"
	"net/http"
	"strconv"
)

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
