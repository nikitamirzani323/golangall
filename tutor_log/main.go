package main

import (
	"log"
	"os"

	logrus "github.com/sirupsen/logrus"
)

var (
	warningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	// InfoLogger.Println("This is some info")
	// warningLogger.Println("This is some warning")
	// ErrorLogger.Println("This is some error")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(
		logrus.Fields{
			"Day":  "Foo",
			"Time": "bar",
		},
	).Info("This is a json message")
	logrus.Info("ini info logrus")
	logrus.Error("ini info logrus")
}
