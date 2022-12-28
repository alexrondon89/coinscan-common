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

func New(method string, host string, path string) *request {
	return &request{req: &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: "http",
			Host:   host,
			Path:   path,
		},
	},
	}
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
