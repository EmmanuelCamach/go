package comm

import (
	"net"
)

// Connect crea una conexión TCP compatible con IPv4/IPv6
func Connect(ip, port string) (net.Conn, error) {
	addr := net.JoinHostPort(ip, port)
	return net.Dial("tcp", addr)
}
