package textParser

import (
	"path/filepath"

	"github.com/clh021/text-parser/config"
	"github.com/clh021/text-parser/lib"
	"github.com/clh021/text-parser/log"
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

func Run() {
	log.SetLog()
	currentPath, err := lib.GetProgramPath()
	if err != nil {
		log.Panic(err)
	}
	file := filepath.Join(currentPath, "config.json")
	if WaitFile(file) {
		configs, err := config.LoadConfig(file)
		if err != nil {
			log.Panic(err)
		}
		// fmt.Printf("%+v", configs)
		ParseText(*configs)
	} else {
		log.Error("not exist file")
	}
}
