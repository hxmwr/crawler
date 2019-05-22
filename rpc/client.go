package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	c, _ := jsonrpc.Dial("tcp", ":1234")
	args := struct {
		A int
		B int
	}{
		A: 123,
		B: 456,
	}
	var reply int
	err := c.Call("Rpcsrv.Add", args, &reply)
	fmt.Println(err)

	print(reply)
}
