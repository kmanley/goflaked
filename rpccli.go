package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var client *rpc.Client
	var err error
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:10001")
	if err != nil {
		log.Fatal("dialing: ", err)
		return
	}
	var reply int
	err = client.Call("Node.DoNothing", 1, &reply)
	if err != nil {
		log.Fatal("DoNothing error:", err)
		return
	}
	fmt.Printf("DoNothing: %d", reply)
}

