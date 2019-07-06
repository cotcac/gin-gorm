package models
// import "github.com/jinzhu/gorm"
// users type details
type Article struct {
	ID   int
	Title string
	UserID int
	User User
	Categories []Category `gorm:"many2many:category_article;association_autoupdate:false;association_autocreate:false"`
  }