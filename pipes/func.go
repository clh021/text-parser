package pipes

import (
	"fmt"
	"strings"
)

func (p *PipeObj) SC(params []string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}

func (p *PipeObj) Join(params []string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}

func (p *PipeObj) Contain(params []string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}

func (p *PipeObj) Split(params []string) error {
	p.lastArr = strings.Split(p.lastStr, params[0])
	// fmt.Printf("params: %+v \n", params)
	return nil
}
