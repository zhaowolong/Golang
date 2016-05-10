package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "14.17.104.56:8430")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	time.Sleep(time.Second)
	b := []byte("<policy-file-request/>")
	conn.Write(b)
	fmt.Println("send ok %s", string(b))

	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		b := []byte("<policy-file-request/>")
		conn.Write(b)
		fmt.Println("send ok %s", string(b))
		time.Sleep(time.Second)
	}
}

