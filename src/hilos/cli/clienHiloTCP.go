package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

 //var salirC bool = false

func main() {
	fmt.Println("\n--------------------------------------")
	fmt.Println("*     Client Emmanuel - OPERATIVOS       *")
	fmt.Println("---------------------------------------\n")

	serverAddr := "localhost:1510"
	dirTCP, _ := net.ResolveTCPAddr("tcp4", serverAddr)
	fmt.Println("Conectándose al server....")
	var socketC *net.TCPConn
	socketC, _ = net.DialTCP("tcp4", nil, dirTCP)
	fmt.Println("Conectado al server....", socketC.RemoteAddr())

	quitC := make(chan bool)
	go envMensajeC(socketC, quitC)
	go recMensajeC(socketC, quitC)

	limpiarPantalla()
	<-quitC
	socketC.Close()
	fmt.Println("Cliente TCP finalizado")
}

func envMensajeC(socketC net.Conn, quitC chan bool) {
	for {
		fmt.Print("\n\nEscriba el mensaje a enviar: ")
		msjEnv, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		env := bufio.NewWriter(socketC)
		env.WriteString(msjEnv)
		env.Flush()
		fmt.Println("Mensaje enviado!")

		if msjEnv == "bye\n" {
			quitC <- true
			return
		}
	}
}

func recMensajeC(socketC net.Conn, quitC chan bool) {
	for {
		time.Sleep(10 * time.Second)
		msjRec, _ := bufio.NewReader(socketC).ReadString('\n')
		fmt.Println("Mensaje recibido: ", msjRec)

		if msjRec == "bye\n" {
			quitC <- true
			return
		}
	}
}

func limpiarPantalla() {
	print("\033[H\033[2J")
}
