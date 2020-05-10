package server

import (
	"net"
)

//Handle 处理连接
func Handle(conn *net.TCPConn) {
	defer conn.Close()
}