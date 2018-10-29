package api

import (
	"github.com/riona/blog/model"
	"time"
)

type articleGetter interface {
	GetIndex() (map[time.Time]string, error)
	GetArticle(postTime time.Time) ([]model.Article, error)
}

type publisher struct {}

func (p *publisher) GetIndex() (map[time.Time]string, error) {
	return nil, nil
}

func (p *publisher) GetArticle(postTime time.Time) ([]model.Article, error) {
	return nil, nil
}
