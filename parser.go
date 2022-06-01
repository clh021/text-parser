package textParser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/clh021/text-parser/config"
	"github.com/clh021/text-parser/pipes"
	"github.com/linakesi/lnksutils"
)

func ParseTextFormFile(source string) string {
	if !lnksutils.IsFileExist(source) {
		fmt.Printf("Error: file not found: %+v\n", source)
		os.Exit(-1)
	}
	content, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
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
		fmt.Println("Error: cmd run error:", err)
		os.Exit(-1)
	}
	return outbuf.String()
}

// fmt.printf
func ParseText(configs config.Config) {
	// fmt.Printf("%+v \n", text)
	for name, conf := range configs {
		fmt.Println(name, "------>>>")
		fmt.Printf("%+v\n", conf)
		switch conf.FormType {
		case "file":
			conf.Text = ParseTextFormFile(conf.FormSource)
		case "command":
			conf.Text = ParseTextFormCommand(conf.FormSource)
		default:
			fmt.Printf("Error: Do not support formType '%s'.\n", conf.FormType)
		}

		txtArr := []string{conf.Text}

		for _, p := range conf.Pipes {
			// fmt.Printf("%T %+v \t", p.Cmd, p.Cmd)
			// fmt.Printf("%T %+v \n", p.Params, p.Params)
			po := &pipes.PipeObj{}
			meth := reflect.ValueOf(po).MethodByName(p.Cmd)
			if !meth.IsValid() {
				fmt.Printf("Error: Do not Support PipeMethod '%+v'.\n", p.Cmd)
			}
			result := meth.Call([]reflect.Value{
				reflect.ValueOf(p.Params),
				reflect.ValueOf(txtArr),
			})
			fmt.Printf("%+v", result)
			err := result[0].Interface() // 返回的是多个值
			if err == nil {
				fmt.Println("No error returned by", p.Cmd)
			} else {
				fmt.Printf("Error calling %s: %v", p.Cmd, err)
			}
		}
	}
}
