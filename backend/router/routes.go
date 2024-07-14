package router

import (
	"goTodo/controller"
	"net/http"
)

var todoListRoutes = Routes{
	Route{"Get all tasks", http.MethodGet, "/tasks", controller.GetAllTasks},
	Route{"Delete all tasks", http.MethodDelete, "/tasks", controller.DeleteAllTasks},
	Route{"Create task", http.MethodPost, "/task", controller.CreateTask},
	Route{"Get task", http.MethodGet, "/tasks/:id", controller.GetTask},
	Route{"Delete task", http.MethodDelete, "/task/:id", controller.DeleteTask},
	// Route{"Edit task", http.MethodPut, "/task/:id", controller.EditTask},
	// Route{"Update Status", http.MethodPatch, "/task/:id", controller.UpdateStatus},
}
