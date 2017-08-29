package models

// User - represents user account information
type User struct {
	Model

	Name       string `gorm:"size:255"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}
