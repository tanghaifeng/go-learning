package net_02

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

/**
net-02主要是进行了分装
*/

//func ServerRun() {
//	var (
//		err            error
//		streamListener net.Listener
//		conn           net.Conn
//	)
//	if streamListener, err = net.Listen("tcp", ":8080"); err != nil {
//		panic("listen happen error")
//	}
//	for {
//		if conn, err = streamListener.Accept(); err != nil {
//			fmt.Println(err)
//		}
//		go handler(conn)
//	}
//
//}
//
//// hander
//func handler(conn net.Conn) {
//	var (
//		err     error
//		recvLen int
//	)
//	defer conn.Close()
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("conn close!")
//			conn.Close()
//		}
//	}()
//	buffer := make([]byte, 1024)
//	for {
//		if recvLen, err = conn.Read(buffer); err != nil {
//			fmt.Println("recv error: ", err)
//			break
//		}
//		msg := string(buffer[:recvLen])
//		fmt.Println(msg)
//	}
//}

// server
type server struct {
	listener net.Listener
	lock     sync.Mutex
	addr     string
}

// listenTCP
func (s *server) listenTCP() error {
	if len(s.addr) == 0 {
		return errors.New("addr nil")
	}
	var (
		err error
	)
	s.listener, err = net.Listen("tcp", s.addr)
	if s.listener, err = net.Listen("tcp", s.addr); err != nil {
		fmt.Println("listen: ", err)
		return err
	}
	return nil
}
