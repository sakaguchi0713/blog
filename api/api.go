package api

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

func (f *finder) FindArticleForIndex() map[time.Time][]string {
	var articles []model.Article
	f.db.Find(&articles)

	index := map[time.Time][]string{}
	for _, art := range articles {
		_, ok := index[art.Date]
		if !ok {
			index[art.Date] = []string{art.Title}
		} else {
			index[art.Date] = append(index[art.Date], art.Title)
		}
	}

	return index
}

func (f *finder) FindArticleByDate(postTime time.Time) []model.Article {
	var articles []model.Article
	f.db.Find(&articles, "date=?", postTime)
	return articles
}

func (f *finder) FindAllArticles() []model.Article {
	var articles []model.Article
	f.db.Find(&articles)
	return articles
}
