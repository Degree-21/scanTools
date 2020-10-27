package main

import (
	"blast/task"
	"flag"
	"fmt"
)

const (
	PortScan = "port"
	CBlockScan = "c"
	DomainScan = "domain"
	Cracker = "cracker"
)



var ScanType string
var ScanAddr string
var CrackerType string

func initFlag(){
	fmt.Println("task start")
	//Scan = ScanFlag{}
	flag.StringVar(&ScanType,"t",PortScan,"Input you scan type")
	flag.StringVar(&ScanAddr,"u","192.168.0.1","Input you scan addr")
	flag.StringVar(&CrackerType,"p","mysql","请输入你要破解功能")
	//flag.StringVar(&ScanAddr,"p","mysql","请输入你要破解功能")
	//flag.StringVar(&ScanFile,"f","","Input you scan file")
}

func main(){
	initFlag()
	flag.Parse()
	//defer
	// todo 收集错误信息
	fmt.Println(ScanType)
	switch ScanType {
		case PortScan :
			portTask := task.NewPortTask(ScanAddr)
			portTask.Run()
		case DomainScan:
			domainTask := task.NewDomainTask(ScanAddr)
			domainTask.Run()
		case Cracker:
			crackerTask := task.NewCrackerTask(ScanAddr,CrackerType)
			crackerTask.Run()
	}

	//// 端口扫描
	//if ScanType == PortScan && ScanAddr != "" {
	//
	//	os.Exit(1)
	//}else if ScanType == CBlockScan && ScanAddr != "" {
	//	//cTask := task.NewCBlock(ScanAddr)
	//}else if ScanType == DomainScan && ScanAddr != ""{
	//	domainTask := task.NewDomainTask(ScanAddr)
	//	domainTask.Run()
	//}
	//
	////lib.Ping("192.187.0.100")
	//
	////子网存活扫描

	


}
