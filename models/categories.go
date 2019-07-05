package models
// import "github.com/jinzhu/gorm"
// users type details
type Category struct {
	Id   int
	Title string
	Articles []*Article `gorm:"many2many:category_article;"`
  }