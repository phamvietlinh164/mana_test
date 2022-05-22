package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id   primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name string             `json: "name" bson: "user_name"`
	Age  int                `json: "age" bson: "user_age"`
	Task int                `json: "task" bson: "user_task"`
}
