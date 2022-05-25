package main

import (
	"lavless.com/chat/pkg/controller"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1010")

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer ln.Close()

	controller.HandleIncomingSockets(ln)

	

}
