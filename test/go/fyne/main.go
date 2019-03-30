package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("没啥好说的"),
		widget.NewButton("退出", func() {
			app.Quit()
		})))
}
