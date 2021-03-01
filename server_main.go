package main

import (
	"go-22study/server"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	//rpc.HandleHTTP()
	//listen, err := net.Listen("tcp", ":1234")
	http.HandleFunc(rpc.DefaultRPCPath, func(rw http.ResponseWriter, r *http.Request) {
		conn, _, err := rw.(http.Hijacker).Hijack()
		if err != nil {
			log.Print("rpc hijacking ", r.RemoteAddr, ": ", err.Error())
			return
		}
		var connected = "200 Connected to JSON RPC"
		io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n")
		jsonrpc.ServeConn(conn)
	})
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	//rpc.Accept(listen)
	http.Serve(listen, nil)
	//for {
	//	conn, err := listen.Accept()
	//	if err != nil {
	//		log.Println("jsonrpc.Serve：accept：", err.Error())
	//	}
	//	go jsonrpc.ServeConn(conn)
	//}
}
