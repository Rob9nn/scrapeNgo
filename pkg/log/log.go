package log

import (
	"log"
	"os"
)

var (
  warnLog *log.Logger
  infoLog *log.Logger
  errorLog *log.Logger
)

func init() {
  infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
  warnLog = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
  errorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  infoLog.Println("Logger init done.")
}

func InfoLog(v ...any) {
  infoLog.Println(v...)
}

func ErrorLog(v ...any) {
  errorLog.Panicln(v...)
}
