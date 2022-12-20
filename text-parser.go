package textParser

import (
	"path/filepath"

	"github.com/clh021/text-parser/lib"
	"github.com/clh021/text-parser/log"
	parseconf "github.com/clh021/text-parser/parse-conf"
	"github.com/linakesi/lnksutils"
)

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
		configs, err := parseconf.LoadConfig(file)
		if err != nil {
			log.Panic(err)
		}
		// fmt.Printf("%+v", configs)
		ParseText(*configs)
	} else {
		log.Error("not exist file:", file)
	}
}
