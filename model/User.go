package model

import "time"

// https://gorm.io/docs/models.html

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `form:"name"`
	Surname   string `form:"surname"`
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	Level     string `gorm:"default:normal"`
	CreatedAt time.Time
}
