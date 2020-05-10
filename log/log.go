package log

import (
	"log"
	"os"
)

//SLog 服务器日志
var SLog *log.Logger

func init() {
	SLog = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
}
