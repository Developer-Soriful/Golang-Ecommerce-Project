package main

import (
	"api-with-golang/configs"
	"api-with-golang/internal/controllers"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	app := http.NewServeMux()

	configs.ConnectMongo("mongodb://localhost")

	client := configs.Client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	fmt.Println("Connected to MongoDB successfully!")

	db := client.Database("shopsphere")

	controllers.InitDB(db)

	//
	app.HandleFunc("/api/register", controllers.Register)
	app.HandleFunc("/api/users", controllers.User)

	products := configs.Collection("shopsphere", "products")
	fmt.Println("Collection ready:", products.Name())

	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", app)
}
