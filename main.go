package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mxnuchim/golang-crud-api/controllers"
	"github.com/mxnuchim/golang-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreatePost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetSinglePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
} 