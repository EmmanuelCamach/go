package comm

import (
	"bufio"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Receiver struct {
	output *widget.Entry
}

// Crea el Entry de salida
func NewReceiver() (*Receiver, *widget.Entry) {
	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("Mensajes del servidor...")

	return &Receiver{
		output: entry,
	}, entry
}

// Agrega mensajes al Entry usando CallOnMainThread
func (r *Receiver) Append(msg string) {
	fyne.CurrentApp().Driver().CallOnMainThread(func() {
		r.output.SetText(r.output.Text + msg + "\n")
	})
}

// Loop que recibe mensajes desde el servidor
func Receive(conn net.Conn, r *Receiver) {
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			r.Append("Conexión cerrada.")
			return
		}
		r.Append(msg)
	}
}
