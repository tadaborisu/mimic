package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func gui() {
	mainApp := app.New()
	window := mainApp.NewWindow("Mimic")
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	window.SetIcon(icon)
	window.Resize(fyne.NewSize(220, 250))
	window.CenterOnScreen()
	window.SetFixedSize(true)
	label := widget.NewLabel("Crypt and hide files in photo")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter key")
	window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		label,
		input,
		widget.NewButton("Encrypt", func() {
			key = input.Text
			hide(key)
			label.SetText(massage)
		}),
		widget.NewButton("Decrypt", func() {
			key = input.Text
			dcrypt()
			label.SetText(massage)
		}),
		widget.NewButton("Clean", func() {
			label.SetText("Cleaned")
			clean()
		}),
	))

	window.ShowAndRun()
}
