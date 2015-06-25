package main

import (
	"fmt"
	packet "nettask/package"
	tcplib "nettask/tcplib"
	tcpsession "nettask/tcpsession"
)

//--------------------------回调函数----------------------------------//
// send 回调函数
func OnSendFinish(s interface{}, Wpacket *packet.Wpacket) {
	//session := s.(*tcpsession.Tcpsession)
	fmt.Printf("数据发送完毕\n")

}

// 收到客户端消息
func OndataReceive(session *tcpsession.Tcpsession, rpk *packet.Rpacket) {
	msg, _ := rpk.String()
	wpk := packet.NewWpacket(packet.NewByteBuffer(64), false)
	wpk.PutString("helloworld")
	fmt.Println("收到消息", msg)
	session.Send(wpk, OnSendFinish)
	//fmt.Println(string(rpk.Buffer().Bytes())[:rpk.Buffer().Len()])
}

// tcp关闭
func OnClientClose(session *tcpsession.Tcpsession) {
	fmt.Println("玩家关闭", session.Conn.RemoteAddr())
	session.Close()
}

// 新玩家进入
func OnNewUserCome(session *tcpsession.Tcpsession) {
	fmt.Println("新玩家进入:", session.Conn.RemoteAddr())
}

//---------------------------开始-----------------------------------------//
func main() {
	tcplib.TcpSvrCreate(":8000", OndataReceive, OnClientClose, OnNewUserCome)
}
