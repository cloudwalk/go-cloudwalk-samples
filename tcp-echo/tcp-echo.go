package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":800"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		c, err := listener.Accept()
		if err != nil {
			continue
		}
		go client(c)
	}
}

func client(c net.Conn) {
	defer c.Close()

	var buf [64]byte
	for {
		n, err := c.Read(buf[0:])
		if err != nil {
			return
		}
		if string(n) != "" {
			_, err := c.Write([]byte("CLOUDWALK" + string(buf[0:n])))
			if err != nil {
				return
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error, %v\n", err.Error())
		os.Exit(1)
	}
}

// Para monitorar esse código se está on line
// pode instalar o monit ou mesmo fazer um script
// para verificar a porta 800 e colocar essa verificação
// no cron para cada 2min
