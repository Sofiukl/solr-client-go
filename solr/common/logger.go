package solr

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Trace - This is TRACE log
// Info -  This is INFO log
// Debug - This is DEBUG log
// Warning - This is Warning log
// Error - This is Error log
var (
	Trace   *log.Logger
	Info    *log.Logger
	Debug   *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// InitLogger - This instantiates the logger instance for this
// module
func InitLogger(loglevel string) {

	switch loglevel {
	case "TRACE":
		createLogger(os.Stdout, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	case "INFO":
		createLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stdout, os.Stderr)
		break
	case "DEBUG":
		createLogger(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	case "ERROR":
		createLogger(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	}

}

func createLogger(traceHandle io.Writer,
	infoHandle io.Writer,
	debugHandler io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(debugHandler,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// GetInfoLogger - This return the Info logger
func GetInfoLogger() *log.Logger {
	return Info
}

// GetDebugLogger - This return the Debug logger
func GetDebugLogger() *log.Logger {
	return Debug
}

// GetErrorLogger - This return the Error logger
func GetErrorLogger() *log.Logger {
	return Error
}

// GetWarningLogger - This return the Warning logger
func GetWarningLogger() *log.Logger {
	return Warning
}

// GetTraceLogger - This return the Trace logger
func GetTraceLogger() *log.Logger {
	return Trace
}
