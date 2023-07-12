package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/adasarpan404/simpleEchoServer/config"
	"github.com/adasarpan404/simpleEchoServer/utils"
)

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil

}
func RunServer() {
	fmt.Println("starting a synchronous TCP server on", config.Host, config.Port)

	var concurrentClients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))

	utils.CheckNilError(err, "panic")

	for {
		c, err := lsnr.Accept()
		utils.CheckNilError(err, "panic")
		concurrentClients += 1

		log.Println("Client Connected With Address", c.RemoteAddr(), "Concurrent Clients", concurrentClients)

		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				concurrentClients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", concurrentClients)
				if err == io.EOF {
					break
				}
				log.Println("Error", err)
			}
			log.Println("Command", cmd)
			if err = respond(cmd, c); err != nil {
				log.Print("err write:", err)
			}
		}
	}
}
