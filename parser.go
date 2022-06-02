package textParser

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/clh021/text-parser/config"
	"github.com/clh021/text-parser/lib"
	"github.com/clh021/text-parser/log"
	"github.com/clh021/text-parser/pipes"
	"github.com/linakesi/lnksutils"
)

func ParseTextFormFile(source string) string {
	currentPath, err := lib.GetProgramPath()
	if err != nil {
		log.Panic(err)
	}
	file := filepath.Join(currentPath, source)
	if !lnksutils.IsFileExist(file) {
		log.Panicf("Error: file not found: %+v\n", file)
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	return string(content)
}

func ParseTextFormCommand(cmdStr string) string {
	var outbuf, errbuf strings.Builder
	cmd := exec.Command(cmdStr)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	if err != nil {
		log.Panic("Error: cmd run error:", err)
	}
	return outbuf.String()
}

func ParseText(configs config.Config) {
	for name, conf := range configs {
		log.Debug("conf: ", conf)
		log.Infof("name(%+v): %+v\n", name, conf.FormType)
		switch conf.FormType {
		case "file":
			conf.Text = ParseTextFormFile(conf.FormSource)
		case "command":
			conf.Text = ParseTextFormCommand(conf.FormSource)
		default:
			log.Errorf("Error: Do not support formType '%s'.\n", conf.FormType)
		}
		log.Warnf("name(%+v): %+v\n", name, conf)
		log.Errorf("name(%+v): %+v\n", name, conf)

		po := &pipes.PipeObj{}
		po.Start(conf.Text)

		for _, p := range conf.Pipes {
			log.Debugf("%T %+v \t", p.Cmd, p.Cmd)
			log.Debugf("%T %+v \n", p.Params, p.Params)
			meth := reflect.ValueOf(po).MethodByName(p.Cmd)
			if !meth.IsValid() {
				log.Errorf("Error: Do not Support PipeMethod '%+v'.\n", p.Cmd)
			}
			result := meth.Call([]reflect.Value{
				reflect.ValueOf(p.Params),
			})
			log.Debugf("%+v", result)
			err := result[0].Interface() // result 返回的是多个值
			if err == nil {
				log.Errorf("No error returned by", p.Cmd)
			} else {
				log.Errorf("Error calling %s: %v", p.Cmd, err)
			}
			if conf.Debug {
				// log.Debugf("%+v", po.GetStr())
				log.Debugf("%+v", po.GetArr())
			}
		}
	}
}
