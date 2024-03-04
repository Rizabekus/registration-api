package loggers

import (
	"log"
	"os"
)

var (
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	DebugLog = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime|log.Lshortfile)
)
