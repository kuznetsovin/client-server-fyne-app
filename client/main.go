package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()

	//create new window
	w := a.NewWindow("Client")

	// create form for send
	form := widget.NewForm(
		widget.NewFormItem("План. мероприятия", widget.NewEntry()),
		widget.NewFormItem("Исполнители", widget.NewEntry()),
		widget.NewFormItem("Количество обуч.", widget.NewEntry()),
	)

	btnSave := widget.NewButton("OK", func() {
		//TODO add server send information
	})

	btnExit := widget.NewButton("Выход", func() {
		a.Quit()
	})

	w.SetContent(
		widget.NewVBox(
			form,
			layout.NewSpacer(), // add empty block
			widget.NewHBox(layout.NewSpacer(), btnSave, btnExit), // add button block
		))

	// change theme and window size
	w.Resize(fyne.NewSize(480, 200))
	a.Settings().SetTheme(theme.LightTheme())

	w.ShowAndRun()
}
