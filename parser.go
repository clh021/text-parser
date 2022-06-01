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
			fmt.Printf("Error: Do not support formType '%s'.\n", conf.FormType)
		}

		txtArr := []string{conf.Text}

		for _, p := range conf.Pipes {
			fmt.Printf("%T %+v \t", p.Cmd, p.Cmd)
			fmt.Printf("%T %+v \n", p.Params, p.Params)
			po := pipes.PipeObj{}
			meth := reflect.ValueOf(po).MethodByName(p.Cmd)
			if meth.IsValid() {
				res := meth.Call([]reflect.Value{
					reflect.ValueOf(p.Params),
					reflect.ValueOf(txtArr)},
				)
				fmt.Printf("%+v", res)
			} else {
				fmt.Printf("Error: Do not Support PipeMethod '%+v'.\n", p.Cmd)
			}
		}
		// callback FuncHandle by string(funcName)
	}
}
