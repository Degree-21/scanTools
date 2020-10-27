package lib

import (
	"net/http"
)

type Req struct {
	url string
	header map[string]string
	params map[interface{}]interface{}
}


func NewReq(url string) *Req  {
	req := Req{url:url}
	return  &req
}

func(r *Req) setHeader(header map[string]string)  {
	r.header = header
}

func (r *Req) setParams(params map[interface{}]interface{})  {
	r.params = params
}


func (r *Req) SendGetMethod() (res *http.Response , err error) {
	client := &http.Client{}
	request , _ := http.NewRequest("GET",r.url,nil)

	if len(r.header) > 0 {
		for key , val := range r.header{
			request.Header.Add(key,val)
		}
	}

	res ,err =client.Do(request)
	if err != nil{
		return res , nil
	}
	defer res.Body.Close()

	return res,err
}
