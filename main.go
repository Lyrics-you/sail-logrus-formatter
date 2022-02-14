package main

import (
	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		CharStampFormat: "yyyy-MM-dd HH:mm:ss.SSS zzz",
		FieldsOrder:     []string{"name"},
		Position:        true,
		Colors:          true,
		FieldsColors:    true,
		FieldsSpace:     false,
	})
	logrus.SetReportCaller(false)
	logrus.WithFields(logrus.Fields{
		"name": "sailor",
	}).Info("info msg")

	logrus.WithFields(logrus.Fields{
		"name": "sailor",
	}).Error("error msg")

	logrus.WithFields(logrus.Fields{
		"name": "sailor",
	}).Debug("Debug msg")

	logrus.WithFields(logrus.Fields{
		"name": "sailor",
	}).Warn("error msg")

	l := logrus.New()
	// l.SetOutput(output)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&nested.Formatter{
		TimeStampFormat: "-",
		CallerFirst:     true,
		FieldsSpace:     false,
	})
	l.SetReportCaller(true)

	l.Debug("test1")
	l.Debug("test2")
}
