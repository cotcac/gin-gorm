package models

// users type details
type User struct {
	ID   int
	Name string
	Article []Article `gorm:"association_autoupdate:false;association_autocreate:false"`
  }