package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	Id       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	TaskName string             `json: "taskname" bson: "taskname"`
	Owner    primitive.ObjectID `json: "owner" bson: "owner"`
	Status   string             `json: "status" bson: "status"`
	Date     string             `json: "date" bson: "date"`
}
