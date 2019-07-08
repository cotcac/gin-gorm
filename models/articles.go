package models

type Article struct {
	ID         int
	Title      string
	UserID     int
	User       User       `json:"user,omitempty"`
	Categories []Category `gorm:"many2many:category_article;association_autoupdate:false;association_autocreate:false" json:"categories" binding:"required"`
}

//aaa
type Articles struct {
	ID         int
	Title      string
	User       ShortUser
	Categories []Category
}

func (a Article) ArticleToArticles() Articles {

	return Articles{
		ID:    a.ID,
		Title: a.Title,
		User: ShortUser{
			ID:   a.User.ID,
			Name: a.User.Name,
		},
		Categories: a.Categories,
	}
}
