package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	prog := app.New()
	ventPrinc := prog.NewWindow("Proyecto Operativos 2025")

	textComando := widget.NewEntry()
	textComando.SetPlaceHolder("Digite comando a ejecutar")
	labelRes := widget.NewLabel("Esperando comando...")
	btn := widget.NewButton("Ejecutar", func() {
		labelRes.SetText("Ejecutando " + textComando.Text)
	})

	textComando.OnSubmitted = func(s string) {
		labelRes.SetText("Ejecutando " + s)
	}

	labelReporte := widget.NewLabel("Esperando reporte...")

	ventPrinc.SetContent(container.NewVBox(
		textComando,
		btn,
		labelRes,
		labelReporte,
	))

	ventPrinc.Resize(fyne.NewSize(420, 200))

	go recReporte(labelReporte)

	ventPrinc.ShowAndRun()
}

func recReporte(labelReporte *widget.Label) {
	cont := 0
	for {
		time.Sleep(5 * time.Second)
		sReporte := "Reporte [" + strconv.Itoa(cont) + "]: DD=45% - RAM=23% - nProc=1234"

		fyne.Do(func() {
			labelReporte.SetText(sReporte)
		})
		cont++
	}
}
