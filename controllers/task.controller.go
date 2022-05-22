package controllers

import (
	"mana_test/models"
	"mana_test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService services.TaskService
}

func NewTask(taskservice services.TaskService) TaskController {
	return TaskController{
		TaskService: taskservice,
	}
}

func (uc *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.TaskService.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *TaskController) GetTask(ctx *gin.Context) {
	taskid := ctx.Param("id")
	user, err := uc.TaskService.GetTask(&taskid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *TaskController) GetAll(ctx *gin.Context) {
	tasks, err := uc.TaskService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (uc *TaskController) UpdateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.TaskService.UpdateTask(&task)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.TaskService.DeleteTask(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *TaskController) RegisterTaskRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/task")
	userroute.POST("/create", uc.CreateTask)
	userroute.GET("/get/:id", uc.GetTask)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateTask)
	userroute.DELETE("/delete/:id", uc.DeleteTask)
}
