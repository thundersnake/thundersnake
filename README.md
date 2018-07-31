[![pipeline status](https://gitlab.com/ThunderSnake/thundersnake/badges/develop/pipeline.svg)](https://gitlab.com/ThunderSnake/thundersnake/commits/develop)
[![coverage report](https://gitlab.com/ThunderSnake/thundersnake/badges/develop/coverage.svg)](https://gitlab.com/ThunderSnake/thundersnake/commits/develop)

# ThunderSnake

ThunderSnake is an application framework designed to bootstrap a Golang application easily
with all required tooling to make it work properly in a production environment.

# Example

Here is how to bootstrap your application

```go
// Define appServer singleton
var appServer *thundersnake.AppServer

func main() {
	configFile := "/etc/myapp/myapp.yml"
	appServer = thundersnake.NewAppServer(AppName, configFile, onStart)
	appServer.Config.Custom = &common.GConfig
	if appServer != nil {
		// Appserver will be started
		// * Logging manager will start
		// * Configuration will be loaded
		// * AppServer hooks are launched
		// * onStart callback function is called
		appServer.Start()
	}
}

// When appServer is ready it will call this callback function
func onStart() error {
	appServer.Log.Info("Application started")
	// Do your main code here
	appServer.Log.Info("Application Ended")
}
```
