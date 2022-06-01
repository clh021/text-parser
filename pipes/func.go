package pipes

import (
	"fmt"
	"strings"
)

func (p *PipeObj) Split(params []string, txt []string) error {
	p.lastArr = strings.Split("a,b,c", params[0])
	fmt.Printf("params: %+v \n", params)
	// fmt.Printf("txt: %+v \n", txt)
	return nil
}
