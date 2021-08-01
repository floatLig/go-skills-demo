package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":2717")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, _ := listen.Accept()
		go handConn(conn)
	}
}

func handConn(conn net.Conn) {
	var data []byte
	_, _ = conn.Read(data)
	fmt.Println(string(data))
}
