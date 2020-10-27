package tools


type Toolser interface {
	Conn() (err error)
}


type  ConnInfo struct {
	Ip string
	User string
	Password string
	Port int
}

func NewPlus(info ConnInfo) Toolser {
	 return NewDbConn(info)
	//.Conn()
}

