package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OutletModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone_number"`
	CreatedAt time.Time          `bson:"created_at"`
}
