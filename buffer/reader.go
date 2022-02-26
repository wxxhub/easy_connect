package buffer

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type ReaderBuffer struct {
	conn     *net.TCPConn
	reader   *bufio.Reader
	buff     []byte
	length   int
	index    int
	sequence uint8
}

func (r *ReaderBuffer) Read() ([]byte, error) {
	var tempData []byte
	for {
		head, err := r.ReadNext(4)
		if err != nil {
			return nil, err
		}

		pkgLen := int(uint32(head[0]) | uint32(head[1]) | uint32(head[2]))

		fmt.Println("pkgLen:", pkgLen)
		if head[3] != r.sequence {
			return nil, fmt.Errorf("Error sequence")
		}

		r.sequence += 1
		// 结尾包
		if pkgLen == 0 {
			return tempData, nil
		}

		data, err := r.ReadNext(pkgLen)
		fmt.Println("data:", data)
		if err != nil {
			return nil, err
		}

		if pkgLen < maxPacketSize {
			if tempData == nil {
				return data, nil
			}

			return append(tempData, data...), nil
		}

		tempData = append(tempData, data...)
	}
}

func (r *ReaderBuffer) fill(need int) error {
	n := r.length

	dest := make([]byte, need-n)

	if n > 0 && r.index < n {
		copy(dest[:n], r.buff[r.index:])
	}

	r.buff = dest
	r.index = 0

	for {
		nn, err := r.reader.Read(r.buff[n:])

		n += nn
		switch err {
		case nil:
			if n < need {
				continue
			}

			r.length = n
			return nil
		case io.EOF:
			if n >= need {
				r.length = n
				return nil
			}

			return io.ErrUnexpectedEOF
		default:
			return err
		}
	}

}

func (r *ReaderBuffer) ReadNext(need int) ([]byte, error) {
	if need > r.length {
		if err := r.fill(need); err != nil {
			return nil, err
		}
	}

	offset := r.index
	r.index += need
	r.length -= need

	return r.buff[offset:r.index], nil
}
