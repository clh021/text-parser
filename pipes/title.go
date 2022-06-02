package pipes

import "fmt"

func (p *PipeObj) TitleByStartWith(params []string, txt *[]string) error {
	// res := make(map[string]string)
	fmt.Printf("params: %+v \n", params)
	// fmt.Printf("txt: %+v \n", txt)
	return nil
}

func (p *PipeObj) TitleByNextLineStartWith(params []string, txt *[]string) error {
	fmt.Printf("params: %+v \n", params)
	return nil
}
