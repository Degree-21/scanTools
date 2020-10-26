package lib

import "os"

func AutoLog(fileName string,info string)  {
	filePath := "log/" +fileName +".log"
	_,err :=os.Stat(filePath)
	var f *os.File
	if err != nil{
		f, _ =os.Create(filePath)
	}else {
		f ,_ = os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	}
	f.Write([]byte(info))
}
