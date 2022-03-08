package sever

import (
	"log"
	"net"
)

func handleoConn(c net.Conn) {
	defer c.Close()

	for {
	}
}

func main() {
	log.Println("Begin Dial...")
	addr := ""
	conn, err := net.Dial("TCP", addr)
	if err != nil {
		log.Println("Dial Error: check address:" + addr)
		return
	}

	defer conn.Close()
	log.Println("Dial Successful")
}
