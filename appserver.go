package thundersnake

import (
	"github.com/op/go-logging"
	"gitlab.com/thundersnake/thundersnake/utils"
	"os"
	"os/signal"
	"syscall"
)

// AppServerVersion application version (from git tag)
var AppServerVersion = "[unk]"

// AppServerName application server name
var AppServerName = "ThunderSnake"

type AppServer struct {
	name            string
	version         string
	buildDate       string
	configPath      string
	logManager      *LogManager
	Log             *logging.Logger
	Config          Config
	onStartCallBack func() error
}

func NewAppServer(appName string, configPath string, onStartCallBack func() error) *AppServer {
	a := &AppServer{
		name:            appName,
		configPath:      configPath,
		logManager:      NewLogManager(appName),
		onStartCallBack: onStartCallBack,
		version:         "[unk]",
		buildDate:       "[unk]",
	}

	a.Log = a.logManager.Log

	if len(appName) == 0 {
		a.Log.Errorf("[%s] appName not defined, cannot create AppServer.", AppServerName)
		return nil
	}

	if a.onStartCallBack == nil {
		a.Log.Errorf("[%s] onStartCallback not defined, cannot create AppServer.", AppServerName)
		return nil
	}

	return a
}

func (app *AppServer) Start() error {
	app.logManager.Start()

	app.Log.Infof("Starting %s version %s (%s/%s).", app.name, app.version, AppServerName, AppServerVersion)
	app.Log.Infof("Build date: %s.", app.buildDate)
	if utils.IsInDocker() {
		app.Log.Infof("Application is running in a Docker container.")
	}

	app.Config.loadConfiguration()

	if app.Config.EnableSigHUPReload {
		app.listenSigHUPReloadConfig()
	}

	return app.onStartCallBack()
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
