package blog

import (
	"net/http"
	"time"
)

type Article struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	TagIds    []int     `db:"tag_ids"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *Article) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	}
}
