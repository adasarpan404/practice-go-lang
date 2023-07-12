package main

import (
	"flag"
	"fmt"

	"github.com/adasarpan404/simpleEchoServer/config"
	"github.com/adasarpan404/simpleEchoServer/server"
)

func setUpFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host for Echo Server")
	flag.IntVar(&config.Port, "port", 4099, "It is port for Echo Server")
	flag.Parse()
}
func main() {
	setUpFlags()
	fmt.Println("Server started running for 4099")
	server.RunServer()

}
