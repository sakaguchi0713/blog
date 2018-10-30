package model

import "time"

type Article struct {
	Title string    `db:"title"`
	Body  string    `db:"body"`
	Date  time.Time `db:"date"`
}
