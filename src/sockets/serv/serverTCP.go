package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var salirS bool = false

func main() {
	fmt.Println("\n****************************************")
	fmt.Println("*    Programa de Servidor TCP          *")
	fmt.Println("*    Autor: Emmanuel Camacho           *")
	fmt.Println("****************************************\n")

	dirTCP, _ := net.ResolveTCPAddr("tcp4", ":1510")
	socketServer, _ := net.ListenTCP("tcp4", dirTCP)
	fmt.Println("Esperando conexiones...")
	socketS, _ := socketServer.Accept()
	fmt.Println("Conexion aceptada!", socketS.RemoteAddr())

	for {
		if recMensajeS(&socketS) || envMensajeS(&socketS) {
			break
		}
	}
}

func recMensajeS(socketS *net.Conn) bool {
	msjRec, _ := bufio.NewReader(*socketS).ReadString('\n')
	fmt.Println("\tServ# From:[", (*socketS).RemoteAddr(), " -> ", msjRec, "]")

	if msjRec == "bye\n" {
		salirS = true
	}
	return salirS
}

func envMensajeS(socketS *net.Conn) bool {
	fmt.Println("\n\nServ# Digite el mensaje a enviar:")
	msjEnv, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	env := bufio.NewWriter(*socketS)
	env.WriteString(msjEnv)
	env.Flush()
	fmt.Println("Mensaje enviado!")

	if msjEnv == "bye\n" {
		salirS = true
	}
	return salirS
}
