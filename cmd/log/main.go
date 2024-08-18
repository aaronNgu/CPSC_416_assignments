package main

import (
	"io"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)

	logger := log.New(mw, "Main - ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Hey there")
}
