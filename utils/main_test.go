package utils

import (
	"math/rand"
	"os"
	"testing"
)

// TestMain unit tests ramp up
func TestMain(m *testing.M) {
	rand.Seed(4269)

	code := m.Run()

	// Deinit code
	os.Exit(code)
}
