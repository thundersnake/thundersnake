package thundersnake

import (
	"github.com/op/go-logging"
	"github.com/pborman/getopt/v2"
	"gitlab.com/thundersnake/thundersnake/httpserver"
	"gitlab.com/thundersnake/thundersnake/utils"
	"os"
	"os/signal"
	"syscall"
)

// AppServerVersion application version (from git tag)
var AppServerVersion = "[unk]"

// AppServerName application server name
var AppServerName = "ThunderSnake"

// AppServer application server object
type AppServer struct {
	name            string
	version         string
	buildDate       string
	logManager      *LogManager
	Log             *logging.Logger
	Config          *Config
	onStartCallBack func() error
	HTTP            *httpserver.HTTPServer
}

// NewAppServer creates AppServer object if basic prerequisites are satisfied
func NewAppServer(appName string, onConfigFlagInitCallback func(), onStartCallBack func() error) *AppServer {
	app := &AppServer{
		name:            appName,
		logManager:      NewLogManager(appName),
		onStartCallBack: onStartCallBack,
		version:         "[unk]",
		buildDate:       "[unk]",
	}

	app.Log = app.logManager.Log

	if len(appName) == 0 {
		app.Log.Errorf("[%s] appName not defined, cannot create AppServer.", AppServerName)
		return nil
	}

	if app.onStartCallBack == nil {
		app.Log.Errorf("[%s] onStartCallback not defined, cannot create AppServer.", AppServerName)
		return nil
	}

	app.Config = &Config{}

	getopt.FlagLong(&app.Config.path, "config", 'c', "Configuration file")
	if onConfigFlagInitCallback != nil {
		onConfigFlagInitCallback()
	}
	return app
}

// Start starts the AppServer
// It will load the configuration, enable AppServer utils & run the onStartCallback function
func (app *AppServer) Start() error {
	app.logManager.start()

	app.Log.Infof("Starting %s version %s (%s/%s).", app.name, app.version, AppServerName, AppServerVersion)
	app.Log.Infof("Build date: %s.", app.buildDate)
	if utils.IsInDocker() {
		app.Log.Infof("Application is running in a Docker container.")
	}

	getopt.Parse()

	app.Config.loadConfiguration()

	if app.Config.EnableSigHUPReload {
		app.listenSigHUPReloadConfig()
	} else {
		app.Log.Info("Configuration reload on SIGHUP is disabled.")
	}

	app.Log.Infof("Application node ID: %s", app.Config.NodeName)

	if app.Config.HTTP.Port > 0 {
		app.HTTP = httpserver.New(app.Log, app.Config.HTTP)
	}

	ret := app.onStartCallBack()
	app.Log.Infof("Exiting %s", app.name)
	return ret
}

func (app *AppServer) listenSigHUPReloadConfig() {
	app.Log.Info("Listening for SIGHUP signal to reload configuration.")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP)

	go func() {
		for sig := range sigs {
			app.Log.Infof("SIGHUP(%s) received, reloading configuration", sig)
			app.Config.loadConfiguration()

			// if sighup reload has been disabled by this reload, close the event listener
			if !app.Config.EnableSigHUPReload {
				close(sigs)
			}
		}
	}()
}
