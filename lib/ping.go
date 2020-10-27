package lib

import (
	"errors"
	"os/exec"
)

import (
	"bytes"
)


func Ping(host string)  (err error) {
	var buf bytes.Buffer
	//关键代码
	cmd := exec.Command("ping","-c","1",host)
	cmd.Stdout = &buf
	err = cmd.Run()
	if err != nil {
		return  err
	}
	if buf.String() == ""{
		return  errors.New("ping 错误")
	}
	return  nil

}

