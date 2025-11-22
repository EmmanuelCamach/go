package ui

import (
	"cli/comm"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func StartClient(ip, port string) {
	a := app.New()
	w := a.NewWindow("Cliente Operativos")

	entryCmd := widget.NewEntry()
	entryCmd.SetPlaceHolder("Escribe comando...")

	receiver, entryOutput := comm.NewReceiver()

	btn := widget.NewButton("Enviar", func() {
		cmd := entryCmd.Text
		if cmd == "" {
			return
		}

		comm.SendCommand(globalConn, cmd)
		entryCmd.SetText("")
	})

	// Conectar
	conn, err := comm.Connect(ip, port)
	if err != nil {
		entryOutput.SetText("Error conectando: " + err.Error())
	} else {
		globalConn = conn
		go comm.Receive(conn, receiver)
	}

	w.SetContent(container.NewVBox(
		entryCmd,
		btn,
		entryOutput,
	))

	w.Resize(fyne.NewSize(500, 350))
	w.ShowAndRun()
}

var globalConn net.Conn
