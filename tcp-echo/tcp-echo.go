package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

var addr *string = flag.String("h", "0.0.0.0:2000", "address")
var delay *int = flag.Int("d", 0, "delay in secs to process POST")
var random *int = flag.Int("r", 0, "random delay range, -d must be 0 or not setted")

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", *addr)
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
		d := *delay
		if *delay == 0 && *random != 0 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			d = r.Intn(*random)
		}
		time.Sleep(time.Duration(d) * time.Second)
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
