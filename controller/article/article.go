package article

import (
	"github.com/jinzhu/gorm"
	"github.com/riona/blog/model"
	"github.com/riona/blog/shared/db"
)

type ArticleFinder interface {
	FindArticleForIndex() []*model.Article
	findArticleByArticle(id int) model.Article
	FindAllArticles() ([]model.Article, error)
}

type finder struct {
	db *gorm.DB
}

func NewFinder() (*finder, error) {
	c, err := db.NewConnection()
	if err != nil {
		return nil, err
	}

	db, err := c.Connect()
	if err != nil {
		return nil, err
	}

	f := &finder{db}
	return f, nil
}

func (f *finder) FindArticleForIndex() []*model.Article {
	var articles []*model.Article
	f.db.Find(&articles)
	return articles
}

func (f *finder) FindArticleByArticle(id int) model.Article {
	var articles model.Article
	f.db.Find(&articles, "id=?", id)
	return articles
}

func (f *finder) findAllArticles() []model.Article {
	var articles []model.Article
	f.db.Find(&articles)
	return articles
}
