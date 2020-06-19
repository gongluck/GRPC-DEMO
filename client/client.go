/*
 * @Author: gongluck
 * @Date: 2020-06-19 09:25:56
 * @Last Modified by: gongluck
 * @Last Modified time: 2020-06-19 10:54:56
 */

package main

import (
	"fmt"
	"log"

	"rpc-demo/api"
)

func main() {
	client, err := api.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
