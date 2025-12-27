package internal

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ORDER POST SCHEMA
type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId    string             `bson:"userId" json:"userId"`
	TotalFees float64            `bson:"totalFees" json:"totalFees"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}
