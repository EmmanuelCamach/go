package comm

import (
	"bufio"
	"net"
)

func SendCommand(conn net.Conn, cmd string) error {
	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString(cmd + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}
