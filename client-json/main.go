/*
 * @Author: gongluck
 * @Date: 2020-06-19 11:03:39
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-19 11:07:53
 */

package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// {"method":"HelloService.Hello","params":["hello"],"id":0}
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
