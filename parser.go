package textParser

import (
	"fmt"

	"github.com/clh021/text-parser/config"
)

// fmt.printf
func ParseText(configs config.Config) {
	// fmt.Printf("%+v \n", text)

	for name, conf := range configs {
		fmt.Println(name, "------>>>")
		fmt.Printf("%+v\n", conf)
		// switch conf.FormType {
		// case "file":
		// 	if lnksutils.IsFileExist(conf.FormSource) {
		// 		content, err := ioutil.ReadFile(conf.FormSource)
		// 		if err != nil {
		// 			log.Fatal(err)
		// 		}
		// 		conf.Text = string(content)
		// 	} else {
		// 		fmt.Printf("%+v 文件不存在\n", conf.FormSource)
		// 	}
		// case "command":
		// 	lnksutils.IsFileExist(conf.FormSource)
		// 	// conf.Text = string(content)
		// default:
		// 	fmt.Printf(" Do not support formType '%s'.\n", conf.FormType)
		// }
		// ParseText(conf.Text)
	}
}
