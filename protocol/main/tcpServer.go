package main

import (
	"encoding/json"
	"fmt"
	"go-skills-demo/protocol"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":1999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		var str string
		data := protocol.RpcReceive(conn)
		_ = json.Unmarshal(data, &str)
		fmt.Println(str)
	}
}
