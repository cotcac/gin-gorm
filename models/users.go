package models

// users type details
type User struct {
	ID   int
	Name string `json:"name" binding:"required,min=3,max=15"`
	Email string `json:"email" binding:"required,email"`
	Article []Article `gorm:"association_autoupdate:false;association_autocreate:false" json:",omitempty"`
  }