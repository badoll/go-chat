package server

import (
	"net"

	"github.com/badoll/go-chat/data"
	"github.com/badoll/go-chat/log"
)

//Server 服务器
type Server struct {
	Dao  data.Dao
	Conf Config
}

//Init ...
func (s *Server) Init() {
	addr, err := net.ResolveTCPAddr("tcp", ":"+s.Conf.Port)
	if err != nil {
		log.SLog.Panicf("ResolveTCPAddr error: %s", err.Error())
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.SLog.Panicf("ListenTCP error: %s", err.Error())
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.SLog.Printf("AcceptTCP error: %s", err.Error())
			continue
		}
		go Handle(conn)
	}
}
