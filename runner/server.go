package runner

import (
	"fmt"
	"net"
	"os"
	"time"
)

type CtlServer struct {
	socket  string
	bufsize int
	handler func(string) string
}

func NewCtlSocket(socketPath string) *CtlServer {
	us := CtlServer{socket: socketPath, bufsize: 1024}
	return &us
}

func (this *CtlServer) createServer() {
	os.Remove(this.socket)
	addr, err := net.ResolveUnixAddr("unix", this.socket)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	listener, err := net.ListenUnix("unix", addr)
	defer listener.Close()
	if err != nil {
		panic("Cannot listen to unix domain socket: " + err.Error())
	}
	fmt.Println("Listening on", listener.Addr())
	for {
		c, err := listener.Accept()
		if err != nil {
			panic("Accept: " + err.Error())
		}
		go this.HandleServerConn(c)
	}

}

func (this *CtlServer) HandleServerConn(c net.Conn) {
	defer c.Close()
	buf := make([]byte, this.bufsize)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	result := this.HandleServerContext(string(buf[0:nr]))
	_, err = c.Write([]byte(result))
	if err != nil {
		panic("Writes failed.")
	}
}

func (this *CtlServer) SetContextHandler(f func(string) string) {
	this.handler = f
}

//接收内容并返回结果
func (this *CtlServer) HandleServerContext(context string) string {
	if this.handler != nil {
		return this.handler(context)
	}
	now := time.Now().String()
	return now
}

func (this *CtlServer) StartServer() {
	this.createServer()
}