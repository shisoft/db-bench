package ruler

import "net"

func Send(socketPath string, content string) string {
	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		panic("Cannot resolve unix addr: " + err.Error())
	}
	c, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		panic("DialUnix failed.")
	}
	_, err = c.Write([]byte(content))
	if err != nil {
		panic("Writes failed.")
	}
	buf := make([]byte, 1024)
	nr, err := c.Read(buf)
	if err != nil {
		panic("Read: " + err.Error())
	}
	return string(buf[0:nr])
}