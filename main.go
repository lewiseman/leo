package main

import (
	"lavless.com/chat/pkg/controller"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":2020")

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer ln.Close()

	controller.HandleIncomingSockets(ln)

	

}
