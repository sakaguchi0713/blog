package db

import (
	"github.com/jinzhu/gorm"
	"github.com/riona/blog/config"
	"log"
)

type Connection struct {
	config *config.Config
}

func NewConnection() (*Connection, error) {
	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	c := &Connection{config: config}
	return c, nil
}

func (c *Connection) Connect() (*gorm.DB, error) {
	config := c.config
	data := "%s:%s@/%s?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(config.Adapter, data, config.User, config.Password, config.Database)
	if err != nil {
		log.Fatalf("Error can't Connection db err: %v", err)
		return nil, err
	}
	return db, nil
}
