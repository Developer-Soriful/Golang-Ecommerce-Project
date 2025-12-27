package main

import (
	"api-with-golang/configs"
	authservice "api-with-golang/internal/authService"
	"api-with-golang/internal/controllers"
	"api-with-golang/internal/database"
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

	app.HandleFunc("/api/register", authservice.Register).Methods("POST")
	app.HandleFunc("/api/users", controllers.User).Methods("GET")
	app.HandleFunc("/api/login", controllers.LoginUser).Methods("POST")
	app.HandleFunc("/api/user-delete/{id}", authservice.UserDelete).Methods("DELETE")

	products := configs.Collection("shopsphere", "products")
	fmt.Println("Collection ready:", products.Name())

	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
