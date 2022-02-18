package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	_ "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func test1() {
	l := logrus.New()
	l.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		CharStampFormat: "yyyy-MM-dd HH:mm:ss.SSS zzz",
		FieldsOrder:     []string{"name"},
		Position:        true,
		Colors:          true,
		FieldsColors:    true,
		FieldsSpace:     false,
		CallerFirst:     true,
	})
	l.SetReportCaller(false)
	l.Info("test 11")
	l.Info("test 12")
}
func test2() {
	l := logrus.New()
	l.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		CharStampFormat: "yyyy-MM-dd HH:mm:ss.SSS zzz",
		FieldsOrder:     []string{"name"},
		Position:        true,
		Colors:          true,
		FieldsColors:    true,
		FieldsSpace:     false,
	})
	l.SetReportCaller(false)
	l.Info("test 21")
	l.Info("test 22")
}
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
		Position:        false,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf("[%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	l.SetReportCaller(true)

	l.Debug("test1")
	l.Debug("test2")

	test1()
	test2()
}
