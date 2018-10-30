package article

import (
	"github.com/jinzhu/gorm"
	"github.com/riona/blog/model"
	"github.com/riona/blog/shared/db"
	"time"
)

type ArticleFinder interface {
	FindArticleForIndex() (map[time.Time][]string, error)
	FindArticleByDate(postTime time.Time) ([]model.Article, error)
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

func (f *finder) FindArticleForIndex() map[int][]model.Article {
	var articles []model.Article
	f.db.Find(&articles)

	index := map[int][]model.Article{}
	for _, art := range articles {
		_, ok := index[art.ID]
		if !ok {
			index[art.ID] = []model.Article{art}
		} else {
			index[art.ID] = append(index[art.ID], art)
		}
	}

	return index
}

func (f *finder) findArticleByDate(postTime time.Time) []model.Article {
	var articles []model.Article
	f.db.Find(&articles, "date=?", postTime)
	return articles
}

func (f *finder) findAllArticles() []model.Article {
	var articles []model.Article
	f.db.Find(&articles)
	return articles
}
