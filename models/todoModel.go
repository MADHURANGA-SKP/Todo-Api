package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    UserID    uint
    Title     string
	Time 	string
	Date string
    Completed bool
}


