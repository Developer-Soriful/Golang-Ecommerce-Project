package database

import "go.mongodb.org/mongo-driver/mongo"

var UserCollection *mongo.Collection

func InitDB(db *mongo.Database) {
	UserCollection = db.Collection("users")
}
