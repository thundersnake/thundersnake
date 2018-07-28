package thundersnake

import (
	"github.com/op/go-logging"
	"gitlab.com/thundersnake/thundersnake/utils"
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
	onStartCallBack func() error
}

func NewAppServer(appName string, configPath string, onStartCallBack func() error) *AppServer {
	a := &AppServer{
		name:            appName,
		configPath:      configPath,
		logManager:      NewLogManager(appName),
		onStartCallBack: onStartCallBack,
	}

	a.Log = a.logManager.Log

	if len(appName) == 0 {
		a.Log.Errorf("[%s] appName not defined, cannot create AppServer.", AppServerName)
		return nil
	}

	if len(configPath) == 0 {
		a.Log.Errorf("[%s] configPath not defined, cannot create AppServer.", AppServerName)
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

	return app.onStartCallBack()
}
