/* Copyright 2020 DXM Inc. All Rights Reserved. */

/* File : rpc.go */
/*
modification history
--------------------
2020-05-28 09:19 , by liuchang, create
*/
/*
DESCRIPTION
*/
package jsonrpc

import "errors"

type DemoService struct {}

type Args struct {
	A,B int
}

func (DemoService)Div(args Args,result *float64) error{
	if args.B == 0{
		return errors.New("division by zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}