package main

import (
	"bytes"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
name: steve
hobbies:
- go
clothing:
  jacket: leather
  trousers: denim
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	viper.Get("name")
}
