package models

// User is type details
type User struct {
	ID       int
	Name     string    `json:"name" binding:"required,min=3,max=15"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=6"`
	Article  []Article `gorm:"association_autoupdate:false;association_autocreate:false" json:",omitempty"`
}

//ShortUser is ...
type ShortUser struct {
	ID   int
	Name string
}
