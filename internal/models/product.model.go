package internal

// PRODUCT SCHEMA MODEL

type PRODUCT struct {
	Title       string  `bson:"title" json:"title"`
	Description string  `bson:"description" json:"description"`
	Price       float64 `bson:"price" json:"price"`
	Quantity    int32   `bson:"quantity" json:"quantity"`
	Images      string  `bson:"images" json:"images"`
	ID          string  `bson:"_id" json:"_id"`
}
