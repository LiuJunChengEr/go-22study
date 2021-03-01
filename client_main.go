package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-22study/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	//client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	//client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	client, err := DialHTTP("tcp",  "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := server.Args{A: 7, B: 8}
	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add errors:", err)
	}
	fmt.Printf("MathService.Add：%d+%d=%d", args.A, args.B, reply)
}

func DialHTTP(network, address string) (*rpc.Client, error) {
	return DialHTTPPath(network, address, rpc.DefaultRPCPath)
}

func DialHTTPPath(network, address, path string) (*rpc.Client, error) {
	var err error
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	io.WriteString(conn, "GET "+path+" HTTP/1.0\n\n")
	// Require successful HTTP response
	// before switching to RPC protocol.
	resp, err := http.ReadResponse(bufio.NewReader(conn), &http.Request{Method: "GET"})
	connected := "200 Connected to JSON RPC"
	if err == nil && resp.Status == connected {
		return jsonrpc.NewClient(conn), nil
	}
	if err == nil {
		err = errors.New("unexpected HTTP response: " + resp.Status)
	}
	conn.Close()
	return nil, &net.OpError{
		Op:   "dial-http",
		Net:  network + " " + address,
		Addr: nil,
		Err:  err,
	}
}