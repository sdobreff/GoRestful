package main

import (
	"log"
	"os"
)

func FilePrintln(logger *log.Logger, msg string) {
	f, err := os.OpenFile(cfg.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()

	logger.SetOutput(f)
	logger.Println(msg)
}
