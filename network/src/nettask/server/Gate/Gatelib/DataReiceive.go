package Gatelib

import (
	"fmt"
	packet "nettask/package"
	tcpsession "nettask/tcpsession"
	util "nettask/util"
)

const (
	FLAG_NON = 0
	FLAG_SVR = 1
	FLAG_CLT = 2
)

type Key struct {
	uuid    string
	nFlag   uint
	uuidSvr string
	gameid  string
}

// 1->client 2->server
func NewKey(id string) *Key {
	key := &Key{
		uuid:   id,
		nFlag:  FLAG_NON,
		gameid: "",
	}
	return key
}

func (self *Key) Destroy() {

}

var (
	UuidToSession   = make(map[string]*tcpsession.Tcpsession)
	GameidToSession = make(map[string]*tcpsession.Tcpsession)
)

// tcp关闭
func OnClientClose(session *tcpsession.Tcpsession) {
	fmt.Println("玩家关闭", session.Conn.RemoteAddr())
	session.Close()
}

// 新玩家进入
func OnNewUserCome(session *tcpsession.Tcpsession) {
	uuid := util.Rand()
	strUuid := uuid.Hex()
	session.SetKey(NewKey(strUuid))
	UuidToSession[strUuid] = session

	//test
	key := session.GetKey().(*Key)
	fmt.Println(session.Conn.RemoteAddr(), "新玩家进入, uuid:", key.uuid)

}

// 收到客户端消息
func OnDataReceive(session *tcpsession.Tcpsession, rpk *packet.Rpacket) {
	key := session.GetKey().(*Key)
	// 注册
	if key.nFlag == FLAG_NON {
		msg, _ := rpk.String()
		gameid, _ := rpk.String()
		key.gameid = gameid
		// server
		if msg == "server" {
			key.nFlag = FLAG_SVR
			GameidToSession[gameid] = session
			fmt.Println("server RegisterOk   ", gameid)
		} else {
			// client
			key.nFlag = FLAG_CLT
			fmt.Println("client RegisterOk  To ", gameid)
		}
		wpk := packet.NewWpacket(packet.NewByteBuffer(64), false)
		wpk.PutString("registerOk")
		session.Send(wpk, OnSendFinish)
		return
	}

	if key.nFlag == FLAG_CLT {
		SendToServer(key, session, rpk)
		return
	}

	if key.nFlag == FLAG_SVR {
		SendToClient(key, session, rpk)
	}

}

// send 回调函数
func OnSendFinish(s interface{}, Wpacket *packet.Wpacket) {
	//session := s.(*tcpsession.Tcpsession)
	fmt.Printf("数据发送完毕\n")

}

func SendToServer(key *Key, session *tcpsession.Tcpsession, rpk *packet.Rpacket) {
	//
	fmt.Println(session.Conn.RemoteAddr(), "发过来消息,Gate转发,")
	server, ok := GameidToSession[key.gameid]
	if !ok {
		fmt.Println("服务器未注册", key.gameid)
		return
	}
	msg, _ := rpk.String()
	wpk := packet.NewWpacket(packet.NewByteBuffer(64), false)
	wpk.PutString(key.uuid)
	wpk.PutString(msg)
	fmt.Println("转发消息", msg)
	server.Send(wpk, OnSendFinish)
}

func SendToClient(key *Key, session *tcpsession.Tcpsession, rpk *packet.Rpacket) {
	fmt.Println("收到了服务器要分发的消息")
	wpk := packet.NewWpacket(packet.NewByteBuffer(64), false)
	uuid, err := rpk.String()
	if err != nil {
		return
	}
	msg, _ := rpk.String()
	wpk.PutString(msg)
	client, ok := UuidToSession[uuid]
	if !ok {
		fmt.Println("玩家已下线", key.uuid)
		return
	}
	client.Send(wpk, OnSendFinish)
}
