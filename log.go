package thundersnake

import (
	"github.com/op/go-logging"
	"gitlab.com/thundersnake/thundersnake/utils"
	"os"
)

// LogManager AppServer logging manager object
type LogManager struct {
	name string
	fmt  logging.Formatter
	Log  *logging.Logger
}

// Global singleton for the library
var gLog *logging.Logger

// NewLogManager initialize logger
func NewLogManager(name string) *LogManager {
	l := &LogManager{
		name: name,
		fmt: logging.MustStringFormatter(
			`%{color}%{time:15:04:05.000} %{shortfunc} - %{level:.5s} %{color:reset} %{message}`,
		),
		Log: logging.MustGetLogger(name),
	}

	// Initiale the singleton
	gLog = l.Log
	return l
}

func (l *LogManager) start() {
	// @TODO use a configuration file
	l.addSyslogBackend()
}

func (l *LogManager) addSyslogBackend() {
	if !utils.IsInDocker() {
		stderrLog := logging.NewLogBackend(os.Stderr, "", 0)
		syslogBackend, err := logging.NewSyslogBackend(l.name)
		if err != nil {
			l.Log.Error("Failed to setup logs syslog backend, ignoring. Error was: ", err.Error())
			return
		}
		logging.SetBackend(logging.NewBackendFormatter(stderrLog, l.fmt), syslogBackend)
	}
}
