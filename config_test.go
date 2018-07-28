package thundersnake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestConfig struct {
	CustomConfig
	MustBeTrueByDefault bool `yaml:"must-be-true-by-default"`
}

func (tc *TestConfig) loadDefaults() {
	tc.MustBeTrueByDefault = true
}

func TestConfigDefaultsNoChild(t *testing.T) {
	c := &Config{}
	c.loadDefaults()
	assert.True(t, c.EnableSigHUPReload)
}

func TestConfigDefaultsWithChild(t *testing.T) {
	c := &Config{}
	tc := &TestConfig{}
	c.Custom = tc
	c.loadDefaults()
	assert.True(t, tc.MustBeTrueByDefault)
}
