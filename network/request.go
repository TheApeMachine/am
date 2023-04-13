package network

import (
	"strings"

	"github.com/dgrr/http2"
	"github.com/valyala/fasthttp"
	"github.com/wrk-grp/errnie"
)

type MethodType string

const (
	GET  MethodType = "GET"
	POST MethodType = "POST"
)

type Request struct {
	method   MethodType
	addr     string
	endpoint string
	headers  map[string]string
	handle   *fasthttp.Request
	response *fasthttp.Response
}

func NewRequest(t MethodType, endpoint string) *Request {
	errnie.Trace()

	return &Request{
		method:   t,
		endpoint: endpoint,
		headers:  make(map[string]string),
		handle:   fasthttp.AcquireRequest(),
		response: fasthttp.AcquireResponse(),
	}
}

func (request *Request) AddHeader(key, value string) {
	errnie.Trace()
	errnie.Debugs("AddHeader <-", key, value)
	request.headers[key] = value
}

func (request *Request) Do(payload []byte) []byte {
	errnie.Trace()

	hc := &fasthttp.HostClient{
		Addr:  request.getAddr(),
		IsTLS: true,
	}

	errnie.Handles(http2.ConfigureClient(hc, http2.ClientOpts{}))
	request.handle.Reset()
	request.handle.Header.SetMethod(string(request.method))
	request.handle.URI().Update(request.endpoint)

	for key, value := range request.headers {
		request.handle.Header.Add(key, value)
	}

	request.handle.SetBody(payload)
	errnie.Handles(hc.Do(request.handle, request.response))

	res := request.response.Body()

	return res
}

func (request *Request) getAddr() string {
	errnie.Trace()
	return strings.Split(request.endpoint, "/")[2] + ":443"
}
