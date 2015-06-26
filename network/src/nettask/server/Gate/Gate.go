package main

import (
	Gatelib "nettask/server/Gate/Gatelib"
	tcplib "nettask/tcplib"
)

//---------------------------开始-----------------------------------------//
// start serer

func main() {
	tcplib.TcpSvrCreate(":8000", Gatelib.OnDataReceive, Gatelib.OnClientClose, Gatelib.OnNewUserCome)
}
