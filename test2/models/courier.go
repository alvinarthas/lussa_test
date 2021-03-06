package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Courier collection
type Courier struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	Slug        string             `bson:"slug" json:"slug"`
	Description string             `bson:"description" json:"description"`
	Cost        int32              `bson:"cost" json:"cost"`
}
