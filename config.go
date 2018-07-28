package thundersnake

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	EnableSigHUPReload bool         `yaml:"enable-sighup-reload"`
	Custom             CustomConfig `yaml:"custom"`
	path               string
}

type CustomConfig interface {
	loadDefaults()
}

func (c *Config) loadDefaults() {
	// Place default appserver configurations here
	c.EnableSigHUPReload = true

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
