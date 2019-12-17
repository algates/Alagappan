package request

import (
	"crypto/tls"
	"errors"
	"github.com/valyala/fasthttp"
	"net"
	"net/http"
	"strconv"
	"time"
)


const (
	DefaultTimeout = 10 * time.Second
	DefaultConnections       = 10000
	DefaultConnectionTimeout = 300 * time.Second
)

var (
	FasthttpClient = &fasthttp.Client{
		MaxConnsPerHost:     DefaultConnections,
		MaxIdleConnDuration: DefaultConnectionTimeout,
		ReadTimeout:         DefaultTimeout,
		WriteTimeout:        DefaultTimeout,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			ClientSessionCache: tls.NewLRUClientSessionCache(0),
		},
		Dial: func(addr string) (net.Conn, error) {
			var dialer = net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 1 * time.Second,
			}
			return dialer.Dial("tcp", addr)
		},
	}
	errMissingLocation = errors.New("missing Location header for http redirect")

	HTTPClient = &http.Client{
		Timeout:DefaultTimeout,

	}
)


func GetHttpClientObj() *fasthttp.Client {
	if FasthttpClient == nil {
		return FasthttpClient
	}
	return FasthttpClient
}

//Assuming only GET is the input
func MakeHTTPRequest(url string) ([]byte,error) {

	client := GetHttpClientObj()

	reqValidate := fasthttp.AcquireRequest()
	respValidate := fasthttp.AcquireResponse()

	reqValidate.SetRequestURI(url)
	reqValidate.Header.SetMethod("GET") // assuming GET is the only input

	errValidate := client.DoTimeout(reqValidate, respValidate, 10* time.Second)
	if errValidate == nil && respValidate.StatusCode() == 200 {
		//do nothing
	} else if errValidate == nil && respValidate.StatusCode() != 200 {
		errValidate = errors.New("Response code not 200 "+strconv.Itoa(respValidate.StatusCode()))
	} else if errValidate != nil {
		//do nothing
	}

	body:=respValidate.Body()
	err:=errValidate
	//fasthttp.ReleaseResponse(respValidate)
	//fasthttp.ReleaseRequest(reqValidate)
	return body,err
}
