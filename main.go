package main

import (
	datalist "api-with-golang/data_list"
	"api-with-golang/databse"
	"api-with-golang/handlers"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Main function to start the server and define routes
	app := http.NewServeMux()
	app.HandleFunc("GET /getProduct", handlers.GetProduct)
	app.HandleFunc("POST /use-post", handlers.UserPost)
	app.HandleFunc("GET /getProduct/{id}", handlers.GetProductByID)
	app.HandleFunc("PUT /updateProduct/{id}", handlers.UpdateProductByID)
	app.HandleFunc("DELETE /deleteProduct/{id}", handlers.DeleteProductByID)
	app.HandleFunc("PATCH /updateProductTitle/{id}", handlers.UpdateTitleByID)
	fmt.Println(`server is running 8080 port`)
	http.ListenAndServe(":8080", app)
}

func init() {
	item1 := databse.ItemList{
		ID:      1,
		Title:   "Test title 1",
		Content: "This is test content",
		Created: time.Now().Format(time.RFC3339),
		Updated: time.Now().Format(time.RFC3339),
	}
	item2 := databse.ItemList{
		ID:      2,
		Title:   "Test title 1",
		Content: "This is test content",
		Created: time.Now().Format(time.RFC3339),
		Updated: time.Now().Format(time.RFC3339),
	}
	item3 := databse.ItemList{
		ID:      3,
		Title:   "Test title 1",
		Content: "This is test content",
		Created: time.Now().Format(time.RFC3339),
		Updated: time.Now().Format(time.RFC3339),
	}

	datalist.Items = append(datalist.Items, item1, item2, item3)

}
