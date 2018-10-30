package model

import "time"

type Article struct {
	ID    int       `db:"id"`
	Title string    `db:"title"`
	Body  string    `db:"body"`
	Date  time.Time `db:"date"`
}
