package textParser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/clh021/text-parser/config"
	"github.com/linakesi/lnksutils"
)

type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

// fmt.printf
func ParseText(configs config.Config) {
	// fmt.Printf("%+v \n", text)

	for name, conf := range configs {
		fmt.Println(name, "------>>>")
		fmt.Printf("%+v\n", conf)
		switch conf.FormType {
		case "file":
			if lnksutils.IsFileExist(conf.FormSource) {
				content, err := ioutil.ReadFile(conf.FormSource)
				if err != nil {
					log.Fatal(err)
				}
				conf.Text = string(content)
			} else {
				fmt.Printf("%+v 文件不存在\n", conf.FormSource)
			}
		case "command":
			var outbuf, errbuf strings.Builder
			var cmd *exec.Cmd
			cmd = exec.Command(conf.FormSource)
			cmd.Stdout = &outbuf
			cmd.Stderr = &errbuf
			err := cmd.Run()
			if err != nil {
				fmt.Println("RUN error:", err)
				os.Exit(-1)
				return
			}
			conf.Text = outbuf.String()
		default:
			fmt.Printf(" Do not support formType '%s'.\n", conf.FormType)
		}
		for _, p := range conf.Pipes {
			fmt.Printf("%+v \t", p.Cmd)
			fmt.Printf("%+v \n", p.Params)
		}
		// callback FuncHandle by string(funcName)
	}
}
