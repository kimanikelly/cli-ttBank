package home

import (
	// "fmt"

	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TestHome() {
	app := app.New()

	homeWindow := app.NewWindow("TT Bank: Home")

	homeWindow.Resize(fyne.NewSize(800, 600))

	heading1 := canvas.NewText("TT Bank", color.Black)
	heading1.Alignment = fyne.TextAlignCenter
	heading1.TextSize = 24

	heading2 := canvas.NewText("Welcome to Decentralized Banking", color.Black)
	heading2.Alignment = fyne.TextAlignCenter
	heading2.TextSize = 24

	openAccountBtn := widget.NewButton("Open Account", func() {
		log.Println("Open account button")
	})

	bankingBtn := widget.NewButton("Banking", func() {
		log.Println("Banking button")
	})

	content := container.New(layout.NewVBoxLayout(), heading1, heading2, openAccountBtn, bankingBtn)
	homeWindow.SetContent(content)

	homeWindow.ShowAndRun()

}
