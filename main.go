package main

import (
	"todo_api/controllers"
	"todo_api/initializers"
	"todo_api/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	router := gin.Default()

	//user defined endpoint routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth ,controllers.Validate)
	router.PUT("/edituser/:id", controllers.UpdateUser)
	router.DELETE("/deleteuser/:id", controllers.DeleteUser)

	//todo defined endpoint routes
	router.POST("/createtodo", controllers.CreateTodo)
	router.GET("/gettodo", controllers.GetTodo)
	router.PUT("updatetodo", controllers.UpdateTodo)
	router.DELETE("/deletetodo", controllers.DeleteTodo)
 

 	router.Run()
}