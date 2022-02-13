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
		FieldsSpace:     true,
	})

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
}
