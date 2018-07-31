[![pipeline status](https://gitlab.com/ThunderSnake/thundersnake/badges/develop/pipeline.svg)](https://gitlab.com/ThunderSnake/thundersnake/commits/develop)
[![coverage report](https://gitlab.com/ThunderSnake/thundersnake/badges/develop/coverage.svg)](https://gitlab.com/ThunderSnake/thundersnake/commits/develop)

# ThunderSnake

ThunderSnake is an application framework designed to bootstrap a Golang application easily
with all required tooling to make it work properly in a production environment.

# Example

## Basic bootstrap

Here is how to bootstrap your application

```go
// Define appServer singleton
var appServer *thundersnake.AppServer

func main() {
	configFile := "/etc/myapp/myapp.yml"
	appServer = thundersnake.NewAppServer(AppName, configFile, onStart)
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

## Custom configuration

We provide an interface if you want to include custom configuration in application configuration file.

First define a configuration structure depending on thundersnake CustomConfig object and implement loadDefaults() function:

```go
type myConfig struct {
	thundersnake.CustomConfig
	foo int `yaml:"foo"`
	bar string `yaml:"bar"`
}

func (c *myConfig) loadDefaults() {
	c.foo = 1
	c.bar = "test"
}
```