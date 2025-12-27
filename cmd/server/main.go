package main

import (
	"api-with-golang/configs"
	"api-with-golang/internal/controllers"
	"api-with-golang/internal/database"
	"api-with-golang/internal/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()

	configs.ConnectMongo("mongodb://localhost")

	client := configs.Client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to MongoDB successfully!")

	db := client.Database("shopsphere")
	database.InitDB(db)

	app.HandleFunc("/api/register", services.Register).Methods("POST")
	app.HandleFunc("/api/users", controllers.User).Methods("GET")
	app.HandleFunc("/api/login", services.LoginUser).Methods("POST")
	app.HandleFunc("/api/user-delete/{id}", services.UserDelete).Methods("DELETE")
	app.HandleFunc("/api/orders", services.OrderHandler).Methods("POST")

	products := configs.Collection("shopsphere", "products")
	fmt.Println("Collection ready:", products.Name())

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, app); err != nil {
		log.Fatal("Server error:", err)
	}
}
