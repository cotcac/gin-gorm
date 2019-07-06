package models
// import "github.com/jinzhu/gorm"
// users type details
type Category struct {
	Id   int
	Title string `json:"title" binding:"required,min=3,max=15"`
	Articles []Article `gorm:"many2many:category_article;"`
  }