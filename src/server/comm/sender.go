package comm

import (
	"bufio"
	"net"
)

func Send(conn net.Conn, msg string) error {
	w := bufio.NewWriter(conn)
	_, err := w.WriteString(msg + "\n")
	if err != nil {
		return err
	}
	return w.Flush()
}
