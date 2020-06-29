package net_01

import (
	"fmt"
	"net"
)

/**
简单的tcp服务

*/
//
func ServerRun() {
	var (
		err            error
		streamListener net.Listener
		conn           net.Conn
	)
	if streamListener, err = net.Listen("tcp", ":8080"); err != nil {
		panic("listen happen error")
	}
	for {
		if conn, err = streamListener.Accept(); err != nil {
			fmt.Println(err)
		}
		go handler(conn)
	}

}

// hander
func handler(conn net.Conn) {
	var (
		err     error
		recvLen int
	)
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("conn close!")
			conn.Close()
		}
	}()
	buffer := make([]byte, 1024)
	for {
		if recvLen, err = conn.Read(buffer); err != nil {
			fmt.Println("recv error: ", err)
			break
		}
		msg := string(buffer[:recvLen])
		fmt.Println(msg)
	}
}
