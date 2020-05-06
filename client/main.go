package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//create log file
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// get server add from file
	srv, err := ioutil.ReadFile("server.ini")
	if err != nil {
		log.Fatal(err)
	}

	srvAddr := string(srv)
	log.Println("Server addr: ", srvAddr)

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
		//conn, err := net.Dial("tcp", srvAddr)
		//if err != nil {
		//	log.Fatalf("Error connect to server: %v\n", err)
		//}
		//
		//_ = conn.Close()
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

	// start app
	w.ShowAndRun()
}
