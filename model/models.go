package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // "owner", "tenant", "admin"
}

type Property struct {
	gorm.Model
	Title     string
	Location  string
	Latitude  float64
	Longitude float64
	OwnerID   uint
	Views     int
}

type Booking struct {
	gorm.Model
	PropertyID uint
	TenantID   uint
	Status     string
}

type Message struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint
	Content    string
}
