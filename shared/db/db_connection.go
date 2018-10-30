package db

import (
	"github.com/jinzhu/gorm"
	"github.com/riona/blog/config"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	config *config.Config
}

func NewConnection() (*Connection, error) {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error cant't read config file err: %v", err)
		return nil, err
	}
	c := &Connection{config: config}
	return c, nil
}

func (c *Connection) Connect() (*gorm.DB, error) {
	config := c.config
	data := "%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(config.DB.Adapter, fmt.Sprintf(data, config.DB.User, config.DB.Password, config.DB.Database))
	if err != nil {
		log.Fatalf("Error can't Connection db err: %v", err)
		return nil, err
	}
	return db, nil
}
