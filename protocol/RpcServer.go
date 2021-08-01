package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func RpcReceive(conn net.Conn) []byte {
	buf := make([]byte, 65542)
	read, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err.Error())
		return nil
	}
	newBuffer := bytes.NewBuffer(nil)
	newBuffer.Write(buf[:read])

	scanner := bufio.NewScanner(newBuffer)
	scanner.Split(packageSplitFunc)
	for scanner.Scan() {
		return scanner.Bytes()[6:]
	}
	return nil
}

func packageSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if !atEOF && len(data) > 6 && binary.BigEndian.Uint32(data[:4]) == 0x123456 {
		l := binary.BigEndian.Uint16(data[4:6])
		sumLen := int(l) + 6
		return sumLen, data[:sumLen], nil
	}
	return
}
