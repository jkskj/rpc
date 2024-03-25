package main

import (
	"log"
	"net"
	test "rpc/kitex_gen/test/test"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:7000")
	svr := test.NewServer(new(TestImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
