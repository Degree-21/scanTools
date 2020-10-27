package task

type CBlocker interface {

	Run()
}

type CBlock struct {
	ip string
}

func NewCBlock(addr string) CBlocker {
	c := CBlock{ip:addr}
	return  &c
}

func (c *CBlock) Run()  {

}


