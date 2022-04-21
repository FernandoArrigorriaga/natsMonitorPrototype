package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	log "github.com/sirupsen/logrus"

	js "monitor/jetstream"
)

// Possible endpoints for HTTP queries
// varz, connz, routez, subsz, gatewayz, leafz, accountz, and jsz

func main() {
	// nc, err := nats.Connect("nats://172.21.199.202:4222")
	// if err != nil {
	// 	log.Error(err.Error())
	// 	panic(err)
	// }

	a := app.New()
	w := a.NewWindow("Bus Monitor")
	// msg, err := nc.Request("$JS.API.STREAM.NAMES", []byte(""), 10*time.Millisecond)
	// if err != nil {
	// 	log.Error(err.Error())
	// 	panic(err.Error())
	// }

	resp, err := http.Get("http://localhost:8222/jsz?streams=1&consumers=1")
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	defer resp.Body.Close()

	info := widget.NewLabel("INFO desde puerto de Monitoreo")
	var bodyBytes []byte
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err = io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}

	busData := js.GlobalJetStream{}
	err = json.Unmarshal(bodyBytes, &busData)
	if err != nil {
		log.Fatal(err)
	}
	busTotalMessages := widget.NewLabel(fmt.Sprintf("Cantidad de Mensajes en el Bus: %d", busData.Messages))
	busTotalBytes := widget.NewLabel(fmt.Sprintf("Cantidad de Bytes en el Bus: %d", busData.Bytes))
	busTotalStreams := widget.NewLabel(fmt.Sprintf("Cantidad de Streams en el Bus: %d", busData.Streams))
	busTotalConsumers := widget.NewLabel(fmt.Sprintf("Cantidad de Consumidores en el Bus: %d", busData.Consumers))
	var streams []string
	for i := 0; i < busData.Streams; i++ {
		streams = append(streams, busData.AccountDetails[0].StreamDetail[i].Name)
	}

	streamNameLabel := widget.NewLabel(fmt.Sprintf("Streams: %s", streams))
	// onHttp := false
	w.SetContent(container.NewVBox(
		info,
		busTotalMessages,
		busTotalBytes,
		busTotalStreams,
		busTotalConsumers,
		streamNameLabel,
		// widget.NewButton("read data", func() {
		// 	// if onHttp {
		// 	// 	dataLabel.SetText(string(msg.Data))
		// 	// 	info.SetText("INFO desde TCP")
		// 	// } else {
		// 	dataLabel.SetText(httpString)
		// 	info.SetText("INFO desde server HTTP")
		// 	// }
		// 	// onHttp = !onHttp
		// }),
		widget.NewButton("Salir", func() {
			a.Quit()
		}),
	))
	w.ShowAndRun()
}
