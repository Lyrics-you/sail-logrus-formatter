package main

import (
	"fmt"

	nested "github.com/Lyrics-you/sail-logrus-formatter"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("\n--- nested-logrus-formatter ---\n\n")
	f := &nested.Formatter{
		HideKeys:        true,
		CharStampFormat: "yyyy-MM-dd HH:mm:ss.SSS zzz",
		FieldsOrder:     []string{"name"},
		Position:        true,
		Colors:          true,
		FieldsColors:    true,
		FieldsSpace:     false,
	}

	printDemo(f, "nested-logrus-formatter")

	fmt.Print("\n--- default logrus formatter ---\n\n")
	printDemo(nil, "default logrus formatter")
}

func printDemo(f logrus.Formatter, title string) {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)

	if f != nil {
		l.SetFormatter(f)
	}

	// enable/disable file/function name
	l.SetReportCaller(false)

	l.Infof("this is %v demo", title)

	lWebServer := l.WithField("component", "web-server")
	lWebServer.Info("starting...")

	lWebServerReq := lWebServer.WithFields(logrus.Fields{
		"req":   "GET /api/stats",
		"reqId": "#1",
	})

	lWebServerReq.Info("params: startYear=2048")
	lWebServerReq.Error("response: 400 Bad Request")

	lDbConnector := l.WithField("category", "db-connector")
	lDbConnector.Info("connecting to db on 10.10.10.13...")
	lDbConnector.Warn("connection took 10s")

	l.Info("demo end.")
}
