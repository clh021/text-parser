package textParser

import (
	"github.com/go-playground/validator"
	"github.com/linakesi/lnksutils"
)

func LoadConfig(fpath string) (*Config, error) {
	var conf Config
	err := lnksutils.FileToJSON(fpath, &conf)
	if err != nil {
		return nil, err
	}
	v := validator.New()
	v.SetTagName("binding")
	err = v.Struct(conf)
	return &conf, err
}

type Pipe struct {
	Cmd    string   `json:cmd`
	Params []string `json:params`
}

type Config struct {
	File  string `json:"file"`
	Pipes []Pipe `json:"pipes"`
	Debug bool   `json:"debug"`
}
