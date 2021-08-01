package main

import (
	"encoding/json"
	"fmt"
	"go-skills-demo/protocol"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", ":1999", time.Second*100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, _ := json.Marshal("这是一个完整的包")
	for i := 0; i < 100; i++ {
		protocol.RpcSend(conn, data)
	}
}
