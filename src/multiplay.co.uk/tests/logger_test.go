package tests

import (
	"log"
	"os"
	"testing"

	"../main"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func TestDefault(t *testing.T) {
	Warning := log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	main.FilePrintln(Warning, "asd")
}
