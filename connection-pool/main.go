package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	InitConnectionPool(10, func() (conn net.Conn, err error) {
		return net.DialTimeout("tcp", "127.0.0.1:2717", time.Second*30)
	})

	go tcpClient()

	tcpServer()

}

func tcpServer() {
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
	time.Sleep(1 * time.Second)
	conn.Write([]byte("return"))
}

func tcpClient() {
	for i := 0; i < 20; i++ {
		go func() {
			conn, _ := Get()
			conn.Write([]byte("hello"))
			var read []byte
			conn.Read(read)
			fmt.Println(string(read))
			Put(conn)
		}()
	}
}
