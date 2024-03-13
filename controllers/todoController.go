package controllers

import (
	"net/http"
	"todo_api/initializers"
	"todo_api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
		// Get todo fields from  body
	var body struct {
		UserID    uint
		Title     string
		Time 	string
		Date string
		Completed bool
	}
    
    // get todo data from body and check wheather it successfull or not
    if c.Bind(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":"Failed to read body"})
        return
    }

	var newTodo models.Todo
    // bind the todo with the logged-in user
    newTodo.UserID = body.UserID
	newTodo.Title = body.Title
	newTodo.Time = body.Time
	newTodo.Date = body.Date

    // Save the new todo to the database
    initializers.DB.Create(&newTodo)

    c.JSON(http.StatusCreated, gin.H{"message": "Todo created successfully"})
}

func GetTodo(c *gin.Context) {
		// Get todo fields from  body
	var body struct {
		UserID    uint
		Title     string
		Time 	string
		Date string
		Completed bool
	}
	// get todo data from body and check wheather it successfull or not
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Failed to read body"})
	}


    //Get todos by user ID
    var todos []models.Todo
    initializers.DB.Where("user_id = ?", body.UserID).Find(&todos)

    // Return todos as JSON response
    c.JSON(http.StatusOK, gin.H{"message":"Todo successfully got"})
}

func UpdateTodo(c *gin.Context) {
	// Get todo fields from  body
	var body struct {
		UserID    uint
		Title     string
		Time 	string
		Date string
		Completed bool
	}
	// get todo data from body and check wheather it successfull or not
    if c.Bind(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
        return
    }

	todoID := c.Param("id")
    // Find the existing todo by ID
    var todo models.Todo
    result := initializers.DB.First(&todo, "userID = ?", todoID)
    if result != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    // Verify that the user owns this todo
    userID := c.MustGet("userID").(uint)
    if todo.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "User ont found"})
        return
    }

    // Update todo
    todo.Title = body.Title
	todo.Time = body.Time
	todo.Date = body.Date
    todo.Completed = body.Completed

    // Save changes to the database
    initializers.DB.Save(&todo)

    c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func DeleteTodo(c *gin.Context) {
	// Find the existing todo by ID
    todoID := c.Param("id")
    var todo models.Todo
    result := initializers.DB.First(&todo, todoID)
    if result != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    // Verify that the user owns this todo
    userID := c.MustGet("userID").(uint)
    if todo.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "User ont found"})
        return
    }

    // Delete the todo from the database
    initializers.DB.Delete(&todo)

    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}