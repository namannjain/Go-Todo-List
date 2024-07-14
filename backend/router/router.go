package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) TodoList(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/api")
	for _, route := range todoListRoutes {
		switch route.Method {
		case http.MethodGet:
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid HTTP method with this route",
				})
			})
		}
	}
}

func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}

	// Add CORS middleware allowing all origins (for development)
	r.router.Use(cors.Default())

	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.TodoList(v1)

	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
