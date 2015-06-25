package tcplib

import (
	"fmt"
	"net"
	packet "nettask/package"
	Tcp "nettask/tcpsession"
)

func TcpClientCreate(server string, OndataReceive func(*Tcp.Tcpsession, *packet.Rpacket), OnClientClose func(*Tcp.Tcpsession)) *Tcp.Tcpsession {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Printf("ResolveTCPAddr error")

	}

	cnn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("DialTCP error")
	}
	fmt.Println("connect success")
	session := Tcp.NewTcpSession(cnn, false)
	go Tcp.ProcessSession(session, OndataReceive, OnClientClose)
	return session
}
