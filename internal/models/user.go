package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Role      string             `bson:"role" json:"role"` // "admin", "vendor", "customer"
	CreatedAt string             `bson:"created_at" json:"created_at"`
	UpdatedAt string             `bson:"updated_at" json:"updated_at"`
}
type Login struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
