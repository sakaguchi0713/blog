package model

import "time"

type Article struct {
	Title string
	Body  string
	Date  time.Time
}
