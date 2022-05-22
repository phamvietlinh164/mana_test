package services

import (
	"context"
	"errors"
	"mana_test/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskServiceImpl struct {
	mongoclient *mongo.Client
	ctx         context.Context
}

func NewTaskService(mongoclient *mongo.Client, ctx context.Context) TaskService {
	return &TaskServiceImpl{
		mongoclient: mongoclient,
		ctx:         ctx,
	}
}

func (u *TaskServiceImpl) CreateTask(task *models.Task) error {

	var taskNum int
	var taskToday int
	var user models.User
	query := bson.D{bson.E{Key: "_id", Value: task.Owner}}

	err := u.mongoclient.Database("mana_test").Collection("user").FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		return err
	}

	taskNum = user.Task
	today := time.Now()
	today_format := today.Format("20060102")

	query = bson.D{bson.E{Key: "owner", Value: task.Owner}, bson.E{Key: "date", Value: today_format}}
	cursor, _ := u.mongoclient.Database("mana_test").Collection("task").Find(u.ctx, query)
	for cursor.Next(u.ctx) {
		taskToday = taskToday + 1
	}
	// log.Println(taskNum)
	// log.Println(taskToday)
	if taskToday >= taskNum {
		return errors.New("task today is exceed task number")
	}
	task.Id = primitive.NewObjectID()
	task.Date = today_format
	_, err = u.mongoclient.Database("mana_test").Collection("task").InsertOne(u.ctx, task)
	return err
}

func (u *TaskServiceImpl) GetTask(id *string) (*models.Task, error) {
	var task models.Task
	objID, _ := primitive.ObjectIDFromHex(*id)
	// log.Println(objID)
	query := bson.D{bson.E{Key: "_id", Value: objID}}

	err := u.mongoclient.Database("mana_test").Collection("task").FindOne(u.ctx, query).Decode(&task)

	return &task, err
}

func (u *TaskServiceImpl) GetAll() ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := u.mongoclient.Database("mana_test").Collection("task").Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(tasks) == 0 {
		return nil, errors.New("documents not found")
	}
	return tasks, nil
}

func (u *TaskServiceImpl) UpdateTask(task *models.Task) error {

	filterCount := bson.D{bson.E{Key: "_id", Value: task.Owner}}
	count, err := u.mongoclient.Database("mana_test").Collection("user").CountDocuments(u.ctx, filterCount)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("User not exist")
	}

	filter := bson.D{bson.E{Key: "_id", Value: task.Id}}
	// if user.Age == 0 {
	// 	log.Println("aaaaaaa")
	// }

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "taskname", Value: task.TaskName}, bson.E{Key: "owner", Value: task.Owner}, bson.E{Key: "status", Value: task.Status}, bson.E{Key: "date", Value: task.Date}}}}
	result, err := u.mongoclient.Database("mana_test").Collection("task").UpdateOne(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found")
	}
	return err
}

func (u *TaskServiceImpl) DeleteTask(id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "_id", Value: objID}}
	result, _ := u.mongoclient.Database("mana_test").Collection("task").DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched docoment found")
	}
	return nil
}
