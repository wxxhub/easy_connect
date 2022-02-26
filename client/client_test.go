package client

import (
	"net"
	"strings"
	"testing"

	"github.com/wxxhub/easy_connect/buffer"
)

func TestClient(t *testing.T) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		t.Log(err)
		return
	}

	defer conn.Close()

	writer := buffer.NewWriterBuffer(conn)
	inputInfo := "Hello World"

	for {

		writer.Write([]byte(inputInfo))
		t.Log("write:", inputInfo)

		if strings.ToUpper(inputInfo) == "EXIT" {
			return
		}

		inputInfo = "EXIT"
	}
}
