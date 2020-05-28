/* Copyright 2018 Inc. All Rights Reserved. */

/* File : main.go */
/*
modification history
--------------------
2018-05-28 09:19 , by o0TigerLiu0o, create
*/
/*
DESCRIPTION
*/
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"tmp/crawler/learngo/rpcdemo"
)

func main(){
	rpc.Register(rpcdemo.DemoService{})

	listener,err := net.Listen("tcp",":1234")

	if err != nil {
		panic(err)
	}

	for{
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v",err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

/*
  test：
baidudeMacBook-Pro-52:study liuchang$ telnet 1234
Trying 0.0.4.210...
telnet: connect to address 0.0.4.210: No route to host
telnet: Unable to connect to remote host
baidudeMacBook-Pro-52:study liuchang$ telnet localhost 1234
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
{"method":"abc.def"}
{"id":null,"result":null,"error":"rpc: can't find service abc.def"}
{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
{"id":1,"result":0.75,"error":null}
*/