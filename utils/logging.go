package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// ログのフォーマット指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// ログの出力先指定
	log.SetOutput(multiLogFile)
}
