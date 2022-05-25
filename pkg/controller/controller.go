package controller

import (
	"encoding/json"
	"log"
	"net"

	"lavless.com/chat/pkg/controller/presence"
	"lavless.com/chat/pkg/models"
)

func HandleIncomingSockets(ln net.Listener) {
	conn, err := ln.Accept()

	if err != nil {
		log.Fatalln(err)
		return
	}

	for {
		buffer := make([]byte, 1024)
		size, err := conn.Read(buffer)
		buffer = buffer[:size]

		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(buffer))

		go handleRequest(conn, buffer)
	}
}

func handleRequest(conn net.Conn, data []byte) {
	var in models.Incoming
	err := json.Unmarshal(data, &in)
	if err != nil {
		log.Println("Can;t unmarshal the byte array", err)
		return
	}

	if in.Action == "presence" {
		presence.PresenceHandler(string(data), conn)
	} else {
		log.Println("fail: ", data)
		conn.Write([]byte("fail"))
	}
}

// d := models.Message{
// 	Text:       "Hello, World!",
// 	SenderID:   1,
// 	ReceiverID: 2,
// }
// c := d.CreateMessage()
// fmt.Println(c)
