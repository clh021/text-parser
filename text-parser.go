package textParser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/clh021/text-parser/config"
	"github.com/linakesi/lnksutils"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

/**
 * WaitFile
 * TODO: 等待文件生成
 * 此操作是为了保证文件生成后第一时间分析出结果
 */
func WaitFile(file string) bool {
	return lnksutils.IsFileExist(file)
}

func GetProgramPath() string {
	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	return filepath.Dir(exReal)
}

func ParseParam() string {
	defaultConfPath := filepath.Join(GetProgramPath(), "config.json")
	return defaultConfPath
}

func ParseText(text string) {
	fmt.Printf("%+v \n", text)
}

func Run() {
	file := ParseParam()
	if WaitFile(file) {
		configs, err := config.LoadConfig(file)
		if err != nil {
			fmt.Println(err)
		}
		for name, conf := range *configs {
			fmt.Println(name, "------>>>")
			// fmt.Printf("%+v\n", conf)
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
				lnksutils.IsFileExist(conf.FormSource)
				// conf.Text = string(content)
			default:
				fmt.Printf(" Do not support formType '%s'.\n", conf.FormType)
			}
			ParseText(conf.Text)
		}
		// 	fmt.Println("exist file", conf, err)
	} else {
		fmt.Println("not exist file")
	}
}
