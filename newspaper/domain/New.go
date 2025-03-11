package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type News struct {
	Id              primitive.ObjectID `bson:"_id,omitempty"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	Content         string             `json:"content"`
	Topic           string             `json:"topic"`
	Author          string             `json:"author"`
	PublicationDate string             `json:"PublicationDate"`
}
