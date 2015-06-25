package tcpsession

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	packet "nettask/package"
)

var (
	ErrUnPackError = fmt.Errorf("TcpSession: UnpackError")
)

type Tcpsession struct {
	Conn       net.Conn
	Packet_que chan interface{}
	Send_que   chan *packet.Wpacket
	raw        bool
	send_close bool
	ud         interface{}
}

func (this *Tcpsession) IsRaw() bool {
	return this.raw
}

func (this *Tcpsession) SetUd(ud interface{}) {
	this.ud = ud
}

func (this *Tcpsession) Ud() interface{} {
	return this.ud
}

func dorecv(session *Tcpsession) {
	for {
		header := make([]byte, 4)
		n, err := io.ReadFull(session.Conn, header)
		if n == 0 && err == io.EOF {
			close(session.Packet_que)
			break
		} else if err != nil {
			close(session.Packet_que)
			break
		}
		size := binary.LittleEndian.Uint32(header)
		if size > packet.Max_bufsize {
			close(session.Packet_que)
			break
		}

		body := make([]byte, size)
		n, err = io.ReadFull(session.Conn, body)
		if n == 0 && err == io.EOF {
			close(session.Packet_que)
			break
		} else if err != nil {
			close(session.Packet_que)
			break
		}
		fmt.Println("header is ", header)
		fmt.Println("size is ", size)
		fmt.Println("body is ", body)
		pkbuf := make([]byte, size+4)
		copy(pkbuf[:], header[:])
		copy(pkbuf[4:], body[:])
		rpk := packet.NewRpacket(packet.NewBufferByBytes(pkbuf), false)
		session.Packet_que <- rpk
	}
}

func dorecv_raw(session *Tcpsession) {
	for {
		recvbuf := make([]byte, packet.Max_bufsize)
		_, err := session.Conn.Read(recvbuf)
		if err != nil {
			session.Packet_que <- "rclose"
			return
		}
		rpk := packet.NewRpacket(packet.NewBufferByBytes(recvbuf), true)
		session.Packet_que <- rpk
	}
}

func dosend(session *Tcpsession) {
	for {
		wpk, ok := <-session.Send_que
		if !ok {
			return
		}
		begidx := 0
		for {
			n, err := session.Conn.Write(wpk.Buffer().Bytes()[begidx:wpk.Buffer().Len()])
			if err != nil || n == 0 {
				session.send_close = true
				return
			}
			begidx += n
			if begidx >= int(wpk.Buffer().Len()) {
				break
			}
		}
		if wpk.Fn_sendfinish != nil {
			wpk.Fn_sendfinish(session, wpk)
		}
	}
}

func ProcessSession(tcpsession *Tcpsession, process_packet func(*Tcpsession, *packet.Rpacket), session_close func(*Tcpsession)) {
	for {
		msg, ok := <-tcpsession.Packet_que
		if !ok {
			fmt.Printf("client disconnect\n")
			return
		}
		switch msg.(type) {
		case *packet.Rpacket:
			rpk := msg.(*packet.Rpacket)
			process_packet(tcpsession, rpk)
		case string:
			str := msg.(string)
			if str == "rclose" {
				session_close(tcpsession)
				close(tcpsession.Packet_que)
				close(tcpsession.Send_que)
				tcpsession.Conn.Close()
				return
			}
		}
	}
}

func NewTcpSession(conn net.Conn, raw bool) *Tcpsession {
	session := &Tcpsession{Conn: conn, Packet_que: make(chan interface{}, 1024), Send_que: make(chan *packet.Wpacket, 1024), raw: raw, send_close: false}
	if raw {
		go dorecv_raw(session)
	} else {
		go dorecv(session)
	}
	go dosend(session)
	return session
}

func (this *Tcpsession) Send(wpk *packet.Wpacket, send_finish func(interface{}, *packet.Wpacket)) error {
	if !this.send_close {
		wpk.Fn_sendfinish = send_finish
		this.Send_que <- wpk
	}
	return nil
}

func (this *Tcpsession) Close() {
	this.Conn.Close()
}
