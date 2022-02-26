package buffer

import "net"

type WriterBuffer struct {
	conn     *net.TCPConn
	sequence uint8
}

func (w *WriterBuffer) Write(data []byte) error {
	for {
		tempData := []byte{0, 0, 0, 0}
		lenData := len(data)
		tempData[3] = w.sequence
		tempData[0] = byte(lenData >> 16)
		tempData[1] = byte(lenData >> 8)
		tempData[2] = byte(lenData)
		w.sequence += 1

		tempData = append(tempData, data...)
		_, err := w.conn.Write(tempData)
		if err != nil {
			return err
		}

		break

	}
	return nil
}
