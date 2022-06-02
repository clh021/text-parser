package pipes

import (
	"strings"
)

func (p *PipeObj) Split(params []string) error {
	p.lastArr = strings.Split(p.lastStr, params[0])
	// fmt.Printf("params: %+v \n", params)
	return nil
}
