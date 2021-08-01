package connection_pool

import (
	"fmt"
	"net"
	"sync"
)

// connectionPool
// 这里使用的是"指针"。
// 且不对外暴露，只外只暴露公共方法
var connectionPool *ConnectionPool

type ConnectionPool struct {
	connectionNum     int64
	connections       chan net.Conn
	connectionCreator func() (conn net.Conn, err error)
	isClosed          bool
	mu                *sync.Mutex
}

func InitConnectionPool(connectionNum int64, connectionCreator func() (conn net.Conn, err error)) {
	connectionPool = &ConnectionPool{
		connectionNum:     connectionNum,
		connections:       make(chan net.Conn, connectionNum),
		connectionCreator: connectionCreator,
		isClosed:          false,
		mu:                &sync.Mutex{},
	}
	for i := 0; i < int(connectionNum); i++ {
		func() {
			conn, err := connectionPool.connectionCreator()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			connectionPool.connections <- conn
		}()
	}
}

func (connectionPool *ConnectionPool) get() (net.Conn, error) {
	var closed bool
	connectionPool.mu.Lock()
	closed = connectionPool.isClosed
	connectionPool.mu.Unlock()

	if closed {
		return nil, fmt.Errorf("[%s]: [%s]", "Get connections fail", "ConnectionPool is already closed")
	}
	return <-connectionPool.connections, nil
}

func (connectionPool *ConnectionPool) put(conn net.Conn) error {
	var closed bool
	connectionPool.mu.Lock()
	closed = connectionPool.isClosed
	connectionPool.mu.Unlock()

	if closed {
		return fmt.Errorf("[%s]:[%s]", "Put connections fail", "ConnectionPool is already closed")
	}
	connectionPool.connections <- conn
	return nil
}

func (connectionPool *ConnectionPool) close() {
	if connectionPool.isClosed {
		return
	}
	connectionPool.mu.Lock()
	if !connectionPool.isClosed {
		connectionPool.isClosed = true
		for i := 0; i < int(connectionPool.connectionNum); i++ {
			conn := <-connectionPool.connections
			err := conn.Close()
			if err != nil {
				fmt.Printf("[%s]:[%s]", "Close connections fail", err.Error())
				continue
			}
		}
	}
	connectionPool.mu.Unlock()
}

func (connectionPool *ConnectionPool) len() int {
	return len(connectionPool.connections)
}

func Get() (net.Conn, error) {
	if connectionPool == nil {
		return nil, fmt.Errorf("[%s]", "Please init pool first")
	}
	fmt.Println("Get connection from connection pool, len():", connectionPool.len())
	return connectionPool.get()
}

func Put(conn net.Conn) error {
	if connectionPool == nil {
		return fmt.Errorf("[%s]", "Please init pool first")
	}
	fmt.Println("Put connection from connection pool, len():", connectionPool.len())
	return connectionPool.put(conn)
}

func Close() {
	if connectionPool == nil {
		return
	}
	connectionPool.close()
}

func Len() int {
	return connectionPool.len()
}
