package task

import (
	"blast/lib/tools"
	"fmt"
)

type Crackerer interface {
	Run()
}

var defaultPort =  map[string]int{
	"mysql":3306,
}

type CrackerTask struct {
	ip string
	port int
	function string
}

func NewCrackerTask(ip string , function string) Crackerer {
	c := CrackerTask{ip:ip}
	if v , ok := defaultPort[function] ; ok{
		c.function = function
		c.port = v
	}else {
		panic("协议暂未完善。")
	}
	return  &c
}

func (c *CrackerTask) Run()  {

	//todo 协议补充
	info := tools.ConnInfo{
		c.ip,
		"root",
		"root",
		c.port,
	}
	plus := tools.NewPlus(info)
	err := plus.Conn()
	fmt.Println(err)
}
