package textParser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/clh021/text-parser/config"
	"github.com/clh021/text-parser/lib"
	"github.com/clh021/text-parser/log"
	"github.com/clh021/text-parser/pipes"
	"github.com/linakesi/lnksutils"
)

func ParseTextFormFile(source string) string {
	currentPath, err := lib.GetProgramPath()
	if err != nil {
		log.Panic(err)
	}
	file := filepath.Join(currentPath, source)
	if !lnksutils.IsFileExist(file) {
		log.Panicf("Error: file not found: %+v\n", file)
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	return string(content)
}

func ParseTextFormCommand(pathBin string, args ...string) string {
	var outbuf, errbuf strings.Builder
	cmd := exec.Command(pathBin, args...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "TEXT_PARSER=1")
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	if err != nil {
		log.Panic("Error: cmd run error:", err)
	}
	return outbuf.String()
}

func ParseText(configs config.Config) {
	for name, conf := range configs {
		log.Infof("name(%+v): %+v\n", name, conf)

		// 根据配置的获取不同类型的文本来源
		switch conf.FormType {
		case "file":
			filepath := conf.FormSource[0]
			conf.Text = ParseTextFormFile(filepath)
		case "command":
			pathBin, args := conf.FormSource[0], conf.FormSource[1:]
			conf.Text = ParseTextFormCommand(pathBin, args...)
		default:
			log.Errorf("Error: Do not support formType '%s'.\n", conf.FormType)
		}

		// 初始化文本处理中心
		po := &pipes.PipeObj{}

		// 传入预处理文本 最终也通过po.Get** 等方式取出处理后的文本
		po.Start(conf.Text)

		// 循环出所有文本处理配置
		for _, p := range conf.Pipes {
			log.Debugf("Pipes:cmd: type=%T, value=%+v \t", p.Cmd, p.Cmd)
			log.Debugf("Pipes:params: type=%T, value=%+v \n", p.Params, p.Params)

			// 验证配置的处理方式是否已经支持
			meth := reflect.ValueOf(po).MethodByName(p.Cmd)
			if !meth.IsValid() {
				log.Errorf("Error: Do not Support PipeMethod '%+v'.\n", p.Cmd)
			}

			// 调用配置的处理方式，得到结果
			calledResult := meth.Call([]reflect.Value{
				reflect.ValueOf(p.Params),
			})

			// 确认处理没有发生错误
			err := calledResult[0].Interface() // calledResult 返回的是多个值(虽然函数就一个值)
			if err != nil {
				log.Errorf("Pipes: Error calling %s: %v", p.Cmd, err)
			}

			// 根据配置判断是否输出日志
			if conf.Debug {
				// log.Debugf("%+v", po.GetStr())
				lastArrJSON, err := json.Marshal(po.GetArr())
				if err != nil {
					log.Errorf("Error: %s", err.Error())
				}
				log.Debugf("Pipes: lastArr: %+v", string(lastArrJSON))
			}
		}
	}
}
