package buffer

import (
	"bufio"
	"net"
)

const maxPacketSize = 1<<24 - 1

func NewReaderBuff(conn *net.TCPConn) *ReaderBuffer {
	reader := bufio.NewReader(conn)

	return &ReaderBuffer{
		conn:   conn,
		reader: reader,
	}
}

func NewWriterBuffer(conn *net.TCPConn) *WriterBuffer {
	return &WriterBuffer{
		conn: conn,
	}
}
