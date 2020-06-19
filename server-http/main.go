/*
 * @Author: gongluck
 * @Date: 2020-06-19 11:11:06
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-19 11:12:52
 */

package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	//localhost:1234/jsonrpc
	//{"method":"HelloService.Hello","params":["hello"],"id":0}
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
