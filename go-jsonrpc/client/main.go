/* Copyright 2020 DXM Inc. All Rights Reserved. */

/* File : main.go */
/*
modification history
--------------------
2020-05-28 09:20 , by liuchang, create
*/
/*
DESCRIPTION
*/
package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"tmp/crawler/learngo/rpcdemo"
)

func main(){
	conn,err := net.Dial("tcp",":1234")
	if nil != err{
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64

	err = client.Call("DemoService.Div",rpcdemo.Args{10,3},&result)
	if nil != err{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div",rpcdemo.Args{10,0},&result)
	if nil != err{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}
