package main

import (
	"context"
	"fmt"
	"log"
	"mana_test/controllers"
	"mana_test/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	taskservice    services.TaskService
	usercontroller controllers.UserController
	taskcontroller controllers.TaskController
	ctx            context.Context
	usercollection *mongo.Collection
	taskcollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	// mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoconn := options.Client().ApplyURI("mongodb+srv://vietlinhco:Ankedalinhco1@cluster0.qvtpl.mongodb.net/?retryWrites=true&w=majority")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection established")
	usercollection = mongoclient.Database("mana_test").Collection("user")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)

	// taskcollection = mongoclient.Database("mana_test").Collection("task")
	taskservice = services.NewTaskService(mongoclient, ctx)
	taskcontroller = controllers.NewTask(taskservice)

	server = gin.Default()
}
func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
	taskcontroller.RegisterTaskRoutes(basepath)
	log.Fatal(server.Run(":8000"))
}
