package blog

import (
	"os"
	"github.com/jinzhu/gorm"
	"fmt"
)

func connect() error {
	var c Config
	switch os.Getenv("GO_ENV") {
	default:
		conf, err := SetConfig()
		if err != nil {
			return err
		}
		fmt.Sprintf("%s:%s@/gorm?charset=utf8&parseTime=True&loc=Local",
			conf.Development.Host,
			conf.Development.Password)
		gorm.Open("mysql", )
	}

}

func FindAll() {

}
