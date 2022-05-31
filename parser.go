package textParser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/clh021/text-parser/config"
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
			fmt.Printf(" Do not support formType '%s'.\n", conf.FormType)
		}
		// ParseText(conf.Text)
	}
}
