package controller

import (
	"errors"
	"goTodo/constant"
	"goTodo/database"
	"goTodo/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTasks handler to retrieve all tasks
func GetAllTasks(c *gin.Context) {
	data, err := database.Mgr.FetchAll(constant.CollectionName)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

// DeleteAllTasks handler to delete all tasks
func DeleteAllTasks(c *gin.Context) {
	err := database.Mgr.DeleteAll(constant.CollectionName)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// CreateTask handler to create a new task
func CreateTask(c *gin.Context) {
	var newTask types.ListItemDao
	err := c.BindJSON(&newTask) // Use BindJSON for decoding request body
	if err != nil {
		handleError(c, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	newTask.Status = false
	newTask.Task = "Bakchodi1"

	insertedID, err := database.Mgr.Insert(newTask, constant.CollectionName)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	newTask.ID = insertedID.(primitive.ObjectID)

	c.JSON(http.StatusCreated, gin.H{"error": false, "message": "success", "data": newTask})
}

// GetTask handler to retrieve a specific task by ID
func GetTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, errors.New("invalid task ID"))
		return
	}

	data, err := database.Mgr.Fetch(id, constant.CollectionName)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			handleError(c, http.StatusNotFound, errors.New("task not found"))
		} else {
			handleError(c, http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, data)
}

// DeleteTask handler to delete a specific task by ID
func DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		handleError(c, http.StatusBadRequest, errors.New("invalid task ID"))
		return
	}

	err = database.Mgr.Delete(id, constant.CollectionName)
	if err != nil {
		handleError(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func handleError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}
