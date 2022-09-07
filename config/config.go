package config

import (
	"github.com/linakesi/lnksutils"
)

func LoadConfig(fpath string) (*Config, error) {
	var conf Config
	err := lnksutils.FileToJSON(fpath, &conf)
	return &conf, err
}

type Pipe struct {
	Cmd    string   `json:"cmd"`
	Params []string `json:"params"`
}

type TextForm struct {
	FormType   string   `json:"formType" validate:"required,oneof=command file2"`
	FormSource []string `json:"formSource" validate:"required,min=3,max=260"`
	Debug      bool     `json:"debug"`
	Pipes      []Pipe   `json:"pipes" validate:"required"`
	Text       string
}

type Config map[string]TextForm
