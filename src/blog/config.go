package blog

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/sksksks/blog"
)

type Config struct {
	Development struct {
		Adapter  string `db:"adapter"`
		Encoding string `db:"encoding"`
		Database string `db:"database"`
		Pool     string `db:"pool"`
		Host     string `db:"host"`
		Port     string `db:"port"`
		User     string `db:"root"`
		Password string `db:"password"`
	}
}

func SetConfig() (*Config, error) {
	f, err := asset.Assets.Open("config.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var conf Config
	yaml.Unmarshal(b, &conf)
	return &conf, nil
}
