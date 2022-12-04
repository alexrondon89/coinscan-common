package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type response struct {
	Body     []byte
	Response *http.Response
}

type request struct {
	req *http.Request
}

func New() *request {
	return &request{req: &http.Request{}}
}

func (r *request) AddMethod(method string) *request {
	r.req.Method = method
	return r
}

func (r *request) AddUrl(host, path string) *request {
	r.req.URL = &url.URL{
		Host: host,
		Path: path,
	}
	return r
}

func (r *request) AddHeader(header map[string][]string) *request {
	r.req.Header = header
	return r
}

func (r *request) Exec() (response, error) {
	client := http.Client{}
	respReq, err := client.Do(r.req)
	if err != nil {
		return response{}, err
	}

	defer respReq.Body.Close()
	bodyByte, err := ioutil.ReadAll(respReq.Body)
	if err != nil {
		return response{}, err
	}

	return response{
		Body:     bodyByte,
		Response: respReq,
	}, nil
}
