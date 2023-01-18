package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/clh021/text-parser/parse-conf/df"
	"github.com/clh021/text-parser/parse-conf/env"
	"github.com/spf13/viper"
)

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func bindOSArgs() {
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "--") {
			format := os.Args[1]
			viper.SetDefault("format", format[2:])
			// viper.SetDefault("command", os.Args[2])
		}
	}
}

func printJson(v any) {
	jsonStr, _ := json.Marshal(v)
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, jsonStr, "", "\t")
	if error != nil {
		fmt.Println("JSON parse error: ", error)
	} else {
		fmt.Println(prettyJSON.String())
	}
}

func main() {
	fmt.Printf("Build: %s\n", build)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	var conf config
	if err != nil {
		fmt.Println(err)
		bindOSArgs()
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(viper.AllSettings())
	// fmt.Println("config:", conf)
	switch conf.Format {
	case "df":
		printJson(df.GetFileSystem())

	case "env":
		printJson(env.GetEnv())

	default:
		fmt.Printf("parsing this format '%v' is not currently supported.", conf.Format)
		fmt.Println()
	}
}
