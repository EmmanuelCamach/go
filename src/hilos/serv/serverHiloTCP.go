package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

//var salirS bool = false

func main() {
	fmt.Println("\n****************************************")
	fmt.Println("*    Programa de Servidor TCP          *")
	fmt.Println("*    Autor: Emmanuel Camacho           *")
	fmt.Println("****************************************\n")

	addr := "localhost:1510"
	dirTCP, _ := net.ResolveTCPAddr("tcp4", addr)
	socketServer, _ := net.ListenTCP("tcp4", dirTCP)
	fmt.Println("Esperando conexiones...")
	socketS, _ := socketServer.Accept()
	fmt.Println("Conexion aceptada!", socketS.RemoteAddr())

	quitS := make(chan bool)

	go recMensajeS(&socketS, quitS)
	go envMensajeS(&socketS, quitS)
	<-quitS

	limpiarPantalla()
	socketS.Close()
	fmt.Println("Servidor TCP finalizado")

}

func recMensajeS(socketS *net.Conn, quitS chan bool) {
	for {
		msjRec, _ := bufio.NewReader(*socketS).ReadString('\n')
		fmt.Println("\tServ# From:[", (*socketS).RemoteAddr(), " -> ", msjRec, "]")

		if msjRec == "bye\n" {
			quitS <- true
			return
		}
	}
}

func envMensajeS(socketS *net.Conn, quitS chan bool) {
	num := 0
	for {
		//fmt.Println("\n\nServ# Digite el mensaje a enviar:")
		//msjEnv, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		time.Sleep(10 * time.Second)
		num++
		msjEnv := "\t[Report " + strconv.Itoa(num) + "]: RAM = 37% - DD = 78% - PROC = 55%\n"
		env := bufio.NewWriter(*socketS)
		env.WriteString(msjEnv)
		env.Flush()
		fmt.Println("Mensaje enviado!")

		if msjEnv == "bye\n" {
			quitS <- true
			return
		}
	}
}

func limpiarPantalla() {
	fmt.Print("\033[H\033[2J")
}
