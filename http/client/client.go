package client

import (
	"io"
	"io/ioutil"
	"net/http"
)

type response struct {
	Body     []byte
	Response *http.Response
}

type request struct {
	req *http.Request
}

func New(method string, host, path string, body io.Reader) (*request, error) {
	req, err := http.NewRequest(method, host+path, body)
	if err != nil {
		return nil, err
	}
	return &request{
		req: req,
	}, nil
}

func (r *request) AddHeader(header map[string][]string) *request {
	r.req.Header = header
	return r
}

func (r *request) AddHost(host string) *request {
	r.req.Host = host
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
