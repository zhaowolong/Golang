package main

import (
	"fmt"
	packet "nettask/package"
	tcplib "nettask/tcplib"
	tcpsession "nettask/tcpsession"
	"time"
)

var session *tcpsession.Tcpsession

// send 回调函数
func OnSendFinish(s interface{}, Wpacket *packet.Wpacket) {
	//session := s.(*tcpsession.Tcpsession)
	fmt.Printf("数据发送完毕\n")

}

// 收到客户端消息
func OnDataReceive(session *tcpsession.Tcpsession, rpk *packet.Rpacket) {
	//session.Send(packet.NewWpacket(rpk.Buffer(), rpk.IsRaw()), OnSendFinish)
	uuid, _ := rpk.String()
	msg, _ := rpk.String()
	fmt.Println("收到uuid:", uuid, "转发过来的消息：", msg)
	echo := packet.NewWpacket(packet.NewByteBuffer(1024), false)
	echo.PutString(uuid)
	echo.PutString(msg)
	session.Send(echo, OnSendFinish)
}

// tcp关闭
func OnClientClose(session *tcpsession.Tcpsession) {
	fmt.Println("玩家关闭", session.Conn.RemoteAddr())
	session.Close()
}

func main() {

	session = tcplib.TcpClientCreate("127.0.0.1:8000", OnDataReceive, OnClientClose)
	wpk := packet.NewWpacket(packet.NewByteBuffer(1024), false)
	wpk.PutString("server")
	wpk.PutString("502301")
	session.Send(wpk, OnSendFinish)

	fmt.Println("...")
	for {
		time.Sleep(1 * 1e9)
	}
}
