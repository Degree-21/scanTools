package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDb struct {
	conn ConnInfo
	table string
}

func NewDbConn(info ConnInfo)  Toolser {
	m :=MysqlDb{conn:info,table:"mysql"}
	return &m
}


func (m *MysqlDb) Conn() (err error)  {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	connUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",m.conn.User,m.conn.Password,m.conn.Ip,m.conn.Port,m.table)
	db ,err := sql.Open("mysql",connUrl)
	if err != nil{
		return err
	}
	if  err := db.Ping() ; err != nil{
		return err
	}
	return  nil
}
