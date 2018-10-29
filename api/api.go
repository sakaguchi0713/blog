package api

import (
	"github.com/riona/blog/model"
)

type articleGetter interface {
	GetIndexs() ([]model.Article, error)
}



type publisher struct {}

func (p *publisher) GetIndex() ([]model.Article, error) {

}
