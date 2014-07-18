package main

import (
	"fmt"
	"log/syslog"
	"net"
	"os"
	"strconv"
	"strings"
)

const hexLookup = "0123456789ABCDEF"

var (
	param1, param2, param3 string
	count                  int
)

func main() {
	l, err := syslog.New(syslog.LOG_INFO, "transaction-sample")
	checkError(err)
	service := ":2000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	l.Info(strconv.Itoa(os.Getpid()))
	for {
		c, err := listener.Accept()
		if err != nil {
			continue
		}
		l.Info("Client Connected")
		fmt.Println("Client Connected")
		go client(c, l)
	}
}

func client(c net.Conn, l *syslog.Writer) {
	defer c.Close()
	var buf [1024]byte
	for {
		n, err := c.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Printf("%s\n\n", buf[0:n])
		hexa := strings.Split(string(buf[0:n]), "#")
		count += 1

		param3 = strconv.Itoa(count)
		if hexToInt(hexa[0]) != 0 && hexToInt(hexa[1]) != 0 && hexToInt(hexa[2]) != 0 {
			if hexToInt(hexa[1]) < 10000 { //Aproved Credit till 100,00
				if hexToInt(hexa[2]) == 1234123412341234 { //If card number is correct
					if hexToInt(hexa[0]) == 1 { //If is credit
						//Aproved transaction
						param1 = "ACCEPT"
						param2 = "OK"
						l.Info("CREDIT - ACCEPT Value: #{hexa[1]}")
					} else { //If its debit
						fmt.Printf("%s", hexa[3])
						if hexToInt(hexa[3]) == 1234 {
							//password ok, aproved
							param1 = "ACCEPT"
							param2 = "OK"
							l.Info("DEBIT - ACCEPT Value: #{hexa[1]}")
						} else {
							//wrong password
							param1 = "REFUSED"
							param2 = "WRONG_PASSWORD"
							l.Info("DEBIT - REFUSED WRONG_PASSWORD")
						}
					}
				} else {
					// Card or account invalid
					param1 = "REFUSED"
					param2 = "WRONG_CARD/ACCOUNT"
					l.Info("DEBIT - REFUSED WRONG_CARD/ACCOUNT")
				}
			} else {
				//INSUFFICIENT_FOUNDS
				param1 = "REFUSED"
				param2 = "INSUFFICIENT_FOUNDS"
				l.Info("DEBIT - REFUSED INSUFFICIENT_FOUNDS")
			}
		}
		_, err = c.Write([]byte(param1 + "#" + param2 + "#" + param3 + "\n"))
		checkError(err)
	}
	l.Info("CLOSED CONNECTION")
}

func hexToInt(hex string) int {
	value := 0
	multiplier := 1
	for i := len(hex) - 1; i >= 0; i-- {
		value += multiplier * strings.Index(hexLookup, hex[i:i+1])
		multiplier *= len(hexLookup)
	}
	return value
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
