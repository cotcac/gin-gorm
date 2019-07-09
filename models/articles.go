package models

// Article is gorm model
type Article struct {
	ID         int
	Title      string     `json:"title" binding:"required,min=3,max=15"`
	Content    string     `sql:"type:text;" json:"content" binding:"required,min=5"`
	UserID     int        `json:"user_id" binding:"required"`
	User       User       `json:"user,omitempty" binding:"-"`
	Categories []Category `gorm:"many2many:category_article;association_autoupdate:false;association_autocreate:false" json:"categories" binding:"required"`
}

// Articles is Article that display as json.
type Articles struct {
	ID         int
	Title      string
	User       ShortUser
	Categories []Category
}

// SingleArticle struct
type SingleArticle struct {
	ID         int
	Title      string
	Content    string
	User       ShortUser
	Categories []Category
}

// ArticleToSingleArticle is use to convert Article to Single Article
func (a Article) ArticleToSingleArticle() SingleArticle {
	return SingleArticle{
		ID:      a.ID,
		Title:   a.Title,
		Content: a.Content,
		User: ShortUser{
			ID:   a.User.ID,
			Name: a.User.Name,
		},
		Categories: a.Categories,
	}
}

// ArticleToArticles is use to convert Article struct to Articles struct
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
