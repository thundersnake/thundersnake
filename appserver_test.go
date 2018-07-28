package thundersnake

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCallBackOk bool = false

func testCallback() error {
	testCallBackOk = true
	return nil
}

func TestNewAppServer(t *testing.T) {
	assert.Nil(t, NewAppServer("", "test.yml", testCallback))
	assert.Nil(t, NewAppServer("tests", "", testCallback))
	assert.Nil(t, NewAppServer("tests", "test.yml", nil))
	assert.Nil(t, NewAppServer("tests", "", nil))
	assert.Nil(t, NewAppServer("", "test.yml", nil))
	assert.Nil(t, NewAppServer("", "", nil))
	assert.NotNil(t, NewAppServer("tests", "test.yml", testCallback))
}

func TestAppServer_Start(t *testing.T) {
	app := NewAppServer("tests", "test.yml", testCallback)
	// Ignore later tests in that case, and don't do any tests, it's handled by TestNewAppServer
	if app == nil {
		return
	}

	testCallBackOk = false

	app.Start()
	assert.Equal(t, true, testCallBackOk)
}
