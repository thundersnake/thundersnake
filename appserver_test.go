package thundersnake

import (
	"github.com/pborman/getopt/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func reinitAppServerGlobals() {
	getopt.CommandLine = getopt.New()
	// -t is launched by unittests, just add a fake -t, it's not good but it's the only solution there
	var unitTestOpt string
	getopt.FlagLong(&unitTestOpt, "tests", 't', "UnitTests fake")
}

var testOnStartCallbackOk = false
var testOnOptCallbackOk = false

func testOnStartCallback() error {
	testOnStartCallbackOk = true
	return nil
}

func testOptCallback() {
	testOnOptCallbackOk = true
}

func TestNewAppServer(t *testing.T) {
	// Invalid use cases
	reinitAppServerGlobals()
	assert.Nil(t, NewAppServer("", nil, testOnStartCallback))

	reinitAppServerGlobals()
	assert.Nil(t, NewAppServer("", testOptCallback, nil))

	reinitAppServerGlobals()
	assert.Nil(t, NewAppServer("tests", nil, nil))

	reinitAppServerGlobals()
	assert.Nil(t, NewAppServer("tests", testOptCallback, nil))

	reinitAppServerGlobals()
	assert.Nil(t, NewAppServer("", nil, nil))

	// Valid use cases
	reinitAppServerGlobals()
	assert.NotNil(t, NewAppServer("tests", nil, testOnStartCallback))

	reinitAppServerGlobals()
	a := NewAppServer("tests", testOptCallback, testOnStartCallback)
	assert.NotNil(t, a)
	assert.NotEmpty(t, a.buildDate)
	assert.NotEmpty(t, a.version)
}

func TestAppServer_Start(t *testing.T) {
	reinitAppServerGlobals()

	testOnOptCallbackOk = false
	testOnStartCallbackOk = false

	app := NewAppServer("tests", testOptCallback, testOnStartCallback)
	// Ignore later tests in that case, and don't do any tests, it's handled by TestNewAppServer
	if app == nil {
		return
	}

	assert.Equal(t, true, testOnOptCallbackOk)

	app.Start()
	assert.Equal(t, true, testOnStartCallbackOk)
}
