package home

import (
	// "fmt"

	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TestHome() {
	app := app.New()

	homeWindow := app.NewWindow("TT Bank: Window")

	homeWindow.Resize(fyne.NewSize(800, 600))

	openAccountBtn := widget.NewButton("Open Account", func() {
		log.Println("Open account button")
	})

	bankingBtn := widget.NewButton("Banking", func() {
		log.Println("Banking button")
	})

	grid := container.New(layout.NewFormLayout(), openAccountBtn, bankingBtn)

	homeWindow.SetContent(grid)

	homeWindow.ShowAndRun()

}
