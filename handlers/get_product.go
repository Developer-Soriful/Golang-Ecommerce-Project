package handlers

import (
	datalist "api-with-golang/data_list"
	"encoding/json"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	json.NewEncoder(w).Encode(datalist.Items)
}
