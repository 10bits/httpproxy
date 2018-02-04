package main

import (
	"log"
	"net/http"

	"github.com/go-httpproxy/httpproxy"
)

func OnError(ctx *httpproxy.Context, when string, err *httpproxy.Error, opErr error) {
	log.Printf("%s %s %s", when, err, opErr)
}

func OnAccept(ctx *httpproxy.Context, w http.ResponseWriter, r *http.Request) bool {
	return false
}

func OnAuth(ctx *httpproxy.Context, user string, pass string) bool {
	if user == "test" && pass == "1234" {
		return true
	}
	return false
}

func OnConnect(ctx *httpproxy.Context, host string) (httpproxy.ConnectAction, string) {
	return httpproxy.ConnectMitm, host
}

func OnRequest(ctx *httpproxy.Context, req *http.Request) (resp *http.Response) {
	log.Printf("Request %s", req.URL.String())
	return
}

func OnResponse(ctx *httpproxy.Context, req *http.Request, resp *http.Response) {
	resp.Header.Add("Via", "test")
}

func main() {
	prx, _ := httpproxy.NewProxy()
	prx.OnError = OnError
	prx.OnAccept = OnAccept
	prx.OnAuth = OnAuth
	prx.OnConnect = OnConnect
	prx.OnRequest = OnRequest
	prx.OnResponse = OnResponse
	prx.MitmChunked = true

	http.ListenAndServe(":8080", prx)
}
