package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

// Test Servers:
// ec2-23-23-49-158.compute-1.amazonaws.com:2000
// 177.106.89.125:15555
// demo.cloudwalk.io:2000
// localhost:2000

var httpListen = flag.String("server", "127.0.0.1:2000", "host:port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *httpListen)
	checkError(err)

	a := "2"
	b := "109900"
	c := "1234123412341234"
	d := "1234"
	_, err = conn.Write([]byte(a + "#" + b + "#" + c + "#" + d))
	checkError(err)

	var buf [1024]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	hexa := strings.Split(string(buf[0:n]), "#")
	param1 := hexa[0]
	param2 := hexa[1]
	param3 := hexa[2]
	fmt.Printf("Param1: %s\nParam2: %s\nParam3: %s", param1, param2, param3)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error, %v\n", err.Error())
		os.Exit(1)
	}
}
