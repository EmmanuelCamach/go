package comm

import (
	"bufio"
	"net"
)

func ReadLine(conn net.Conn) (string, error) {
	r := bufio.NewReader(conn)
	return r.ReadString('\n')
}
