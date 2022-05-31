package config

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
	Cmd    string   `json:"cmd"`
	Params []string `json:"params"`
}

type TextForm struct {
	FormType   string `json:"formType"`
	FormSource string `json:"formSource"`
	Debug      bool   `json:"debug"`
	Pipes      []Pipe `json:"pipes"`
}

type Config map[string]TextForm
