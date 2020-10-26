package task

type CIntserface interface {

	Run()
}

type CBlock struct {
	ip string
}

func NewCBlock(addr string) CIntserface {
	c := CBlock{ip:addr}
	return  &c
}

func (c *CBlock) Run()  {

}


