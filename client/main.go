package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"io/ioutil"
	"log"
	"net"
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
	planEvent := widget.NewEntry()
	executor := widget.NewEntry()
	count := widget.NewEntry()
	form := widget.NewForm(
		widget.NewFormItem("План. мероприятия", planEvent),
		widget.NewFormItem("Исполнители", executor),
		widget.NewFormItem("Количество обуч.", count),
	)

	btnSave := widget.NewButton("OK", func() {
		dataSend := fmt.Sprintf("%s %s %s", planEvent.Text, executor.Text, count.Text)
		if err := sendData(srvAddr, dataSend); err != nil {
			log.Println(err)
		}

		// clear form fields
		planEvent.SetText("")
		executor.SetText("")
		count.SetText("")
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

//sendData send data to server
func sendData(addr, data string) error {
	// connect with server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("Error connect to server: %v\n", err)
	}

	// send data to server
	if _, err := conn.Write([]byte(data)); err != nil {
		return fmt.Errorf("Error send data to server: %v\n", err)
	}

	// close connection
	return conn.Close()
}
