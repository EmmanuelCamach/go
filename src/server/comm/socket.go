package comm

import (
	"fmt"
	"net"
)

func Listen(port string) (net.Listener, error) {
	addr := ":" + port
	fmt.Println("Servidor escuchando en", addr)
	return net.Listen("tcp", addr)
}
