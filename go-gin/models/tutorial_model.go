package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tutorial struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Published   bool               `json:"published"`
	CreatedAt   string             `json:"createdAt"`
	UpdatedAt   string             `json:"updatedAt"`
	Id          primitive.ObjectID `bson:"_id"`
}
