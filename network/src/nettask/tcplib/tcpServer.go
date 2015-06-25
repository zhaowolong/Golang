package tcplib

import (
	"fmt"
	"net"
	packet "nettask/package"
	Tcp "nettask/tcpsession"
)

func TcpSvrCreate(port string, OndataReceive func(*Tcp.Tcpsession, *packet.Rpacket), OnClientClose func(*Tcp.Tcpsession), OnNewUser func(*Tcp.Tcpsession)) {
	svrAddr := port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", svrAddr)
	if err != nil {
		fmt.Printf("ResolveTCPAddr error")

	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("ListenTCP error")
	}

	for {
		cnn, err := listener.Accept()
		if err != nil {
			continue
		}

		session := Tcp.NewTcpSession(cnn, false)
		fmt.Println("new client come:  ", session.Conn.RemoteAddr())
		OnNewUser(session)
		go Tcp.ProcessSession(session, OndataReceive, OnClientClose)
	}
}
