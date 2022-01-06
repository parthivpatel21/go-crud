package models

import "gorm.io/gorm"

// User schema of the user table
type User struct {
	gorm.Model
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}
