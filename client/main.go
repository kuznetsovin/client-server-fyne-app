package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// parse config
	cfg, err := LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	//create log file
	logFile, err := os.OpenFile(cfg.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Server addr: ", cfg.Srv)

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
		if err := sendData(cfg.Srv, dataSend); err != nil {
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

	log.Printf("Send data: %v\n", data)

	resp := make([]byte, 512)
	respLen, err := conn.Read(resp)
	if err != nil && err != io.EOF {
		return fmt.Errorf("Response error: %v\n", err)
	}

	log.Println("Packet processing: ", string(resp[:respLen]))

	// close connection
	return conn.Close()
}
