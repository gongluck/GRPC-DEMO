/*
 * @Author: gongluck
 * @Date: 2020-06-19 09:21:46
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-19 10:51:02
 */

package main

import (
	"log"
	"net"
	"net/rpc"

	"rpc-demo/api"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	api.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
