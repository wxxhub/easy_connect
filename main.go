package main

import (
	"fmt"
	"net"

	"github.com/wxxhub/easy_connect/buffer"
)

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println("Disconneted :", ipStr)
		conn.Close()
	}()

	ReaderBuffer := buffer.NewReaderBuff(conn)
	for {
		data, err := ReaderBuffer.Read()

		if err != nil {
			fmt.Println(err)
			break
		}

		if len(data) == 4 {
			sig := string(data)
			if sig == "exit" {
				fmt.Println("reciver:exit, Exit")
			}
			break
		}

		fmt.Println("recive:", string(data))
	}
}

func main() {
	fmt.Println("Test")
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")

	if nil != err {
		fmt.Println(err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		fmt.Println(err)
		return
	}
	defer tcpListener.Close()
	fmt.Println("Server ready to read ...")

	for {
		conn, err := tcpListener.AcceptTCP()

		if err != nil {
			fmt.Println("Accept failed, ", err.Error())
		} else {
			fmt.Println("Accept success, ", conn.RemoteAddr())
		}

		go tcpPipe(conn)
	}
}
