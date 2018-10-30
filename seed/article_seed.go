package main

import (
	"fmt"
	"time"
	"github.com/riona/blog/model"
	"github.com/riona/blog/shared/db"
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		title := fmt.Sprintf("title%d", i+1)
		body := fmt.Sprintf("おはよう%d", i+1)
		now := time.Now()
		date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

		a := model.Article{
			Title: title,
			Body:  body,
			Date:  date,
		}



		c, err := db.NewConnection()
		if err != nil {
			panic(err)
		}

		db, err := c.Connect()
		if err != nil {
			panic(err)
		}
		err = db.Create(&a).Error

		if err != nil {
			log.Fatalf("Error can't create article err: %v", err)
			panic(err)
		}
	}
}
