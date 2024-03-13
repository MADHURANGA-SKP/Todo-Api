package initializers

import "todo_api/models"

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Todo{})
}