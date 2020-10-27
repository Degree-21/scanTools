package task

import (
	"blast/lib"
	"bufio"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type DomainTasker interface {
	getDomain()
	getThreadNum()(threadNum int, err error)
	getPreDict() (dict []string, err error)
	listen()
	verifyUrl() string
}

type DomainTaskResult struct {
	url string
	exits bool
}


type DomainTask struct {
	url string
	result chan DomainTaskResult
	wg sync.WaitGroup
}

func NewDomainTask(url string) DomainTask {
	task := DomainTask{}
	task.url = task.verifyUrl(url)
	threadNum , err := task.getThreadNum()
	if err != nil{
		panic(err)
	}
	task.result = make(chan DomainTaskResult , threadNum)
	return task
}

func (d *DomainTask) Run()  {
	dict ,err := d.getPreDict()
	if err != nil{
		panic(err)
	}
	go d.listen()

	for _, val := range dict{
		d.wg.Add(1)
		go d.scanDomain(val)
	}
	d.wg.Wait()
	close(d.result)
}

func(d *DomainTask) listen()  {
	for {
		select {
			case info:= <-d.result:{
				if info.exits != true{
					display := fmt.Sprintf("扫描子域名:%s,结果:%s。\r\n",info.url,"不存在")
					fmt.Println(display)
				}else {
					display := fmt.Sprintf("扫描子域名:%s,结果:%s。\r\n",info.url,"存在")
					lib.AutoLog(info.url,display)
					fmt.Println(display)
				}
				d.wg.Done()
			}
		}
	}

}

func(d *DomainTask) scanDomain(pre string)  {
	url :="http://" + pre + "." +d.url
	req := lib.NewReq(url)
	reqRes ,err  := req.SendGetMethod()
	taskResult := DomainTaskResult{url:pre + "." +d.url}

	if err != nil ||reqRes == nil ||reqRes.StatusCode != http.StatusOK{
		taskResult.exits = false
	}else {
		taskResult.exits = true
	}

	d.result <- taskResult
}

func (d *DomainTask) getPreDict() (dict []string, err error) {
	cnf, err :=goconfig.LoadConfigFile("conf/app.ini")
	if err != nil{
		return  dict ,nil
	}
	k , err :=cnf.GetValue("domain","file_path")

	if err != nil{
		return  dict ,nil
	}

	f , err := os.Open(k)
	defer f.Close()

	if err != nil{
		return  dict ,nil
	}

	br := bufio.NewReader(f)

	for {
		line ,_, info:= br.ReadLine()
		if info == io.EOF{
			break
		}else {
			dict = append(dict,string(line))
		}
	}
	return  dict ,nil
}


func (d *DomainTask) verifyUrl(url string) string  {
	if strings.HasPrefix(url,"http://") || strings.HasPrefix(url,"https://"){
		url = strings.Replace(url,"http://","",-1)
		url = strings.Replace(url,"https://","",-1)
	}
	return  url
}

func (d *DomainTask) getThreadNum()(threadNum int, err error)  {
	cnf, err :=goconfig.LoadConfigFile("conf/app.ini")
	if err != nil{
		return  threadNum ,nil
	}
	k , err :=cnf.GetValue("domain","thread")

	if err != nil{
		return  threadNum ,nil
	}
	threadNum , _ = strconv.Atoi(k)
	return   threadNum, nil
}
