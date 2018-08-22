package thundersnake

import (
	"fmt"
	"gitlab.com/thundersnake/thundersnake/httpserver"
	"gitlab.com/thundersnake/thundersnake/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

// Config AppServer configuration object
type Config struct {
	EnableSigHUPReload bool              `yaml:"enable-sighup-reload"`
	NodeName           string            `yaml:"node-name"`
	HTTP               httpserver.Config `yaml:"http"`
	Custom             CustomConfig      `yaml:"custom"`
	path               string
}

// CustomConfig interface permitting to plug a specific configuration object
// for end users apps
type CustomConfig interface {
	loadDefaults()
}

func (c *Config) loadDefaults() {
	// Place default appserver configurations here
	sigHupReload := os.Getenv("SIGHUP_RELOAD")
	if len(sigHupReload) > 0 {
		sigHupReloadTmp, err := strconv.Atoi(sigHupReload)
		if err != nil {
			gLog.Fatalf("SIGHUP_RELOAD environment variable is not an integer. Please fix it.")
		}
		c.EnableSigHUPReload = sigHupReloadTmp > 0
	} else {
		c.EnableSigHUPReload = true
	}

	nodeName := os.Getenv("NODE_NAME")
	if len(nodeName) > 0 {
		c.NodeName = nodeName
	} else {
		c.NodeName = fmt.Sprintf("node-%s", utils.AlphaNumString(8))
	}

	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) > 0 {
		var err error
		c.HTTP.Port, err = strconv.Atoi(httpPort)
		if err != nil {
			gLog.Fatalf("HTTP port must be an integer")
		}
	} else {
		c.HTTP.Port = 8080
	}

	httpEnableAccessLogs := os.Getenv("HTTP_ENABLE_ACCESS_LOGS")
	if len(httpEnableAccessLogs) > 0 {
		eal, err := strconv.Atoi(httpEnableAccessLogs)
		if err != nil || eal < 0 {
			gLog.Fatalf("HTTP enable access logs must be an integer, greater or equal to zero.")
		}

		c.HTTP.EnableAccessLogs = eal > 0
	} else {
		c.HTTP.EnableAccessLogs = true
	}

	httpEnableHealthEndpoint := os.Getenv("HTTP_ENABLE_HEALTH_ENDPOINT")
	if len(httpEnableHealthEndpoint) > 0 {
		ehe, err := strconv.Atoi(httpEnableHealthEndpoint)
		if err != nil || ehe < 0 {
			gLog.Fatalf("HTTP enable health endpoint must be an integer, greater or equal to zero.")
		}

		c.HTTP.EnableHealthEndpoint = ehe > 0
	} else {
		c.HTTP.EnableHealthEndpoint = true
	}

	// then load custom configuration defaults if any
	if c.Custom != nil {
		c.Custom.loadDefaults()
	}
}

func (c *Config) loadConfiguration() {
	// first load defaults
	c.loadDefaults()

	// load configuration from file
	if len(c.path) == 0 {
		gLog.Info("Configuration path is empty, using default configuration.")
		return
	}

	gLog.Infof("Loading configuration from '%s'...", c.path)

	yamlFile, err := ioutil.ReadFile(c.path)
	if err != nil {
		gLog.Fatalf("Failed to read YAML file #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		gLog.Fatalf("Failed to unmashal configuration file. Error was: %v", err)
	}

	gLog.Infof("Configuration loaded from '%s'.", c.path)
}
