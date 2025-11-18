package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var salirC bool = false

func main() {
	fmt.Println("\n--------------------------------------")
	fmt.Println("*     Client Emmanuel - OPERATIVOS       *")
	fmt.Println("---------------------------------------\n")

	ptoServer := "1610"
	ipServer := "10.1.29.171"
	dirTCP, _ := net.ResolveTCPAddr("tcp4", ipServer+":"+ptoServer)
	fmt.Println("Conectándose al server....")
	var socketC *net.TCPConn
	socketC, _ = net.DialTCP("tcp4", nil, dirTCP)
	fmt.Println("Conectado al server....", socketC.RemoteAddr())

	for {
		if envMensajeC(socketC) || recMensajeC(socketC) {
			break
		}
	}
}

func envMensajeC(socketC net.Conn) bool {
	fmt.Print("\n\nEscriba el mensaje a enviar: ")
	msjEnv, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	env := bufio.NewWriter(socketC)
	env.WriteString(msjEnv)
	env.Flush()
	fmt.Println("Mensaje enviado!")

	if msjEnv == "bye\n" {
		salirC = true
	}
	return salirC
}

func recMensajeC(socketC net.Conn) bool {
	msjRec, _ := bufio.NewReader(socketC).ReadString('\n')
	fmt.Println("Mensaje recibido: ", msjRec)

	if msjRec == "bye\n" {
		salirC = true
	}
	return salirC
}
