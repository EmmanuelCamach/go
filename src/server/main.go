package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"server/auth"
	"server/comm"
	"server/config"
	"server/exec"
	"server/monitor"
)

func main() {

	cfg, err := config.Load("data/server.conf")
	if err != nil {
		panic(err)
	}

	ln, err := comm.Listen(cfg.Port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn, cfg)
	}
}

func handleClient(conn net.Conn, cfg *config.Config) {

	fmt.Println("Cliente conectado:", conn.RemoteAddr().String())

	if !strings.Contains(conn.RemoteAddr().String(), cfg.AllowedIP) {
		conn.Close()
		return
	}

	// Autenticación
	attempts := 0
	authOK := false

	for attempts < cfg.MaxAttempts {
		comm.Send(conn, "Usuario:")
		user, _ := comm.ReadLine(conn)
		user = strings.TrimSpace(user)

		comm.Send(conn, "Password:")
		pass, _ := comm.ReadLine(conn)
		pass = strings.TrimSpace(pass)

		if auth.Validate(user, pass, cfg.UsersFile) {
			authOK = true
			break
		}

		attempts++
		comm.Send(conn, "Credenciales incorrectas.")
	}

	if !authOK {
		comm.Send(conn, "Demasiados intentos. Conexión cerrada.")
		conn.Close()
		return
	}

	comm.Send(conn, "OK: Autenticado.")

	// Lanzar monitor
	go func() {
		for {
			rep := monitor.GetReport()
			comm.Send(conn, rep)
			time.Sleep(5 * time.Second)
		}
	}()

	// Loop comandos
	for {
		cmd, err := comm.ReadLine(conn)
		if err != nil {
			return
		}

		cmd = strings.TrimSpace(cmd)

		if cmd == "bye" {
			conn.Close()
			return
		}

		out := exec.Run(cmd)
		comm.Send(conn, out)
	}
}
