package data

import (
	_ "embed"
	"gopkg.in/yaml.v2"
)

//go:embed data.yaml
var fileData []byte

type Data struct {
	Texts map[int]string `yaml:"texts"`
}

func Get() (Data, error) {
	var data Data
	err := yaml.Unmarshal(fileData, &data)
	return data, err
}
