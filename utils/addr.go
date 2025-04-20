package utils

import (
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

var ServerIP = ""

func init() {
	//fixme In the configuration file, ip takes precedence, if not, get the valid network card ip of the machine
	// if config.Config.ServerIP != "" {
	// 	ServerIP = config.Config.ServerIP
	// 	return
	// }

	// see https://gist.github.com/jniltinho/9787946#gistcomment-3019898
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ServerIP = localAddr.IP.String()
}
