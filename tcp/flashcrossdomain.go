package main

import (
	"net/http"
	"fmt"
	"net"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)
var ipstr = "127.0.0.1:8430"
func main() {
	filePath, _ := filepath.Abs(os.Args[0])
	if os.Getppid() != 1 {
		fmt.Printf("server start as daemon:%s,%v", filePath, os.Args[1:])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Start()
		os.Exit(0)
	}
	var tcpAddr *net.TCPAddr
	ip := get_external()
	ipstr = ip + ":8430" 
	tcpAddr, _ = net.ResolveTCPAddr("tcp", ipstr)
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}

}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	fmt.Println("a connect connectd %s", ipStr)
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	msg := `<cross-domain-policy>
	<site-control permitted-cross-domain-policies="all"/>
	<allow-access-from domain="*" to-ports="*"/>
	<allow-http-request-headers-from domain="*" headers="*"/>
	</cross-domain-policy>`
	b := []byte(msg)
	conn.Write(b)
}
func get_external() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data)
}

