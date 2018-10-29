package api

import (
	"github.com/riona/blog/model"
	"time"
)

type articleFinder interface {
	FindArticleForIndex() (map[time.Time]string, error)
	FindArticleByDate(postTime time.Time) ([]model.Article, error)
	FindAllArticles() ([]model.Article, error)
}

type finder struct{}

func (p *finder) FindArticleForIndex() (map[time.Time]string, error) {
	return nil, nil
}

func (p *finder) FindArticleByDate(postTime time.Time) ([]model.Article, error) {
	return nil, nil
}

func (p *finder) FindAllArticles() ([]model.Article, error) {
	return nil, nil
}
