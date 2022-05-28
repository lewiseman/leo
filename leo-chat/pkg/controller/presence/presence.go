package presence

import (

	// "log"
	"net"
	"time"
)

func PresenceHandler(data string, conn net.Conn) {
	time.Sleep(time.Second * 5)
	r := `{"action": "presence", "type": "check"}`
	conn.Write([]byte(r))
}
