package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	//go:embed conf.yaml
	content []byte
	conf    Conf
)

func init() {
	if err := yaml.Unmarshal(content, &conf); err != nil {
		log.Fatalf("parser conf.yaml error: %v", err)
	}
}

func GetConfig() Conf {
	return conf
}

type Conf struct {
	Env   string `yaml:"env"`
	Ports struct {
		Http string `yaml:"http"`
		Grpc string `yaml:"grpc"`
	} `yaml:"ports"`
	Mongo struct {
		Uri    string `yaml:"uri"`
		DbName string `yaml:"dbname"`
	} `yaml:"mongo"`
}
