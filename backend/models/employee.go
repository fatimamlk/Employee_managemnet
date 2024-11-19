package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName   string             `bson:"first_name" json:"first_name"`
	LastName    string             `bson:"last_name" json:"last_name"`
	Email       string             `bson:"email" json:"email"`
	PhoneNumber string             `bson:"phone_number" json:"phone_number"`
	Position    string             `bson:"position" json:"position"`
	Department  string             `bson:"department" json:"department"`
	DateOfHire  string             `bson:"date_of_hire" json:"date_of_hire"`
}
