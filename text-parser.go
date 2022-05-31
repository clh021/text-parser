package textParser

import (
	"fmt"
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

func Run() {
	file := ParseParam()
	if WaitFile(file) {
		configs, err := config.LoadConfig(file)
		if err != nil {
			fmt.Println(err)
		}
		for key, value := range *configs {
			fmt.Println(key, "------>>>")
			fmt.Printf("%+v\n", value)
		}
		// for conf := range *configs {
		// 	switch conf.FormType {
		// 	case "file":
		// 		lnksutils.IsFileExist(conf.FormSource)
		// 	case "command":
		// 		lnksutils.IsFileExist(conf.FormSource)
		// 	default:
		// 		fmt.Printf(" Do not support formType '%s'.\n", conf.FormType)
		// 	}
		// 	// parseText(conf.DataSource)
		// 	fmt.Println("exist file", conf, err)
		// }
	} else {
		fmt.Println("not exist file")
	}
}
