package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func RpcSend(conn net.Conn, data []byte) {
	magicNum := make([]byte, 4)
	binary.BigEndian.PutUint32(magicNum, 0x123456)

	l := len(data)
	lenNum := make([]byte, 2)
	binary.BigEndian.PutUint16(lenNum, uint16(l))

	newBuffer := bytes.NewBuffer(magicNum)
	newBuffer.Write(lenNum)
	newBuffer.Write(data)

	_, err := conn.Write(newBuffer.Bytes())
	if err != nil {
		fmt.Println(err.Error())
	}
}
