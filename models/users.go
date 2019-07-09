package models

// User is type details
type User struct {
	ID       int
	Name     string    `json:"name" binding:"required,min=3,max=15"`
	Email    string    `gorm:"unique_index" json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=6"`
	Article  []Article `gorm:"association_autoupdate:false;association_autocreate:false" json:",omitempty"`
}

//ShortUser is ...
type ShortUser struct {
	ID   int
	Name string
}

//Login is ...
type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// ListUsers is struct
type ListUsers struct {
	ID    int
	Name  string
	Email string
}

// ToListUsers is ...
func (u User) ToListUsers() ListUsers {

	return ListUsers{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
