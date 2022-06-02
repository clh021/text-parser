package pipes

type PipeObj struct {
	lastArr  []string
	lastStr  string
	startStr string
}

func (p *PipeObj) GetArr() []string {
	return p.lastArr
}

func (p *PipeObj) GetStr() string {
	return p.lastStr
}

func (p *PipeObj) Start(str string) {
	p.lastStr = str
	p.startStr = str
}
