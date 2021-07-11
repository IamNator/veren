package logger

import (
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

//Logger logs information
var Logger *log.Logger

// var logglyToken = ""

// var url = "http://logs-01.loggly.com/inputs//tag/internal/"

func init() {

	Logger = log.New()

	log.SetFormatter(&log.JSONFormatter{PrettyPrint: true})

	log.WithFields(logrus.Fields{
		"appname": "veren",
	})

	//log.SetOutput(OpenLogfile())

	Logger.Println("logger object created")
}
