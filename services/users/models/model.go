package models

import "github.com/jinzhu/gorm"

// Model - base model, contains common fields for all models
type Model struct {
	gorm.Model
}
