package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"Ok", "ECHO: " + method + " ~ " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("method", "params1")
	resp2, _ := client2.Call("method", "params2")

	if resp1.Body != "ECHO: method ~ params1" || resp2.Body != "ECHO: method ~ params2" {
		t.Error("IpcClient.call failde. resp1:", resp1, "resp2:", resp2)
	}
}
