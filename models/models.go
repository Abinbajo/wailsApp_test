package models

import (
	
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

type User struct {
	ID			  	primitive.ObjectID		`json:"_id" bson:"_id"`
	Email		  	*string					`json:"email" validate:"email, required"`
	Password	  	*string					`json:"password" validate:"required,min=6"`


}