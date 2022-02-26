package server

import (
	"net"
	"testing"

	"github.com/wxxhub/easy_connect/buffer"
)

func tcpPipe(t *testing.T, conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		t.Log("Disconneted :", ipStr)
		conn.Close()
	}()

	ReaderBuffer := buffer.NewReaderBuff(conn)
	for {
		data, err := ReaderBuffer.Read()

		if err != nil {
			t.Log(err)
			break
		}

		if len(data) == 4 {
			sig := string(data)
			if sig == "exit" {
				t.Log("reciver:exit, Exit")
			}
			break
		}

		t.Log("recive:", string(data))
	}
}

func TestServer(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")

	if nil != err {
		t.Log(err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if nil != err {
		t.Log(err)
		return
	}
	defer tcpListener.Close()
	t.Log("Server ready to read ...")

	for {
		conn, err := tcpListener.AcceptTCP()

		if err != nil {
			t.Log("Accept failed, ", err.Error())
		} else {
			t.Log("Accept success, ", conn.RemoteAddr())
		}

		go tcpPipe(t, conn)
	}
}
