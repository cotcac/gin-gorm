package models

// Category is ...
type Category struct {
	ID       int
	Title    string    `json:"title" binding:"required,min=3,max=15"`
	Articles []Article `gorm:"many2many:category_article" json:",omitempty"`
}
