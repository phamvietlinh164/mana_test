package services

import (
	"context"
	"errors"
	"mana_test/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	user.Id = primitive.NewObjectID()
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(id *string) (*models.User, error) {
	var user models.User
	objID, _ := primitive.ObjectIDFromHex(*id)
	// log.Println(*id)
	query := bson.D{bson.E{Key: "_id", Value: objID}}

	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	return &user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "_id", Value: user.Id}}
	// if user.Age == 0 {
	// 	log.Println("aaaaaaa")
	// }

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "age", Value: user.Age}, bson.E{Key: "task", Value: user.Task}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(id *string) error {
	objID, _ := primitive.ObjectIDFromHex(*id)
	filter := bson.D{bson.E{Key: "_id", Value: objID}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched docoment found")
	}
	return nil
}
