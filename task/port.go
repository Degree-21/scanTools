package task

import (
	"blast/lib"
	"bufio"
	"errors"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
)

type PortScan interface {
	addrToIp(addr string) (string, error)
	getPortList() ([]string, error)
	getScanThreadNum() (num int64, err error)
	waitScan()
	Done()
	Run()
}

type PortScanResult struct {
	ip string
	port int
	result error
}

type PortScanTask struct {
	ip string
	taskChanel chan PortScanResult
	wg sync.WaitGroup
	portList []string
}

func NewPortTask(addr string) PortScan {
	task := PortScanTask{}
	ip, err := task.addrToIp(addr)
	if err != nil {
		panic(err)
	}
	// 获取 限制的并发数量
	threadNum, err := task.getScanThreadNum()
	if err != nil{
		panic(err)
	}
	portlist, err := task.getPortList()
	if err != nil {
		panic(err)
	}

	//生成 阻塞channe 池
	if int(threadNum) > len(portlist) {
		task.taskChanel = make(chan PortScanResult,len(portlist))
	}else {
		task.taskChanel = make(chan PortScanResult,threadNum)
	}
	task.ip = ip
	task.wg = sync.WaitGroup{}
	task.portList = portlist

	return &task
}

// 运行扫描程序
func (p *PortScanTask) Run() {
	fmt.Println("加载完毕，准备开始执行")
	go p.Done()
	//生成一个工作池
	for _ ,val := range p.portList{
		p.wg.Add(1)
		go p.ScanPort(val)
	}
	p.waitScan()
}

// 扫描端口
func (p *PortScanTask) ScanPort(port string)  {
	intPort ,_ :=strconv.Atoi(port)
	tcpAddr := net.TCPAddr{IP:net.ParseIP(p.ip),Port:intPort}
	taskResult:= PortScanResult{
		ip:p.ip,
		port:intPort,
	}
	conn , err :=net.DialTCP("tcp",nil,&tcpAddr)
	if err != nil {
		taskResult.result = err
	}else {
		taskResult.result = nil
		defer conn.Close()
	}
	p.taskChanel  <- taskResult
}




//开启 task 监听
func(p *PortScanTask) Done()  {
	for {
		select {
			case info := <-p.taskChanel:
				var saveInfo string
				if info.result != nil{
					saveInfo = fmt.Sprintf("扫描ip为:%s,端口为%d,结果为:%s\r\n",info.ip,info.port,"关闭")
				}else {
					saveInfo = fmt.Sprintf("扫描ip为:%s,端口为%d,结果为:%s\r\n",info.ip,info.port,"打开")
					lib.AutoLog(info.ip,saveInfo)
				}
				fmt.Println(saveInfo)
				p.wg.Done()
		}
	}
}

// 等待扫描
func (p *PortScanTask) waitScan()  {
	p.wg.Wait()
}

// 获取并发线程数 channel 数量
func (p *PortScanTask) getScanThreadNum() (num int64, err error) {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		return 0, err
	}
	key, err := cfg.GetValue("port", "thread")
	if err != nil {
		return 0, err
	}
	intKey, _ := strconv.ParseInt(key, 10, 10)
	if intKey <= 0 {
		return 1, nil
	} else {
		return intKey, nil
	}
}

// 读取端口列表 返回一个 字符串切片
func (p *PortScanTask) getPortList() ([]string, error) {
	cfg, err := goconfig.LoadConfigFile("conf/app.ini")
	if err != nil {
		return []string{}, err
	}
	key, err := cfg.GetValue("port", "file_path")
	if err != nil {
		return []string{}, err
	}

	file, err := os.Open(key)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	portList := make([]string, 0)

	br := bufio.NewReader(file)
	for {
		info, _, end := br.ReadLine()
		if end == io.EOF {
			break
		} else {
			portList = append(portList, string(info))
		}
	}

	if len(portList) == 0 {
		return []string{}, errors.New("port list length is zero")
	}
	return portList, nil
}

// 地址到ip 只支持ipv4
func (p *PortScanTask) addrToIp(addr string) (string, error) {
	if ip := net.ParseIP(addr); ip != nil {
		return ip.String(), nil
	}
	return "", errors.New("Please input correct host ip ")
}
