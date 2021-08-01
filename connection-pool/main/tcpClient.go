package main

import (
	"fmt"
	connectionPool "go-skills-demo/connection-pool"
	"net"
	"time"
)

func main() {
	connectionPool.InitConnectionPool(10, func() (conn net.Conn, err error) {
		return net.DialTimeout("tcp", "127.0.0.1:2717", time.Second*30)
	})
	fmt.Println(connectionPool.Len())

	for i := 0; i < 20; i++ {
		go func() {
			conn, _ := connectionPool.Get()
			conn.Write([]byte("hello"))
			defer connectionPool.Put(conn)
		}()
	}
	time.Sleep(1 * time.Second)
}
