package server

import (
	"fmt"
	"net"
	"strconv"

	"github.com/adasarpan404/simpleEchoServer/config"
	"github.com/adasarpan404/simpleEchoServer/utils"
)

func RunServer() {
	fmt.Println("starting a synchronous TCP server on", config.Host, config.Port)

	var concurrentClients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	utils.CheckNilError(err, "panic")

}
