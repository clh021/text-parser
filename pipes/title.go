package pipes

import "fmt"

func (p *PipeObj) TitleByStartWith(params []string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}

func (p *PipeObj) TitleByNextLineStartWith(params []string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}
